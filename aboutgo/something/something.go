package something

import "fmt"

// 함수를 만들 때 대문자로 만든 SayHello는 자동 export되서 main.go에서 가져와서 사용할 수 있지만 아래 소문자로 만든 sayBye는 export되지 않아서 main.go에서 가져올 수 없다.
func SayHello(){
	fmt.Println("Hello! something.go")
}

func sayBye(){
	fmt.Println("Bye! something.go")
}