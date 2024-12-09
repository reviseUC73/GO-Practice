package conditionPackage

import "fmt"

func ShowCondition() {
	fmt.Println("===== Condition Examples =====")

	// If-Else Example
	var score int = 85
	if score >= 80 {
		fmt.Println("Grade: A")
	} else if score >= 70 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Switch Example
	var dayOfWeek = 3
	switch dayOfWeek {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid Day")
	}

	// For Loop Example
	fmt.Println("For Loop Example:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// While Loop Example
	fmt.Println("While Loop Example:")
	i := 1
	for i <= 5 {
		fmt.Printf("Number: %d\n", i)
		i++
	}

	// Do-While Loop Example (using infinite loop with break)
	fmt.Println("Do-While Loop Example:")
	i = 1
	for {
		fmt.Printf("Number: %d\n", i)
		i++
		if i > 5 {
			break
		}
	}


	num1 := 5;
	num2 :=10;
	if sumNum := num1+ num2; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}
}
