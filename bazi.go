package cng

import (
	"strings"

	"github.com/festum/chronos"
)

var diIndex = map[string]int{
	"子": 0, "醜": 1, "寅": 2, "卯": 3, "辰": 4, "巳": 5, "午": 6, "未": 7, "申": 8, "酉": 9, "戌": 10, "亥": 11,
}

var tianIndex = map[string]int{
	"甲": 0, "乙": 1, "丙": 2, "丁": 3, "戊": 4, "己": 5, "庚": 6, "辛": 7, "壬": 8, "癸": 9,
}

//天干強度表
var tiangan = [][]int{
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200},
	{1060, 1060, 1000, 1000, 1100, 1100, 1140, 1140, 1100, 1100},
	{1140, 1140, 1200, 1200, 1060, 1060, 1000, 1000, 1000, 1000},
	{1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000},
	{1100, 1100, 1060, 1060, 1100, 1100, 1100, 1100, 1040, 1040},
	{1000, 1000, 1140, 1140, 1140, 1140, 1060, 1060, 1060, 1060},
	{1000, 1000, 1200, 1200, 1200, 1200, 1000, 1000, 1000, 1000},
	{1040, 1040, 1100, 1100, 1160, 1160, 1100, 1100, 1000, 1000},
	{1060, 1060, 1000, 1000, 1000, 1000, 1140, 1140, 1200, 1200},
	{1000, 1000, 1000, 1000, 1000, 1000, 1200, 1200, 1200, 1200},
	{1000, 1000, 1040, 1040, 1140, 1140, 1160, 1160, 1060, 1060},
	{1200, 1200, 1000, 1000, 1000, 1000, 1000, 1000, 1140, 1140},
}

//地支強度表
var dizhi = []map[string][]int{
	{
		"癸": {1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200, 1200, 1060, 1140},
	}, {
		"癸": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"辛": {200, 228, 200, 200, 230, 212, 200, 220, 228, 248, 232, 200},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"丙": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 342, 318},
		"甲": {840, 742, 798, 840, 770, 700, 700, 728, 742, 700, 700, 840},
	}, {
		"乙": {1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060, 1000, 1000, 1200},
	}, {
		"乙": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"癸": {240, 220, 200, 200, 208, 200, 200, 200, 240, 240, 212, 228},
		"戊": {500, 550, 530, 500, 550, 600, 600, 580, 500, 500, 570, 500},
	}, {
		"庚": {300, 342, 300, 300, 330, 300, 300, 330, 342, 360, 348, 300},
		"丙": {700, 700, 840, 840, 742, 840, 840, 798, 700, 700, 728, 742},
	}, {
		"丁": {1000, 1000, 1200, 1200, 1060, 1140, 1200, 1100, 1000, 1000, 1040, 1060},
	}, {
		"丁": {300, 300, 360, 360, 318, 342, 360, 330, 300, 300, 312, 318},
		"乙": {240, 212, 228, 240, 220, 200, 200, 208, 212, 200, 200, 240},
		"己": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"壬": {360, 330, 300, 300, 312, 318, 300, 300, 360, 360, 318, 342},
		"庚": {700, 798, 700, 700, 770, 742, 700, 770, 798, 840, 812, 700},
	}, {
		"辛": {1000, 1140, 1000, 1000, 1100, 1060, 1000, 1100, 1140, 1200, 1160, 1000},
	}, {
		"辛": {300, 342, 300, 300, 330, 318, 300, 330, 342, 360, 348, 300},
		"丁": {200, 200, 240, 240, 212, 228, 240, 220, 200, 200, 208, 212},
		"戊": {500, 550, 530, 500, 550, 570, 600, 580, 500, 500, 570, 500},
	}, {
		"甲": {360, 318, 342, 360, 330, 300, 300, 312, 318, 300, 300, 360},
		"壬": {840, 770, 700, 700, 728, 742, 700, 700, 840, 840, 724, 798},
	},
}

var heavenlyStemToFiveElement = map[string]string{
	"甲": "木",
	"乙": "木",
	"丙": "火",
	"丁": "火",
	"戊": "土",
	"己": "土",
	"庚": "金",
	"辛": "金",
	"壬": "水",
	"癸": "水",
}

