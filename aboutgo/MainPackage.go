// main.go가 아닌 다른 파일들을 실행해서 확인해보려면 파일명을 main.go로 바꾼 후 실행시키기. (go run main.go)

// #1.0 Main Package

// 만약 사용자가 프로젝트를 컴파일하고 싶다면 main.go파일을 만들고 그 안에 package main을 선언해줘야 한다.
// main은 오로지 컴파일을 위해 필요한 파일이다.
// 그래서 만약 컴파일이 필요하지 않은 단순히 기능들을 가지고 있는 파일들이라면 다른 이름의 파일을 만들어서 사용할 수도 있다.
// 예를들면 learning.go파일을 만들고 package learning을 선언해주고 사용하면 된다.
// go컴파일러는 기본적으로 파일이름이 main.go이고 그 안에 package main이 선언되어 있는 파일을 가장 먼저 찾고 그 안에 main함수를 찾아서 컴파일해준다.
// go에서는 기본적으로 파일 내에 어떤 패키지를 사용하는지 선언해줘야 한다.
// main.go에서는 package main을 통해 main을 사용한다고 선언해준 것이다.
package main

import "fmt"

// go에서는 main.go내부에 있는 함수를 실행하면서 실행되기 때문에 main.go안에 반드시 함수를 생성해줘야 한다.
// 현재 이 main 함수가 go프로그램의 시작점이 되는 부분이다.
func main() {
	
	fmt.Println("Hello!!")

}