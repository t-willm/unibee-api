package subscription

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"unibee/internal/consumer/webhook/event"
	"unibee/internal/consumer/webhook/log"
	"unibee/internal/consumer/webhook/message"
	"unibee/internal/logic/subscription/service"
	entity "unibee/internal/model/entity/default"
	"unibee/utility"
)

func SendMerchantSubscriptionPendingUpdateWebhookBackground(one *entity.SubscriptionPendingUpdate, event event.WebhookEvent, metadata map[string]interface{}) {
	go func() {
		ctx := context.Background()
		g.Log().Debugf(ctx, "SendMerchantSubscriptionnPendingUpdateWebhookBackground event:%v", event)
		var err error
		defer func() {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					err = v
				} else {
					err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
				}
				log.PrintPanic(ctx, err)
				return
			}
		}()
		pendingUpdate := service.GetSubscriptionPendingUpdateEventByPendingUpdateId(ctx, one.PendingUpdateId)
		utility.Assert(pendingUpdate != nil, "SendMerchantSubscriptionPendingUpdateWebhookBackground SubscriptionPendingUpdateEvent Not Found")
		message.SendWebhookMessage(ctx, event, one.MerchantId, utility.FormatToGJson(pendingUpdate), "", "", metadata)
	}()
}
