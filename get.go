package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Get 订单信息查询接口
// 如果下单的时候密文/运单号没有接收成功，则可以用此接口查询
// 按照运单号或者订单序列号查询订单的状态，运单号，打印密文或者 json 数据
// 如果需要接口返回 pdf 信息则需要传递参数 print_file
// 如果需要返回打印密文信息或者 json 明文，则传递参数 json_data,调用本地服务
func (c *Client) Get(order *GetOrder) (err error) {
	const url = `interface_order_info.php`

	data, err := xml.Marshal(GetOrders{Order: *order})
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

	response := &GetResponsesBody{}

	err = xml.Unmarshal(resp.Body(), &response)
	if err != nil {
		return err
	}

	if response.Response.Status != "1" {
		return fmt.Errorf("%+v", response.Response)
	}
	return nil
}

type GetOrder struct {
	// 订单唯一序列号(可以是订单号，可 以是序列号，必须保证唯一)
	OrderSerialNo string `xml:"order_serial_no" json:"order_serial_no" bson:"order_serial_no"`

	// 运单号，可选填
	Mailno string `xml:"mailno,omitempty" json:"mailno,omitempty" bson:"mailno,omitempty"`

	// 是否需要 pdf 文件:1 需要，0 不需要
	PrintFile string `xml:"print_file,omitempty" json:"print_file,omitempty" bson:"print_file,omitempty"`

	// 是否需要打印文件的 json 数据，用于 直接生成打印图片:1 需要，0 不需要
	JSONData string `xml:"json_data,omitempty" json:"json_data,omitempty" bson:"json_data,omitempty"`

	// 强制加密参数，1 加密
	JSONEncrypt string `xml:"json_encrypt,omitempty" json:"json_encrypt,omitempty" bson:"json_encrypt,omitempty"`
}

type GetOrders struct {
	XMLName xml.Name `xml:"orders" json:"orders" bson:"orders"`
	Order   GetOrder `xml:"order" json:"order" bson:"order"`
}

type GetResponsesBody struct {
	XMLName  xml.Name        `xml:"responses" json:"responses" bson:"responses"`
	Response GetResponseBody `xml:"response" json:"response" bson:"response"`
}

type GetResponseBody struct {
	Status      string `xml:"status" json:"status" bson:"status"`
	Msg         string `xml:"msg" json:"msg" bson:"msg"`
	GetResponse `xml:"innerxml" json:"innerxml" bson:"innerxml"`
}

type GetResponse struct {
	// 订单唯一序列号(可以是订单号，可 以是序列号，必须保证唯一)
	OrderSerialNo string `xml:"order_serial_no" json:"order_serial_no" bson:"order_serial_no"`

	// 运单号，可选填
	Mailno string `xml:"mailno" json:"mailno" bson:"mailno"`
}
