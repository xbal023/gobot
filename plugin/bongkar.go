package plugin

import (
	"encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Bongkar(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner", true)
		return
	}
	jsonRes, _ := json.MarshalIndent(m, "", " ")
	ball.Reply(string(jsonRes), true)
}