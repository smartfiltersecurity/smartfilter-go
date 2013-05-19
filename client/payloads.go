package client

type SmartFilterVerification struct {
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}

type SmartFilterInformation struct {
	Maximum   int    `json:"maximum"`
	Used      int    `json:"used"`
	Remaining int    `json:"remaining"`
	Message   string `json:"message"`
}

type SmartFilterRuleVerification struct {
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}

type SmartFilterResult struct {
	Message    string                 `json:"message"`
	Output     string                 `json:"output"`
	Statistics *SmartFilterStatistics `json:"statistics"`
}

type SmartFilterStatistics struct {
	InvalidAttributes    int `json:"invalid_attributes"`
	InvalidProtocols     int `json:"invalid_protocols"`
	InvalidTags          int `json:"invalid_tags"`
	JavaScriptAttributes int `json:"javascript_attributes"`
	JavaScriptProtocols  int `json:"javascript_protocols"`
	JavaScriptTags       int `json:"javascript_tags"`
	TagsBalanced         int `json:"tags_balanced"`
	Transformations      int `json:"transformations"`
}
