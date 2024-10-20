package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func init() {
	log.Println("==============gin package init")
}

// gin web start
func Start() {
	// 1.创建路由
	r := gin.Default()
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	handler(r)
	log.Printf("gin web start on: %d\n", 8080)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run("127.0.0.1:8080")

}

// hander
func handler(r *gin.Engine) {
	handlerBaseDemo(r)
	handlerUploadx(r)
	handlerGroup(r)
}

// http://127.0.0.1:8080/v1/login?name=rain
// http://127.0.0.1:8080/1/submit?name=rain
func handlerGroup(r *gin.Engine) {
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/login", login)
		v1.GET("submit", submit)
	}
	v2 := r.Group("/v2")
	{
		//POST MODE
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func handlerUploadx(r *gin.Engine) {
	r.POST("/uploadx", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
}

func handlerBaseDemo(r *gin.Engine) {
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	//API参数
	//http://127.0.0.1:8080/user/rain/run
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action)
	})
	//URL参数
	//http: //127.0.0.1:8080/user?name=rain
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	//表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./video/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
}
