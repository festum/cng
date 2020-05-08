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

func (f *fateImpl) MakeName(ctx context.Context) error {
	_logger.Info("MakeName processing...")

	_logger.Debug("Do Head")
	if e := f.out.Head(f.config.FileOutput.Heads...); e != nil {
		return errorWith(e, "write head failed")
	}
	_logger.Debug("Do RunInit")
	if e := f.RunInit(); e != nil {
		return errorWith(e, "init failed")
	}
	_logger.Debug("Do CountWuGeLucky")
	n, e := f.db.CountWuGeLucky()
	if e != nil || n == 0 {
		return errorWith(e, "count total error")
	}
	_logger.Debug("Do getLastCharacter")
	if e := f.getLastCharacter(); e != nil {
		return errorWith(e, "get char failed")
	}
	name := make(chan *Name)
	go func() {
		if e := f.getWugeName(name); e != nil {
			_logger.Error(e)
		}
	}()
	_logger.Debug("got getWugeName")
	var tmpChar []*Character
	//supplyFilter := false
	for n := range name {
		select {
		case <-ctx.Done():
			_logger.Info("end")
			return nil
		default:
		}

		tmpChar = n.FirstName
		tmpChar = append(tmpChar, n.LastName...)
		//filter bazi
		if f.config.SupplyFilterEnabled && !filterXiYong(f.XiYong().Shen(), tmpChar...) {
			// _logger.Debugw("supply", "name", n.String())
			continue
		}
		//filter zodiac
		if f.config.ZodiacFilterEnabled && !filterZodiac(f.born, n.FirstName...) {
			// _logger.Debugw("zodiac", "name", n.String())
			continue
		}
		//filter trigram
		if f.config.TrigramFilterEnabled && !filterYao(n.Trigram(), "凶") {
			// _logger.Debugw("trigram", "name", n.String())
			continue
		}
		ben := n.Trigram().Get(yi.BenGua)
		bian := n.Trigram().Get(yi.BianGua)

		_logger.Debugw("bazi", "born", f.born.LunarDate(), "time", f.born.Lunar().EightCharacter())
		_logger.Debugw("xiyong", "wuxing", n.FiveElements(), "god", f.XiYong().Shen(), "pinheng", f.XiYong())
		_logger.Debugw("ben", "ming", ben.GuaMing, "chu", ben.ChuYaoJiXiong, "er", ben.ErYaoJiXiong, "san", ben.SanYaoJiXiong, "si", ben.SiYaoJiXiong, "wu", ben.WuYaoJiXiong, "liu", ben.ShangYaoJiXiong)
		_logger.Debugw("bian", "ming", bian.GuaMing, "chu", bian.ChuYaoJiXiong, "er", bian.ErYaoJiXiong, "san", bian.SanYaoJiXiong, "si", bian.SiYaoJiXiong, "wu", bian.WuYaoJiXiong, "liu", bian.ShangYaoJiXiong)

		if err := f.out.Write(*n); err != nil {
			return err
		}
		_logger.Debugw(n.String(), "筆畫", n.Strokes(), "拼音", n.PinYin(), "八字", f.born.Lunar().EightCharacter(), "喜用神", f.XiYong().Shen(), "本卦", ben.GuaMing, "變卦", bian.GuaMing)
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
		if f.gender == Female && filterGender(l) {
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
				n.eightChars = NewBazi(f.born)
				name <- n
			}
		}
	}
	return nil
}

func filterGender(lucky *WuGeLucky) bool {
	return lucky.NeutralGender == true
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
