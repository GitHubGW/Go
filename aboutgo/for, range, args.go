// #1.5 for, range, ...args

package main

import "fmt"

func Add(numbers... int) int {
	// for range문 
	// for와 range를 이용해서 go에서 반복문을 사용할 수 있는 for문을 만들 수 있다.
	// range는 기본적으로 index를 줘서 순번을 확인할 수 있다. index를 안 쓸 때는 _로 비워두면 된다.
	// for index, 변수:= range 배열{} 형태로 사용할 수 있다.
	// for index, number:= range numbers{
	// 	fmt.Println(index, number)
	// }

	// for문 
	for i:=0; i<len(numbers); i++{
		fmt.Println(numbers[i])
	}

	return 1;
}

func main() {
	total:=Add(1,2,3,4)
	fmt.Println(total)
}