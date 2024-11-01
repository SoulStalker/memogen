package telegram

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

type Service struct {
	b *tele.Bot
}

func NewService(token string) *Service {
	pref := tele.Settings{
		//Token: os.Getenv("TELEGRAM_TOKEN"),
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/hello", func(context tele.Context) error {
		return context.Send("Hello!")

	})

	b.Start()
}
