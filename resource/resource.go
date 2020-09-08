package resource

import (
	"strings"
	"time"
)

var (
	s = "🍏🍎🍐🍊🍋🍌🍉🍇🍓🍈🍒🍑🍍🥝🍅🍆🥑🥒🌶🌽🥕🥔🍠🥐🍞🥖🧀🥚🍳🥞🥓🍗🍖🌭🍔🍟🍕🥙🌮🌯🥗🥘🍝" +
		"🍜🍲🍛🍣🍱🍤🍙🍚🍘🍥🍢🍡🍧🍨🍦🍰🎂🍮🍭🍬🍫🍿🍩🍪🌰🥜🍯🥛🍼☕🍶🍺🍻🥂🍷🥃🍸🍹🍾🥄🍴🍽"
	Set = strings.Split(s, "")
)

const (
	GameName = "最后那个"
	HelpInfo = `每次吃掉一份、两份或三分食物
吃掉最后一份食物的人获胜

你将和一个非常聪明的机器人比赛
但幸运的是你可以在游戏开始前设置食物总数
同时决定自己是否先手`
	Robert   = "🤖"
	Player   = "😶"
	WinFace  = "☺"
	LoseFace = "🙁"
	WinInfo  = WinFace + " 赢了!"
	LoseInfo = LoseFace + " 输了~"

	Start = "开始"

	FirstHand  = "先手"
	SecondHand = "后手"

	InitialTotal = "45"

	Padding int32 = 10

	//  每次最少吃 1 个，最多吃 Limited 个
	Limited = 3

	WaiteTime = 700 * time.Millisecond
)
