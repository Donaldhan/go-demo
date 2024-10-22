package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	zhTWTranslations "github.com/go-playground/validator/v10/translations/zh_tw"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
	log.Println("==============gin package init")
}

// 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//before1
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		// 执行函数， 执行下个中间件，如果有的情况下
		c.Next()
		//after2
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 定义中间
func myTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		//before2
		start := time.Now()
		c.Next()
		//after2
		// 统计时间
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}
}
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

// 初始化一个cookie存储对象
// something-very-secret应该是一个你自己的密匙，只要不被别人知道就行
var store = sessions.NewCookieStore([]byte("something-very-secret"))

// 1、自定义的校验方法
func nameNotNullAndAdmin(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		// 字段不能为空，并且不等于  admin
		return value != "" && !("admin" == value)
	}
	return true
}

// Booking contains binded and validated data.
type Booking struct {
	//定义一个预约的时间大于今天的时间
	CheckIn time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	//gtfield=CheckIn退出的时间大于预约的时间
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	//field.Interface().(time.Time)获取参数值并且转换为时间格式
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}
	return true
}

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

type User struct {
	Username string `form:"user_name" validate:"required"`
	Tagline  string `form:"tag_line" validate:"required,lt=10"`
	Tagline2 string `form:"tag_line2" validate:"required,gt=1"`
}

// gin web start
func Start() {
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件, 多个中间先执行中间的before1， before2，..., after2, after1
	//r.Use(MiddleWare())
	r.Use(myTime())
	//限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	//gin.SetMode(gin.ReleaseMode)
	//r.LoadHTMLGlob("gin/**/*")
	//r.Static("/assets", "./assets")
	// 3、将我们自定义的校验方法注册到 validator中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 这里的 key 和 fn 可以不一样最终在 struct 使用的是 key
		v.RegisterValidation("NotNullAndAdmin", nameNotNullAndAdmin)
		v.RegisterValidation("bookabledate", bookableDate)
	}
	en := en.New()
	zh := zh.New()
	zh_tw := zh_Hant_TW.New()
	Uni = ut.New(en, zh, zh_tw)
	Validate = validator.New()
	handler(r)
	log.Printf("gin web start on: %d\n", 8080)
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Gin failed to start: %v", err)
	}
}

// handler
func handler(r *gin.Engine) {
	handlerBaseDemo(r)
	handlerUploadx(r)
	handlerGroup(r)
	handlerJsonBind(r)
	handlerFormBind(r)
	handlerUrlBind(r)
	handlerResponseFmt(r)
	handlerResponseHtml(r)
	handlerResponseAsync(r)
	handlerMiddleWare(r)
	handlerMiddleWareV2(r)
	handlerCookieBaseMiddleWare(r)
	handlerSession(r)
	handlerDataVerify(r)
	handlerLocale(r)
}
func handlerLocale(r *gin.Engine) {
	r.GET("/5lmhV4", startPage)
	r.POST("/5lmhV4", startPage)
}

/*
http://localhost:8080/5lmhV4?user_name=枯藤&tag_line=9&tag_line2=33&locale=zh
http://localhost:8080/5lmhV4?user_name=枯藤&tag_line=9&tag_line2=3&locale=en
http://localhost:8080/5lmhV4?user_name=枯藤&tag_line=9&tag_line2=3&locale=zh_tw
*/
func startPage(c *gin.Context) {
	//这部分应放到中间件中
	locale := c.DefaultQuery("locale", "zh")
	trans, _ := Uni.GetTranslator(locale)
	switch locale {
	case "zh":
		zhTranslations.RegisterDefaultTranslations(Validate, trans)
		break
	case "en":
		enTranslations.RegisterDefaultTranslations(Validate, trans)
		break
	case "zh_tw":
		zhTWTranslations.RegisterDefaultTranslations(Validate, trans)
		break
	default:
		zhTranslations.RegisterDefaultTranslations(Validate, trans)
		break
	}

	//自定义错误内容
	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	//这块应该放到公共验证方法中
	user := User{}
	c.ShouldBind(&user)
	fmt.Println(user)
	err := Validate.Struct(user)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		c.String(200, fmt.Sprintf("%#v", sliceErrs))
	}
	c.String(200, fmt.Sprintf("%#v", "user"))
}

/*
对绑定解析到结构体上的参数，自定义验证功能
比如我们要对 name 字段做校验，要不能为空，并且不等于 admin ，类似这种需求，就无法 binding 现成的方法
需要我们自己验证方法才能实现 官网示例（https://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Custom_Functions）
这里需要下载引入下 gopkg.in/go-playground/validator.v8
*/
type PersonV2 struct {
	Age int `form:"age" binding:"required,gt=10"`
	// 2、在参数 binding 上使用自定义的校验方法函数注册时候的名称
	Name    string `form:"name" binding:"NotNullAndAdmin"`
	Address string `form:"address" binding:"required"`
}

