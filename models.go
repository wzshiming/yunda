package yunda

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"net/url"
)

// RequestBody 主要请求封装
type RequestBody struct {
	PartnerID  string  // 合作商 ID，如:partnerid = YUNDA(合作商 ID 待正式上线前更改 为韵达指定的 ID)
	Request    Request // 数据请求类型，如 request=data;其中 data 表示下单
	Version    string  // 版本
	XMLdata    string  // XML 数据内容
	Validation string  // 效验码
}

// NewRequestBody 创建新的请求
func NewRequestBody(partnerID, password string, request Request, data []byte) *RequestBody {
	r := &RequestBody{
		PartnerID: partnerID,
		Request:   request,
		Version:   "1.0",
		XMLdata:   base64.StdEncoding.EncodeToString(data),
	}
	r.validation(password)
	return r
}

func (r *RequestBody) validation(password string) {
	sum := r.XMLdata + r.PartnerID + password
	hash := md5.Sum([]byte(sum))
	r.Validation = hex.EncodeToString(hash[:])
}

// Encode 编码
func (r *RequestBody) Encode() string {
	return `partnerid=` + r.PartnerID + `&version=` + r.Version + `&request=` + string(r.Request) + `&xmldata=` + url.QueryEscape(r.XMLdata) + `&validation=` + r.Validation
}
