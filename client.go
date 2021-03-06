package yunda

import (
	"github.com/wzshiming/requests"
)

const (
	TestHost = `http://orderdev.yundasys.com:10110/cus_order/order_interface/`
	Host     = `http://order.yundasys.com:10235/cus_order/order_interface/`
)

// Client 韵达客户端
type Client struct {
	PartnerID string
	Password  string
	Host      string
	cli       *requests.Request
}

// NewClient 创建一个新的韵达客户端
func NewClient(host, partnerID, password string) *Client {
	c := &Client{
		PartnerID: partnerID,
		Password:  password,
		Host:      host,
	}
	c.cli = requests.NewClient().
		SetLogLevel(requests.LogError).
		NewRequest().
		AddHeaderIfNot(requests.HeaderContentType, requests.MimeURLEncoded).
		SetMethod(requests.MethodPost).
		SetURLByStr(c.Host)
	return c
}
