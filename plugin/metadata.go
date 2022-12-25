package plugin

import (
	"fmt"
	"encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Metadata(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner")
		return
	}
	val := ball.GetMetadata(*m.Chat)
	jsonRes, err := json.MarshalIndent(val,""," ")
	if err != nil {
		fmt.Println(err)
	}
	ball.Reply(string(jsonRes))
}