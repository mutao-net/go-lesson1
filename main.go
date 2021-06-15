package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

/*
* パラメータ
 */
type Values struct {
	Value1 int
	Value2 int
	Result int
	Value3 string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("./template/index.html.tpl"))
	// 書き方汚いけどあとで直す。
	values := Values{
		Value1: 0,
		Value2: 0,
		Result: 0,
	}
	// error handlingもしっかり実装すること
	if len(r.PostFormValue("param1")) > 0 {
		values.Value1, _ = strconv.Atoi(r.PostFormValue("param1"))
	}
	if len(r.PostFormValue("param2")) > 0 {
		values.Value2, _ = strconv.Atoi(r.PostFormValue("param2"))
	}
	if len(r.PostFormValue("param3")) > 0 {
		values.Value3 = r.PostFormValue("param3")
	}

	values.Result = calc(values.Value1, values.Value2, values.Value3)
	if err := template.ExecuteTemplate(w, "index.html.tpl", values); err != nil {
		log.Fatal(err)
	}
}

// {{ if eq .Type 1 }} selected {{ end }}
func calc(val1, val2 int, val3 string) int {
	switch val3 {
	case "1":
		return val1 + val2
	case "2":
		return val1 - val2
	case "3":
		return val1 * val2
	case "4":
		return val1 / val2
	}
	return 0
}
func main() {
	http.HandleFunc("/index/", indexHandler)
	http.ListenAndServe(":8081", nil)
}
