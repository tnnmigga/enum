```
package main

import (
	"fmt"

	"github.com/tnnmigga/enum"
)

var EnumTestInt = enum.New[struct {
	ZERO int
	ONE  int
	TWO  int
}]()

var EnumTestStr = enum.New[struct {
	ZERO string
	ONE  string
	TWO  string
}]()

func main() {
	fmt.Println(EnumTestInt.ZERO, EnumTestInt.ONE, EnumTestInt.TWO)
	fmt.Println(EnumTestStr.ZERO, EnumTestStr.ONE, EnumTestStr.TWO)
}

// 运行结果
// 0 1 2
// ZERO ONE TWO

```
