package behaviour

import (
	"fmt"
	"strings"
	"testing"
)

// 示例：数字处理器 - 处理偶数
type EvenHandler struct {
	BaseChainHandler[int]
}

func (h *EvenHandler) Name() string {
	return "EvenHandler"
}

func (h *EvenHandler) Handle(req int) bool {
	if req%2 == 0 {
		fmt.Printf("%s 处理了偶数: %d\n", h.Name(), req)
		return true
	}
	return h.BaseChainHandler.Handle(req)
}

// 示例：数字处理器 - 处理奇数
type OddHandler struct {
	BaseChainHandler[int]
}

func (h *OddHandler) Name() string {
	return "OddHandler"
}

func (h *OddHandler) Handle(req int) bool {
	if req%2 != 0 {
		fmt.Printf("%s 处理了奇数: %d\n", h.Name(), req)
		return true
	}
	return h.BaseChainHandler.Handle(req)
}

// 示例：字符串处理器 - 处理大写字符串
type UpperCaseHandler struct {
	BaseChainHandler[string]
}

func (h *UpperCaseHandler) Name() string {
	return "UpperCaseHandler"
}

func (h *UpperCaseHandler) Handle(req string) bool {
	if strings.ToUpper(req) == req && req != "" {
		fmt.Printf("%s 处理了大写字符串: %s\n", h.Name(), req)
		return true
	}
	return h.BaseChainHandler.Handle(req)
}

// 示例：字符串处理器 - 处理小写字符串
type LowerCaseHandler struct {
	BaseChainHandler[string]
}

func (h *LowerCaseHandler) Name() string {
	return "LowerCaseHandler"
}

func (h *LowerCaseHandler) Handle(req string) bool {
	if strings.ToLower(req) == req && req != "" {
		fmt.Printf("%s 处理了小写字符串: %s\n", h.Name(), req)
		return true
	}
	return h.BaseChainHandler.Handle(req)
}

// 测试数字处理责任链
func TestIntegerChain(t *testing.T) {
	fmt.Println("=== 测试数字处理责任链 ===")

	// 创建处理器
	evenHandler := &EvenHandler{}
	oddHandler := &OddHandler{}

	// 使用构建器创建责任链
	chain := NewChainBuilder[int]().
		Add(evenHandler).
		Add(oddHandler).
		Build()

	// 测试数据
	testNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range testNumbers {
		handled := chain.Handle(num)
		if !handled {
			t.Errorf("数字 %d 未被任何处理器处理", num)
		}
	}
}

// 测试字符串处理责任链
func TestStringChain(t *testing.T) {
	fmt.Println("\n=== 测试字符串处理责任链 ===")

	// 创建处理器
	upperHandler := &UpperCaseHandler{}
	lowerHandler := &LowerCaseHandler{}

	// 使用构建器创建责任链
	chain := NewChainBuilder[string]().
		Add(upperHandler).
		Add(lowerHandler).
		Build()

	// 测试数据
	testStrings := []string{"HELLO", "world", "MixedCase", "golang", "TEST"}

	for _, str := range testStrings {
		handled := chain.Handle(str)
		if !handled {
			fmt.Printf("字符串 '%s' 未被任何处理器处理\n", str)
		}
	}
}

// 测试手动设置责任链
func TestManualChain(t *testing.T) {
	fmt.Println("\n=== 测试手动设置责任链 ===")

	// 手动创建和连接处理器
	evenHandler := &EvenHandler{}
	oddHandler := &OddHandler{}

	// 手动设置链条
	evenHandler.SetNext(oddHandler)

	// 测试
	testNumbers := []int{2, 3, 4, 5}
	for _, num := range testNumbers {
		handled := evenHandler.Handle(num)
		if !handled {
			t.Errorf("数字 %d 未被任何处理器处理", num)
		}
	}
}

// 测试空链条
func TestEmptyChain(t *testing.T) {
	fmt.Println("\n=== 测试空链条 ===")

	// 创建空的构建器
	chain := NewChainBuilder[int]().Build()

	if chain != nil {
		t.Error("空构建器应该返回 nil")
	}
}

// 基准测试
func BenchmarkChainHandler(b *testing.B) {
	evenHandler := &EvenHandler{}
	oddHandler := &OddHandler{}

	chain := NewChainBuilder[int]().
		Add(evenHandler).
		Add(oddHandler).
		Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chain.Handle(i)
	}
}

// 示例函数 - 演示完整用法
func ExampleChainBuilder() {
	// 创建处理器
	evenHandler := &EvenHandler{}
	oddHandler := &OddHandler{}

	// 构建责任链
	chain := NewChainBuilder[int]().
		Add(evenHandler).
		Add(oddHandler).
		Build()

	// 处理请求
	numbers := []int{1, 2, 3, 4}
	for _, num := range numbers {
		chain.Handle(num)
	}

	// Output:
	// OddHandler 处理了奇数: 1
	// EvenHandler 处理了偶数: 2
	// OddHandler 处理了奇数: 3
	// EvenHandler 处理了偶数: 4
}
