package cng

import "github.com/xormsharp/xorm"

const sanCai = "水木木火火土土金金水"
const yinYang = "陰陽"

// SanCai ...
type SanCai struct {
	tianCai        string `bson:"tian_cai"`
	tianCaiYinYang string `bson:"tian_cai_yin_yang"`
	renCai         string `bson:"ren_cai"`
	renCaiYinYang  string `bson:"ren_cai_yin_yang"`
	diCai          string `bson:"di_cai"`
	diCaiYingYang  string `bson:"di_cai_ying_yang"`
	fortune        string `bson:"fortune"` //吉凶
	comment        string `bson:"comment"` //說明
}

//NewSanCai 新建一個三才物件
func NewSanCai(tian, ren, di int) *SanCai {
	return &SanCai{
		tianCai:        sanCaiAttr(tian),
		tianCaiYinYang: yinYangAttr(tian),
		renCai:         sanCaiAttr(ren),
		renCaiYinYang:  yinYangAttr(ren),
		diCai:          sanCaiAttr(di),
		diCaiYingYang:  yinYangAttr(di),
	}
}

//Check 檢查三才屬性
func Check(engine *xorm.Engine, cai *SanCai, point int) bool {
	wx := FindWuXing(engine, cai.tianCai, cai.renCai, cai.diCai)
	if wx.Luck.Point() >= point {
		return true
	}
	return false
}

// GenerateThreeTalent 計算字元的三才屬性
// 1-2木：1為陽木，2為陰木   3-4火：3為陽火，4為陰火   5-6土：5為陽土，6為陰土   7-8金：7為陽金，8為陰金   9-10水：9為陽水，10為陰水
func sanCaiAttr(i int) string {
	return string([]rune(sanCai)[i%10])
}

func yinYangAttr(i int) string {
	return string([]rune(yinYang)[i%2])
}
