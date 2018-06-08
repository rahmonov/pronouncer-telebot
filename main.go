package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
	"log"
	"fmt"
)

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token: "486118859:AAFYyxr-OYEfn4Pcm4z3jw-UIqYf1vHdpAs",
		Poller: &tb.LongPoller{Timeout:10*time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/read", func(m *tb.Message) {
		handleRead(bot, m)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		bot.Send(m.Sender, "Я не понимаю простой текст. Воспользуйтесь какой нибудь командой.")
	})

	bot.Start()
}

func handleRead(bot *tb.Bot, m *tb.Message) {
	wordToPronounce := m.Payload

	if wordToPronounce == "" {
		bot.Send(m.Sender, "Напишите слово...")
		return
	}

	pronunciationURL := fmt.Sprintf(
		"https://ssl.gstatic.com/dictionary/static/sounds/oxford/%s--_gb_1.mp3", wordToPronounce)
	pronunciationAudio := &tb.Audio{File: tb.FromURL(pronunciationURL)}

	_, err := bot.Send(m.Sender, pronunciationAudio)
	if err != nil {
		bot.Send(m.Sender, "Что-то пошло не так. Убедитесь, что слово правильно написано")
		log.Println(err)
	}
}



