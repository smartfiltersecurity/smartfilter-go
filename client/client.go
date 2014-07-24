package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SmartFilterClient struct {
	apiKey string
	base   string
}

func NewSmartFilterClient(apiKey string) *SmartFilterClient {
	smartFilter := new(SmartFilterClient)
	smartFilter.apiKey = apiKey
	smartFilter.base = "http://api.prevoty.com/1"
	return smartFilter
}

func (self *SmartFilterClient) Verify() (bool, error) {
	verifyUrl := fmt.Sprintf("%s/key/verify?api_key=%s", self.base, self.apiKey)
	response, err := http.Get(verifyUrl)
	defer response.Body.Close()
	if err == nil {
		switch response.StatusCode {
		case 200:
			return true, nil
		case 400:
			return false, &SmartFilterBadInputParameter{}
		case 403:
			return false, &SmartFilterBadAPIKey{}
		case 500:
			return false, &SmartFilterInternalError{}
		}
	}
	return false, err
}

func (self *SmartFilterClient) Info() (*SmartFilterInformation, error) {
	infoUrl := fmt.Sprintf("%s/key/info?api_key=%s", self.base, self.apiKey)
	response, err := http.Get(infoUrl)
	defer response.Body.Close()
	information := new(SmartFilterInformation)
	if err == nil {
		switch response.StatusCode {
		case 200:
			decodingErr := json.NewDecoder(response.Body).Decode(information)
			if decodingErr != nil {
				return information, decodingErr
			}
			return information, nil
		case 400:
			return information, &SmartFilterBadInputParameter{}
		case 403:
			return information, &SmartFilterBadAPIKey{}
		case 500:
			return information, &SmartFilterInternalError{}
		}
	}
	return information, err
}

func (self *SmartFilterClient) VerifyRule(ruleKey string) (bool, error) {
	verifyUrl := fmt.Sprintf("%s/rule/verify?api_key=%s&rule_key=%s", self.base, self.apiKey, ruleKey)
	response, err := http.Get(verifyUrl)
	defer response.Body.Close()
	if err == nil {
		switch response.StatusCode {
		case 200:
			return true, nil
		case 400:
			return false, &SmartFilterBadInputParameter{}
		case 403:
			return false, &SmartFilterBadAPIKey{}
		case 500:
			return false, &SmartFilterInternalError{}
		}
	}
	return false, err
}

func (self *SmartFilterClient) Filter(input string, ruleKey string) (*SmartFilterResult, error) {
	filterUrl := fmt.Sprintf("%s/xss/filter", self.base)
	response, err := http.PostForm(filterUrl, url.Values{"api_key": {self.apiKey}, "rule_key": {ruleKey}, "input": {input}})
	defer response.Body.Close()
	result := new(SmartFilterResult)
	if err == nil {
		switch response.StatusCode {
		case 200:
			decodingErr := json.NewDecoder(response.Body).Decode(result)
			if decodingErr != nil {
				return result, decodingErr
			}
			return result, nil
		case 400:
			return result, &SmartFilterBadInputParameter{}
		case 403:
			return result, &SmartFilterBadAPIKey{}
		case 413:
			return result, &SmartFilterRequestTooLarge{}
		case 500:
			return result, &SmartFilterInternalError{}
		case 507:
			return result, &SmartFilterAccountQuotaExceeded{}
		}
	}
	return result, err
}
