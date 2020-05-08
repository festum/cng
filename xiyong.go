package cng

//XiYong 喜用神
type XiYong struct {
	FiveElementsScore  map[string]int
	Similar            []string //同类
	SimilarScore       int
	Heterogeneous      []string //异类
	HeterogeneousScore int
}

var sheng = []string{"木", "火", "土", "金", "水"}
var ke = []string{"木", "土", "水", "火", "金"}

//AddFiveElementsScore 五行分
func (xy *XiYong) AddFiveElementsScore(s string, point int) {
	if xy.FiveElementsScore == nil {
		xy.FiveElementsScore = make(map[string]int)
	}

	if v, b := xy.FiveElementsScore[s]; b {
		xy.FiveElementsScore[s] = v + point
	} else {
		xy.FiveElementsScore[s] = point
	}
}

//GetFiveElementsScore 取得分
func (xy *XiYong) GetFiveElementsScore(s string) (point int) {
	if xy.FiveElementsScore == nil {
		return 0
	}
	if v, b := xy.FiveElementsScore[s]; b {
		return v
	}
	return 0
}

func (xy *XiYong) minFiveElementsScore(ss ...string) (wx string) {
	// concatenate 5e for firstnames
	min := 9999
	for _, s := range ss {
		if xy.FiveElementsScore[s] < min {
			min = xy.FiveElementsScore[s] //FIXME: score always 2000
			wx = s
		} else if xy.FiveElementsScore[s] == min {
			wx += s
		}
	}
	return
}

//Shen 喜用神
func (xy *XiYong) Shen() string {
	if !xy.IsEightCharStrong() {
		return xy.minFiveElementsScore(xy.Similar...)
	}
	return xy.minFiveElementsScore(xy.Heterogeneous...)
}

//IsEightCharStrong 八字偏强（true)弱（false）
func (xy *XiYong) IsEightCharStrong() bool {
	return xy.SimilarScore > xy.HeterogeneousScore
}

func filterXiYong(yong string, cs ...*Character) (b bool) {
	for _, c := range cs {
		// _logger.Debugw("filterXiYong", "5E", c.FiveElements, "yong", yong)
		if c.FiveElements == yong {
			return true
		}
	}
	return false
}
