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

type SubscriptionPendingUpdateCancelListener struct {
}

func (t SubscriptionPendingUpdateCancelListener) GetTopic() string {
	return redismq2.TopicSubscriptionPendingUpdateCancel.Topic
}

func (t SubscriptionPendingUpdateCancelListener) GetTag() string {
	return redismq2.TopicSubscriptionPendingUpdateCancel.Tag
}

func (t SubscriptionPendingUpdateCancelListener) Consume(ctx context.Context, message *redismq.Message) redismq.Action {
	utility.Assert(len(message.Body) > 0, "body is nil")
	utility.Assert(len(message.Body) != 0, "body length is 0")
	g.Log().Infof(ctx, "SubscriptionPendingUpdateCancelListener Receive Message:%s", utility.MarshalToJsonString(message))
	one := query.GetSubscriptionPendingUpdateByPendingUpdateId(ctx, message.Body)
	if one != nil {
		subscription.SendMerchantSubscriptionPendingUpdateWebhookBackground(one, event.UNIBEE_WEBHOOK_EVENT_SUBSCRIPTION_PENDING_UPDATE_CANCELLED, message.CustomData)
	}
	return redismq.CommitMessage
}

func init() {
	redismq.RegisterListener(NewSubscriptionPendingUpdateCancelListener())
	fmt.Println("SubscriptionPendingUpdateCancelListener RegisterListener")
}

func NewSubscriptionPendingUpdateCancelListener() *SubscriptionPendingUpdateCancelListener {
	return &SubscriptionPendingUpdateCancelListener{}
}
