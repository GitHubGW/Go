// #1.4 Functions Part Two

package main

import (
	"fmt"
	"strings"
)

// lengthAndUppercase함수를 아래와 같이 다르게 바꿔서 만들 수도 있다.
// (length int, uppercase string)를 통해 이 함수가 리턴할 값이 어떤 변수인지 알려주게 되면 return에 쓰지 않아도 알아서 Go가 리턴할 변수를 찾아서 리턴한다.
// 이것을 naked return이라고 부른다. (리턴할 변수를 꼭 명시하지 않아도 된다는 의미에서 나왔다.)
// 개인적으로 선호하는 방식은 아님 (리턴되는 곳에 어떤 값이 리턴되는지 적어주는게 좀 더 명시적이라고 생각함)
func lengthAndUppercase2(name string) (nameLength int, nameUppercase string){
	// defer을 이용하면 현재 이 함수의 실행이 끝나고 난 후 추가적으로 실행시킬 수 있다.
	defer fmt.Println("defer execute!")

	nameLength = len(name)
	nameUppercase = strings.ToUpper(name);
	return
}

func main() {
	nameLength, nameUppercase := lengthAndUppercase2("javascript")
	fmt.Println(nameLength, nameUppercase)

}