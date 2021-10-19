package review

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *ServiceReviewCommander) Default(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "/help__service__review to view available commands.")
	c.bot.Send(msg)
}
