package main

import (
	"github.com/zheng199512/node-usage-update/pkg/cpu"
	"time"
)

func main() {

	for i := 0; i < 1; i++ {
		usage := cpu.GetUsage()
		println(usage)
		time.Sleep(5 * time.Second)
	}
}
