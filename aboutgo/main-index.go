// main.go가 아닌 다른 파일들을 테스트 해보려면 이름을 main.go로 바꾼 후 테스트 해보면 됨.

package main

import (
	"fmt"
)

func canIDrink(age int) bool{
	// go에서의 if문 
	// go에서는 if문에 조건을 ()괄호로 묶어줄 필요없이 바로 쓰면 된다.
	// go의 if문에서는 if문 안에서 변수를 생성하고 바로 가져와 쓸 수 있다.  
	// if koreanAge := age+2; koreanAge <18{
	// 	return false;
	// }else {
	// 	return true;
	// }

	// go에서의 switch문 
	// switch {
	// 	case age < 18:
	// 		return false;
	// 	case age >=18:
	// 		return true;
	// 	default:
	// 		return false;
	// }

	switch koreanAge:=age+2; koreanAge{
		case 10:
			return false;
		case 18:
			return true;
		default:
			return false;
	}

}

func superAdd(numbers ...int) int {
	// fmt.Println("numbers",numbers)

	total:=0;

	// for와 range를 이용해서 자바스크립트의 for처럼 반복문을 만들 수 있다.
	// range는 기본적으로 index를 준다.(순번을 확인할 수 있음)
	// for index,number := range numbers{ }
	// index를 안 쓸 때는 _로 비워두면 된다.
	for _,number := range numbers{
		total=total+number;
	}

	// 위의 코드를 for문만을 이용해서도 만들 수 있다. 
	// for i:=0; i<len(numbers); i++{
	// 	fmt.Println(numbers[i])
	// }

	return total;
}

func main() {
	result:=superAdd(1,2,3,4,5,6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))

	// a:=2;
	// b:=a;
	// a=10;
	// a와 b앞에 &a, &b를 붙이게 되면 해당 변수의 메모리 주소를 볼 수 있다.
	// fmt.Println(&a,&b)

	// 아래와 같이 a의 메모리 주소를 b에 넣어주게 되면 a와 b는 같은 메모리 주소를 가지게 된다.
	// a:=2;
	// b:=&a;
	// fmt.Println(&a,b)

	// 그리고 *는 메모리 주소를 통해 해당 주소 안에 담긴 값을 살펴볼 수 있다.
	// b:=&a를 통해 a의 메모리 주소를 b에 넣었고, a와 b를 연결했기 떄문에 a의 값이 바뀌면 b의 값도 바뀌게 된다.
	// 즉 b는 a의 메모리를 살펴보고 그 안의 값을 가진다. (b는 a의 포인터가 된다.)
	// *b는 메모리 주소 안의 값에 접근해서 값을 변경할 수도 있다. 
	a:=2;
	b:=&a;
	a=5;
	*b=20;
	fmt.Println(a,b)
}