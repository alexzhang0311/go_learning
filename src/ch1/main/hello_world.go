package main //包，表明代码所在模块（包）

import (
	"fmt" //引入代码依赖
	"os"
)

//功能实现
func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World!", os.Args[1])
		os.Exit(0)
	}
	fmt.Println("Hello Alex!")
}
