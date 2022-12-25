package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func GcOpen(ball *y.S, m *x.Parse)  {
	if !m.IsGc {
		ball.Reply("Khusus di group")
		return
	}
	if !m.IsBotAdmin {
		ball.Reply("Bot bukan admin")
		return
	}
	if !m.IsAdmin {
		ball.Reply("Kamu bukan admin!")
		return
	}
	ball.SetGcChat(*m.Chat, false)
	ball.Reply("sukses membuka chat group!")
}