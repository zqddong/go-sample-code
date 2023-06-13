package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type BankBusinessHandler interface {
	// TakeRowNumber 排队拿号
	TakeRowNumber()
	// WaitInHead 等位
	WaitInHead()
	// HandlerBusiness 处理具体业务
	HandlerBusiness()
	// Commentate 对服务评价
	Commentate()
	// CheckVipIdentity 钩子方法 判断是否Vip
	CheckVipIdentity() bool
}

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// ExecutorBankBusiness 模板方法 处理银行业务
func (b *BankBusinessExecutor) ExecutorBankBusiness() {
	// 适用于与客户端单次交互的流程
	// 如果需要与客户端多次交互才能完成整个流程，
	// 每次交互的操作去对应模板里定义的方法就好，
	// 并不需要一个调用所有方法的模板方法
	b.handler.TakeRowNumber()
	if !b.handler.CheckVipIdentity() {
		b.handler.WaitInHead()
	}
	b.handler.HandlerBusiness()
	b.handler.Commentate()
}

type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (dh *DepositBusinessHandler) HandlerBusiness() {

}

func (*DepositBusinessHandler) HandleBusiness() {
	fmt.Println("账户存储很多人民币...")
}

func (dh *DepositBusinessHandler) CheckVipIdentity() bool {
	return dh.userVip
}

type DefaultBusinessHandler struct{}

func (*DefaultBusinessHandler) TakeRowNumber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (dh *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (*DefaultBusinessHandler) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (*DefaultBusinessHandler) CheckVipIdentity() bool {
	return false
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}

func main() {
	dh := &DepositBusinessHandler{userVip: false}
	bbe := NewBankBusinessExecutor(dh)
	bbe.ExecutorBankBusiness()
}
