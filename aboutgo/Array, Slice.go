// #1.9 Array and Slice

package main

import "fmt"

func main() {

	// Array 
	// go에서는 배열을 만들 때 배열의 길이와 타입을 정의해줘야 한다.
	// [배열 길이]타입{배열 안에 들어갈 데이터}
	name :=[5]string{"pizza","chicken","cola"}
	name[3]="fanta"
	name[4]="salad"
	fmt.Println(name)

	// 위처럼 배열의 길이를 지정해줄 수도 있고 배열의 길이를 지정하고 싶지 않을 때는 slice를 사용할 수 있다.
	// slice에 아이템을 추가하기 위해서 사용하는 메서드가 append()이다. 
	// append는 두 개의 인자를 받는데 첫 번째는 담을 변수(slice)이고 두 번째는 사용자가 추가하고 싶은 값이다. ex) append(변수명, "40")
	// 주의할 점은 append는 새로운 값이 추가된 새로운 slice를 리턴해준다. 
	age :=[]string{"10","20","30"}
	age2:=append(age, "40")

	fmt.Println(age)
	fmt.Println(age2)

	// names := []string{"pizza", "chicken", "cola"} 
	// names := [...]string{"pizza", "chicken", "cola"}
	// []는 길이 제한이 없는것이고 [...]은 초기에 작성한 변수 갯수만큼의 리스트 길이를 자동으로 설정해주는 것 같다.
	// 만약, [...]string('값1','값2','값3')으로 하면 길이 제한이 3개인 배열이 생성되는 것이다.
	// [...]으로 설정한 배열은 append나 names[인덱스]="값4" 로 배열에 값을 추가할 수 없다.
}