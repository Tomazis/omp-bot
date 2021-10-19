package review

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ServiceReviewCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	review, err := c.reviewService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get review with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		review.String(),
	)

	c.bot.Send(msg)
}
