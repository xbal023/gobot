package plugin

import (
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Join(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	if m.Query == "" { ball.Reply(fmt.Sprintf(a.TLink, m.CmdP, "*LinkGroup*"), true); return }
	ball.Joining(m.Query);
}