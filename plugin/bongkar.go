package plugin

import (
	"encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Bongkar(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	jsonRes, _ := json.MarshalIndent(m, "", " ")
	ball.Reply(string(jsonRes), true)
}