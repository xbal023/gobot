package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func SetDesc(ball *y.S, m *x.Parse)  {
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
		ball.Reply("Masukan isi deskripsi setelah command", true)
		return
	}
	ball.SetDesc(*m.Chat, m.Query)
	ball.Reply("Sukses mengganti Deskripsi dengan yg baru", true)
}