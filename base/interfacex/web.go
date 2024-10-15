package interfacex

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Go语言里面提供了一个完善的 net/http 包，通过 net/http 包我们可以很方便的搭建一个可以运行的 Web 服务器。
// 同时使用 net/http 包能很简单地对 Web 的路由，静态文件，模版，cookie 等数据进行设置和操作。
// http://localhost:8000/index
func WebTest() {
	http.HandleFunc("/", home) // index 为向 url发送请求时，调用的函数
	// 在/后面加上 index ，来指定访问路径
	http.HandleFunc("/index", index)
	log.Println("WebTest ListenAndServe localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go demo home")
}

func index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("./interfacex/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		// log.Fatal(err)
	} else {
		w.Write(content)
	}
}
