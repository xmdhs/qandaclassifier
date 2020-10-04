package spider

import "encoding/json"

type Fidinfo struct {
	Variables Variables
}

type Variables struct {
	ForumThreadlist []tidinfo `json:"forum_threadlist"`
}

type tidinfo struct {
	Tid string `json:"tid"`
}

func jsondecoder(jsondata []byte) (Fidinfo, error) {
	fidinfo := Fidinfo{}
	err := json.Unmarshal(jsondata, &fidinfo)
	if err != nil {
		return Fidinfo{}, err
	}
	return fidinfo, nil
}
