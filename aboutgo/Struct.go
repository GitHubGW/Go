// #1.11 Struct

package main

import "fmt"

// struct는 일종의 구조체 같은 것으로 struct을 만들기 위해서는 먼저 타입 정의가 필요하다.
// favFood의 []string은 일종의 배열형태로 만든다고 정의한 것이다.
// go에서는 constructor가 없다. 그래서 사용자 스스로 constructor를 실행해야 한다.
type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood:=[]string{"pizza", "chicken", "cola"}
	fmt.Println(favFood)

	// person에게 value를 주는 방법엔 2가지가 있다.
	// 위의 type에서 지정한 것처럼 차례대로 줄 수 있다. 
	// 하지만 아래 방법은 value가 어떤 타입인지 정확히 인지하기가 어려워서 추천되는 방법은 아니다.
	// gw := person{"gw",20,favFood}

	gw:=person{name:"gw", age:20, favFood:favFood}
	fmt.Println(gw)
	fmt.Println(gw.age)
	fmt.Println(gw.favFood)
}