package main

import (
	"bots/discord"
	"bots/utils"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/PuerkitoBio/goquery"
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

type Stat struct {
	Service            string
	Url                string
	EnvName            string
	NormalStatMsg      string
	ErrorStatMsg       string
	MaintenanceStatMsg string
}

func (stat *Stat) Bot() {
	if msg, err := stat.ServerStat(); msg != "" && err == nil {
		var (
			imgUrl string
		)
		if msg == "CAUTION" {
			imgUrl = cautionImg_url
			msg = "なにかが起きているようです"
		} else if msg == "REPAIRED" {
			imgUrl = repairImg_url
			msg = "復旧したようです"
		}
		dw := &discord.DiscordWebhook{
			UserName:  fmt.Sprintf("%s_監視ちゃん", stat.Service),
			AvatarURL: "",
			Content:   msg,
			Embeds:    []discord.DiscordEmbed{},
			TTS:       false,
		}
		dw.Embeds = []discord.DiscordEmbed{
			discord.DiscordEmbed{
				Title: fmt.Sprintf("%s サーバーステータス", stat.Service),
				Desc:  "",
				URL:   stat.Url,
				Color: 0x550000,
				Image: discord.DiscordImg{URL: imgUrl},
			},
		}
		if err := discord.SendWebhook(discord_webhookUrl, dw); err != nil {
			utils.ErrLog.Printf("%s bot sending error: %s", stat.Service, err)
		}
	}
}

func (stat *Stat) ServerStat() (string, error) {
	var (
		ret string
		err error
	)
	doc, err := goquery.NewDocument(stat.Url)
	if err != nil {
		utils.ErrLog.Printf("error at access status of %s : %v\n", stat.Service, err)
	}
	switch stat.Service {
	case "loilo":
		target := doc.Find(".description-text").Text()
		if os.Getenv(stat.EnvName) == "OK" && !strings.Contains(target, stat.NormalStatMsg) {
			// 正常な状態から異常な状態(種類は不明（3種類くらいはある？）)に遷移した
			utils.InfoLog.Printf("-- !something accidnet occured on SERVICE [ %s ]\n ", stat.Service)
			if err = os.Setenv(stat.EnvName, "NOT_OK"); err != nil {
				utils.ErrLog.Println(err)
			}
			ret = "CAUTION"
		} else if os.Getenv(stat.EnvName) != "OK" && strings.Contains(target, stat.NormalStatMsg) {
			// 異常な状態(種類は不明（3種類くらいはある？）)から正常な状態に戻った
			utils.InfoLog.Printf("-- !changed status of [ %s ] : SERVICE [ %s ]\n ", os.Getenv(stat.EnvName), stat.Service)
			if err = os.Setenv(stat.EnvName, "OK"); err != nil {
				utils.ErrLog.Println(err)
			}
			ret = "REPAIRED"
		} else {
		}
	case "l-gate":
	case "miraiseed":

	}
	return ret, err
}

func main() {
	// loilo

	stat := &Stat{
		Service:       "loilo",
		Url:           "https://status.loilonote.app/ja",
		EnvName:       "LOILO_STATUS",
		NormalStatMsg: "すべてのサービスが正常に稼働しています",
	}
	stat.Bot()

	// l-gate

	// l-gate

	// miraiseed

	// 環境変数を変更してもプロセスを抜けるとなかったことになるので
	defer syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

}
