package main

import (
	"github.com/zrcoder/lastone/common"
	"log"
	"os"
	"time"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	"github.com/zrcoder/lastone/resource"
	"github.com/zrcoder/lastone/text"
)

var logger = log.New(os.Stdout, "Apple", log.Lshortfile|log.Ltime)

type PlayForm struct {
	*vcl.TForm
	Content *vcl.TStaticText
	Robert  *vcl.TStaticText
	Player  *vcl.TStaticText
	Btn1    *vcl.TButton
	Btn2    *vcl.TButton
	Btn3    *vcl.TButton

	total int
}

var (
	mainForm *PlayForm

	content *text.Text
)

const (
	//  每次最少吃 1 个，最多吃 limited 个
	limited = 3

	robert   = "🤖"
	player   = "😶"
	winFace  = "😃"
	loseFace = "🙁"
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *PlayForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("THE LAST ONE")
	f.ScreenCenter()

	bound := f.BoundsRect()
	padding := common.Padding
	w, h := bound.Width(), bound.Height()

	playerWidth := int32(20)

	f.Content = vcl.NewStaticText(f)
	f.Content.SetParent(f)
	f.Content.SetBounds(padding+playerWidth, padding, w-(padding+playerWidth)*2, h/5*4-3*padding)

	f.Robert = vcl.NewStaticText(f)
	f.Robert.SetParent(f)
	f.Robert.SetBounds(padding, f.Content.Top(), playerWidth, f.Content.Height())
	f.Robert.SetAlignment(types.TaCenter)
	f.Robert.SetTextBuf(robert)

	f.Player = vcl.NewStaticText(f)
	f.Player.SetParent(f)
	f.Player.SetBounds(w-(padding+playerWidth), f.Content.Top(), playerWidth, f.Content.Height())
	f.Player.SetAlignment(types.TaCenter)
	f.Player.SetTextBuf(player)

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(padding, h/5*4-padding, (w-4*padding)/3, h/5)
	f.Btn1.SetCaption("1")
	f.Btn1.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn1, 1)
	})
	bound1 := f.Btn1.BoundsRect()
	f.Btn2 = vcl.NewButton(f)
	f.Btn2.SetParent(f)
	f.Btn2.SetBounds(bound1.Right+padding, bound1.Top, bound1.Width(), bound1.Height())
	f.Btn2.SetCaption("2")
	f.Btn2.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn2, 2)
	})
	f.Btn3 = vcl.NewButton(f)
	f.Btn3.SetParent(f)
	f.Btn3.SetBounds(f.Btn2.BoundsRect().Right+padding, bound1.Top, bound1.Width(), bound1.Height())
	f.Btn3.SetCaption("3")
	f.Btn3.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn3, 3)
	})

	f.total = 40
	f.reset(true)
}

func (f *PlayForm) reset(isPlayerFirst bool) {
	content = text.NewWithSet(f.total, 6, resource.Set)
	f.Content.SetTextBuf(content.String())
	f.Content.SetAlignment(types.TaCenter)
	if !isPlayerFirst {
		f.autoEat()
	}
}

func (f *PlayForm) onButtonClicked(btn *vcl.TButton, n int) {
	logger.Println(btn.ToString())
	f.eat(n, true)
	if f.total == 0 {
		return
	}
	f.autoEat()
}

func (f *PlayForm) autoEat() {
	f.Btn1.SetEnabled(false)
	f.Btn2.SetEnabled(false)
	f.Btn2.SetEnabled(false)
	time.Sleep(time.Second)
	f.eat(f.total%(limited+1)+1, false)
	f.Btn1.SetEnabled(true)
	f.Btn2.SetEnabled(true)
	f.Btn2.SetEnabled(true)
}

func (f *PlayForm) eat(c int, isPlayer bool) {
	if f.total < c {
		c = f.total
	}
	f.total -= c
	content.Remove(c)
	f.Content.SetTextBuf(content.String())

	if f.total == 0 {
		if isPlayer {
			vcl.ShowMessage(player + " win! " + winFace)
		} else {
			vcl.ShowMessage(robert + " win! " + loseFace)
		}
	}
}
