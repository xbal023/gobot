package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func SetPPGC(ball *y.S, m *x.Parse)  {
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
	if m.Quoted.Media != nil {
		pp := ball.DL(m.Quoted.Media)
		ball.SetGcPP(*m.Chat, pp)
		ball.Reply("Sukses mengganti Photo group dengan yg baru")
	} else if m.Media != nil {
		pp := ball.DL(m.Media)
		ball.SetGcPP(*m.Chat, pp)
		ball.Reply("Sukses mengganti Photo group dengan yg baru")
	} else {
		ball.Reply("Reply / kirim media dengan caption command")
	}
}