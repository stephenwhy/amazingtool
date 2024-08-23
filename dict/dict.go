package dict

import "context"

type RegisterFunc func(ctx context.Context) map[string]any

var RegisterFuncList []RegisterFunc

func Dict(ctx context.Context) map[string]any {
	result := make(map[string]any)
	for _, registerFunc := range RegisterFuncList {
		kvs := registerFunc(ctx)
		for k, v := range kvs {
			result[k] = v
		}
	}
	return result
}

type Enum[T comparable] struct {
	Key  T
	Desc string
}

func (e Enum[T]) Equal(key T) bool {
	return e.Key == key
}

type EnumList[T comparable] []Enum[T]

func (e EnumList[T]) Search(key T) (string, bool) {
	for _, item := range e {
		if item.Key == key {
			return item.Desc, true
		}
	}
	return "", false
}

func (e EnumList[T]) ToMap() map[T]string {
	result := make(map[T]string, len(e))
	for _, item := range e {
		result[item.Key] = item.Desc
	}
	return result
}

func (e EnumList[T]) ToKvList() []map[string]any {
	result := make([]map[string]any, 0, len(e))
	for _, item := range e {
		result = append(result, map[string]any{
			"key":  item.Key,
			"desc": item.Desc,
		})
	}
	return result
}
