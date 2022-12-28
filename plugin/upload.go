package plugin

import (
	"encoding/json"
	
	"go.mau.fi/whatsmeow"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Upload(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	if m.Media != nil {
		down := ball.DL(m.Media)
		res := ball.Up(down, whatsmeow.MediaImage)
		jsonRes, _ := json.MarshalIndent(res, "", " ")
		ball.Reply(string(jsonRes), true)
	} else if m.Quoted.Media != nil {
		down := ball.DL(m.Quoted.Media)
		res := ball.Up(down, whatsmeow.MediaImage)
		jsonRes, _ := json.MarshalIndent(res, "", " ")
		ball.Reply(string(jsonRes), true)
	} else { ball.Reply(a.QMedia, true) }
}