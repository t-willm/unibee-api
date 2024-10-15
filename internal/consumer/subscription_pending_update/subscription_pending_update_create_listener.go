package subscription

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	redismq "github.com/jackyang-hk/go-redismq"
	redismq2 "unibee/internal/cmd/redismq"
	"unibee/internal/consumer/webhook/event"
	subscription "unibee/internal/consumer/webhook/subscription_pending_update"
	"unibee/internal/query"
	"unibee/utility"
)

type SubscriptionPendingUpdateCreateListener struct {
}

func (t SubscriptionPendingUpdateCreateListener) GetTopic() string {
	return redismq2.TopicSubscriptionPendingUpdateCreate.Topic
}

func (t SubscriptionPendingUpdateCreateListener) GetTag() string {
	return redismq2.TopicSubscriptionPendingUpdateCreate.Tag
}

func (t SubscriptionPendingUpdateCreateListener) Consume(ctx context.Context, message *redismq.Message) redismq.Action {
	utility.Assert(len(message.Body) > 0, "body is nil")
	utility.Assert(len(message.Body) != 0, "body length is 0")
	g.Log().Debugf(ctx, "SubscriptionPendingUpdateCreateListener Receive Message:%s", utility.MarshalToJsonString(message))
	one := query.GetSubscriptionPendingUpdateByPendingUpdateId(ctx, message.Body)
	if one != nil {
		subscription.SendMerchantSubscriptionPendingUpdateWebhookBackground(one, event.UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_PENDING_UPDATE_CREATE, message.CustomData)
	}
	return redismq.CommitMessage
}

func init() {
	redismq.RegisterListener(NewSubscriptionPendingUpdateCreateListener())
	fmt.Println("SubscriptionPendingUpdateCreateListener RegisterListener")
}

func NewSubscriptionPendingUpdateCreateListener() *SubscriptionPendingUpdateCreateListener {
	return &SubscriptionPendingUpdateCreateListener{}
}