var earthlyBranchToFiveElement = map[string]string{
	"子": "水",
	"醜": "土",
	"寅": "木",
	"卯": "木",
	"辰": "土",
	"巳": "火",
	"午": "火",
	"未": "土",
	"申": "金",
	"酉": "金",
	"戌": "土",
	"亥": "水",
}

//GetFiveElementByHeavenlyStem 五行天干 10 days as a circle
func GetFiveElementByHeavenlyStem(s string) string {
	return heavenlyStemToFiveElement[s]
}

//GetFiveElementByEarthlyBranch 五行地支
func GetFiveElementByEarthlyBranch(s string) string {
	return earthlyBranchToFiveElement[s]
}

// BaZi ...
type BaZi struct {
	eightChars []string
	fiveEles   []string
	xiyong     *XiYong
}

//NewBazi 建立八字
func NewBazi(calendar chronos.Calendar) *BaZi {
	ec := calendar.Lunar().EightCharacter()
	return &BaZi{
		eightChars: ec,
		fiveEles:   baziToWuXing(ec),
	}
}

// String ...
func (z *BaZi) String() string {
	return strings.Join(z.eightChars, "")
}

//RiZhu 日主
func (z *BaZi) RiZhu() string {
	return z.eightChars[4]
}

func (z *BaZi) calcXiYong() {
	z.xiyong = &XiYong{}
	//TODO:need fix
	z.score().calcSimilar().calcHeterogeneous() //.yongShen().xiShen()
}

//XiYong 喜用神
func (z *BaZi) XiYong() *XiYong {
	if z.xiyong == nil {
		z.calcXiYong()
	}
	return z.xiyong
}

//XiYongShen 平衡用神
func (z *BaZi) XiYongShen() string {
	return z.XiYong().Shen()
}

func (z *BaZi) score() *BaZi {
	di := diIndex[z.eightChars[3]]
	for idx, v := range z.eightChars {
		if idx%2 == 0 {
			z.xiyong.AddFiveElementsScore(GetFiveElementByHeavenlyStem(v), tiangan[di][tianIndex[v]])
		} else {
			dz := dizhi[diIndex[v]]
			for k := range dz {
				z.xiyong.AddFiveElementsScore(GetFiveElementByHeavenlyStem(k), dz[k][di])
			}
		}
	}
	return z
}

func baziToWuXing(bazi []string) []string {
	var wx []string
	for idx, v := range bazi {
		if idx%2 == 0 {
			wx = append(wx, GetFiveElementByHeavenlyStem(v))
		} else {
			wx = append(wx, GetFiveElementByEarthlyBranch(v))
		}
	}
	return wx
}

//計算同類
func (z *BaZi) calcSimilar() *BaZi {
	for i := range sheng {
		if heavenlyStemToFiveElement[z.RiZhu()] == sheng[i] {
			z.xiyong.Similar = append(z.xiyong.Similar, sheng[i])
			z.xiyong.SimilarScore = z.xiyong.GetFiveElementsScore(sheng[i])
			if i == 0 {
				i = len(sheng) - 1
				z.xiyong.Similar = append(z.xiyong.Similar, sheng[i])
				z.xiyong.SimilarScore += z.xiyong.GetFiveElementsScore(sheng[i])
			} else {
				z.xiyong.Similar = append(z.xiyong.Similar, sheng[i-1])
				z.xiyong.SimilarScore += z.xiyong.GetFiveElementsScore(sheng[i-1])
			}
			break
		}
	}
	return z
}

//計算異類
func (z *BaZi) calcHeterogeneous() *BaZi {
	for i := range sheng {
		for ti := range z.xiyong.Similar {
			if z.xiyong.Similar[ti] == sheng[i] {
				goto EndSimilar
			}
		}
		z.xiyong.Heterogeneous = append(z.xiyong.Heterogeneous, sheng[i])
		z.xiyong.HeterogeneousScore += z.xiyong.GetFiveElementsScore(sheng[i])
	EndSimilar:
		continue

	}
	return z
}
