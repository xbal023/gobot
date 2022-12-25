package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func SetName(ball *y.S, m *x.Parse)  {
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
	if len(m.Query) < 1 {
		ball.Reply("Masukan nama groupnya setelah command")
		return
	}
	ball.SetGcName(*m.Chat, m.Query)
	ball.Reply("Sukses mengganti nama group dengan yg baru")
}