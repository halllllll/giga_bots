package main

import (
	"bots/discord"
	"bots/lgate"
	"bots/loilo"
	"bots/miraiseed"
	"bots/utils"
	"os"
)

func init() {
	utils.LoggingSetting("bots.log")
}

func main() {
	// 環境変数読み取り
	loilo_statusUrl := os.Getenv("LOILO_STATUSURL")
	if loilo_statusUrl == "" {
		utils.ErrLog.Println("not found env value of loilo_statusurl")
		return
	}
	discord_webhook := os.Getenv("DISCORD_WEBHOOKURL")
	if discord_webhook == "" {
		utils.ErrLog.Println("not found env value of discord_webhookurl")
		return
	}
	dw := &discord.DiscordWebhook{
		UserName:  "test",
		AvatarURL: "",
		Content:   "にゃ〜",
		Embeds:    []discord.DiscordEmbed{},
		TTS:       false,
	}
	discord.SendWebhook(discord_webhook, dw)
	loilo.Bot()
	lgate.Bot()
	miraiseed.Bot()
}
