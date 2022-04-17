package main

import (
	"fmt"
	"goConser"
)

func main() {
	// const block one
	block1 := goConser.NewConstBlock().AddConstOne(
		goConser.ConstOne{ConstName: "Language1", ConstValue: "Chinese"},
		goConser.ConstOne{ConstName: "Language2", ConstValue: "English", Annotate: "English Language"}, // block annotate
	)
	// const block two
	block2 := goConser.NewConstBlock().AddConstOne(
		goConser.ConstOne{ConstName: "Location1", ConstValue: "Chengdu"},
		goConser.ConstOne{ConstName: "Location2", ConstValue: "Beijing"},
	).SetBlockNote("the places where have airports") // line annotate

	constFile := goConser.NewConst2File().
		SetSaveDir("./model").        // save file dir
		SetSaveFile("file.go").       //save file name
		SetPackageName("model").      // package name
		SetFlag(goConser.FileCreate). // save data by append or create a file
		AddConstBlock(block1, block2)

	err := constFile.Run() // generate file
	if err != nil {
		fmt.Println(err.Error())
	}
}
