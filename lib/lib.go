package lib

type commandable struct{}

func (c *commandable) EmbeddedMethod() {}

type Client struct {
	commandable
}

func (c *Client) DirectMethod() {}
