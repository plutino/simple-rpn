package main

import (
	"fmt"

	"github.com/fatih/color"

	"simple-rpn/calc"
)

func printErr(err error) {
	fmt.Println(color.RedString("Error: ") + err.Error())
}

func main() {
	c, err := calc.NewCalculator(calc.MODE_DEC, 128)
	if err != nil {
		fmt.Println(err)
		return
	}

	tokenCt := 0

	for {
		fmt.Printf(color.CyanString("In [ ")+"%d"+color.CyanString(" ]")+": ", tokenCt)
		var token string
		if _, err := fmt.Scan(&token); err != nil {
			printErr(err)
			continue
		}
		if err := c.Insert(token); err != nil {
			printErr(err)
			continue
		}
		tokenCt++
		str, err := c.PrintStack()
		if err != nil {
			printErr(err)
			continue
		}
		fmt.Printf(color.YellowString("Stack: [ ")+"%s"+color.YellowString(" ]")+"\n", str)
	}
}
