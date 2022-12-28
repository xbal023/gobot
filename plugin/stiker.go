package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	"github.com/arugaz/go-wasticker"
	a "gobot/constanta"
	)
	
func Stiker(ball *y.S, m *x.Parse)  {
	if m.Quoted.Media != nil {
		val := ball.DL(m.Quoted.Media)
		st, _ := wasticker.NewSticker(val).ToByte()
		ball.SendStik(*m.Chat, st)
	} else if m.Media != nil {
		val := ball.DL(m.Media)
		st, _ := wasticker.NewSticker(val).ToByte()
		ball.SendStik(*m.Chat, st)
	} else {
		ball.Reply(a.QMedia, true)
	}
}
