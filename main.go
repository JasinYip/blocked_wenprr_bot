package main

import (
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("token"),
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnUserJoined, func(m *tb.Message) {
		user := m.UserJoined
		log.Println("debug: user join", user.FirstName, user.LastName)
		if strings.Contains(user.FirstName, "wenprr") || strings.Contains(user.LastName, "wenprr") {
			member, err := b.ChatMemberOf(m.Chat, user)
			if err != nil {
				log.Println("chat member error", err)
				return
			}
			if err = b.Delete(m); err != nil {
				log.Println("delete error", err)
				return
			}
			if err = b.Ban(m.Chat, member); err != nil {
				log.Println("ban error", err)
				return
			}

			log.Println("success ban user")
		}
	})
	b.Start()
}
