package behavioral

import "fmt"

/*	策略模式
 *	优点：
 *		1. 策略模式提供了对“开闭原则”的完美支持，用户可以在不修改原有系统的基础上选择算法或行为，
 *		   也可以灵活地增加新的算法或行为。
 * 		2. 使用策略模式可以避免多重条件选择语句。多重条件选择语句不易维护，
 *		   它把采取哪一种算法或行为的逻辑与算法或行为本身的实现逻辑混合在一起，
 *		   将它们全部硬编码(Hard Coding)在一个庞大的多重条件选择语句中，
 *		   比直接继承环境类的办法还要原始和落后。
 *		3. 策略模式提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，因此不同的环境类可以方便地复用这些策略类。
 *
 *	缺点：
 *		1. 策略模式提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，
 *		   因此不同的环境类可以方便地复用这些策略类。
 *		2. 策略模式将造成系统产生很多具体策略类，任何细小的变化都将导致系统要增加一个新的具体策略类。
 *
 *	适用场景：
 * 		1. 准备一组算法，并将每一个算法封装起来，使得它们可以互换。
 */

// --------------------------------------------------------------

// 商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景

// 销售策略
type SellStrategy interface {
	//根据原价得到售卖价
	GetPrice(price float64) float64
}

// --------------------------------------------------------------

type StrategyA struct{}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A, 所有商品打八折")
	return price * 0.8
}

type StrategyB struct{}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B, 所有商品满200 减100")

	if price >= 200 {
		price -= 100
	}

	return price
}

// --------------------------------------------------------------

// 环境类
type Goods struct {
	Price    float64
	Strategy SellStrategy
}

// 设置策略
func (g *Goods) SetStrategy(s SellStrategy) {
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("原价值 ", g.Price, " .")
	return g.Strategy.GetPrice(g.Price)
}
