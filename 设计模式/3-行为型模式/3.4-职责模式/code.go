package main

import "fmt"

// 挂号—>诊室看病—>收费处缴费—>药房拿药

type PatientHandler interface {
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

// Next 充当抽象类型，实现公共方法，抽象方法不实现留给实现类自己实现
type Next struct {
	nextHandler PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) (err error) {
	// 调用不到外部类型的 Do 方法，所以 Next 不能实现 Do 方法
	if n.nextHandler != nil {
		if err = n.nextHandler.Do(patient); err != nil {
			return
		}

		return n.nextHandler.Execute(patient)
	}

	return
}

type patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

// Reception 挂号处理器
type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) (err error) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	return
}

// Clinic 诊室处理器--用于医生给病人看病
type Clinic struct {
	Next
}

func (d *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	return
}

// Cashier 收费处处理器
type Cashier struct {
	Next
}

func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
		return
	}
	fmt.Println("Cashier getting money from patient patient")
	p.PaymentDone = true
	return
}

// Pharmacy 药房处理器
type Pharmacy struct {
	Next
}

func (m *Pharmacy) Do(p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		return
	}
	fmt.Println("Pharmacy giving medicine to patient")
	p.MedicineDone = true
	return
}

// StartHandler 不做操作，作为第一个Handler向下转发请求
type StartHandler struct {
	Next
}

// Do 空Handler的Do
func (h *StartHandler) Do(c *patient) (err error) {
	// 空Handler 这里什么也不做 只是载体 do nothing...
	return
}

func main() {
	patientHealthHandler := StartHandler{}
	//
	patient := &patient{Name: "abc"}
	// 设置病人看病的链路
	patientHealthHandler.SetNext(&Reception{}). // 挂号
							SetNext(&Clinic{}).  // 诊室看病
							SetNext(&Cashier{}). // 收费处交钱
							SetNext(&Pharmacy{}) // 药房拿药
	// 还可以扩展，比如中间加入化验科化验，图像科拍片等等

	// 执行上面设置好的业务流程
	if err := patientHealthHandler.Execute(patient); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")
}
