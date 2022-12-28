package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func SetPPGC(ball *y.S, m *x.Parse)  {
	if !m.IsGc { ball.Reply(a.FGroup, true); return }
	if !m.IsBotAdmin { ball.Reply(a.FBotAdmin, true); return }
	if !m.IsAdmin { ball.Reply(a.FAdmin, true); return }
	if m.Quoted.Media != nil {
		pp := ball.DL(m.Quoted.Media)
		ball.SetGcPP(*m.Chat, pp)
		ball.Reply(a.GCPP, true)
	} else if m.Media != nil {
		pp := ball.DL(m.Media)
		ball.SetGcPP(*m.Chat, pp)
		ball.Reply(a.GCPP, true)
	} else {
		ball.Reply(a.QMedia, true)
	}
}