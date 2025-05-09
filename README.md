
# 安装
```bash
go get github.com/DicoderCn/ginp
```
# 快速开始
```go
package main
import (
	"github.com/DicoderCn/ginp"
)
//1.router.go:register router
func main() {
	r.GET("/index", ginp.RegisterHandler(controller.Index))
}

//2.controller/index.go: use
func AccountConfigCreate(ctx *ginp.ContextPlus) {
	//your code here 
    //The usage is exactly the same as gin
    //ctx.Success() Success
}


//extra:
//You can customize the status codes for success and failure, such as:
// ginp.SetHttpCode(code int)
```