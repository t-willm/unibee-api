package subscription

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	redismq "github.com/jackyang-hk/go-redismq"
	"strconv"
	redismq2 "unibee/internal/cmd/redismq"
	"unibee/internal/consumer/webhook/event"
	"unibee/internal/consumer/webhook/user"
	"unibee/utility"
)

type UserAccountCreateListener struct {
}

func (t UserAccountCreateListener) GetTopic() string {
	return redismq2.TopicUserAccountCreate.Topic
}

func (t UserAccountCreateListener) GetTag() string {
	return redismq2.TopicUserAccountCreate.Tag
}

func (t UserAccountCreateListener) Consume(ctx context.Context, message *redismq.Message) redismq.Action {
	utility.Assert(len(message.Body) > 0, "body is nil")
	utility.Assert(len(message.Body) != 0, "body length is 0")
	g.Log().Infof(ctx, "UserAccountCreateListener Receive Message:%s", utility.MarshalToJsonString(message))
	if len(message.Body) > 0 {
		userId, _ := strconv.ParseUint(message.Body, 10, 64)
		if userId > 0 {
			user.SendMerchantUserWebhookBackground(userId, event.UNIBEE_WEBHOOK_EVENT_USER_CREATED)
		}
	}

	return redismq.CommitMessage
}

func init() {
	redismq.RegisterListener(NewUserAccountCreateListener())
	fmt.Println("UserAccountCreateListener RegisterListener")
}

func NewUserAccountCreateListener() *UserAccountCreateListener {
	return &UserAccountCreateListener{}
}
