package cng

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/festum/chronos"
	"github.com/festum/cng/config"
	"github.com/festum/cng/logger"
	"github.com/festum/yi"
	"github.com/xormsharp/xorm"
)

var _logger = logger.NewLogger()

type HandleOutputFunc func(name Name)

type Gender int

const (
	Male Gender = iota + 1
	Female
)

const HelpContent = "正在使用Fate生成姓名列表，如遇到問題請訪問項目地址：https://github.com/godcong/fate獲取幫助!"

type Fate interface { //FIXME: nonsense interface
	MakeName(ctx context.Context) (e error)
	XiYong() *XiYong
	RunInit() (e error)
	RegisterHandle(outputFunc HandleOutputFunc)
}

type fateImpl struct { //FIXME: simplify it
	config   *config.Config
	db       Database
	out      Information
	born     chronos.Calendar
	last     []string
	lastChar []*Character
	names    []*Name
	nameType int
	gender   Gender
	debug    bool
	baZi     *BaZi
	zodiac   *Zodiac
	handle   HandleOutputFunc
}

// RunInit ...
func (f *fateImpl) RunInit() (e error) {
	if f.config.RunInit {
		if err := f.db.Sync(WuGeLucky{}); err != nil {
			return err
		}

		lucky := make(chan *WuGeLucky)
		go initWuGe(lucky)
		for la := range lucky {
			_, e = f.db.InsertOrUpdateWuGeLucky(la)
			if e != nil {
				return errorWith(e, "insert failed")
			}
		}
	}
	return nil
}

type Options func(f *fateImpl)

func ConfigOption(cfg *config.Config) Options {
	return func(f *fateImpl) {
		f.config = cfg
	}
}

func SetGender(gender Gender) Options {
	return func(f *fateImpl) {
		f.gender = gender
	}
}

func Debug() Options {
	return func(f *fateImpl) {
		f.debug = true
	}
}

func NewFate(lastName string, born time.Time, options ...Options) Fate {
	f := &fateImpl{
		last: strings.Split(lastName, ""),
		born: chronos.New(born),
	}
	f.lastChar = make([]*Character, len(f.last))
	if len(f.last) > 2 {
		panic("last name was bigger than 2 characters")
	}

	for _, op := range options {
		op(f)
	}

	f.init()

	return f
}

// RegisterHandle ...
func (f *fateImpl) RegisterHandle(outputFunc HandleOutputFunc) {
	f.handle = outputFunc
}

func (f *fateImpl) getLastCharacter() error {
	size := len(f.last)
	if size == 0 {
		return errors.New("last name was not inputted")
	} else if size > 2 {
		return fmt.Errorf("%d characters last name was not supported", size)
	} else {
		//ok
	}

	for i, c := range f.last {
		character, e := f.db.GetCharacter(Char(c))
		if e != nil {
			return e
		}
		f.lastChar[i] = character
	}
	return nil
}

func (f *fateImpl) MakeName(ctx context.Context) (e error) {
	_logger.Info(HelpContent)
	e = f.out.Head(f.config.FileOutput.Heads...)
	if e != nil {
		return errorWith(e, "write head failed")
	}
	e = f.RunInit()
	if e != nil {
		return errorWith(e, "init failed")
	}
	n, e := f.db.CountWuGeLucky()
	if e != nil || n == 0 {
		return errorWith(e, "count total error")
	}

	e = f.getLastCharacter()
	if e != nil {
		return errorWith(e, "get char failed")
	}
	name := make(chan *Name)
	go func() {
		e := f.getWugeName(name)
		if e != nil {
			_logger.Error(e)
		}
	}()

	var tmpChar []*Character
	//supplyFilter := false
	for n := range name {
		select {
		case <-ctx.Done():
			_logger.Info("end")
			return
		default:
		}

		tmpChar = n.FirstName
		tmpChar = append(tmpChar, n.LastName...)
		//filter bazi
		if f.config.SupplyFilter && !filterXiYong(f.XiYong().Shen(), tmpChar...) {
			//_logger.Infow("supply", "name", n.String())
			continue
		}
		//filter zodiac
		if f.config.ZodiacFilter && !filterZodiac(f.born, n.FirstName...) {
			//_logger.Infow("zodiac", "name", n.String())
			continue
		}
		//filter bagua
		if f.config.BaguaFilter && !filterYao(n.BaGua(), "凶") {
			//_logger.Infow("bagua", "name", n.String())
			continue
		}
		ben := n.BaGua().Get(yi.BenGua)
		bian := n.BaGua().Get(yi.BianGua)
		if f.debug {
			_logger.Infow("bazi", "born", f.born.LunarDate(), "time", f.born.Lunar().EightCharacter())
			_logger.Infow("xiyong", "wuxing", n.WuXing(), "god", f.XiYong().Shen(), "pinheng", f.XiYong())
			_logger.Infow("ben", "ming", ben.GuaMing, "chu", ben.ChuYaoJiXiong, "er", ben.ErYaoJiXiong, "san", ben.SanYaoJiXiong, "si", ben.SiYaoJiXiong, "wu", ben.WuYaoJiXiong, "liu", ben.ShangYaoJiXiong)
			_logger.Infow("bian", "ming", bian.GuaMing, "chu", bian.ChuYaoJiXiong, "er", bian.ErYaoJiXiong, "san", bian.SanYaoJiXiong, "si", bian.SiYaoJiXiong, "wu", bian.WuYaoJiXiong, "liu", bian.ShangYaoJiXiong)
		}

		if err := f.out.Write(*n); err != nil {
			return err
		}
		if f.debug {
			_logger.Infow(n.String(), "筆畫", n.Strokes(), "拼音", n.PinYin(), "八字", f.born.Lunar().EightCharacter(), "喜用神", f.XiYong().Shen(), "本卦", ben.GuaMing, "變卦", bian.GuaMing)
		}
	}
	return nil
}

