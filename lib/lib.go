package lib

type baseClient struct{}

func (b *baseClient) Foo() error {
	return nil
}

type commandable struct{}

func (c *commandable) Del(key string) error {
	return nil
}

type Client struct {
	baseClient
	commandable
}

func (c *Client) Bar() error {
	return nil
}

type AnotherCmd struct {
	commandable
}
