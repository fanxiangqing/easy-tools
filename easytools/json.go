package easytools

import (
	"encoding/json"
	"log"
	"time"
)

type JSONObject map[string]interface{}

// Get 获取 interface 类型 value
func (o *JSONObject) Get(key string) interface{} {
	return (*o)[key]
}

// Put 添加 value
func (o *JSONObject) Put(key string, value interface{}) {
	(*o)[key] = value
}

// Remove 删除指定 key 的 value
func (o *JSONObject) Remove(key string) {
	(*o)[key] = nil
}

// Clear 清空所有 key
func (o *JSONObject) Clear() {
	*o = JSONObject{}
}

// GetString 获取 string 类型 value
func (o *JSONObject) GetString(key string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(string)
}

// GetInt 获取 int 类型 value
func (o *JSONObject) GetInt(key string) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(int)
}

// GetInt64 获取 int64 类型 value
func (o *JSONObject) GetInt64(key string) int64 {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(int64)
}

// GetUint 获取 uint 类型 value
func (o *JSONObject) GetUint(key string) uint {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	return o.Get(key).(uint)
}

// GetUint64 获取 uint64 类型 value
func (o *JSONObject) GetUint64(key string) uint64 {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(uint64)
}

// GetFloat64 获取 float64 类型 value
func (o *JSONObject) GetFloat64(key string) float64 {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(float64)
}

// GetBool 获取 bool 类型 value
func (o *JSONObject) GetBool(key string) bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	return o.Get(key).(bool)
}

// GetTime 获取 time.Time 类型 value
func (o *JSONObject) GetTime(key string) time.Time {
	return o.Get(key).(time.Time)
}

// ToJsonString 转 json 字符串
func ToJsonString(source interface{}) (string, error) {
	bytes, err := json.Marshal(source)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Parse 解析 json 字符串返回 interface{} 类型
func Parse(text string) (interface{}, error) {
	var v interface{}
	err := json.Unmarshal([]byte(text), &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ParseJsonObject 解析 json 字符串返回 JSONObject 类型
func ParseJsonObject(text string) (JSONObject, error) {
	var v JSONObject
	err := json.Unmarshal([]byte(text), &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// ParseObject 解析 json 字符串返回指定类型
func ParseObject[T any](text string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(text), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
