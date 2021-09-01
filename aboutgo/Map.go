// #1.10 Maps

package main

import "fmt"

func main() {
	// map[키 타입]벨류 타입
	// go에서는 map메서드를 사용하기 위해 키와 벨류가 필요하고 각각의 키와 벨류에는 타입을 지정해줘야 한다.
	// map[string]string{"name":"GW", "age":"12"}을 통해 map의 키는 string타입이고 벨류 또한 string타입이라고 지정했다.
	person := map[string]string{"name":"GW", "age":"12"}
	fmt.Println(person)

	// 아래와 같이 map메서드는 for, range와 같이 이용해서 쓸 수도 있다. 
	for key, value := range person {
		fmt.Println(key, value);
	}
}