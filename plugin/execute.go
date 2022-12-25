package plugin 

import (
	"os/exec"
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Execute(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("Fitur ini khusus owner")
		return
	}
	res, err := exec.Command("bash", "-c", m.Query).Output()
	if err != nil {
		ball.Reply(fmt.Sprintf("%v", err))
		return
	}
	ball.Reply(string(res))
}