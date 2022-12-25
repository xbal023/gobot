package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Out(ball *y.S, m *x.Parse)  {
	if !m.IsGc {
		ball.Reply("Khusus di group")
		return
	}
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner")
		return
	}
	ball.LeaveGc(*m.Chat);
}