// Person ..
type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

/*
http://localhost:8080/5lmh?age=11&name=枯藤&birthday=2006-01-02
*/
func handlerDataVerify(r *gin.Engine) {
	r.GET("/5lmh", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, fmt.Sprint(err))
			return
		}
		c.String(200, fmt.Sprintf("%#v", person))
	})
	/*
	   curl -X GET "http://127.0.0.1:8080/5lmhV2?name=&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/5lmhV2?name=admin&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/5lmhV2?name=lmh&age=12&address=beijing"
	   curl -X GET "http://127.0.0.1:8080/5lmhV2?name=adz&age=12&address=beijing"
	*/
	r.GET("/5lmhV2", func(c *gin.Context) {
		var person PersonV2
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})
	// curl -X GET "http://localhost:8080/5lmhV3?check_in=2024-11-07&check_out=2024-11-20"
	// curl -X GET "http://localhost:8080/5lmhV3?check_in=2024-09-07&check_out=2024-11-20"
	// curl -X GET "http://localhost:8080/5lmhV3?check_in=2024-11-07&check_out=2024-11-01"
	r.GET("/5lmhV3", getBookable)
}
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/*
http://localhost:8080/session/save
http://localhost:8080/session/get
*/
// https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E4%BC%9A%E8%AF%9D%E6%8E%A7%E5%88%B6/Sessions.html
func handlerSession(r *gin.Engine) {
	sessionGroup := r.Group("/session")
	sessionGroup.GET("/save", func(c *gin.Context) {
		SaveSession(c.Writer, c.Request)
	})
	sessionGroup.GET("/get", func(c *gin.Context) {
		GetSession(c.Writer, c.Request)
	})
}
func SaveSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	//　获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("SaveSession:done")
}
func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println("GetSession:", foo)
}

/*
http://localhost:8080/login
http://localhost:8080/home
*/
func handlerCookieBaseMiddleWare(r *gin.Engine) {
	r.GET("/login", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("abc", "123", 60, "/",
			"localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	//检查cookie
	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
}

/*
中间件类似aop
http://localhost:8080/shopping/index
http://localhost:8080/shopping/home
*/
func handlerMiddleWareV2(r *gin.Engine) {
	// {}为了代码规范
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

/*
中间件类似aop
http://localhost:8080/ce
http://localhost:8080/ceV2
*/
func handlerMiddleWare(r *gin.Engine) {
	r.GET("/ce", func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
	//局部中间键使用
	r.GET("/ceV2", MiddleWare(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
}

/*
同步异步处理请求
http://localhost:8080/long_async
http://localhost:8080/long_sync
*/
func handlerResponseAsync(r *gin.Engine) {
	// 1.异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		// 异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + copyContext.Request.URL.Path)
		}()
	})
	// 2.同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步执行：" + c.Request.URL.Path)
	})

}

/*
页面渲染
http://localhost:8080/index
http://localhost:8080/indexV1
*/
func handlerResponseHtml(r *gin.Engine) {
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})
	})
	r.GET("/indexV1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
	})
}

/*
响应数据格式
curl -X GET http://localhost:8080/someJSON
curl -X GET http://localhost:8080/someStruct
curl -X GET http://localhost:8080/someXML
curl -X GET http://localhost:8080/someYAML
curl -X GET http://localhost:8080/someProtoBuf --output output.pb
*/
func handlerResponseFmt(r *gin.Engine) {
	// 1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	// 3.XML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	// 4.YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "zhangsan"})
	})
	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	// 数组？切片？如果自己构建一个传输格式，应该是什么格式？
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})

}

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

/*
数据绑定
curl http://localhost:8080/root/admin
*/

func handlerUrlBind(r *gin.Engine) {
	// JSON绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		// 声明接收的变量
		var login Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if login.User != "root" || login.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// form数据绑定
func handlerFormBind(r *gin.Engine) {
	// JSON绑定
	// JSON绑定
	r.POST("/loginForm", func(c *gin.Context) {
		// 声明接收的变量
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if form.User != "root" || form.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

/*
json数据绑定
curl -X POST http://localhost:8080/loginJSON \
-H "Content-Type: application/json" \
-d '{"user": "root", "password": "admin"}'

curl -X POST http://localhost:8080/loginJSON \
-H "Content-Type: application/json" \
-d '{"user": "root", "password": "123"}'
*/
func handlerJsonBind(r *gin.Engine) {
	// JSON绑定
	r.POST("loginJSON", func(c *gin.Context) {
		// 声明接收的变量
		var json Login
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&json); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if json.User != "root" || json.Pssword != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
}

// 路由组
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

// 文件上传
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
