package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	//var s int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b%8 == 0 && b > 9 && b < 100 {
			c += b
		}
	}
	fmt.Println(c)
}

//fmt.Println(a % 100 / 10)

/*s := strconv.Itoa(a)
if len(s) < 2 {
	fmt.Println("0")
} else {
	fmt.Println(s[len(s)-2 : len(s)-1])
}*/

//Задание 2
/*hours := a / 30
minutes := a % 30 * 2
fmt.Printf("It is %d hours %d minutes.", hours, minutes)*/

//задание 3
/*switch {
case a < 0:
fmt.Println("Число отрицательное")
case a > 0:
fmt.Println("Число положительное")
default:
fmt.Println("Ноль")
}*/
