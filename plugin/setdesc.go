package plugin

import (
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func SetDesc(ball *y.S, m *x.Parse)  {
	if !m.IsGc { ball.Reply(a.FGroup, true); return }
	if !m.IsBotAdmin { ball.Reply(a.FBotAdmin, true); return }
	if !m.IsAdmin { ball.Reply(a.FAdmin, true); return }
	if len(m.Query) < 1 { ball.Reply(fmt.Sprintf(a.TDesc, m.CmdP, "balbalabalaabalab"), true); return }
	ball.SetDesc(*m.Chat, m.Query)
	ball.Reply(a.GCDesc, true)
}