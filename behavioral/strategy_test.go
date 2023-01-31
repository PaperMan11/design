package behavioral

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	nike := Goods{Price: 200.0}

	//上午 ，商场执行策略A
	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	//下午， 商场执行策略B
	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午nike鞋卖", nike.SellPrice())
}
