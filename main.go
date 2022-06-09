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

type Loilo struct {
	Stat
}

func (loilo *Loilo) ServerStat() (Stat, bool) {
	doc, err := goquery.NewDocument(loilo.Url)
	if err != nil {
		utils.ErrLog.Printf("error at access status of %s : %v\n", loilo.Service, err)
	}
	target := doc.Find(".description-text").Text()
	// 現在が正常な場合かどうかの判定
	fmt.Println("target... : ", target, strings.Contains(target, loilo.NormalStatMsg))
	return loilo.Stat, strings.Contains(target, loilo.NormalStatMsg)
}

type LGate struct {
	Stat
}

type MiraiSeed struct {
	Stat
}

type Bot interface {
	ServerStat() (Stat, bool)
}

func Go(bot Bot) error {
	// 正常な場合はOK==true
	stat, ok := bot.ServerStat()
	pre := os.Getenv(stat.EnvName) // 今回取得する前の状態
	var (
		err    error
		imgUrl string
		msg    string
	)

	if pre == "OK" && !ok {
		// 前回まで正常な状態から今回は異常な状態(種類は不明（3種類くらいはある？）)に遷移した
		utils.InfoLog.Printf("-- !something accidnet occured on SERVICE [ %s ]\n ", stat.Service)
		imgUrl = cautionImg_url
		msg = "なにかが起きているようです"
		if err = os.Setenv(stat.EnvName, "NOT_OK"); err != nil {
			utils.ErrLog.Println(err)
		}
	} else if pre == "NOT_OK" && ok {
		// 異常 -> 正常
		utils.InfoLog.Printf("-- repaired : SERVICE [ %s ]\n ", stat.Service)
		imgUrl = repairImg_url
		msg = "復旧したようです"
		if err = os.Setenv(stat.EnvName, "OK"); err != nil {
			utils.ErrLog.Println(err)
		}
	} else {
		// 前回と変わらない値のはず
		utils.StdLog.Printf("[Negative] pre: %s\n", pre)
		return nil
	}

	// このへんはあとでわけるのをかんがえます
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
	if err = discord.SendWebhook(discord_webhookUrl, dw); err != nil {
		utils.ErrLog.Printf("%s bot sending error: %s", stat.Service, err)
	}
	return err
}

func main() {
	loilo := &Loilo{
		Stat: Stat{
			Service:       "loilo",
			Url:           "https://status.loilonote.app/ja",
			EnvName:       "LOILO_STATUS",
			NormalStatMsg: "すべてのサービスが正常に稼働しています",
		},
	}
	var bot Bot = loilo
	Go(bot)

	// l-gate

	// miraiseed

	// 環境変数を変更してもプロセスを抜けるとなかったことになるので
	defer syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

}

func Say() {
	fmt.Println("env:")
	fmt.Println(os.Getenv("LOILO_STATUS"))
}
