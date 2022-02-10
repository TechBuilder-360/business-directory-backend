package models

import (
	"github.com/TechBuilder-360/business-directory-backend.git/utility"
	"time"
)

type Client struct {
	ClientID    string    `json:"ClientID"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Secret      string    `json:"Secret"`
	AESKey      string    `json:"AESKey"`
	NeedAES     bool      `json:"NeedAES"`
	Created     time.Time `json:"Created"`
}

func (c *Client) ValidateClient(cs string, body string) bool {
	hash:=utility.ComputeHmac256(body, c.Secret)
	return hash==cs
}