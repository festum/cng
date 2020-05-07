package cng

//Strokes 推薦筆畫數
type Strokes struct {
	SimplifiedChinese  int //簡體中文
	TraditionalChinese int //繁體中文
	KangxiDictionary   int //康熙字典
	GodBook            int //神冊
}

//FindCharacterStrokes 通過文字查詢推薦筆畫數
func FindCharacterStrokes(char string) int {
	//TODO:find
	s := &Strokes{}
	if s.GodBook > 0 {
		return s.GodBook
	}
	if s.KangxiDictionary > 0 {
		return s.KangxiDictionary
	}
	if s.TraditionalChinese > 0 {
		return s.TraditionalChinese
	}
	if s.SimplifiedChinese > 0 {
		return s.SimplifiedChinese
	}
	return 0
}
