package creational

import "fmt"

/* 	抽象工厂模式
 *
 *	优点：
 *		1. 拥有工厂方法模式的优点。
 *		2. 当一个产品族中的多个对象被设计成一起工作时，
 *		   它能够保证客户端始终只使用同一个产品族中的对象。
 *		3. 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。
 *
 *	缺点：
 *		1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，
 *		   甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。
 *
 * 	适用场景：
 *		1. 系统中有多于一个的产品族。而每次只使用其中某一产品族。
 *	       可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
 *		2. 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。
 */

/*	总结
 *	简单工厂模式：一个工厂负责创建所有产品
 *  	- 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂
 *
 * 	工厂方法模式：一个工厂创建一个产品
 *		- 系统的可扩展性也就变得非常好，无需修改接口和原类
 * 		- 增加系统中类的个数，复杂度和理解度增加（一个具体产品就需要对应一个具体工厂）
 *
 * 	抽象方法模式：一个工厂创建一系列（一个产品族）的产品
 * 		- 增加新的产品族很方便，无须修改已有系统，符合“开闭原则”
 * 		- 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，违背了“开闭原则”
 * 		- 相当于在工厂方法模式的基础下进行了折中
 *			■ 对于产品族来说遵循了开闭原则
 * 			■ 对于产品等级结构来说没有遵循开闭原则
 *			■ 如果产品结构等级稳定，那么就相当于完全遵循开闭原则
 */

// --------------------------------------------------------------

// 抽象层
// OrderMainDAO 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 订单详情
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// --------------------------------------------------------------

// DAOFactory DAO 抽象工厂
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// --------------------------------------------------------------

// 实现
// RDBMainDAO 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory 是 RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// --------------------------------------------------------------

// XMLMainDAO 为关系型数据库的OrderMainDAO实现
type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO 为关系型数据库的OrderDetailDAO实现
type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

// RDBDAOFactory 是 RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
