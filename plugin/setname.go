package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func SetName(ball *y.S, m *x.Parse)  {
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
	if len(m.Query) < 1 {
		ball.Reply("Masukan nama groupnya setelah command", true)
		return
	}
	ball.SetGcName(*m.Chat, m.Query)
	ball.Reply("Sukses mengganti nama group dengan yg baru", true)
}