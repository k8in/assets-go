package behaviour

// ChainHandler 责任链处理器接口
type ChainHandler[T any] interface {
	// SetNext 设置下一个处理器
	SetNext(handler ChainHandler[T]) ChainHandler[T]
	// Handle 处理请求
	Handle(req T) bool
	// Name 获取处理器名称
	Name() string
}

// BaseChainHandler 基础责任链处理器
type BaseChainHandler[T any] struct {
	next ChainHandler[T]
}

// SetNext 设置下一个处理器
func (h *BaseChainHandler[T]) SetNext(handler ChainHandler[T]) ChainHandler[T] {
	h.next = handler
	return handler
}

// Handle 处理请求，如果当前处理器无法处理，则传递给下一个处理器
func (h *BaseChainHandler[T]) Handle(req T) bool {
	if h.next != nil {
		return h.next.Handle(req)
	}
	return false
}

// Name 获取处理器名称
func (h *BaseChainHandler[T]) Name() string {
	return "BaseChainHandler"
}

// ChainBuilder 责任链构建器
type ChainBuilder[T any] struct {
	first ChainHandler[T]
	last  ChainHandler[T]
}

// NewChainBuilder 创建新的责任链构建器
func NewChainBuilder[T any]() *ChainBuilder[T] {
	return &ChainBuilder[T]{}
}

// Add 添加处理器到责任链
func (b *ChainBuilder[T]) Add(handler ChainHandler[T]) *ChainBuilder[T] {
	if b.first == nil {
		b.first = handler
		b.last = handler
	} else {
		b.last.SetNext(handler)
		b.last = handler
	}
	return b
}

// Build 构建责任链并返回第一个处理器
func (b *ChainBuilder[T]) Build() ChainHandler[T] {
	return b.first
}
