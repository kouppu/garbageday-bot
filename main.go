package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/suhrr/garbageday-bot/schedule"
)

var env = map[string]string{
	"LINE_CHANNEL_SECRET": "",
	"LINE_CHANNEL_TOKEN":  "",
}

func main() {
	fmt.Println("start")

	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	s, err := schedule.ReadJSON("./schedule.json")
	if err != nil {
		log.Fatal(err)
	}
	gds := schedule.ExtractTodayGarbagedays(s.Garbagedays)
	if len(gds) == 0 {
		fmt.Println("today is not garbage day.")
		os.Exit(0)
	}

	m := schedule.CreateGarbagedayMessage(gds)
	bot, err := linebot.New(env["LINE_CHANNEL_SECRET"], env["LINE_CHANNEL_TOKEN"])
	if err != nil {
		log.Fatal(err)
	}
	if _, err := bot.BroadcastMessage(linebot.NewTextMessage(m)).Do(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("send messages.")
}

func loadEnv() error {
	godotenv.Load(".env")

	for index := range env {
		_env := os.Getenv(index)
		if _env == "" {
			return fmt.Errorf("can not find %s", index)
		}
		env[index] = _env
	}

	return nil
}
