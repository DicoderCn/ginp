package ginp

// 定义一个接口，用于实现自定义的扩展方法
type IContextPlus interface {
	Success()
	Fail(strs ...string)
	FailData(msg string, data any, extra ...any)
	SuccessData(data any)
	SuccessHtml(path string)
	respJson(code int, obj any)
	Log(data any)
}

// 强制Context实现IContext接口
// 由于 *Context 未实现 IContext 接口的 GetToken 方法，此声明暂时无法通过编译。
// 需确保 Context 结构体实现了 IContext 接口的所有方法后再使用。
// 目前仅作占位，实际使用时需完成方法实现。
var _ IContextPlus = (*ContextPlus)(nil)
