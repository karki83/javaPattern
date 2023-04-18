package main

import (
"context"
"fmt"
"log"
"os"
"github.com/shomali11/slacker")


func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel {
		fmt.Println("command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4210161750929-4589236074784-Zg7KJIq3VbxSmbiTlVIVemdm")
    os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04GZVCL37B-4589061699088-8ccc35c3b1989cc7b11d7b4eb5c441d7e7eec614ac4954a2820f3c2540dd23c8")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}