# U.S.G.

Basic package to output Unicode symbols for golang



> This package is a port of [https://github.com/sindresorhus/figures](htts://github.com/sindresorhus/figures) in golang.


## Install

```
go get -u github.com/julienroland/usg
```


## Usage

```
import "github.com/julienroland/usg"

// Output: ✔ Tests
fmt.Println(usg.Get.Tick, "Tests") 

```


## Symbols available

| Name               | Real OSes | Windows |
| ------------------ | :-------: | :-----: |
| Tick               |     ✔     |    √    |
| Cross              |     ✖     |    ×    |
| CrossGraph         |     ✘     |    ×    |
| CrossThin          |     ✕     |    ×    |
| Star               |     ★     |    *    |
| ExclamationMark    |     ❗    |    !    |
| QuestionMark       |     ❓    |    ?    |
| QuoteStart         |     ❝     |    "    |
| QuoteEnd           |     ❞     |    "    |
| WhiteDiamond       |     ◇     |    []   |
| BlackDiamond       |     ◆     |    █    |
| WrapDiamond        |     ◈     |   [█]   |
| Square             |     ▇     |    █    |
| SquareSmall        |     ◻     |   [ ]   |
| SquareSmallFilled  |     ◼     |   [█]   |
| Play               |     ▶     |    ►    |
| Circle             |     ◯     |   ( )   |
| CircleFilled       |     ◉     |   (*)   |
| CircleDotted       |     ◌     |   ( )   |
| CircleDouble       |     ◎     |   ( )   |
| CircleCircle       |     ⓞ     |   (○)   |
| CircleCross        |     ⓧ     |   (×)   |
| CirclePipe         |     Ⓘ     |   (│)   |
| CircleQuestionMark |     ?⃝    |   (?)   |
| Bullet             |     ●     |    *    |
| Dot                |     ․     |    .    |
| Line               |     ─     |    ─    |
| Ellipsis           |     …     |   ...   |
| Pointer            |     ❯     |    >    |
| PointerSmall       |     ›     |    »    |
| Info               |     ℹ     |    i    |
| Warning            |     ⚠     |    ‼    |
| Hamburger          |     ☰     |    ≡    |
| Smiley             |     ㋡     |    ☺    |
| Mustache           |     ෴     |   ┌─┐   |
| Heart              |     ♥     |    ♥    |
| ArrowUp            |     ↑     |    ↑    |
| ArrowDown          |     ↓     |    ↓    |
| ArrowLeft          |     ←     |    ←    |
| ArrowRight         |     →     |    →    |
| RadioOn            |     ◉     |   (*)   |
| RadioOff           |     ◯     |   ( )   |
| CheckboxOn         |     ☒     |   [×]   |
| CheckboxOff        |     ☐     |   [ ]   |
| CheckboxCircleOn   |     ⓧ     |   (×)   |
| CheckboxCircleOff  |     Ⓘ     |   ( )   |
| QuestionMarkPrefix |     ?⃝    |    ？    |
| OneHalf            |     ½     |   1/2   |
| OneThird           |     ⅓     |   1/3   |
| OneQuarter         |     ¼     |   1/4   |
| OneFifth           |     ⅕     |   1/5   |
| OneSixth           |     ⅙     |   1/6   |
| OneSeventh         |     ⅐     |   1/7   |
| OneEighth          |     ⅛     |   1/8   |
| OneNinth           |     ⅑     |   1/9   |
| OneTenth           |     ⅒     |   1/10  |
| TwoThirds          |     ⅔     |   2/3   |
| TwoFifths          |     ⅖     |   2/5   |
| ThreeQuarters      |     ¾     |   3/4   |
| ThreeFifths        |     ⅗     |   3/5   |
| ThreeEighths       |     ⅜     |   3/8   |
| FourFifths         |     ⅘     |   4/5   |
| FiveSixths         |     ⅚     |   5/6   |
| FiveEighths        |     ⅝     |   5/8   |
| SevenEighths       |     ⅞     |   7/8   |
