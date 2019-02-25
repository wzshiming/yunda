package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Update 订单更新接口 【接口功能描述】
// 下单后如果订单信息需要更新，请用此接口
// 更新订单后会将运单号，打印信息(密文或者明文 json)返回，错误信息
// 错误信息类型有:
// 1. 超过下单时间 15 天，不允许更新
// 2. 已经有物流信息，不允许更新
func (c *Client) Update(order *CreateOrder) (r *CreateResponse, err error) {
	const url = `interface_modify_order.php`

	data, err := xml.Marshal(CreateOrders{Order: *order})
	if err != nil {
		return nil, err
	}
	d := NewRequestBody(c.PartnerID, c.Password, data)

	resp, err := c.cli.Clone().
		SetURLByStr(url).
		SetBody(strings.NewReader(d.Encode())).
		Do()

	if err != nil {
		return nil, err
	}

	response := &CreateResponsesBody{}

	err = xml.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}

	if response.Response.Status != "1" {
		return nil, fmt.Errorf("%+v", response.Response)
	}
	return &response.Response.CreateResponse, nil
}
