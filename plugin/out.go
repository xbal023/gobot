package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Out(ball *y.S, m *x.Parse)  {
	if !m.IsGc {
		ball.Reply("Khusus di group", true)
		return
	}
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner", true)
		return
	}
	ball.LeaveGc(*m.Chat);
}