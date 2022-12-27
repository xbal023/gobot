package plugin 

import (
	"os/exec"
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Getplugin(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	if len(m.Query) < 1 { ball.Reply(fmt.Sprintf(a.TPath, m.CmdP, "*plugin/menu*"), true); return }
	res, err := exec.Command("bash", "-c", `cat `+m.Query+ `.go`).Output()
	if err != nil {
		ball.Reply(fmt.Sprintf("%v", err), true)
		return
	}
	ball.Reply(string(res), true)
}