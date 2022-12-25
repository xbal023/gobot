package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func GcUnlock(ball *y.S, m *x.Parse)  {
	if !m.IsGc {
		ball.Reply("Khusus di group", true)
		return
	}
	if !m.IsBotAdmin {
		ball.Reply("Bot bukan admin", true)
		return
	}
	if !m.IsAdmin {
		ball.Reply("Kamu bukan admin!", true)
		return
	}
	ball.SetGcLock(*m.Chat, false)
	ball.Reply("sukses membuka edit info group!", true)
}