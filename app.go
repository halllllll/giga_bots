package main

import (
	"bots/lgate"
	"bots/loilo"
	"bots/miraiseed"
	"bots/utils"
	"fmt"
	"os"

	ua "github.com/wux1an/fake-useragent"
)

func init() {
	utils.LoggingSetting("bots.log")
}

func main() {
	fmt.Println("Yo")
	// 環境変数読み取り
	fmt.Println(os.Getenv("LOILO_STATUS"))
	loilo.Bot()
	lgate.Bot()
	miraiseed.Bot()
	for i := 0; i < 10; i++ {
		fmt.Println(ua.Random())
	}
}
