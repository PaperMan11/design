package structural

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}

	// 如果不使用代理来完成从韩国购买任务
	var shopping = new(KoreaShopping) //具体的购买主题

	// 1-先验货
	if g1.Fact == true {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪.")
		// 2-去韩国购买
		shopping.Buy(&g1)
		// 3-海关安检
		fmt.Println("对[", g1.Kind, "]进行了海关检查，成功的带回祖国")
	}

	fmt.Println("---------------以下是使用代理模式-------")
	var overseasProxy = NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)
}
