package plugin 

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Revoke(ball *y.S, m *x.Parse)  {
	if !m.IsBotAdmin {
		ball.Reply("Bot bukan admin")
		return
	}
	if !m.IsAdmin {
		ball.Reply("Kamu bukan admin!")
		return
	}
	val := ball.SetLink(ball.M.Info.Chat, true)
	ball.Reply(val)
}