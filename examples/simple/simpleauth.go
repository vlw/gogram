package main

import (
	tg "github.com/vlw/gogram/telegram"
)

func main() {
	client, _ := tg.NewClient(tg.ClientConfig{
		AppID:   6,
		AppHash: "",
	})
	client.Start() // this calls client.Connect()
	// Then asks for userInput for phone num or botToken
	// Further authing with that, fully automatic the auth code flow.
	// fmt.Println(client.ExportSession()) -> "Bwjiaw27sbss..."
	// client Do anything
	// client.Idle()
}
