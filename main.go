package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	v "github.com/TwiN/go-color"
	// Library Whatsmeow
	"go.mau.fi/whatsmeow"
	db "go.mau.fi/whatsmeow/store/sqlstore"
	ngelog "go.mau.fi/whatsmeow/util/log"
	// Sqlite
	_ "github.com/mattn/go-sqlite3"
	// Print to terminal code
	qrCode "github.com/mdp/qrterminal"
	// Module local
	handler "gobot/utils/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error saat meload file .env")
	}
	dbName := os.Getenv("SESSION_NAME")
	typeDb := os.Getenv("DB_TYPE")
	dbLog := ngelog.Stdout("Database", "ERROR", true)
	container, err := db.New(typeDb, "file:gobot.Session/"+dbName+".db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	ds, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	logs := ngelog.Stdout("Client", "ERROR", false)
	conn := whatsmeow.NewClient(ds, logs)
	msgUp := handler.MessageUpsert(conn)
	conn.SetForceActiveDeliveryReceipts(true)
	conn.AddEventHandler(msgUp)
	
	if conn.Store.ID == nil {
		qrChan, _ := conn.GetQRChannel(context.Background())
		err = conn.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrCode.GenerateHalfBlock(evt.Code, qrCode.M, os.Stdout)
				fmt.Println("Silahkan SCAN QR Code dibawah ini:\n")
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err = conn.Connect()
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(v.InCyan("Koneksi telah tersambung whatsapp web..."))
		fmt.Println(v.InPurple("------------------------------------------------"))
		if err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	conn.Disconnect()
}