// XiYong ...
func (f *fateImpl) XiYong() *XiYong {
	if f.baZi == nil {
		f.baZi = NewBazi(f.born)
	}
	return f.baZi.XiYong()
}

func (f *fateImpl) init() {
	if f.config == nil {
		f.config = config.DefaultConfig()
	}

	if f.config.FileOutput.Heads == nil {
		f.config.FileOutput.Heads = config.DefaultHeads
	}

	f.db = initDatabaseWithConfig(f.config.Database)
	f.out = initOutputWithConfig(f.config.FileOutput)
}

//SetBornData 設定生日
func (f *fateImpl) SetBornData(t time.Time) {
	f.born = chronos.New(t)
}

func (f *fateImpl) getWugeName(name chan<- *Name) (e error) {
	defer func() {
		close(name)
	}()
	lucky := make(chan *WuGeLucky)
	go func() {
		e = f.db.FilterWuGe(f.lastChar, lucky)
		if e != nil {
			_logger.Error(e)
			return
		}
	}()
	var f1s []*Character
	var f2s []*Character
	for l := range lucky {
		if f.config.FilterMode == config.FilterModeCustom {
			//TODO
		}

		if f.gender == Female && filterSex(l) {
			continue
		}

		if f.config.HardFilter && hardFilter(l) {
			sc := NewSanCai(l.TianGe, l.RenGe, l.DiGe)
			if !Check(f.db.Database().(*xorm.Engine), sc, 5) {
				continue
			}
		}

		if f.config.StrokeMin > l.FirstStroke1 || f.config.StrokeMin > l.FirstStroke2 || f.config.StrokeMax < l.FirstStroke1 || f.config.StrokeMax < l.FirstStroke2 {
			continue
		}

		if f.debug {
			_logger.Infow("lucky", "l1", l.LastStroke1, "l2", l.LastStroke2, "f1", l.FirstStroke1, "f2", l.FirstStroke2)
		}
		if f.config.Regular {
			f1s, e = f.db.GetCharacters(Stoker(l.FirstStroke1, Regular()))
		} else {
			f1s, e = f.db.GetCharacters(Stoker(l.FirstStroke1))
		}

		if e != nil {
			return errorWith(e, "first stroke1 error")
		}

		if f.config.Regular {
			f2s, e = f.db.GetCharacters(Stoker(l.FirstStroke2, Regular()))
		} else {
			f2s, e = f.db.GetCharacters(Stoker(l.FirstStroke2))
		}

		if e != nil {
			return errorWith(e, "first stoke2 error")
		}

		for _, f1 := range f1s {
			if len(f1.PinYin) == 0 {
				continue
			}
			for _, f2 := range f2s {
				if len(f2.PinYin) == 0 {
					continue
				}
				n := createName(f, f1, f2)
				n.baZi = NewBazi(f.born)
				name <- n
			}
		}
	}
	return nil
}

func filterSex(lucky *WuGeLucky) bool {
	return lucky.ZongSex == true
}

func isLucky(s string) bool {
	if strings.Compare(s, "吉") == 0 || strings.Compare(s, "半吉") == 0 {
		return true
	}
	return false
}

func hardFilter(lucky *WuGeLucky) bool {
	if !isLucky(GetDaYan(lucky.DiGe).Lucky) ||
		!isLucky(GetDaYan(lucky.RenGe).Lucky) ||
		!isLucky(GetDaYan(lucky.WaiGe).Lucky) ||
		!isLucky(GetDaYan(lucky.ZongGe).Lucky) {
		return true
	}
	return false
}
