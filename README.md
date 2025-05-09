
# 安装
```bash
go get github.com/DicoderCn/ginp
```
# 快速开始
```go
package main
import (
	"github.com/DicoderCn/ginp"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	ginp.SetSuccessCode(200)
	r.GET("/", ginp.RegisterHandler(Index))
	r.Run(":8082")
}
func Index(c *ginp.ContextPlus) {
	//your code here
    //...
	c.Success()
}



//extra:
//You can customize the status codes for success and failure, such as:
// ginp.SetHttpCode(code int)
```

# Customize 自定义
```go

//1.mygin/main.go  Define your context 
import (
	"github.com/DicoderCn/ginp"
)

type MyContextPlus struct {
	ginp.ContextPlus
}
func (c *MyContextPlus) Success() {
	//your code here
    //...
}
func (c *MyContextPlus) NewMethod() {
	//your code here
    //...
}

// 
func MyRegisterHandler(handler func(c *MyContextPlus)) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(&MyContextPlus{
			Context: c,
		})
	}
}

//2. Use
import (
	"myproject/mygin"
    "myproject/controller"
)
func main() {
    r := gin.Default()
	r.GET("/index", MyRegisterHandler("/", controller.Index))
}

//3.controller/index.go: use

func Index(ctx *MyContextPlus) {
	//your code here
    //The usage is exactly the same as gin
    //ctx.Success() Success
}

```