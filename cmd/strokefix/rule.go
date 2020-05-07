package main

import (
	"strings"

	"github.com/festum/cng"
)

var charCharList = []string{
	`阧障隴陾階阾降隄隓陞隂陳隯陖隗陠阱陲阼隩阬陳陌隀阠陿隉陘隥隉陁陱階陂陽陼陪陗陊阪阞陑隭阬陸隡陫隯陂陽防阽陔隄防隕隕隴陑陚隬陘隣陝隝` +
		`阺阫隢陹限附陰阮隚隥阪阭陏陳陜陸阦阷阞隫隞隘陒隴隍陶隀隔陀阤隤陬隨除陲降陮阽阩陰陋隊隳隊隊隭陮隋隝隴陘阥隆陡陏陫陙陳隧阸陯隑隍` +
		`陡陛阸隔阻隑隯陘陃陊除陖隘隫阨阨隌陔阧陵附陟隞隊隨隌陙陃`,
	`邤酆那鄟鄶鄺邯酃邸邢郟鄧郤郕邚郫鄀郠郂鄝邩酈邽部郚鄽鄹鄚邲鄳郕郥邶邪鄡邡郿邠鄱邞邩邗鄯鄤邯郎鄐鄲鄁鄧酇郲鄽郩鄴郔郭郜鄻邦鄺郘鄋酈郡` +
		`鄷鄇都鄲鄂鄥郣鄶郻鄧鄕邒郭郳鄣郷鄆邞鄶邿鄮鄗鄏郗鄒鄔郪邷郲郈鄼鄄郙邰鄛邦郱鄆邜鄈鄁鄵郵郮邨郠鄌鄲阿邾郈鄝郛郄郂鄸邨郴鄲邴邵邥部` +
		`邴鄖郞郟邘邟鄍郹郖郍鄿鄴郰鄘鄶郙邭邛郟郡鄙鄺鄭邸鄗鄉邠郀郬邼邖邼鄄郹鄤酁邽鄈鄩邔酈鄰鄇邳郴郝郀邲郃郞鄜邙邭郯鄺郜鄑郝邟郣鄐邱鄛` +
		`鄒郢鄰鄷郆郅酇酄鄊郊鄧郟郇郊邗郖鄭鄫鄖郉酆酄郵鬱郋邧都邶邡`,
}

func CharChar(ch *cng.Character) bool {
	if i := strings.Index(charCharList[0], ch.Ch); i != -1 {
		if ch.Stroke != 0 {
			ch.ScienceStroke = ch.Stroke + 8
			return true
		}
		if ch.SimpleTotalStroke != 0 {
			ch.ScienceStroke = ch.SimpleTotalStroke + 8
			return true
		}
	}

	if i := strings.Index(charCharList[0], ch.Ch); i != -1 {
		if ch.Stroke != 0 {
			ch.ScienceStroke = ch.Stroke + 7
			return true
		}
		if ch.SimpleTotalStroke != 0 {
			ch.ScienceStroke = ch.SimpleTotalStroke + 7
			return true
		}
	}

	return false
}

var numberCharList = `一二三四五六七八九十`

func NumberChar(ch *cng.Character) bool {
	if i := strings.Index(numberCharList, ch.Ch); i != -1 {
		if ch.Stroke != 0 {
			ch.ScienceStroke = ch.Stroke + i
			return true
		}
		if ch.SimpleTotalStroke != 0 {
			ch.ScienceStroke = ch.SimpleTotalStroke + i
			return true
		}
	}
	return false
}

var radicalCharList = map[string]int{
	"扌": 1,
	"忄": 1,
	"氵": 1,
	"犭": 1,
	"礻": 1,
	"王": 1,
	"艹": 3,
	"衤": 1,
	"月": 3,
	"辶": 4,
}

func RadicalChar(ch *cng.Character) bool {
	for k, v := range radicalCharList {
		if strings.Compare(ch.Radical, k) == 0 {
			ch.ScienceStroke = ch.Stroke + v
			return true
		}
		if strings.Compare(ch.SimpleRadical, k) == 0 {
			ch.ScienceStroke = ch.SimpleTotalStroke + v
			return true
		}
	}
	return false
}

var fixTable = []func(character *cng.Character) bool{
	RadicalChar,
	NumberChar,
	CharChar,
	DefaultChar,
}

func fixChar(character *cng.Character) bool {
	for _, f := range fixTable {
		if f(character) {
			return true
		}
	}
	return false
}

func DefaultChar(character *cng.Character) bool {
	if character.KangXiStroke != 0 {
		character.ScienceStroke = character.KangXiStroke
	} else if character.TraditionalTotalStroke != 0 {
		character.ScienceStroke = character.TraditionalTotalStroke
	} else if character.Stroke != 0 {
		character.ScienceStroke = character.Stroke
	} else if character.SimpleTotalStroke != 0 {
		character.ScienceStroke = character.SimpleTotalStroke
	} else {
		return false
	}
	return true
}
