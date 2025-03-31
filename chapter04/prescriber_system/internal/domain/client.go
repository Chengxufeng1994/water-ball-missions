package domain

type Client struct {
	name string
}

var _ Subscriber = (*Client)(nil)

func NewClient(name string) *Client {
	return &Client{
		name: name,
	}
}
func (c *Client) Name() string {
	return c.name
}

func (c *Client) HandlePrescription(prescription Prescription) {
}
