package easytools

import "encoding/json"

// CopyProperties 结构体拷贝
// 使用 json 相互转换结构体
// 示例:
//
//	type A struct {
//	  	ID        int64     `json:"id"`
//		Name      string    `json:"name"`
//		CreatedAt time.Time `json:"created_at"`
//		UpdatedAt time.Time `json:"updated_at"`
//	}
//
//	type B struct {
//	  	ID        int64     `json:"id"`
//		Name      string    `json:"name"`
//	}
//
// CopyProperties(A{}, &B{})
func CopyProperties(source interface{}, target interface{}) error {
	bytes, err := json.Marshal(source)

	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, target)

	if err != nil {
		return err
	}

	return nil
}
