package main

import "fmt"

// 中介者--机场指挥接口定义
type mediator interface {
	canLanding(airplane airplane) bool
	notifyAboutDeparture()
}

// 组件--飞行棋接口定义
type airplane interface {
	landing()
	takeOff()
	permitLanding()
}

// 组件1--波音飞机
type boeingPlane struct {
	mediator
}

func (b *boeingPlane) landing() {
	if !b.mediator.canLanding(b) {
		fmt.Println("Airplane Boeing: 飞机跑道正在被占用，无法降落")
		return
	}
	fmt.Println("Airplane Boeing: 已经成功降落")
}

func (b *boeingPlane) takeOff() {
	fmt.Println("Airplane Boeing: 正在起飞离开跑道")
	b.mediator.notifyAboutDeparture()
}

func (b *boeingPlane) permitLanding() {
	fmt.Println("Airplane Boeing: 收到指挥塔信号，允许降落，正在降落")
	b.landing()
}

// 组件2--空客飞机
type airBusPlane struct {
	mediator
}

func (a *airBusPlane) landing() {
	if !a.mediator.canLanding(a) {
		fmt.Println("Airplane AirBus: 飞机跑道正在被占用，无法降落")
		return
	}
	fmt.Println("Airplane AirBus: 已经成功降落")
}

func (a *airBusPlane) takeOff() {
	fmt.Println("Airplane AirBus: 正在起飞离开跑道")
	a.mediator.notifyAboutDeparture()
}

func (a *airBusPlane) permitLanding() {
	fmt.Println("Airplane AirBus: 收到指挥塔信号，允许降落，正在降落")
	a.landing()
}

// 中介者实现--指挥塔
type manageTower struct {
	isRunwayFree bool
	airportQueue []airplane
}

func (t *manageTower) canLanding(ap airplane) bool {
	if t.isRunwayFree {
		// 跑道空闲 允许降落 同时把状态更新为繁忙
		t.isRunwayFree = false
		return true
	}
	// 跑道繁忙 把飞机加入等待通知队列
	t.airportQueue = append(t.airportQueue, ap)
	return false
}

func (t *manageTower) notifyAboutDeparture() {
	if !t.isRunwayFree {
		t.isRunwayFree = false
	}

	if len(t.airportQueue) > 0 {
		firstPlaneInWaitingQueue := t.airportQueue[0]
		t.airportQueue = t.airportQueue[1:]
		firstPlaneInWaitingQueue.permitLanding()
	}
}

func newManageTower() *manageTower {
	return &manageTower{
		isRunwayFree: true,
	}
}

func main() {
	t := newManageTower()
	boeing := &boeingPlane{
		mediator: t,
	}
	airbus := &airBusPlane{
		mediator: t,
	}

	boeing.landing()
	airbus.landing()
	boeing.takeOff()

}
