package review

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	model "github.com/ozonmp/omp-bot/internal/model/service"
	"log"
)

func (c *ServiceReviewCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	review := model.Review{}
	err := json.Unmarshal([]byte(args), &review)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, `Use json format { "key1": "value1", "key2": 1 }`)
		c.bot.Send(msg)
		log.Printf("wrong args: %s with error %v", args, err)
		return
	}

	res, err := c.reviewService.Create(review)
	if err != nil {
		log.Printf("Can't create new review")
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, fmt.Sprintf("Created new review: entry №%d", res))
	c.bot.Send(msg)
}
