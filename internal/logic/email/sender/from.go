package sender

type Sender struct {
	Name    string `json:"name" description:"name"`
	Address string `json:"address" description:"address"`
}

func GetDefaultSender() *Sender {
	return &Sender{
		Name:    "no-reply",
		Address: "no-reply@unibee.dev",
	}
}
