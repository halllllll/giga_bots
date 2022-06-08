package loilo

import (
	"bots/utils"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// そのうち構造体やらにして大元で管理すればいいと思うよ
var (
	service            string = "loilo"
	url                string = "https://status.loilonote.app/ja"
	envName            string = "LOILO_STATUS"
	normalStatMsg      string = "すべてのサービスが正常に稼働しています"
	errorStatMsg       string
	maintenanceStatMsg string
	normalStat         string = "OK"
	abNormalStat       string = "NOT_OK" // これは今後増える可能性が多いにある
)

/**
* 状態が前回と変わっていたらtrue
 */
func ServerStat() (ret string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		utils.ErrLog.Printf("error at access status of %s : %v\n", service, err)
	}
	target := doc.Find(".description-text").Text()
	if os.Getenv(envName) == normalStat && !strings.Contains(target, normalStatMsg) {
		// 正常な状態から異常な状態(種類は不明（3種類くらいはある？）)に遷移した
		utils.InfoLog.Printf("-- !something accidnet occured on SERVICE [ %s ]\n ", service)
		os.Setenv(envName, abNormalStat)
		ret = "CAUTION"
	} else if os.Getenv(envName) != normalStat && strings.Contains(target, normalStatMsg) {
		// 異常な状態(種類は不明（3種類くらいはある？）)から正常な状態に戻った
		utils.InfoLog.Printf("-- !changed status of [ %s ] : SERVICE [ %s ]\n ", os.Getenv(envName), service)
		os.Setenv(envName, normalStat)
		ret = "REPAIRED"
	}
	return
}
