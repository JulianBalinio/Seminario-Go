package client

type Client struct{
	id int
}

func (c *Client) SetId(n int) {
	c.id = n
}

func (c *Client) ID() int {
	return c.id
}