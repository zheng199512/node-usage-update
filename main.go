package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"k8s.io/klog"
	"time"
)

func GetUsage() float64 {
	usage, err := cpu.Percent(300, false)
	if err != nil {
		klog.Error("error get cpu percent")
	}
	return usage[0]
}

func main() {
	for i := 0; i < 1; i++ {
		test := GetUsage()
		println(test)
		time.Sleep(5 * time.Second)
	}
}
