package review

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *ServiceReviewCommander) Edit(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Error! Command not implemented!")
	c.bot.Send(msg)
}
