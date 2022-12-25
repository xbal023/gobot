package plugin

import (
	// "fmt"
	// "encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Delete(ball *y.S, m *x.Parse)  {
	if m.Quoted.Sender == nil {
		ball.Reply("Reply pesan yg ingin kamu lenyapkan", true)
		return
	}
	ball.DelMsg(*m.Quoted.Sender, m.Quoted.Id)
}