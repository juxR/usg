package usg

import "runtime"

type symbols struct {
	Tick               string
	Cross              string
	CrossGraph         string
	CrossThin          string
	Star               string
	QuoteStart         string
	QuoteEnd           string
	QuestionMark       string
	ExclamationMark    string
	WhiteDiamond       string
	BlackDiamond       string
	WrapDiamond        string
	Square             string
	SquareSmall        string
	SquareSmallFilled  string
	Play               string
	Circle             string
	CircleFilled       string
	CircleDotted       string
	CircleDouble       string
	CircleCircle       string
	CircleCross        string
	CirclePipe         string
	CircleQuestionMark string
	Bullet             string
	Dot                string
	Line               string
	Ellipsis           string
	Pointer            string
	PointerSmall       string
	Info               string
	Warning            string
	Hamburger          string
	Smiley             string
	Mustache           string
	Heart              string
	ArrowUp            string
	ArrowDown          string
	ArrowLeft          string
	ArrowRight         string
	RadioOn            string
	RadioOff           string
	CheckboxOn         string
	CheckboxOff        string
	CheckboxCircleOn   string
	CheckboxCircleOff  string
	QuestionMarkPrefix string
	OneHalf            string
	OneThird           string
	OneQuarter         string
	OneFifth           string
	OneSixth           string
	OneSeventh         string
	OneEighth          string
	OneNinth           string
	OneTenth           string
	TwoThirds          string
	TwoFifths          string
	ThreeQuarters      string
	ThreeFifths        string
	ThreeEighths       string
	FourFifths         string
	FiveSixths         string
	FiveEighths        string
	SevenEighths       string
}

var Get symbols

func init() {
	if runtime.GOOS == "windows" {
		Get = setWindowsSymbols()
		return
	}
	Get = setDefaultSymbols()
}

func setDefaultSymbols() symbols {
	return symbols{
		Tick:               "✔",
		Cross:              "✖",
		CrossGraph:         "✘",
		CrossThin:          "✕",
		Star:               "★",
		ExclamationMark:    "❗",
		QuestionMark:       "❓",
		QuoteStart:         "❝",
		QuoteEnd:           "❞",
		WhiteDiamond:       "◇",
		BlackDiamond:       "◆",
		WrapDiamond:        "◈",
		Square:             "▇",
		SquareSmall:        "◻",
		SquareSmallFilled:  "◼",
		Play:               "▶",
		Circle:             "◯",
		CircleFilled:       "◉",
		CircleDotted:       "◌",
		CircleDouble:       "◎",
		CircleCircle:       "ⓞ",
		CircleCross:        "ⓧ",
		CirclePipe:         "Ⓘ",
		CircleQuestionMark: "?⃝",
		Bullet:             "●",
		Dot:                "․",
		Line:               "─",
		Ellipsis:           "…",
		Pointer:            "❯",
		PointerSmall:       "›",
		Info:               "ℹ",
		Warning:            "⚠",
		Hamburger:          "☰",
		Smiley:             "㋡",
		Mustache:           "෴",
		Heart:              "♥",
		ArrowUp:            "↑",
		ArrowDown:          "↓",
		ArrowLeft:          "←",
		ArrowRight:         "→",
		RadioOn:            "◉",
		RadioOff:           "◯",
		CheckboxOn:         "☒",
		CheckboxOff:        "☐",
		CheckboxCircleOn:   "ⓧ",
		CheckboxCircleOff:  "Ⓘ",
		QuestionMarkPrefix: "?⃝",
		OneHalf:            "½",
		OneThird:           "⅓",
		OneQuarter:         "¼",
		OneFifth:           "⅕",
		OneSixth:           "⅙",
		OneSeventh:         "⅐",
		OneEighth:          "⅛",
		OneNinth:           "⅑",
		OneTenth:           "⅒",
		TwoThirds:          "⅔",
		TwoFifths:          "⅖",
		ThreeQuarters:      "¾",
		ThreeFifths:        "⅗",
		ThreeEighths:       "⅜",
		FourFifths:         "⅘",
		FiveSixths:         "⅚",
		FiveEighths:        "⅝",
		SevenEighths:       "⅞",
	}
}

func setWindowsSymbols() symbols {
	return symbols{
		Tick:               "√",
		Cross:              "×",
		CrossGraph:         "×",
		CrossThin:          "×",
		Star:               "*",
		ExclamationMark:    "!",
		QuestionMark:       "?",
		QuoteStart:         "\"",
		QuoteEnd:           "\"",
		WhiteDiamond:       "[]",
		BlackDiamond:       "█",
		WrapDiamond:        "[█]",
		Square:             "█",
		SquareSmall:        "[ ]",
		SquareSmallFilled:  "[█]",
		Play:               "►",
		Circle:             "( )",
		CircleFilled:       "(*)",
		CircleDotted:       "( )",
		CircleDouble:       "( )",
		CircleCircle:       "(○)",
		CircleCross:        "(×)",
		CirclePipe:         "(│)",
		CircleQuestionMark: "(?)",
		Bullet:             "*",
		Dot:                ".",
		Line:               "─",
		Ellipsis:           "...",
		Pointer:            ">",
		PointerSmall:       "»",
		Info:               "i",
		Warning:            "‼",
		Hamburger:          "≡",
		Smiley:             "☺",
		Mustache:           "┌─┐",
		Heart:              "♥",
		ArrowUp:            "↑",
		ArrowDown:          "↓",
		ArrowLeft:          "←",
		ArrowRight:         "→",
		RadioOn:            "(*)",
		RadioOff:           "( )",
		CheckboxOn:         "[×]",
		CheckboxOff:        "[ ]",
		CheckboxCircleOn:   "(×)",
		CheckboxCircleOff:  "( )",
		QuestionMarkPrefix: "？",
		OneHalf:            "1/2",
		OneThird:           "1/3",
		OneQuarter:         "1/4",
		OneFifth:           "1/5",
		OneSixth:           "1/6",
		OneSeventh:         "1/7",
		OneEighth:          "1/8",
		OneNinth:           "1/9",
		OneTenth:           "1/10",
		TwoThirds:          "2/3",
		TwoFifths:          "2/5",
		ThreeQuarters:      "3/4",
		ThreeFifths:        "3/5",
		ThreeEighths:       "3/8",
		FourFifths:         "4/5",
		FiveSixths:         "5/6",
		FiveEighths:        "5/8",
		SevenEighths:       "7/8",
	}
}
