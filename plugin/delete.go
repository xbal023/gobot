package plugin

import (
	// "fmt"
	// "encoding/json"
	
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Delete(ball *y.S, m *x.Parse)  {
	if m.Quoted.Sender == nil {
		ball.Reply(a.QDel, true)
		return
	}
	ball.DelMsg(*m.Quoted.Sender, m.Quoted.Id)
}