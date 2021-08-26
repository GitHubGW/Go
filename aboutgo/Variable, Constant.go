// #1.2 Variables and Constants

package main

import "fmt"

func main() {
	// go에서도 자바스크립트처럼 앞에 const와 var를 붙여서 변수를 만들 수 있다. 
	// 단 타입스크립트처럼 뒤에 타입을 지정해줘야 한다.
	// 그리고 마찬가지로 const는 변수의 값을 재할당할 수 없다. (var로 선언한 변수는 재할당이 가능하다.)
	// const name string="gw"
	// name="gn"

	// var name2 string="cody"
	// name2="sugar"

	// 변수명 := 값
	// 위의 var와 const로 변수를 만드는 대신 아래와 같이 축약형으로도 만들 수 있다. (단, 축약형은 func인 함수 내부에서만 사용가능하다.) 
	// 함수 밖에서는 var와 const형태로 변수를 만들 수 있다.
	// 만약 아래와 같이 작성하면 위처럼 타입을 따로 지정해줄 필요없이 Go가 자동으로 값에 대한 타입을 찾아준다.
	// 또한 처음 초기화 할 때 정해진 타입은 나중에 다른 타입의 값을 재할당할 수 없다.
	// (처음에 값을 string 타입으로 만들었다면 변수에 다른 값을 재할당할 때도 타입이 string이어야 한다.)
	name3 := "orange"
	name3="apple"
	
	// 위에서 name3의 타입을 string으로 초기화했기 때문에 boolean인 true로 변경할 수 없다.
	// name3=true 

	fmt.Println(name3)
}