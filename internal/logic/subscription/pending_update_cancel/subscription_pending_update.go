package pending_update_cancel

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	redismq "github.com/jackyang-hk/go-redismq"
	redismq2 "unibee/internal/cmd/redismq"
	"unibee/internal/consts"
	dao "unibee/internal/dao/default"
	"unibee/internal/logic/payment/service"
	entity "unibee/internal/model/entity/default"
	"unibee/internal/query"
	"unibee/utility"
)

func SubscriptionPendingUpdateCancel(ctx context.Context, pendingUpdateId string, reason string) error {
	one := query.GetSubscriptionPendingUpdateByPendingUpdateId(ctx, pendingUpdateId)
	if one != nil {
		if one.Status == consts.PendingSubStatusFinished {
			return nil
		}
		if one.Status == consts.PendingSubStatusCancelled {
			return nil
		}

		if len(one.InvoiceId) > 0 {
			invoice := query.GetInvoiceByInvoiceId(ctx, one.InvoiceId)
			if invoice != nil {
				payment := query.GetPaymentByPaymentId(ctx, invoice.PaymentId)
				if payment != nil {
					err := service.PaymentGatewayCancel(ctx, payment)
					if err != nil {
						g.Log().Errorf(ctx, "PaymentGatewayCancel Error:%s", err.Error())
					}
				}
			}
		}
		_, err := dao.SubscriptionPendingUpdate.Ctx(ctx).Data(g.Map{
			dao.SubscriptionPendingUpdate.Columns().Status:    consts.PendingSubStatusCancelled,
			dao.SubscriptionPendingUpdate.Columns().GmtModify: gtime.Now(),
		}).Where(dao.SubscriptionPendingUpdate.Columns().Id, one.Id).
			WhereIn(dao.SubscriptionPendingUpdate.Columns().Status, []int64{consts.PendingSubStatusInit, consts.PendingSubStatusCreate}).OmitNil().Update()
		if err != nil {
			return err
		} else {
			//send mq message
			_, _ = redismq.Send(&redismq.Message{
				Topic:      redismq2.TopicSubscriptionPendingUpdateCancel.Topic,
				Tag:        redismq2.TopicSubscriptionPendingUpdateCancel.Tag,
				Body:       pendingUpdateId,
				CustomData: map[string]interface{}{"CreateFrom": utility.ReflectCurrentFunctionName()},
			})
		}
	}
	return nil
}

func CancelOtherUnfinishedPendingUpdatesBackground(subscriptionId string, pendingUpdateId string, reason string) {
	go func() {
		var err error
		ctx := context.Background()
		defer func() {
			if exception := recover(); exception != nil {
				if v, ok := exception.(error); ok && gerror.HasStack(v) {
					err = v
				} else {
					err = gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception)
				}
				g.Log().Errorf(ctx, "CancelOtherUnfinishedPendingUpdatesBackground Panic Error:%s", err.Error())
				return
			}
		}()
		var mainList []*entity.SubscriptionPendingUpdate
		err = dao.SubscriptionPendingUpdate.Ctx(ctx).
			Where(dao.SubscriptionPendingUpdate.Columns().SubscriptionId, subscriptionId).
			WhereNot(dao.SubscriptionPendingUpdate.Columns().PendingUpdateId, pendingUpdateId).
			WhereLT(dao.SubscriptionPendingUpdate.Columns().Status, consts.PendingSubStatusFinished).
			Limit(0, 100).
			OmitEmpty().Scan(&mainList)
		if err != nil {
			g.Log().Errorf(ctx, "CancelOtherUnfinishedPendingUpdatesBackground Search List Error:%s", err.Error())
		}
		for _, one := range mainList {
			err = SubscriptionPendingUpdateCancel(ctx, one.PendingUpdateId, reason)
			if err != nil {
				g.Log().Errorf(ctx, "CancelOtherUnfinishedPendingUpdatesBackground SubscriptionPendingUpdateCancel Error:%s", err.Error())
			}
		}
	}()
}
