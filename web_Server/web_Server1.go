package web_Server

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoginContoor struct {
}

func Http_Server_Init() {
	http.HandleFunc("/login/", F1) //目前不会直接从文件路径中链接html文件。
	err := http.ListenAndServe("127.0.0.1", nil)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
}
func (l *LoginContoor) F1(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/login.html")
	if err != nil {
		fmt.Println("open html file failed err :", err)
	}
	t.Execute(w, nil)
}
