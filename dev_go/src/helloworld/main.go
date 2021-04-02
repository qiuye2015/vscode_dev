package main

import (
	"github.com/qiuye2015/vscode_dev/dev_go/src/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
}
