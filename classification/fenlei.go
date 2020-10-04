package classification

import (
	"net/http"

	"github.com/xmdhs/bayesian-classifier/classifier"
)

func fenlei(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q["fenlei"]) == 0 || len(q["pid"]) == 0 {
		e(w, `缺少必要参数`)
		return
	}
	addtemp.Lock()
	defer addtemp.Unlock()
	setcategory := func(pid, category string) {
		addtemp.temp[pid] = pidtemp{
			text:     addtemp.temp[pid].text,
			category: category,
		}
	}
	pid := q["pid"][0]
	switch q["fenlei"][0] {
	case `Yuanban`:
		setcategory(pid, "原版")
	case `Lianji`:
		setcategory(pid, "联机")
	case `Mod`:
		setcategory(pid, "mod")
	case `Be`:
		setcategory(pid, "基岩")
	case `Water`:
		setcategory(pid, "灌水")
	case `Fuifu`:
		setcategory(pid, "回复")
	case `Tiaoguo`:
	}
	addtemp.i = addtemp.i - 1
	if addtemp.i == 0 {
		for _, v := range addtemp.temp {
			handler.Training(v.text, v.category)
		}
	}
	http.Redirect(w, r, `/classification/add`, 302)
}

func e(w http.ResponseWriter, err string) {
	http.Error(w, err, 500)
}

var handler = classifier.NewClassifier(map[string]interface{}{
	"defaultProb":   0.5,     // 默认概率
	"defaultWeight": 1.0,     // 默认概率的权重，假定与一个单词相当
	"debug":         false,   // 开启调试
	"http":          false,   // 开启HTTP服务
	"httpPort":      ":8812", // HTTP服务端口
	"storage": map[string]string{
		"adapter":   "file",         // 存储引擎，接受 file,redis，目前只支持file
		"path":      "storage.data", // 文件存储引擎的存储路径
		"frequency": "20",           // 自动存储的频率, 单位: 秒，0 表示不自动存储
	},
})
