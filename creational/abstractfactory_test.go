package creational

import "testing"

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestExampleRDBFactory(t *testing.T) {
	var factory DAOFactory = &RDBDAOFactory{}
	getMainAndDetail(factory)
}

func TestExampleXMLFactory(t *testing.T) {
	var factory DAOFactory = &XMLDAOFactory{}
	getMainAndDetail(factory)
}
