package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"

	res "github.com/zrcoder/lastone/resource"
	"github.com/zrcoder/lastone/text"
)

var (
	mainForm *PlayForm
	content  *text.Text
	total    int

	logger = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *PlayForm) OnFormCreate(sender vcl.IObject) {
	logger.Println("start to create views", sender.ToString())
	f.createSubViews()
	f.initSubViews()
}

func (f *PlayForm) onButtonClicked(btn *vcl.TButton, n int) {
	logger.Println("button clicked, must eat", n)
	f.eat(n, true)
	if total == 0 {
		return
	}
	f.autoEat()
}

func (f *PlayForm) autoEat() {
	mod := total % (res.Limited + 1)
	if mod != 0 {
		f.eat(mod, false)
	} else {
		f.eat(rand.Intn(res.Limited)+1, false)
	}
}

func (f *PlayForm) eat(c int, isPlayer bool) {
	f.Btn1.SetEnabled(false)
	f.Btn2.SetEnabled(false)
	f.Btn2.SetEnabled(false)
	if total < c {
		c = total
	}
	total -= c
	content.Remove(c)
	f.Content.SetTextBuf(content.String())
	if isPlayer {
		f.PlayerLabel.SetTextBuf(res.Player + "\n-" + strconv.Itoa(c))
	} else {
		f.RobertLabel.SetTextBuf(res.Robert + "\n-" + strconv.Itoa(c))
	}
	time.Sleep(res.WaiteTime)
	f.Btn1.SetEnabled(true)
	f.Btn2.SetEnabled(true)
	f.Btn2.SetEnabled(true)
	if total == 0 {
		f.ConfigBtn.SetEnabled(true)
		f.Box.SetEnabled(true)
		f.FirstRadio.SetEnabled(true)
		f.SecondRadio.SetEnabled(true)
		f.Btn1.SetEnabled(false)
		f.Btn2.SetEnabled(false)
		f.Btn3.SetEnabled(false)
		if isPlayer {
			f.Content.SetTextBuf(res.WinInfo)
		} else {
			f.Content.SetTextBuf(res.LoseInfo)
		}
	}
}
