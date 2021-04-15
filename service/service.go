package service

import (
	"context"
	"log"
	"strings"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/discord"
)

type discordInfo struct {
	Token   string
	Message string
	Channel string
	Title   string
}

//DoNotify
func DoNotify(msg string, token string) {
	var discordOpts discordInfo
	discordOpts.Channel = "340876853785001985"
	discordOpts.Token = token
	discordOpts.Message = msg
	// Create our notifications distributor.
	notifier := notify.New()

	// Create a telegram service. Ignoring error for demo simplicity.
	discordService := discord.New()

	discordService.AuthenticateWithBotToken(token)

	chn := strings.Split(discordOpts.Channel, ",")
	for _, v := range chn {
		if len(v) <= 0 {
			log.Println("EmptyChannel")
		}

		discordService.AddReceivers(v)
	}

	notifier.UseServices(discordService)

	err := notifier.Send(
		context.Background(),
		discordOpts.Title,
		discordOpts.Message,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully sent!")
}
