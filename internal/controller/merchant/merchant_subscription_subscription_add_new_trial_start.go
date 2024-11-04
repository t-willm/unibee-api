package merchant

import (
	"context"
	"unibee/api/merchant/subscription"
	"unibee/internal/logic/subscription/service"
	"unibee/utility"
)

func (c *ControllerSubscription) AddNewTrialStart(ctx context.Context, req *subscription.AddNewTrialStartReq) (res *subscription.AddNewTrialStartRes, err error) {
	utility.Assert(req.AppendTrialEndHour > 0, "invalid AppendTrialEndTime, should greater than 0")
	err = service.SubscriptionAddNewTrialEnd(ctx, req.SubscriptionId, req.AppendTrialEndHour)
	if err != nil {
		return nil, err
	}
	return &subscription.AddNewTrialStartRes{}, nil
}
