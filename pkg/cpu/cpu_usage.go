package cpu

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func GetUsage() (idle uint64) {
	contents, err := ioutil.ReadFile("/porc/stat")
	if err != nil {
		println("readfile /proc/stat  encounter err msg: ", err)
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
		}
	}
	return 100-idle
}

