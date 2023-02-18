package discord

import (
	"github.com/JesusKian/WriteUp/src/config"
	"github.com/gtuk/discordwebhook"

	"fmt"
)

func Send(_title, _link, _pubDate string) {
	var (
		err      error                  = nil
		url      string                 = config.DiscordWebhook
		username string                 = "🔒 Bug Bounty WriteUps"
		content  string                 = fmt.Sprintf("> ```yaml\n> - 📰 %s\n> - 📅 %s\n> ```\n> **🔗 %s**", _title, _pubDate, _link)
		message  discordwebhook.Message = discordwebhook.Message{
			Username: &username,
			Content:  &content,
		}
	)

	err = discordwebhook.SendMessage(url, message)
	if err != nil {
		config.SetLog("E", "discord.Send() -> Couldn't Send Message To Webhook")
		config.SetLog("D", err.Error())
	}
}
