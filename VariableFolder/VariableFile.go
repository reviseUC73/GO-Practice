package variablePackage

import (
	"fmt"
)

func ShowGoVariable() {
	var booleanVar bool = true
	var intVar int = 10
	var floatVar float64 = 3.14
	var stringVar string = "Hello Go"
  
	fmt.Println("Boolean:", booleanVar)
	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)
	fmt.Println("String:", stringVar)

	var fullName string = "rew"
	var lastName = "kub"
	var age int  = 12
	fmt.Printf("Hello %s %s Yay!, age =%d \n",fullName,lastName,age)

	// fullName := "rew" // ตัวแปรไม่สามารถประกาษชื่อซ่ำได้
	fullNameShort := "rew"
	lastNameShort := "kub"
	ageShort := 12
	fmt.Printf("Hello %s %s Yay!, age =%d \n",fullNameShort, lastNameShort, ageShort)
	ageShort = 13
	fmt.Printf("New >> Hello %s %s Yay!, age =%d \n",fullNameShort, lastNameShort, ageShort)
	
	const lock = "============ lock ==========" //or let
	println(lock)
}