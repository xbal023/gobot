package plugin

import (
	x "gobot/utils/message"
	y "gobot/utils/simple"
	)
	
func GetInfoGroup(ball *y.S, m *x.Parse)  {
	if !m.IsOwn {
		ball.Reply("fitur ini khusus owner")
		return
	}
	if len(m.Query) < 1 {
		ball.Reply("Masukan link yang ingin di ambil info");
		return
	}
	val := ball.GetInfoLink(m.Query)
	var kirimPesan string
	if val.IsAnnounce {
		kirimPesan = "Boleh"
	} else {
		kirimPesan = "Tidak boleh"
	}
	var editInfo string
	if val.IsLocked {
		editInfo = "Boleh"
	} else {
		editInfo = "Tidak Boleh"
	}
	ball.Reply(`INFO GROUP\n\nNama Group: `+ string(val.Name)+ `\nID: ` +val.JID.String()+ `\nEdit info: ` +editInfo+ `\nKirim Pesan: `+kirimPesan+ `\nDeskripsi: ` +string(val.Topic))
}