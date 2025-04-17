package bot

type Quota struct {
	Quantity int
	Limit    int
}

func NewQuota(limit int) Quota {
	return Quota{
		Quantity: 0,
		Limit:    limit,
	}
}
