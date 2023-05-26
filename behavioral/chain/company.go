package chain

import "fmt"

const (
	ManagerOK = iota
	DirectorOK
	CeoOK
	_
	ManagerNO
	DirectorNO
	CeoNO
)

// ----------------------------------

type LeaveRequest struct {
	Id      uint
	StaffId uint
	Name    string
	Content string
	Day     int
	State   int
}

// ----------------------------------

// 员工
type Staff struct {
	Id    uint
	Name  string
	Level int
}

func (s *Staff) Interceptor(c IChain) IResponse {
	fmt.Println(1)
	return c.Proceed(c.Request())
}

// 经理
type Manager struct {
	Id    uint
	Name  string
	Level int
}

func (s *Manager) Interceptor(c IChain) IResponse {
	r := c.Request()
	request := r.(*LeaveRequest)
	if request.Day < 3 {
		if request.Id == 0 { // 模拟未成功
			request.State = ManagerNO
		} else {
			request.State = ManagerOK
		}
		return request
	} else {
		request.State = ManagerOK
		return c.Proceed(c.Request())
	}
}

// 总监
type Director struct {
	Id    uint
	Name  string
	Level int
}

func (s *Director) Interceptor(c IChain) IResponse {
	r := c.Request()
	request := r.(*LeaveRequest)
	if request.Day < 7 {
		if request.Id == 1 { // 模拟未成功
			request.State = DirectorNO
		} else {
			request.State = DirectorOK
		}
		return request
	} else {
		request.State = DirectorOK
		return c.Proceed(c.Request())
	}
}

// 总经理
type CEO struct {
	Id    uint
	Name  string
	Level int
}

func (s *CEO) Interceptor(c IChain) IResponse {
	r := c.Request()
	request := r.(*LeaveRequest)

	if request.Id == 2 {
		request.State = CeoNO
	} else {
		request.State = CeoOK
	}
	return c.Proceed(c.Request())
}
