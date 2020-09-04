package main

import (
	"log"
	"os"
	"strings"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

var logger = log.New(os.Stdout, "Apple", log.Lshortfile|log.Ltime)

type MainForm struct {
	*vcl.TForm
	Btn1 *vcl.TButton
	Btn2 *vcl.TButton
	Text *vcl.TStaticText
}

var (
	mainForm *MainForm

	cur = total
)

const (
	total   = 50
	limited = 2

	commonOne = "🍏"
	blackOne  = "🍎"
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *MainForm) OnFormCreate(sender vcl.IObject) {
	const padding = 30
	bounds := f.BoundsRect()
	w, h := bounds.Width(), bounds.Height()
	f.SetBounds(200, 200, w, h)

	f.Text = vcl.NewStaticText(f)
	f.Text.SetParent(f)
	f.Text.SetBounds(padding, padding, w-20, h/5)
	f.Text.SetTextBuf(genText(total))

	f.SetCaption("THE BLACK ONE")
	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(padding, h/4*3-padding, (w-3*padding)/2, h/4)
	f.Btn1.SetCaption("1")
	f.Btn1.SetOnClick(f.OnBtn1Click)

	f.Btn2 = vcl.NewButton(f)
	f.Btn2.SetParent(f)
	f.Btn2.SetBounds(f.Btn1.BoundsRect().Right+padding, f.Btn1.BoundsRect().Top, f.Btn1.BoundsRect().Width(), f.Btn1.BoundsRect().Height())
	f.Btn2.SetCaption("2")
	f.Btn2.SetOnClick(f.OnBtn2Click)
}

func genText(t int) string {
	buf := strings.Builder{}
	for i := 1; i < t; i++ {
		buf.WriteString(commonOne)
	}
	buf.WriteString(blackOne)
	return buf.String()
}

func (f *MainForm) OnBtn1Click(sender vcl.IObject) {
	logger.Println(sender.ToString())
	f.onBtnClick(1)
}

func (f *MainForm) OnBtn2Click(sender vcl.IObject) {
	logger.Println(sender.ToString())
	f.onBtnClick(2)
}

func (f *MainForm) onBtnClick(c int) {
	cur -= c
	if cur < 0 {
		cur = 0
	}
	f.Text.SetTextBuf(genText(cur))

	if cur == 0 {
		vcl.ShowMessage("Oh~ you eaten the black one!")
		f.Text.SetTextBuf("")
	}
}
