package yunda

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Count 面单库存查询
func (c *Client) Count() (r *CountResponse, err error) {
	const url = `interface_txm_remain_num.php`

	d := NewRequestBody(c.PartnerID, c.Password, nil)

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
	TotalNum  string `xml:"total_num"`
	UsedNum   string `xml:"used_num"`
	RemainNum string `xml:"remain_num"`
}

type CountResponsesBody struct {
	XMLName  xml.Name          `xml:"responses"`
	Response CountResponseBody `xml:"response"`
}

type CountResponseBody struct {
	Status        string `xml:"status"`
	Msg           string `xml:"msg"`
	CountResponse `xml:"innerxml"`
}
