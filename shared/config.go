package shared



type Config interface {
	GetValueForKey(key string)string
	SetValueForKey(value, key string)
}

func NewConfigValue(initialMap map[string]string)*ConfigValue {
	return &ConfigValue{_map:initialMap}
}

type ConfigValue struct {
	_map map[string]string
}

func (self *ConfigValue) GetValueForKey(key string) string {
	return self._map[key]
}

func (self *ConfigValue) SetValueForKey(value, key string) {
	self._map[key] = value
}










