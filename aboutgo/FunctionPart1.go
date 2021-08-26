// #1.3 Functions Part One

package main

import (
	"fmt"
	"strings"
)

// go에서는 함수에 파라미터와 리턴하는 값의 타입도 지정해줘야 한다.
// 그래서 파라미터 a와 b의 데이터 타입을 int라고 알려주었고, 이 함수가 리턴하는 값의 타입 또한 int라고 지정해주었다.
// a int, b int를 축약해서 a, b int로 작성할 수도 있다.
func multiply(a int, b int) int{
	return a * b;
}

// 또한 특이하게 go에서는 함수가 여러 개의 값을 리턴할 수 있다.
// 그래서 lengthAndUppercase함수는 name 값을 받아서 len()함수를 통해 그 name 변수의 길이와 strings.ToUpper()을 통해 name을 대문자로 리턴한다.
// 즉, name의 문자열의 길이와 해당 문자열을 대문자로 처리한 두 개의 값을 리턴하는 것이다.
// (int, string): 두 개의 값을 리턴하기 때문에 당연히 return 해주는 타입 또한 2개를 지정해줘야 한다.
func lengthAndUppercase(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

// 만약 함수의 파라미터로 여러 개의 값을 받고 싶다면 아래와 같이 ...을 붙여주면 된다.
// word는 여러 개의 값을 받게 되고 배열 형태로 출력한다.
func repeatMe(word... string) {
	fmt.Println(word)
}

func main() {
	// fmt.Println(multiply(2,2))
	// fmt.Println(lengthAndUppercase("golang"))

	// len(name), strings.ToUpper(name)값을 nameLength, nameUppercase변수로 받아서 아래에서 출력할 수 있다.
	// nameLength, nameUppercase := lengthAndUppercase("golang")

	// 만약 하나의 값만 받고 싶다면 _을 통해 자리를 비워두면 된다. 
	nameLength, _ := lengthAndUppercase("golang")
	fmt.Println(nameLength)

	repeatMe("mon","tue","wed")

}