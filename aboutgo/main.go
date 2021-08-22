// 만약 사용자가 프로젝트를 컴파일하고 싶다면 main.go파일을 만들고 그 안에 main package를 만들어 줘야 한다.
// (main은 오로지 컴파일을 위해 필요한 것임)
// go컴파일러는 패키지 이름이 main인 것을 찾고 그 안에 main함수를 찾아서 컴파일해준다.
// go에서는 package main 구문을 통해 어떤 패키지를 사용하는지 작성해주어야 한다.
package main

// fmt는 포맷팅을 위한 go가 기본적으로 가지고 있는 패키지중 하나로 import를 통해 가져온 것이다.
// import "fmt"
// import "github.com/githubgw/learngo/something"는 go폴더 아래에 있는 해당 폴더 경로에서 import해온다는 의미이다.
import (
	"fmt"
	"strings"
	// "github.com/githubgw/learngo/something"
)

// 만약 함수의 파라미터로 여러 값을 받고 싶다면 아래와 같이 ...을 붙여주면 된다.
func repeatMe(words ...string){
	fmt.Println(words)
}

// lengthAndUppercase함수를 또 다르게 바꿔서 짤 수도 있다.
// (length int, uppercase string)를 통해 이 함수가 리턴할 값이 어떤 변수인지 알려주게 되면 굳이 return에 쓰지 않아도 알아서 Go가 어떤 데이터를 리턴할 것인지 찾는다.
func lengthAndUppercase2(name string) (length int, uppercase string){
	// defer를 이용하면 현재 이 함수가 끝나고 난 후 무언가를 실행시킬 수 있다.
	defer fmt.Println("lengthAndUppercase2 done!")

	length = len(name)
	uppercase = strings.ToUpper(name);
	return
}

// go에서는 특이하게 여러 개의 값을 리턴할 수 있다.
// 그래서 lengthAndUppercase함수는 name 값을 받아서 len()함수를 통해 그 name의 길이를 리턴하고, strings.ToUpper()을 통해 name을 대문자로 리턴한다.
// 즉, name의 문자열의 길이와 해당 문자열을 대문자로 처리해서 두 개를 리턴하는 것이다.
func lengthAndUppercase(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

// go에서는 함수에 파라미터와 리턴하는 값의 타입 또한 알려줘야 한다.
// 그래서 a와 b가 받는 데이터의 타입이 int라고 알려주었고 이 함수가 리턴하는 값이 int라고 알려주었다.
// a int, b int를 a, b int로 작성할 수도 있다.
func multiply(a int, b int) int{
	return a*b;
}

// 또한 go는 main.go를 실행하면서 특정 함수들을 찾는다.
// 그래서 함수를 반드시 생성해줘야 한다. 현재 이 함수가 go프로그램의 시작점이 되는 부분이다.
func main() {
	// 자바스크립트에서는 해당 함수를 내보내려 할 때 export를 사용하지만 go에서는 함수를 대문자로 시작해준다. 
	// 즉 Println()함수는 이미 만들어져서 export되었기 때문에 사용할 수 있는 것이다. 
	// fmt.Println("Hello GO!")

	// something패키지 안에 있는 SayHello함수를 실행한다. 
	// go.mod file not found in current directory or any parent directory; see 'go help modules'
	// 아래 코드를 같이 실행헀을 때 위와 같은 오류가 나면 go mod init을 통해 go.mod파일을 초기화해주면 된다.
	// something.SayHello()

	// go에서도 자바스크립트처럼 앞에 const와 var를 붙여서 변수를 만들수 있다. 
	// 단 타입스크립트처럼 뒤에 타입을 지정해줘야 한다.
	// 그리고 마찬가지로 const는 변수의 값을 재할당할 수 없다.
	// const name string = "gw"
	// name="gn";

	// var name2 string = "cody"
	// name2 = "sugar"

	// 위의 코드를 아래와 같이 축약해서 작성할 수도 있다. 단, 축약형은 func인 함수내부에서만 사용가능하다. 
	// 함수 밖에서는 원래 변수를 만드는 형태로 만들 수 있다.
	// 만약 아래와 같이 작성하면 Go가 알아서 값에 대한 타입을 찾아준다.
	// 또한 처음 초기화 할 때 정해진 타입은 나중에 변수에 다른 타입의 값을 재할당해서 바꿀 수 없다. 
	// (처음에 string으로 만들었다면 변수에 다른 값을 재할당할 때도 타입이 string이어야 함)
	name3 := "orange"
	// name3=true 위에서 name3의 타입을 string으로 초기화했기 때문에 boolean인 true로 변경할 수 없다.

	// fmt.Println(name)
	// fmt.Println(name2)
	fmt.Println(name3)

	fmt.Println(multiply(2,2))

	// totalLength, upperName := lengthAndUppercase("GitHubGW")를 통해 lengthAndUppercase()함수를 실행하고 그 함수가 리턴한 2개의 값을 각각의 변수로 받는다.
	// 만약 totalLength, upperName 2개가 아닌 하나의 값만 받고 싶다면 totalLength, _ 로 해주면 된다. (_로 주게 되면 Go는 해당 value를 무시한다.)
	totalLength, upperName := lengthAndUppercase("GitHubGW")
	fmt.Println(totalLength, upperName)

	repeatMe("Tom", "Cody", "Conor", "Dana")

	// fmt.Println(lengthAndUppercase2("Colby"))
	lengthAndUppercase2("Colby")
}