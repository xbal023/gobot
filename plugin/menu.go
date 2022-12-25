package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func Menu(ball *y.S, m *x.Parse)  {
	menu := `MENU SELF`+`

`+`
CONVERT
` +m.Pref+ `stiker`+`
`+`
GRUP
` +m.Pref+ `setname`+`
` +m.Pref+ `setppgc`+`
` +m.Pref+ `setdesk`+`
` +m.Pref+ `gclink`+`
` +m.Pref+ `gcrevoke`+`
` +m.Pref+ `gcopen`+`
` +m.Pref+ `gcclose`+`
` +m.Pref+ `gclock`+`
` +m.Pref+ `gcunlock`+`
` +m.Pref+ `polling`+`
`+`
OWNER
` +m.Pref+ `getinfo *link*`+`
` +m.Pref+ `metadata *inGroup*`+`
` +m.Pref+ `download`+`
` +m.Pref+ `bongkar`+`
` +m.Pref+ `upload`+`
` +m.Pref+ `out *inGroup*`+`
` +m.Pref+ `exec`+`
` +m.Pref+ `u`+`

`+`Bot ini dibuat dengan Bahasa Golang`+`
`+`Library source:`+`
`+`https://github.com/tulir/whatsmeow`
	ball.Reply(menu)
}