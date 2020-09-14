package keyvalueutil

type KeyValue struct {
	Key   string
	Value string
}

func GetKeyValue(keyValues []KeyValue, key string) KeyValue {
	for _, item := range keyValues {
		if item.Key == key {
			return item
		}
	}
	return KeyValue{}
}

func SetKeyValue(ptKeyValues *[]KeyValue, keyValue KeyValue) {
	for idx, item := range *ptKeyValues {
		if item.Key == keyValue.Key {
			(*ptKeyValues)[idx] = keyValue
			break
		}
	}
}
