package text

import (
	"math/rand"
	"strings"
	"time"
)

type Text struct {
	info         [][]string
	start, width int
}

func NewWithSet(size, width int, set []string) *Text {
	rand.Seed(time.Now().UnixNano())
	height := (size-1)/width + 1
	mod := size % width
	if mod == 0 {
		mod = width
	}
	info := make([][]string, height)
	info[0] = genRandomLine(mod, set)
	for i := 1; i < height; i++ {
		info[i] = genRandomLine(width, set)
	}
	return &Text{info: info, width: width}
}
func genRandomLine(n int, set []string) []string {
	res := make([]string, n)
	for i := range res {
		res[i] = set[rand.Intn(len(set))]
	}
	return res
}

func (t *Text) Remove(n int) {
	for i := t.start; i < len(t.info); i++ {
		width := len(t.info[i])
		if n < len(t.info[i]) {
			t.info[i] = t.info[i][n:]
			break
		}
		t.info[i] = nil
		t.start++
		n -= width
	}
}

func (t *Text) String() string {
	wide := len(t.info[len(t.info)-1])
	buf := &strings.Builder{}
	for _, v := range t.info {
		if v != nil {
			if wide > len(v) {
				buf.WriteString(strings.Repeat(" ", (wide-len(v))*2-1))
			}
			buf.WriteString(strings.Join(v, " "))
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
