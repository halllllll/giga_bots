package discord

import (
	"bots/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/**
* reference: https://blog.narumium.net/2019/08/02/%E3%80%90go%E3%80%91discord%E3%81%AEwebhook%E3%81%A7%E9%80%9A%E7%9F%A5%E3%83%9C%E3%83%83%E3%83%88%E3%82%92%E4%BD%9C%E3%82%8B/
 */
type discordImg struct {
	URL string `json:"url"`
	H   int    `json:"height"`
	W   int    `json:"width"`
}
type discordAuthor struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	Icon string `json:"icon_url"`
}
type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type discordEmbed struct {
	Title  string         `json:"title"`
	Desc   string         `json:"description"`
	URL    string         `json:"url"`
	Color  int            `json:"color"`
	Image  discordImg     `json:"image"`
	Thum   discordImg     `json:"thumbnail"`
	Author discordAuthor  `json:"author"`
	Fields []discordField `json:"fields"`
}

type discordWebhook struct {
	UserName  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []discordEmbed `json:"embeds"`
	TTS       bool           `json:"tts"`
}

func SendWebhook(whurl string, dw *discordWebhook) (err error) {
	j, err := json.Marshal(dw)
	if err != nil {
		utils.ErrLog.Printf("json err: %s\n", err)
		return
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		utils.ErrLog.Printf("new request err: %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.ErrLog.Printf("client err: %s \n", err)
		return
	}
	if resp.StatusCode == 204 {
		utils.StdLog.Printf("successed sending: %v\n", dw) //成功
	} else {
		utils.ErrLog.Printf("missed sending:  %#v\n", resp) //失敗
		return errors.New(fmt.Sprintf("missed sending error: code %d", resp.StatusCode))
	}
	return nil
}
