package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	err := http.ListenAndServe("127.0.0.0:8080", router)
	if err != nil {
		fmt.Println("err:", err)
	}
}

//mux已经是时代的眼泪了要说封装不如直接用gin，要说原生现在go 1.22+可以支持通配符路由和方法识别
//但如果要维护一些老地项目mux还是有点用处的
//mux虽已被时代抛弃但在gorilla中的同胞兄弟还是要学滴
