package plugin 

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Revoke(ball *y.S, m *x.Parse)  {
	if !m.IsGc { ball.Reply(a.FGroup, true); return }
	if !m.IsBotAdmin { ball.Reply(a.FBotAdmin, true); return }
	if !m.IsAdmin { ball.Reply(a.FAdmin, true); return }
	val := ball.SetLink(ball.M.Info.Chat, true)
	ball.Reply(val, true)
}