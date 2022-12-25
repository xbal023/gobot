package plugin

import (
	"time"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Ping(ball *y.S, m *x.Parse)  {
	start := time.Now()
	ball.GetMetadata(*m.Chat)
	to := time.Since(start)
	ball.Reply("Kecepatan bot respon : " +string(to)+ " sekon")
}