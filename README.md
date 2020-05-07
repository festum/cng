# CNG - Chinese Name Generator

A modern science chinese name generation tool

![Go](https://github.com/festum/cng/workflows/Go/badge.svg)
[![GoDoc](https://godoc.org/github.com/festum/cng?status.svg)](http://godoc.org/github.com/festum/cng)
[![license](https://img.shields.io/github/license/festum/cng.svg)](https://github.com/festum/cng/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/festum/cng)](https://goreportcard.com/report/github.com/festum/cng)

## About

Chinese names are actually quite philosophical. A good name can bring you fortune and a amazing life. Some people pay thousands of euros for a single name but you never know how to verify it.

## Features

- Support double lastname or double firstname（Ex: 歐陽XX，李張XX，張XX，王X）
- 周易卦象
- 大衍之數
- 三才五格
- 喜用神（平衡用神）
- 生肖用字
- 八字吉凶

## Example

```go
//Please import data into the database before use (the test font is almost complete,
// you can input some name to verify)
//load configuration (specific parameters refer to example/create_a_name)
cfg := config.Default()
born := chronos.New("2020/01/23 11:31")
lastName := "張"
f := fate.NewFate(lastName, born.Solar().Time(), fate.ConfigOption(cfg))
if err := f.MakeName(context.Background()); err != nil {
  fmt.Errorf(err)
}
```

## Run

```bash
//If you don't have Go installed, please download the zoneinfo file
// (https://github.com/festum/cng/blob/master/zoneinfo.zip) under master and put it together with binary.
//Generate a configuration file (with the ability to modify the database, and some basic parameters).
fate init
//Generate a name
fate name -l 張 -b "2020/02/06 15:04"
```

### Memo

起名演算法:

演算法(進度) 備註

- 五格(95%) 暫用字型檔已補全
- 三才(95%) 暫用字型檔已補全
- 八字(100%) 喜用神完成, 補八字完成
- 卦象(100%) 演算法完成，起卦完成,吉凶過濾完成。
- 生肖(30%) 生肖獲取完成,字型檔篩選待完成
- 天運(TODO)

立春：

2019/02/04 11:14:14
2020/02/04 17:03:12
2021/02/03 22:58:39
2022/02/04 04:50:36
2023/02/04 10:42:21
2024/02/04 16:26:53
2025/02/03 22:10:13
生肖按八字算從立春開始算起，不到時間則為上一年。

1. 配合八字命理的喜忌，是起名字的核心所在。
八字是每個人出生的年、月、日、時，小孩取名的第一步即是分析八字，瞭解命理五行所缺並找出喜用神，並且據此起名，這是最關鍵的核心，所有姓名的吉凶預測與取名，都以此為準。
1. 名字用字字義務必吉祥
中國文字的魅力在於，每個方塊字不僅都有其本身的含義，而且還有其特殊的周易誘導含義，名字在很多時候它還會影響到人性格的形成，正所謂“名如其人”。
所以一個好的名字，務必用字字義吉祥。
寶寶起名字確實需要考慮很多因素，不僅要考慮讀音、字形以及各種禁忌更重要的是要考慮寶寶的生辰八字，因此給孩子起名還是找專家比較好。
現在這方面比較知名的應該是溫雅居士了，溫雅居士在寶寶起名方面有著十幾年的經驗獨創易學起名法：排八字、看五行、測五格、配三才、合屬相、想寓意、聽音律、寫字形，在業界目前應該是比較權威的了。
取名是需要非常系統考慮的，不能只考慮讀音一個方面，溫雅居士採用的排八字、看五行、測五格、配三才、想寓意、聽音律、寫字形的綜合起名法非常喜歡。
1. 五格數理，特別是主格的數理要為吉數。
在姓名學中，數理產生許多福禍吉凶的靈動力，對人生影響很大。
這跟單個姓名用字的筆劃好壞無關，準確的福禍吉凶是按照特殊方式計算數理的。
1. 三才配置一定不可以相剋。
三才配置在姓名學中，佔有很大的分量。三才配置指的是天格、人格、地格之間的關係。中國傳統文化中有順應天時、地利、人和的行事哲理；
測名過程中，有很多姓名數理不錯，但是三才配置不佳，大多表現為運氣反覆，遇事受阻礙，且感情及財運不好。
三才配置相生相剋的關係定吉凶，同樣也影響著一個人事業成功率的高低。
1. 五格配置在姓名學中佔有主要位置。
五格配置是指天格、地格、人格、外格、總格共五格之間的關係。
天格是由祖先流傳而來，單獨出現對人生沒有多大影響；人格是姓名剖象數理的中心所在，對人生的影響最大；
人格與地格結合的數理則為基礎運。地格主要是36歲前的人生，也叫前運力，外格代表人的外圍，吉凶無謂。總格是36歲以後的人生，也是後運力。
1. 小孩取名字時還要結合命主的出生方位、父母資料等因素，以達到事半功倍的效果。
其實任何事情道理都是一樣的，只有適合自己的才是最好的。

Reference of Zhou Yi Trigrams Code:

[Hexagram](https://en.wikipedia.org/wiki/Hexagram_(I_Ching))(六十四卦)
> ䷀ ䷁ ䷂ ䷃ ䷄ ䷅ ䷆ ䷇ ䷈ ䷉ ䷊ ䷋ ䷌ ䷍ ䷎ ䷏
> ䷐ ䷑ ䷒ ䷓ ䷔ ䷕ ䷖ ䷗ ䷘ ䷙ ䷚ ䷛ ䷜ ䷝ ䷞ ䷟
> ䷠ ䷡ ䷢ ䷣ ䷤ ䷥ ䷦ ䷧ ䷨ ䷩ ䷪ ䷫ ䷬ ䷭ ䷮ ䷯
> ䷰ ䷱ ䷲ ䷳ ䷴ ䷵ ䷶ ䷷ ䷸ ䷹ ䷺ ䷻ ䷼ ䷽ ䷾ ䷿

[Bagua](https://en.wikipedia.org/wiki/Bagua)(八卦)
> ☰ ☱ ☲ ☳ ☴ ☵ ☶ ☷

[Four Symbols](https://en.wikipedia.org/wiki/Four_Symbols)(四象)
> ⚌ ⚍ ⚎ ⚏

[Two Forms](https://en.wikipedia.org/wiki/Yin_and_yang)(兩儀)
> ⚋⚊

為什麼要集六大派與一體?
看下下面這個統計,每一派的取名法其實都有其不足之處.
• 筆劃派: 認為筆劃全吉，人生就大吉。其實準確度僅12.5 %
• 三才派: 完全不管筆劃吉凶，只認為天地人三才五行相生，人生就大吉。其實準確度僅56.6 %。
• 補八字: 完全不管筆劃吉凶，只認為名字補到先天八字命盤欠缺，人生就大吉。其實準確度非常低。
• 卦象派: 完全不管筆劃吉凶，只認為名字求出卦象漂亮，人生就大吉。其實準確度僅40.26 %。
• 天運派: 完全不管筆劃吉凶，只認為名字不要被出生年天運五行所剋，人生就大吉。其實準確度僅25.32 %。
• 生肖派: 完全不管筆劃吉凶，只認為生肖用對字形，人生就大吉。其實準確度僅27.55 %。

PS:最近看到有人別出心裁說三才不準,並舉了一些名人的例子.
然後他倒過來算,發現很符合,很正確.
那我也就呵呵了,按準確度來算,非正即反.
你倒過來算,不準的變準了.那原來準的那些不就不準了.
在我看來事分陰陽,而這接近一半的準確度則恰到其好處.

所以,遵照傳統為自己的寶寶起一箇中正平和的名字才是最好的.
從概率論的角度來講,相交得到的最終結果.其準確度最高.
所以,單純得拿一種或兩種方法來取名是不可取的.
儘量符合多種的名字才是最佳,但並不一定需要全中.
Fate的本意是讓起名變得簡單,且能取到一個好的名字.
有人會花個十幾,幾十萬取一個名字(周圍的真人真事),
但是這個名字好不好你卻未必知道.
演算法開源就是為了讓每個人知道,
這個名字取名過程的來龍去脈.
