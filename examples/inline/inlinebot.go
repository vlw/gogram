package examples

import (
	"github.com/vlw/gogram/telegram"
)

const (
	appID    = 6
	appHash  = "YOUR_APP_HASH"
	botToken = "YOUR_BOT_TOKEN"
)

func main() {
	// Create a new client
	client, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    appID,
		AppHash:  appHash,
		LogLevel: telegram.LogInfo,
	})

	// Connect to the server
	if err := client.Connect(); err != nil {
		panic(err)
	}

	// Authenticate the client using the bot token
	if err := client.LoginBot(botToken); err != nil {
		panic(err)
	}

	// Add a inline query handler
	client.AddInlineHandler(telegram.OnInlineQuery, HelloWorld)

	// Start polling
	client.Idle()
}

func HelloWorld(i *telegram.InlineQuery) error {
	builder := i.Builder()
	builder.Article("Hello World", "Hello World", "This is a test article")
	_, err := i.Answer(builder.Results())
	return err
}
