package plugin

import (
	"fmt"
	x "gobot/utils/message"
	y "gobot/utils/simple"
	a "gobot/constanta"
	)
	
func GetInfoGroup(ball *y.S, m *x.Parse)  {
	if !m.IsOwn { ball.Reply(a.FOwner, true); return }
	if len(m.Query) < 1 { ball.Reply(fmt.Sprintf(a.TLink, m.CmdP, "*linkGroup*"), true); return }
	val := ball.GetInfoLink(m.Query)
	var kirimPesan string
	if val.IsAnnounce {
		kirimPesan = "Boleh"
	}	else { 
		kirimPesan = "Tidak boleh"
	}
	var editInfo string
	if val.IsLocked { 
		editInfo = "Boleh"
	} else { 
		editInfo = "Tidak Boleh"
	}
	ball.Reply(fmt.Sprintf("INFO GROUP\n\nNama Group: %s\nID: %s\nDibuat Oleh: %s\nTanggal Pembuatan: %s\nEdit info: %s\nKirim Pesan: %s\nDeskripsi: %s", val.Name, val.JID, val.OwnerJID, val.GroupCreated, editInfo, kirimPesan, val.Topic), true)
}