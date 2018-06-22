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

			b.Delete(m) // delete join service message
			if err = b.Ban(m.Chat, member); err != nil {
				log.Println("ban error", err)
				return
			}

			// del xx remove xx service message
			id3Message, err := b.Send(m.Chat, "Member kicked")
			b.Delete(id3Message)
			if err != nil {
				log.Println("send error", err)
				return
			}
			id1 := m.ID
			id2 := id3Message.ID - 1

			for i := id2; i > id1; i-- {
				m1, err := b.Send(m.Chat, "reply", &tb.SendOptions{ReplyTo: &tb.Message{ID: i}})
				if err != nil {
					log.Println("send reply error", err)
					return
				}
				if m1.ReplyTo.UserLeft.ID == user.ID {
					b.Delete(m1)         // delete remove service message
					b.Delete(m1.ReplyTo) // delete reply message
					return
				}
			}

			log.Println("success ban user")
		}
	})
	b.Start()
}
