package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Menu(ball *y.S, m *x.Parse)  {
	menu := `MENU SELF`+`

` +m.Pref+ `setname`+`
` +m.Pref+ `setdesk`+`
` +m.Pref+ `gclink`+`
` +m.Pref+ `gcrevoke`+`
` +m.Pref+ `gcopen`+`
` +m.Pref+ `gcclose`+`
` +m.Pref+ `gclock`+`
` +m.Pref+ `gcunlock`+`
` +m.Pref+ `polling`+`
` +m.Pref+ `ping`+`
` +m.Pref+ `getinfo *link*`+`
` +m.Pref+ `out`+`
` +m.Pref+ `metadata`+`
` +m.Pref+ `bongkar`+`

`+`Bot ini dibuat dengan Bahasa Golang`+`
`+`Library dari whatsmeow: https://github.com/tulir/whatsmeow`
	ball.Reply(menu)
}