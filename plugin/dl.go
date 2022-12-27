package plugin

import (
	"fmt"
	// "encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Down(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	if m.Quoted.Media != nil {
		val := ball.DL(m.Quoted.Media)
		// jsonRes, err := json.MarshalIndent(val,""," ")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// ball.Reply(string(jsonRes))
		fmt.Println(val)
	} else if m.Media != nil {
		val := ball.DL(m.Media)
		// jsonRes, err := json.MarshalIndent(val,""," ")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// ball.Reply(string(jsonRes))
		fmt.Println(val)
	} else {
		ball.Reply(a.QMedia, true)
	}
}