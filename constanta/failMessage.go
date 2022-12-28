package fail
const (
	g = "*Access denied 📛*\n\n"
		// FOR ACCESS
	FOwner = g+ "```Fitur ini khusus pemilik bot ini```"
	FAdmin = g+ "```Fitur ini di khusus kan oleh admin group```"
	FBotAdmin = g+ "```Bot tidak admin! maaf ngga bisa```"
	FGroup = g+ "```Fitur ini di khusus kan untuk di group```"
)

const (
	q = "*Using Query 🔑*\n\n"
	// FOR QUERY | TEXT
	TPath = q+ "```Masukan Path file nya dengan benar dengan benar contoh : %s %s```"
	TLink = q+ "```Masukan link nya setelah command, contoh : %s %s```"
	TDesc = q+ "```Masukan deskripsi group nya setelah command, contoh : %s %s```"
	TName = q+ "```Masukan nama / subject group nya setelah command, contoh : %s %s```"
	TQuery = q+ "```Masukan Query nya, contoh : %s %s```"
	)
	
const (
	e = "*Using Quoted ⤴️*\n\n"
	// FOR QUOTED MESSAGE | MEDIA 
	QDel = e+ "```Reply pesan yg ingin kamu lenyapkan```"
	QMedia = e+ "```Reply / kirim media dengan caption command```"
	)
const (
	s = "*Succes ✔️*\n\n"
	// FOR SUCCES
	GCClose = s+ "```Sukses menutup chat Group, sekarang hanya admin yg dapat mengirim pesan```"
	GCOpen = s+ "```Sukses membuka chat Group, sekarang seluruh member dapat mengirim pesan```"
	InfoOpen = s+ "```Sukses membuka edit info, Sekarang semua member dapat mengedit info group```"
	InfoClose = s+ "```Sukses menutup edit info, Sekarang hanya admin yang dapat mengedit info group```"
	GCPP = s+ "```Sukses mengganti Photo group dengan yg baru```"
	GCName = s+ "```Sukses mengganti nama group dengan yg baru```"
	GCDesc = s+ "```Sukses mengganti Deskripsi dengan yg baru```"
	)
