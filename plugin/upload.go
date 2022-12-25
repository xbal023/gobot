package plugin

import (
	"encoding/json"
	
	"go.mau.fi/whatsmeow"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Upload(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner", true)
		return
	}
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
	} else {
		ball.Reply("Reply / kirim media", true)
	}
}