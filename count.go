package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Count 面单库存查询
func (c *Client) Count() (r *CountResponse, err error) {
	const url = `interface_txm_remain_num.php`

	d := NewRequestBody(c.PartnerID, c.Password, RequestData, nil)

	resp, err := c.cli.Clone().
		SetURLByStr(url).
		SetBody(strings.NewReader(d.Encode())).
		Do()

	if err != nil {
		return nil, err
	}

	response := &CountResponsesBody{}

	err = xml.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}

	if response.Response.Status != "1" {
		return nil, fmt.Errorf("%+v", response.Response)
	}
	return &response.Response.CountResponse, nil
}

type CountResponse struct {
	TotalNum  string `xml:"total_num" json:"total_num" bson:"total_num"`
	UsedNum   string `xml:"used_num" json:"used_num" bson:"used_num"`
	RemainNum string `xml:"remain_num" json:"remain_num" bson:"remain_num"`
}

type CountResponsesBody struct {
	XMLName  xml.Name          `xml:"responses" json:"responses" bson:"responses"`
	Response CountResponseBody `xml:"response" json:"response" bson:"response"`
}

type CountResponseBody struct {
	Status        string `xml:"status" json:"status" bson:"status"`
	Msg           string `xml:"msg" json:"msg" bson:"msg"`
	CountResponse `xml:"innerxml" json:"innerxml" bson:"innerxml"`
}
