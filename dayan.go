package cng

import (
	_ "github.com/festum/cng/statik"
)

var daYanList [81]DaYan

func init() {
	daYanList = [81]DaYan{
		{Number: 1, Lucky: "吉", SkyNine: "太極之數", Comment: "太極之數，萬物開泰，生髮無窮，利祿亨通。"},
		{Number: 2, Lucky: "凶", SkyNine: "兩儀之數", Comment: "兩儀之數，混沌未開，進退保守，志望難達。"},
		{Number: 3, Lucky: "吉", SkyNine: "三才之數", Comment: "三才之數，天地人和，大事大業，繁榮昌隆。"},
		{Number: 4, Lucky: "凶", SkyNine: "四象之數", Comment: "四象之數，待於生髮，萬事慎重，不具營謀。"},
		{Number: 5, Lucky: "吉", SkyNine: "五行之數", Comment: "五行俱權，迴圈相生，圓通暢達，福祉無窮。"},
		{Number: 6, Lucky: "吉", SkyNine: "六爻之數", Comment: "六爻之數，發展變化，天賦美德，吉祥安泰。"},
		{Number: 7, Lucky: "吉", SkyNine: "七政之數", Comment: "七政之數，精悍嚴謹，天賦之力，吉星照耀。"},
		{Number: 8, Lucky: "半吉", SkyNine: "八卦之數", Comment: "八卦之數，乾坎艮震，巽離坤兌，無窮無盡。"},
		{Number: 9, Lucky: "凶", SkyNine: "大成之數", Comment: "大成之數，蘊涵凶險，或成或敗，難以把握。"},
		{Number: 10, Lucky: "凶", SkyNine: "終結之數", Comment: "終結之數，雪暗飄零，偶或有成，回顧茫然。"},
		{Number: 11, Lucky: "吉", SkyNine: "旱苗逢雨", Comment: "萬物更新，調順發達，恢弘澤世，繁榮富貴。"},
		{Number: 12, Lucky: "凶", SkyNine: "掘井無泉", Comment: "無理之數，發展薄弱，雖生不足，難酬志向。"},
		{Number: 13, Lucky: "吉", SkyNine: "春日牡丹", Comment: "才藝多能，智謀奇略，忍柔當事，鳴奏大功。"},
		{Number: 14, Lucky: "凶", SkyNine: "破兆", Comment: "家庭緣薄，孤獨遭難，謀事不達，悲慘不測。"},
		{Number: 15, Lucky: "吉", SkyNine: "福壽", Comment: "福壽圓滿，富貴榮譽，涵養雅量，德高望重。"},
		{Number: 16, Lucky: "吉", SkyNine: "厚重", Comment: "厚過載德，安富尊榮，財官雙美，功成名就。"},
		{Number: 17, Lucky: "半吉", SkyNine: "剛強", Comment: "權威剛強，突破萬難，如能容忍，必獲成功。"},
		{Number: 18, Lucky: "半吉", SkyNine: "鐵鏡重磨", Comment: "權威顯達，博得名利，且養柔德，功成名就。"},
		{Number: 19, Lucky: "凶", SkyNine: "多難", Comment: "風雲蔽日，辛苦重來，雖有智謀，萬事挫折。"},
		{Number: 20, Lucky: "凶", SkyNine: "屋下藏金", Comment: "非業破運，災難重重，進退維谷，萬事難成。"},
		{Number: 21, Lucky: "吉", Sex: true, SkyNine: "明月中天", Comment: "光風霽月，萬物確立，官運亨通，大搏名利。女性不宜此數。"},
		{Number: 22, Lucky: "凶", SkyNine: "秋草逢霜", Comment: "秋草逢霜，困難疾弱，雖出豪傑，人生波折。"},
		{Number: 23, Lucky: "吉", Sex: true, SkyNine: "壯麗", Comment: "旭日東昇，壯麗壯觀，權威旺盛，功名榮達。女性不宜此數。"},
		{Number: 24, Lucky: "吉", SkyNine: "掘藏得金", Comment: "家門餘慶，金錢豐盈，白手成家，財源廣進。"},
		{Number: 25, Lucky: "半吉", SkyNine: "榮俊", Comment: "資性英敏，才能奇特，克服傲慢，尚可成功。"},
		{Number: 26, Lucky: "凶", SkyNine: "變怪", Comment: "變怪之謎，英雄豪傑，波瀾重疊，而奏大功。"},
		{Number: 27, Lucky: "凶", SkyNine: "增長", Comment: "慾望無止，自我強烈，多受毀謗，尚可成功。"},
		{Number: 28, Lucky: "凶", Sex: true, SkyNine: "闊水浮萍", Comment: "遭難之數，豪傑氣概，四海漂泊，終世浮躁。女性不宜此數。"},
		{Number: 29, Lucky: "吉", SkyNine: "智謀", Comment: "智謀優秀，財力歸集，名聞海內，成就大業。"},
		{Number: 30, Lucky: "半吉", SkyNine: "非運", Comment: "沉浮不定，凶吉難變，若明若暗，大成大敗。"},
		{Number: 31, Lucky: "吉", SkyNine: "春日花開", Comment: "智勇得志，博得名利，統領眾人，繁榮富貴。"},
		{Number: 32, Lucky: "吉", SkyNine: "寶馬金鞍", Comment: "僥倖多望，貴人得助，財帛如裕，繁榮至上。"},
		{Number: 33, Lucky: "吉", Sex: true, SkyNine: "旭日升天", Comment: "旭日升天，鸞鳳相會，名聞天下，隆昌至極。女性不宜此數。"},
		{Number: 34, Lucky: "凶", SkyNine: "破家", Comment: "破家之身，見識短小，辛苦遭逢，災禍至極。"},
		{Number: 35, Lucky: "吉", SkyNine: "高樓望月", Comment: "溫和平靜，智達通暢，文昌技藝，奏功洋洋。"},
		{Number: 36, Lucky: "半吉", SkyNine: "波瀾重疊", Comment: "波瀾重疊，沉浮萬狀，俠肝義膽，捨己成仁。"},
		{Number: 37, Lucky: "吉", SkyNine: "猛虎出林", Comment: "權威顯達，熱誠忠信，宜著雅量，終身榮富。"},
		{Number: 38, Lucky: "半吉", SkyNine: "磨鐵成針", Comment: "意志薄弱，刻意經營，才識不凡，技藝有成。"},
		{Number: 39, Lucky: "半吉", SkyNine: "富貴榮華", Comment: "富貴榮華，財帛豐盈，暗藏險象，德澤四方。"},
		{Number: 40, Lucky: "凶", SkyNine: "退安", Comment: "智謀膽力，冒險投機，沉浮不定，退保平安。"},
		{Number: 41, Lucky: "吉", Max: true, SkyNine: "有德", Comment: "純陽獨秀，德高望重，和順暢達，博得名利。此數為最大好運數。"},
		{Number: 42, Lucky: "凶", SkyNine: "寒蟬在柳", Comment: "博識多能，精通世情，如能專心，尚可成功。"},
		{Number: 43, Lucky: "凶", SkyNine: "散財破產", Comment: "散財破產，諸事不遂，雖有智謀，財來財去。"},
		{Number: 44, Lucky: "凶", SkyNine: "煩悶", Comment: "破家亡身，暗藏慘淡，事不如意，亂世怪傑。"},
		{Number: 45, Lucky: "吉", SkyNine: "順風", Comment: "新生泰和，順風揚帆，智謀經緯，富貴繁榮。"},
		{Number: 46, Lucky: "凶", SkyNine: "浪裡淘金", Comment: "載寶沉舟，浪裡淘金，大難嚐盡，大功有成。"},
		{Number: 47, Lucky: "吉", SkyNine: "點石成金", Comment: "花開之象，萬事如意，禎祥吉慶，天賦幸福。"},
		{Number: 48, Lucky: "吉", SkyNine: "古鬆立鶴", Comment: "智謀兼備，德量榮達，威望成師，洋洋大觀。"},
		{Number: 49, Lucky: "半吉", SkyNine: "轉變", Comment: "吉臨則吉，凶來則凶，轉凶為吉，配好三才。"},
		{Number: 50, Lucky: "半吉", SkyNine: "小舟入海", Comment: "一成一敗，吉凶參半，先得庇廕，後遭悽慘。"},
		{Number: 51, Lucky: "半吉", SkyNine: "沉浮", Comment: "盛衰交加，波瀾重疊，如能慎始，必獲成功。"},
		{Number: 52, Lucky: "吉", SkyNine: "達眼", Comment: "卓識達眼，先見之明，智謀超群，名利雙收。"},
		{Number: 53, Lucky: "凶", SkyNine: "曲捲難星", Comment: "外祥內患，外禍內安，先富後貧，先貧後富。"},
		{Number: 54, Lucky: "凶", SkyNine: "石上栽花", Comment: "石上栽花，難得有活，憂悶煩來，辛慘不絕。"},
		{Number: 55, Lucky: "半吉", SkyNine: "善惡", Comment: "善善得惡，惡惡得善，吉到極限，反生凶險。"},
		{Number: 56, Lucky: "凶", SkyNine: "浪裡行舟", Comment: "歷盡艱辛，四周障礙，萬事齟齷，做事難成。"},
		{Number: 57, Lucky: "吉", SkyNine: "日照春鬆", Comment: "寒雪青松，夜鶯吟春，必遭一過，繁榮白事。"},
		{Number: 58, Lucky: "半吉", SkyNine: "晚行遇月", Comment: "沉浮多端，先苦後甜，寬宏揚名，富貴繁榮。"},
		{Number: 59, Lucky: "凶", SkyNine: "寒蟬悲風", Comment: "寒蟬悲風，意志衰退，缺乏忍耐，苦難不休。"},
		{Number: 60, Lucky: "凶", SkyNine: "無謀", Comment: "無謀之人，漂泊不定，晦暝暗黑，動搖不安。"},
		{Number: 61, Lucky: "吉", SkyNine: "牡丹芙蓉", Comment: "牡丹芙蓉，花開富貴，名利雙收，定享天賦。"},
		{Number: 62, Lucky: "凶", SkyNine: "衰敗", Comment: "衰敗之象，內外不和，志望難達，災禍頻來。"},
		{Number: 63, Lucky: "吉", SkyNine: "舟歸平海", Comment: "富貴榮華，身心安泰，雨露惠澤，萬事亨通。"},
		{Number: 64, Lucky: "凶", SkyNine: "非命", Comment: "骨肉分離，孤獨悲愁，難得心安，做事不成。"},
		{Number: 65, Lucky: "吉", SkyNine: "巨流歸海", Comment: "天長地久，家運隆昌，福壽綿長，事事成就。"},
		{Number: 66, Lucky: "凶", SkyNine: "巖頭步馬", Comment: "進退維谷，艱難不堪，等待時機，一躍而起。"},
		{Number: 67, Lucky: "吉", SkyNine: "順風通達", Comment: "天賦幸運，四通八達，家道繁昌，富貴東來。"},
		{Number: 68, Lucky: "吉", SkyNine: "順風吹帆", Comment: "智慮周密，集眾信達，發明能智，拓展昂進。"},
		{Number: 69, Lucky: "凶", SkyNine: "非業", Comment: "非業非力，精神迫滯，災害交至，遍償痛苦。"},
		{Number: 70, Lucky: "凶", SkyNine: "殘菊逢霜", Comment: "殘菊逢霜，寂寞無礙，慘淡憂愁，晚景淒涼。"},
		{Number: 71, Lucky: "半吉", SkyNine: "石上金花", Comment: "石上金花，內心勞苦，貫徹始終，定可昌隆。"},
		{Number: 72, Lucky: "半吉", SkyNine: "勞苦", Comment: "榮苦相伴，陰雲覆月，外表吉祥，內實凶禍。"},
		{Number: 73, Lucky: "半吉", SkyNine: "無勇", Comment: "盛衰交加，徒有高志，天王福祉，終世平安。"},
		{Number: 74, Lucky: "凶", SkyNine: "殘菊經霜", Comment: "殘菊經霜，秋葉寂寞，無能無智，辛苦繁多。"},
		{Number: 75, Lucky: "凶", SkyNine: "退守", Comment: "退守保吉，發跡甚遲，雖有吉象，無謀難成。"},
		{Number: 76, Lucky: "凶", SkyNine: "離散", Comment: "傾覆離散，骨肉分離，內外不和，雖勞無功。"},
		{Number: 77, Lucky: "半吉", SkyNine: "半吉", Comment: "家庭有悅，半吉半凶，能獲援護，陷落不幸。"},
		{Number: 78, Lucky: "凶", SkyNine: "晚苦", Comment: "禍福參半，先天智慧，中年發達，晚景困苦。"},
		{Number: 79, Lucky: "凶", SkyNine: "雲頭望月", Comment: "雲頭望月，身疲力盡，窮迫不伸，精神不定。"},
		{Number: 80, Lucky: "凶", SkyNine: "遁吉", Comment: "辛苦不絕，早入隱遁，安心立命，化凶轉吉。"},
		{Number: 81, Lucky: "吉", SkyNine: "萬物回春", Comment: "最吉之數，還本歸元，吉祥重疊，富貴尊榮。"},
	}
}

// DaYan ...
type DaYan struct {
	Number  int
	Lucky   string
	Max     bool
	Sex     bool //male(false),female(true)
	SkyNine string
	Comment string
}

func (dy DaYan) isFemale() bool {
	return dy.Sex
}

func (dy DaYan) IsBestPoint() bool {
	return dy.Max
}

//GetDaYan 獲取大衍之數
func GetDaYan(idx int) DaYan {
	if idx <= 0 {
		panic("wrong idx")
	}
	i := (idx - 1) % 81
	return daYanList[i]
}
