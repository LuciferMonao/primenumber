package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var in int64 = 0
	in, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	for ; in < 9223372036854775805; in++ {
		var testnumber int64 = 2
		var prim bool = true
		for ; testnumber < in/testnumber; testnumber++ {
			if in%testnumber == int64(0) {
				prim = false
				break
			}
		}
		if prim == true {
			fmt.Println(in)
		} else {
		}
	}
}
