package model

// SMSConfig provides a config template to record SMS setting
type SMSConfig struct {
	AliyunCNSign                          string `json:"aliyunCNSign"`
	AliyunGlobeSenderID                   string `json:"aliyunGlobeSenderID"`
	AliyunGlobeBasicVerificationTemplate  string `json:"aliyunGlobeBasicVerificationTemplate"`
	AliyunHKMOBasicVerificationTemplate   string `json:"aliyunHKMOBasicVerificationTemplate"`
	AliyunTWBasicVerificationTemplate     string `json:"aliyunTWBasicVerificationTemplate"`
	AliyunJPBasicVerificationTemplate     string `json:"aliyunJPBasicVerificationTemplate"`
	AliyunCNBasicVerificationTemplateCode string `json:"aliyunCNBasicVerificationTemplateCode"`
}

// SystemConfig provides basic system config options
type SystemConfig struct {
	CnidKey string    `json:"cnidKey"`
	SMS     SMSConfig `json:"sms"`
}

var DefaultConfig = SystemConfig{
	CnidKey: "",
	SMS: SMSConfig{
		AliyunCNSign:                          "",
		AliyunGlobeSenderID:                   "",
		AliyunGlobeBasicVerificationTemplate:  "",
		AliyunHKMOBasicVerificationTemplate:   "",
		AliyunTWBasicVerificationTemplate:     "",
		AliyunJPBasicVerificationTemplate:     "",
		AliyunCNBasicVerificationTemplateCode: "",
	},
}
