//
// Author: leafsoar
// Date: 2018-06-12 18:15:46
//

package paylib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"strings"

	"encoding/json"

	"qianuuu.com/kuaigou/domain"
	"qianuuu.com/lib/echotools"
	"qianuuu.com/lib/logs"
	"qianuuu.com/lib/values"
)

type PayHandler interface {
	// 创建订单
	CreateUnifiedOrder(paytype, tradeno string, fee int, ip string, valueMap values.ValueMap) (values.ValueMap, error)
	CheckOrder(order *domain.Orders) (*CheckResult, error)
	NotifyOrder(t *echotools.EchoTools) (*CheckResult, error)
}

type PayItem struct {
	AppName string

	PayName    string
	IsOpen     bool
	AppID      string
	Subject    string
	AppKey     string
	Sign       string
	NotifyURL  string
	BackURL    string
	MchID      string
	DeviceInfo string
	PayTypes   []string
}

func (pi *PayItem) GetRootURL() string {
	return Opts().RootURL
}

func (pi *PayItem) AppIDI() int {
	ret, _ := strconv.Atoi(pi.AppID)
	return ret
}

func (*PayItem) GetMD5(context []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(context)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func (pi *PayItem) GetRequestValues(url string, params values.ValueMap) (values.ValueMap, error) {
	body, err := pi.GetRequest(url, params)
	if err != nil {
		return nil, err
	}
	bvm, err := values.NewValuesFromJSON([]byte(body))
	if err != nil {
		return nil, err
	}
	return bvm, nil
}

// CanPayType 是否启用相应的支付方式
func (pi *PayItem) CanPayType(pt string) bool {
	for _, item := range pi.PayTypes {
		if item == pt {
			return true
		}
	}
	return false
}

func (*PayItem) GetRequestURL(url string, params values.ValueMap) string {
	if len(params) == 0 {
		return url
	}
	ps := ""
	i := 0
	// logs.Info("params: %v: %v", len(params), params)
	sorted_keys := make([]string, 0)
	for k, _ := range params {
		sorted_keys = append(sorted_keys, k)
	}
	// logs.Info("sorted: %v: %v", len(sorted_keys), sorted_keys)
	sort.Strings(sorted_keys)
	for _, k := range sorted_keys {
		i++
		ps = ps + fmt.Sprintf("%s=%v", k, params[k])
		if i != len(params) {
			ps += "&"
		}
	}
	return url + "?" + ps
}

// GetRequest Get 请求
func (pi *PayItem) GetRequest(url string, params values.ValueMap) (string, error) {

	eurl := pi.GetRequestURL(url, params)
	//logs.Info("url: %s", eurl)
	resp, err := http.Get(eurl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (pi *PayItem) PostRequest(url string, params map[string]interface{}) (string, error) {
	content := pi.GetBodyByMap(params)
	logs.Info("--->>>%s", content)
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(string(content)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	//d_v, err := iconv.ConvertString(string(body), "GB2312", "utf-8")
	return string(body), nil
}
func (pi *PayItem) PostRequestJson(url string, params map[string]interface{}) (string, error) {
	//content := pi.GetBodyByMap(params)
	content, err := json.Marshal(params)
	//logs.Info("--->>>%s", content)
	resp, err := http.Post(url, "application/json",
		strings.NewReader(string(content)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}
	//d_v, err := iconv.ConvertString(string(body), "GB2312", "utf-8")
	return string(body), nil
}
func (pi *PayItem) PostRequestString(url, content string) (string, error) {
	defer logs.Info("post request end: %s", url)
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(content))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (pi *PayItem) PostRequestValues(url string, params values.ValueMap) (values.ValueMap, error) {
	body, err := pi.PostRequest(url, params)
	if err != nil {
		return nil, err
	}
	bvm, err := values.NewValuesFromJSON([]byte(body))
	if err != nil {
		return nil, err
	}
	return bvm, nil
}
func (pi *PayItem) PostRequestValuesJson(url string, params values.ValueMap) (values.ValueMap, error) {
	body, err := pi.PostRequestJson(url, params)
	if err != nil {
		return nil, err
	}
	logs.Info("pre body: %v, err: %v", body, err)
	bvm, err := values.NewValuesFromJSON([]byte(body))
	if err != nil {
		return nil, err
	}
	return bvm, nil
}
func (pi *PayItem) GetBodyByMap(params values.ValueMap) string {
	ps := ""
	i := 0
	logs.Info("params: %v: %v", len(params), params)

	sorted_keys := make([]string, 0)
	for k, _ := range params {
		sorted_keys = append(sorted_keys, k)
	}

	logs.Info("sorted: %v: %v", len(sorted_keys), sorted_keys)

	sort.Strings(sorted_keys)

	for _, k := range sorted_keys {
		i++
		ps = ps + fmt.Sprintf("%s=%v", k, params[k])
		if i != len(params) {
			ps += "&"
		}
	}
	return ps
}
func (pi *PayItem) GetBodyByString(params string) values.ValueMap {
	s := strings.Split(params, "&")
	ss := map[string]interface{}{}
	for i := 0; i < len(s); i++ {
		temp := strings.Split(s[i], "=")
		ss[temp[0]] = temp[1]
	}
	return ss
}
func (pi *PayItem) GenAlipaySignString(mapBody map[string]interface{}) (string, error) {
	sorted_keys := make([]string, 0)
	for k, _ := range mapBody {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	index := 0
	for _, k := range sorted_keys {
		//if k == "fund_bill_list" {
		//	billList, _ := json.Marshal(mapBody[k])
		//	mapBody[k] = string(billList)
		//}
		value := fmt.Sprintf("%v", mapBody[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value
		}
		//最后一项后面不要&
		if index < len(sorted_keys)-1 {
			signStrings = signStrings + "&"
		}
		index++
	}
	fmt.Println("生成的待签名---->", signStrings)
	return signStrings, nil
}
