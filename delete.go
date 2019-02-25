package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Delete 订单取消接口 【接口功能描述】
// 下单后如果订单取消，请用此接口
// 取消订单后，如果有运单号，则进入冻结，一个月后释放到库存，可重新使用。错
// 误信息类型有:
// 1. 取消订单后订单要重新生效，则使用下单接口下单即可
func (c *Client) Delete(order *DeleteOrder) (err error) {
	const url = `interface_cancel_order.php`

	data, err := xml.Marshal(DeleteOrders{Order: *order})
	if err != nil {
		return err
	}
	d := NewRequestBody(c.PartnerID, c.Password, data)

	resp, err := c.cli.Clone().
		SetURLByStr(url).
		SetBody(strings.NewReader(d.Encode())).
		Do()

	if err != nil {
		return err
	}

	response := &DeleteResponsesBody{}

	err = xml.Unmarshal(resp.Body(), &response)
	if err != nil {
		return err
	}

	if response.Response.Status != "1" {
		return fmt.Errorf("%+v", response.Response)
	}
	return nil
}

type DeleteOrder struct {
	// 订单唯一序列号(可以是订单号，可 以是序列号，必须保证唯一)
	OrderSerialNo string `xml:"order_serial_no"`

	// 运单号，可选填
	Mailno string `xml:"mailno,omitempty"`
}

type DeleteOrders struct {
	XMLName xml.Name    `xml:"orders"`
	Order   DeleteOrder `xml:"order"`
}

type DeleteResponsesBody struct {
	XMLName  xml.Name           `xml:"responses"`
	Response DeleteResponseBody `xml:"response"`
}

type DeleteResponseBody struct {
	Status         string `xml:"status"`
	Msg            string `xml:"msg"`
	DeleteResponse `xml:"innerxml"`
}

type DeleteResponse struct {
	// 订单唯一序列号(可以是订单号，可 以是序列号，必须保证唯一)
	OrderSerialNo string `xml:"order_serial_no"`

	// 运单号，可选填
	Mailno string `xml:"mailno"`
}
