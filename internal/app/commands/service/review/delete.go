package review

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *ServiceReviewCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	res, err := c.reviewService.Remove(uint64(idx))
	if !res {
		log.Printf("fail to delete review with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Review removed")
	c.bot.Send(msg)
}
