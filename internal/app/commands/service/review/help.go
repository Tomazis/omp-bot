package review

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *ServiceReviewCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__service__review - print list of commands\n"+
			"/get__service__review - get review by id\n"+
			"/list__service__review - get a list of reviews\n"+
			"/delete__service__review - delete review by id\n"+
			"/new__service__review - create a new review\n"+
			"/edit__service__review - edit review",
	)

	c.bot.Send(msg)
}
