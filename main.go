package main

import (
	"clout-bot/internal/handlers"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	dg "github.com/bwmarrin/discordgo"
	env "github.com/joho/godotenv"
)

const (
	botAuthTokenEnv = "BOT_AUTH_TOKEN"
)

var (
	Token string
	Bot *dg.Session
	err error
)

func init() {
	if err := env.Load(".env"); err != nil {
		log.Fatal("failed to load .env file!")
	}
	if tokenEnv, ok := os.LookupEnv(botAuthTokenEnv); !ok {
		log.Fatal("BOT_AUTH_TOKEN environment variable missing!")
	} else {
		Token = tokenEnv
	}
}

func main() {
	if Bot, err = dg.New("Bot " + Token); err != nil {
		log.Fatal("failed to create Discord session: ", err)
	}
	Bot.AddHandler(handlers.ParseEvent)
	Bot.Identify.Intents = dg.IntentsGuildMessages | dg.IntentsGuildMessageReactions
	if err = Bot.Open(); err != nil {
		log.Fatal("error opening connection: ", err)
	}
	fmt.Println("Bot is now running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	Bot.Close()
}

