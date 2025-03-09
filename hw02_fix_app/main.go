package main

import (
	"fmt"

	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/printer"
	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/reader"
	"github.com/sergeyptv/home_work_basic/hw06_testing/hw02/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")

	var err error

	ret, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
