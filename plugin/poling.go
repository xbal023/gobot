package plugin

import (
	"fmt"
	"strings"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func Polling(ball *y.S, m *x.Parse)  {
	if !m.IsGc { ball.Reply(a.FGroup, true); return }
	query := strings.Split(m.Query, "|")
	if query[0] == "" { ball.Reply(fmt.Sprintf(a.TQuery, m.CmdP, "halo guys | isi@isi@isi...."), true); return }
	if query[1] == "" { ball.Reply(fmt.Sprintf(a.TQuery, m.CmdP, "halo guys | isi@isi@isi...."), true); return
	} else {
		isi := strings.Split(query[1], "@")
		ball.PollMsg(*m.Chat, query[0], isi)
	}
}