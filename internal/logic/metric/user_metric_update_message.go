package metric

type UserMetricUpdateMessage struct {
	UserId         uint64 `json:"userId"              description:"user id"`
	SubscriptionId string `json:"subscriptionId" description:"subscription id"`
	Description    string `json:"description"         description:"description"`
}
