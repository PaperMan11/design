package creational

/*	工厂方法模式
 *
 *	优点：
 *		1. 不需要记住具体类名，甚至连具体参数都不用记忆。
 * 		2. 实现了对象创建和使用的分离。
 *		3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
 *		4. 对于新产品的创建，符合开闭原则。
 *
 *	缺点：
 * 		1. 增加系统中类的个数，复杂度和理解度增加。
 *		2. 增加了系统的抽象性和理解难度。
 *
 * 	适用场景：
 * 		1. 客户端不知道它所需要的对象的类。
 *		2. 抽象工厂类通过其子类来指定创建哪个对象。
 */

// --------------------------------------------------------------

// Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// --------------------------------------------------------------

// OperatorBase Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// --------------------------------------------------------------

// OperatorFactory 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// --------------------------------------------------------------

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// --------------------------------------------------------------

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}
