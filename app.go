package main

import (
	"bots/lgate"
	"bots/loilo"
	"bots/miraiseed"
	"bots/utils"
	"fmt"
)

func init() {
	utils.LoggingSetting("bots.log")
}

func main() {
	fmt.Println("Yo")
	loilo.Bot()
	lgate.Bot()
	miraiseed.Bot()
}
