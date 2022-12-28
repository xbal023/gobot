package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func GcOpen(ball *y.S, m *x.Parse)  {
	if !m.IsGc { ball.Reply(a.FGroup, true); return }
	if !m.IsBotAdmin { ball.Reply(a.FBotAdmin, true); return }
	if !m.IsAdmin { ball.Reply(a.FAdmin, true); return }
	ball.SetGcChat(*m.Chat, false)
	ball.Reply(a.GCOpen, true)
}