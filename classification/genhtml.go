package classification

import (
	"html/template"
	"log"
	"net/http"
)

func genhtml(w http.ResponseWriter, data pidinfo) {
	g := genhtmldata{
		Text:    data.message,
		Yuanban: `./fenlei?=Yuanban&pid=` + data.pid,
		Lianji:  `./fenlei?=Lianji&pid=` + data.pid,
		Mod:     `./fenlei?=Mod&pid=` + data.pid,
		Be:      `./fenlei?=Be&pid=` + data.pid,
		Water:   `./fenlei?=Water&pid=` + data.pid,
		Fuifu:   `./fenlei?=Fuifu&pid=` + data.pid,
		Tiaoguo: `./fenlei?=Tiaoguo&pid=` + data.pid,
	}
	err := gentemplate.ExecuteTemplate(w, "pinfenpage", g)
	if err != nil {
		log.Println(err)
	}
}

var gentemplate *template.Template

func init() {
	var err error
	gentemplate, err = template.New("pinfenpage").Parse(htmll)
	if err != nil {
		panic(err)
	}
}

type genhtmldata struct {
	Text    string
	Yuanban string
	Lianji  string
	Mod     string
	Be      string
	Water   string
	Fuifu   string
	Tiaoguo string
}
