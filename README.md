
## ContextPlus API

<div align="right">
  <a href="https://github.com/DicoderCn/ginp">English</a>
  |
  <a href="https://github.com/DicoderCn/ginp/blob/master/README_zh.md">中文</a>
</div>

<a id="english"></a>
### English Version
`ContextPlus` is an extension of `gin.Context` that provides more convenient response methods.

### Method List

#### Success
```go
func (c *ContextPlus) Success(messages ...string)
```
Returns success JSON response
- `messages`: Optional success messages

#### Fail
```go
func (c *ContextPlus) Fail(strs ...string)
```
Returns failure JSON response
- `strs`: Optional error messages

#### SuccessData
```go
func (c *ContextPlus) SuccessData(data any, extra any, messages ...string)
```
Returns success JSON response with data
- `data`: Main data
- `extra`: Additional data
- `messages`: Optional success messages

#### FailData
```go
func (c *ContextPlus) FailData(data any, extra any, messages ...string)
```
Returns failure JSON response with data
- `data`: Main data
- `extra`: Additional data
- `messages`: Optional error messages

#### SuccessHtml
```go
func (c *ContextPlus) SuccessHtml(path string)
```
Returns HTML response
- `path`: HTML template path

#### R
```go
func (c *ContextPlus) R(code int, obj any)
```
General JSON response method
- `code`: HTTP status code
- `obj`: Response object

#### Log
```go
func (c *ContextPlus) Log(data any)
```
Logs request data
- `data`: Log data