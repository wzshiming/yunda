package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Create 订单创建接口
// 接口流程见方案图
// 下单后，接口返回:
// 是否送达
// 运单号(送单时才有)
// 打印文件(1. json 格式，或者 2.密文:使用韵达提供的本地服务可直接生成打印图片)
// 说明:返回打印文件格式可以在韵达系统配置客户属性，json 格式可以支持可以自己生成打印文件
func (c *Client) Create(order *CreateOrder) (r *CreateResponse, err error) {
	const url = `interface_receive_order__mailno.php`

	data, err := xml.Marshal(CreateOrders{Order: *order})
	if err != nil {
		return nil, err
	}
	d := NewRequestBody(c.PartnerID, c.Password, RequestData, data)

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

type CreateItem struct {
	// 商品名称(cod 订单必填)
	Name string `xml:"name" json:"name" bson:"name"`
	// 数量
	Number string `xml:"number" json:"number" bson:"number"`
	// 备注
	Remark string `xml:"remark" json:"remark" bson:"remark"`
}
type CreateItems struct {
	XMLName xml.Name   `xml:"items" json:"items" bson:"items"`
	Item    CreateItem `xml:"item" json:"item" bson:"item"`
}
type CreateOrder struct {

	// 订单唯一序列号(可以是订单号，可 以是序列号，必须保证唯一)
	OrderSerialNo string `xml:"order_serial_no" json:"order_serial_no" bson:"order_serial_no"`

	// 代收货款金额(目前仅用于 cod 订单， cod 订单必填)
	CollectionValue string `xml:"collection_value" json:"collection_value" bson:"collection_value"`

	// 可以自定义显示信息 1，打印在客户自 定义区域 1，换行请用\n
	CusArea1 string `xml:"cus_area1,omitempty" json:"cus_area1,omitempty" bson:"cus_area1,omitempty"`

	// 可以自定义显示信息 2，打印在客户自 定义区域 2，换行请用\n
	CusArea2 string `xml:"cus_area2,omitempty" json:"cus_area2,omitempty" bson:"cus_area2,omitempty"`

	// 可以自定义显示信息 3，打印在客户自 定义区域 3，换行请用\n
	CusArea3 string `xml:"cus_area3,omitempty" json:"cus_area3,omitempty" bson:"cus_area3,omitempty"`

	// 商品种类
	Items CreateItems `xml:"items" json:"items" bson:"items"`

	// 大客户系统订单的订单号
	Khddh string `xml:"khddh" json:"khddh" bson:"khddh"`

	// 内部参考号，供大客户自己使用，可 以是客户的客户编号
	Nbckh string `xml:"nbckh,omitempty" json:"nbckh,omitempty" bson:"nbckh,omitempty"`

	OrderType OrderType `xml:"order_type" json:"order_type" bson:"order_type"`

	// 接口异步回传的时候返回的 ID，客户 方系统使用，此 ID 是客户需求，可以 不填，默认为空
	CallbackID string `xml:"callback_id,omitempty" json:"callback_id,omitempty" bson:"callback_id,omitempty"`

	// 客户波次号，按照此号进行批量打印 校验
	WaveNo string `xml:"wave_no,omitempty" json:"wave_no,omitempty" bson:"wave_no,omitempty"`

	// 人工强制下单(系统筛单不送达的情 况下,会验证此标签,为 1 的话会强制 标记为可送达，使用此参数需要和网 点协商，否则不要使用此参数)
	ReceiverForce string `xml:"receiver_force,omitempty" json:"receiver_force,omitempty" bson:"receiver_force,omitempty"`

	// 送货地
	Receiver CreateLocale `xml:"receiver" json:"receiver" bson:"receiver"`
	// 收货地
	Sender CreateLocale `xml:"sender" json:"sender" bson:"sender"`

	// 商品性质(保留字段，暂时不用)
	Special string `xml:"special,omitempty" json:"special,omitempty" bson:"special,omitempty"`

	// 物品金额
	Value string `xml:"value,omitempty" json:"value,omitempty" bson:"value,omitempty"`

	// 物品重量
	Weight string `xml:"weight,omitempty" json:"weight,omitempty" bson:"weight,omitempty"`

	// 备注
	Remark string `xml:"remark,omitempty" json:"remark,omitempty" bson:"remark,omitempty"`
}
type CreateOrders struct {
	XMLName xml.Name    `xml:"orders" json:"orders" bson:"orders"`
	Order   CreateOrder `xml:"order" json:"order" bson:"order"`
}
type CreateLocale struct {
	// 需要将省市区划信息加上，例如:上 海市,上海市,青浦区盈港东路 7766 号
	Address string `xml:"address" json:"address" bson:"address"`
	Branch  string `xml:"branch" json:"branch" bson:"branch"`
	// 严格按照国家行政区划，省市区三级， 逗号分隔。示例上海市,上海市,青浦区 (cod 订单必填)
	City string `xml:"city" json:"city" bson:"city"`
	// 公司名
	Company string `xml:"company" json:"company" bson:"company"`
	// 移动电话
	Mobile string `xml:"mobile" json:"mobile" bson:"mobile"`
	// 姓名
	Name string `xml:"name" json:"name" bson:"name"`
	// 固定电话
	Phone string `xml:"phone" json:"phone" bson:"phone"`
	// 邮编
	Postcode string `xml:"postcode" json:"postcode" bson:"postcode"`
}

type CreateResponse struct {
	OrderSerialNo string `xml:"order_serial_no" json:"order_serial_no" bson:"order_serial_no"`
	MailNo        string `xml:"mail_no" json:"mail_no" bson:"mail_no"`
	PdfInfo       string `xml:"pdf_info" json:"pdf_info" bson:"pdf_info"`
}

type CreateResponsesBody struct {
	XMLName  xml.Name           `xml:"responses" json:"responses" bson:"responses"`
	Response CreateResponseBody `xml:"response" json:"response" bson:"response"`
}

type CreateResponseBody struct {
	CreateResponse `xml:"innerxml" json:"innerxml" bson:"innerxml"`
	Status         string `xml:"status" json:"status" bson:"status"`
	Msg            string `xml:"msg" json:"msg" bson:"msg"`
}
