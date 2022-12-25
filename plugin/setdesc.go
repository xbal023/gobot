package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func SetDesc(ball *y.S, m *x.Parse)  {
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
		ball.Reply("Masukan isi deskripsi setelah command")
		return
	}
	ball.SetDesc(*m.Chat, m.Query)
	ball.Reply("Sukses mengganti Deskripsi dengan yg baru")
}