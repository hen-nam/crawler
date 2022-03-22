package model

import "encoding/json"

// Profile 用户
type Profile struct {
	Name          string
	Gender        string
	Age           int
	Height        int
	Weight        int
	Income        string
	Marriage      string
	Education     string
	Occupation    string
	Location      string
	Constellation string
	House         string
	Car           string
}

// FromJsonObject 将 JSON 对象转换成用户
func FromJsonObject(object interface{}) (Profile, error) {
	profile := Profile{}

	bytes, err := json.Marshal(object)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(bytes, &profile)
	return profile, err
}
