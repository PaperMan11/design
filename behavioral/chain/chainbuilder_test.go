package chain

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	cb := NewChainBuilder()

	cb.Head(&Manager{})
	cb.AddInterceptor(&Director{})
	cb.Tail(&CEO{})

	res := cb.Execute(&LeaveRequest{
		Id:      2,
		Name:    "tzq",
		Content: "请假",
		Day:     11,
	})

	fmt.Println(res)
}
