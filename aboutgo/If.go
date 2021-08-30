// #1.6 If with a Twist

package main

import "fmt"

func checkDrink(age int) bool {

	// if문 
	// go에서는 if문에 조건을 ()괄호로 묶어줄 필요없이 바로 쓰면 된다.
	// if문 안에서 바로 즉석에서 변수를 생성하고 가져와 쓸 수 있다.
	if koreanAge := age+2; koreanAge < 18{
	 return false
	}else {
		return true
	}
}

func main() {
	BooleanCheckDrink:=checkDrink(15);
	fmt.Println(BooleanCheckDrink)
}