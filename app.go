package main

import (
	"bots/discord"
	"bots/lgate"
	"bots/loilo"
	"bots/miraiseed"
	"bots/utils"
	"fmt"
	"os"
)

var (
	discord_webhookUrl string
	cautionImg_url     string
	repairImg_url      string
)

func init() {
	utils.LoggingSetting("bots.log")
	discord_webhookUrl = os.Getenv("DISCORD_WEBHOOKURL")
	if discord_webhookUrl == "" {
		utils.ErrLog.Println("not found env value of discord_webhookurl")
		return
	}
	cautionImg_url = os.Getenv("CAUTIONIMG_URL")
	if cautionImg_url == "" {
		utils.ErrLog.Println("not found env value of cutionimg_url")
		return
	}
	repairImg_url = os.Getenv("REPAIRIMG_URL")
	if repairImg_url == "" {
		utils.ErrLog.Println("not found env value of repairimg_url")
		return
	}

}

func main() {
	// loilo

	if ret := loilo.ServerStat(); ret == "" {
		var (
			imgUrl string
			msg    string
		)
		if ret == "CAUTION" {
			imgUrl = cautionImg_url
			msg = "なにかが起きているようです"
		} else if ret == "REPAIRED" {
			imgUrl = repairImg_url
			msg = "障害が治ったようです"
		}
		fmt.Println("わーい")
		fmt.Println(imgUrl)
		dw := &discord.DiscordWebhook{
			UserName:  "test",
			AvatarURL: "",
			Content:   msg,
			Embeds:    []discord.DiscordEmbed{},
			TTS:       false,
		}
		dw.Embeds = []discord.DiscordEmbed{
			discord.DiscordEmbed{
				Title:  "",
				Desc:   "",
				URL:    "https://blog.narumium.net/",
				Color:  0x550000,
				Author: discord.DiscordAuthor{Name: "Narumium"},
				Image:  discord.DiscordImg{URL: imgUrl},
			},
		}
		if err := discord.SendWebhook(discord_webhookUrl, dw); err != nil {
			utils.ErrLog.Println("loilo bot sending error: ", err)
		}
	}

	// l-gate

	// miraiseed

	lgate.Bot()
	miraiseed.Bot()
}
