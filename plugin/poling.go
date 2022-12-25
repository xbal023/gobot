package plugin

import (
	"strings"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Polling(ball *y.S, m *x.Parse)  {
	if !m.IsGc {
		ball.Reply("Khusus di group", true)
		return
	}
	query := strings.Split(m.Query, "|")
	if query[0] == "" {
		ball.Reply("masukan teks nya contoh .polling halo guys | isi@isi@isi....", true)
		return
	}
	if query[1] == "" {
		ball.Reply("contoh .polling halo guys | isi@isi@isi....", true)
		return
	} else {
		isi := strings.Split(query[1], "@")
		ball.PollMsg(*m.Chat, query[0], isi)
	}
}