package component

import "github.com/pkg/errors"

type Client struct{}

func (c *Client) DoIt() error {
	return errors.Wrap(errors.New("main error"), "while doing it")
}
