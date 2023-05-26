package chain

type chainBuilder struct {
	head, tail IInterceptor
	body       []IInterceptor
}

func NewChainBuilder() *chainBuilder {
	return &chainBuilder{
		body: make([]IInterceptor, 0),
	}
}

// Head 将拦截器添加到责任链头部
func (ic *chainBuilder) Head(interceptor IInterceptor) {
	ic.head = interceptor
}

// Tail 将拦截器添加到责任链尾部
func (ic *chainBuilder) Tail(interceptor IInterceptor) {
	ic.tail = interceptor
}

// AddInterceptor 顺位添加一个拦截器到责任链中
func (ic *chainBuilder) AddInterceptor(interceptor IInterceptor) {
	ic.body = append(ic.body, interceptor)
}

// Execute 依次执行当前责任链上所有拦截器
func (ic *chainBuilder) Execute(req IRequest) IResponse {

	//将全部拦截器放入Builder中
	var interceptors []IInterceptor
	if ic.head != nil {
		interceptors = append(interceptors, ic.head)
	}
	if len(ic.body) > 0 {
		interceptors = append(interceptors, ic.body...)
	}
	if ic.tail != nil {
		interceptors = append(interceptors, ic.tail)
	}

	//创建一个拦截器责任链，执行每一个拦截器
	chain := NewChain(req, 0, interceptors)

	//进入责任链执行
	return chain.Proceed(req)
}
