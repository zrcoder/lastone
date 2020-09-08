package main

import (
	"strconv"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	res "github.com/zrcoder/lastone/resource"
	"github.com/zrcoder/lastone/text"
)

type PlayForm struct {
	*vcl.TForm
	Content     *vcl.TStaticText
	RobertLabel *vcl.TLabel
	PlayerLabel *vcl.TLabel
	Btn1        *vcl.TButton
	Btn2        *vcl.TButton
	Btn3        *vcl.TButton
	Box         *vcl.TComboBox
	FirstRadio  *vcl.TRadioButton
	SecondRadio *vcl.TRadioButton
	ConfigBtn   *vcl.TButton
}

func (f *PlayForm) createSubViews() {
	f.SetCaption(res.GameName)
	f.ScreenCenter()

	bound := f.BoundsRect()
	padding := res.Padding
	w, h := bound.Width(), bound.Height()
	configWidth := w / 5
	btnWidth := (w - padding*10 - configWidth) / 3
	playersWidth := int32(40)

	left, top := padding, padding
	f.RobertLabel = vcl.NewLabel(f)
	f.RobertLabel.SetParent(f)
	f.RobertLabel.SetBounds(left, top, playersWidth, h/5*4-5*padding)
	f.RobertLabel.Font().SetSize(15)

	f.Content = vcl.NewStaticText(f)
	f.Content.SetParent(f)
	left, top = f.RobertLabel.Left()+f.RobertLabel.Width()+padding, padding*3
	f.Content.SetBounds(left, top, w-padding*10-configWidth-playersWidth*2, f.RobertLabel.Height()-2*padding)
	f.Content.Font().SetSize(10)

	f.PlayerLabel = vcl.NewLabel(f)
	f.PlayerLabel.SetParent(f)
	left, top = f.Content.Left()+f.Content.Width()+padding, f.RobertLabel.Top()
	f.PlayerLabel.SetBounds(left, top, playersWidth, f.RobertLabel.Height())
	f.PlayerLabel.Font().SetSize(15)

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	left, top = padding, h/7*6-padding
	f.Btn1.SetBounds(left, top, btnWidth, h/7)
	f.Btn1.SetCaption("1")
	bound1 := f.Btn1.BoundsRect()

	f.Btn2 = vcl.NewButton(f)
	f.Btn2.SetParent(f)
	left, top = bound1.Right+padding, bound1.Top
	f.Btn2.SetBounds(left, top, bound1.Width(), bound1.Height())
	f.Btn2.SetCaption("2")

	f.Btn3 = vcl.NewButton(f)
	f.Btn3.SetParent(f)
	left, top = f.Btn2.BoundsRect().Right+padding, bound1.Top
	f.Btn3.SetBounds(left, top, bound1.Width(), bound1.Height())
	f.Btn3.SetCaption("3")

	f.ConfigBtn = vcl.NewButton(f)
	f.ConfigBtn.SetParent(f)
	left, top = w-padding-configWidth, f.Btn3.Top()
	f.ConfigBtn.SetBounds(left, top, configWidth, f.Btn3.Height())
	f.ConfigBtn.SetCaption(res.Start)

	radioSize := int32(20)
	f.FirstRadio = vcl.NewRadioButton(f)
	f.FirstRadio.SetParent(f)
	left, top = f.ConfigBtn.Left()+3*padding, f.ConfigBtn.Top()-3*padding-radioSize
	f.FirstRadio.SetBounds(left, top, radioSize, radioSize)
	f.SecondRadio = vcl.NewRadioButton(f)
	f.SecondRadio.SetParent(f)
	left, top = f.ConfigBtn.Left()+btnWidth-3*padding-radioSize, f.FirstRadio.Top()
	f.SecondRadio.SetBounds(left, top, radioSize, radioSize)
	f.FirstRadio.SetCaption(res.FirstHand)
	f.SecondRadio.SetCaption(res.SecondHand)

	f.Box = vcl.NewComboBox(f)
	f.Box.SetParent(f)
	left, top = f.ConfigBtn.Left(), padding
	f.Box.SetBounds(left, top, f.ConfigBtn.Width(), h-f.ConfigBtn.Height()-f.FirstRadio.Height()-4*padding)
}

func (f *PlayForm) initSubViews() {
	f.Content.SetTextBuf(res.HelpInfo)
	f.Content.SetAlignment(types.TaCenter)
	f.Btn1.SetEnabled(false)
	f.Btn2.SetEnabled(false)
	f.Btn3.SetEnabled(false)
	f.FirstRadio.SetChecked(true)
	f.Box.SetText(res.InitialTotal)
	for i := 25; i <= 45; i += 5 {
		f.Box.Items().Add(strconv.Itoa(i))
	}
	f.ConfigBtn.SetOnClick(func(sender vcl.IObject) {
		logger.Println("config button clicked")
		val := f.Box.Text()
		total, _ = strconv.Atoi(val)
		f.reset(f.FirstRadio.Checked())
	})
	f.Btn1.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn1, 1)
	})
	f.Btn2.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn2, 2)
	})
	f.Btn3.SetOnClick(func(sender vcl.IObject) {
		f.onButtonClicked(f.Btn3, 3)
	})
}

func (f *PlayForm) reset(isPlayerFirst bool) {
	f.ConfigBtn.SetEnabled(false)
	f.Box.SetEnabled(false)
	f.FirstRadio.SetEnabled(false)
	f.SecondRadio.SetEnabled(false)
	content = text.NewWithSet(total, 5, res.Set)
	f.Content.SetTextBuf(content.String())
	f.RobertLabel.SetTextBuf(res.Robert)
	f.PlayerLabel.SetTextBuf(res.Player)
	f.Btn1.SetEnabled(true)
	f.Btn2.SetEnabled(true)
	f.Btn3.SetEnabled(true)
	if !isPlayerFirst {
		f.autoEat()
	}
}
