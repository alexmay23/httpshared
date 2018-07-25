package shared



type Config interface {
	GetSharedSecret()string
}

func NewConfigValue(secret string)*ConfigValue {
	return &ConfigValue{secret:secret}
}

type ConfigValue struct {
	secret string
}

func (self *ConfigValue) GetSharedSecret() string {
	return self.secret
}







