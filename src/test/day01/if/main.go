package main

import "fmt"

func main() {
	/* 	age := 18
	   	if age > 18 {
	   		fmt.Println("澳门首家线上赌场开业啦")
	   	} else if age < 18 {
	   		fmt.Println("warning...")
	   	} else {
	   		fmt.Println("you are ok!")
	   	} */
	age := 18
	if age == 18 {
		fmt.Println("you are ok!")
	}
	if age2 := 28; age2 > 18 {
		fmt.Println("you are ok?")
		fmt.Println(age2)
	}
}
