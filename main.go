package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/reviseUC73/go-test1/ControlStructure"
	"github.com/reviseUC73/go-test1/VariableFolder"
	"github.com/reviseUC73/go-test1/file"
)

// fmt show sth from console (print) || import และไม่ใช้ error go จะตรวจสอบตั่งแต่ compile level เลบ runtime level

func main (){
	id := uuid.New() // keep ... data of uuid.New
	fmt.Println("Hello World")
	fmt.Printf("UUID: %s",id)
	file.MyPrint()
	file.TestPrint()

	variablePackage.ShowGoVariable()
	conditionPackage.ShowCondition()

}