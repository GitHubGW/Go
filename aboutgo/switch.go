// #1.7 Switch

package main

import "fmt"

func checkEat(age int) bool {
	
	// switch문 
	// switch문도 if문과 마찬가지로 바로 즉석에서 변수를 생성하고 가져와 쓸 수 있다.
	// ex) switch koreanAge := age + 2; koreanAge{}
	switch {
	case age<= 19:
		return false
	case age >19:
		return true
	default:
		return false
	}
}

func main() {
	BooleanCheckEat:=checkEat(20);
	fmt.Println(BooleanCheckEat)
}