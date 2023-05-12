package main

import (
	"log"
	"time"
)

// 有时，我们希望把一些任务打包进行批量处理。比如，公交车发车场景:
// 公交车每隔5分钟发一班，不管是否已坐满乘客; 已坐满乘客情况下，不足5分钟也发车;
func main() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	maxPassenger := 30
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := GetNewPassenger()
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}

		select {
		case <-ticker.C:
			Launch(passengers)
		default:
			if len(passengers) >= maxPassenger {
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func GetNewPassenger() string {
	//return "New-Passenger"
	return ""
}

func Launch(passengers []string) bool {
	log.Printf("Passengers: %v\n", passengers)
	return true
}
