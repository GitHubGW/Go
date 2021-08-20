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

	"github.com/githubgw/learngo/something"
)

// 또한 go는 main.go를 실행하면서 특정 함수들을 찾는다.
// 그래서 함수를 반드시 생성해줘야 한다. 현재 이 함수가 go프로그램의 시작점이 되는 부분이다.
func main() {
	// 자바스크립트에서는 해당 함수를 내보내려 할 때 export를 사용하지만 go에서는 함수를 대문자로 시작해준다. 
	// 즉 Println()함수는 이미 만들어져서 export되었기 때문에 사용할 수 있는 것이다. 
	fmt.Println("Hello GO!")

	// something패키지 안에 있는 SayHello함수를 실행한다. 
	// go.mod file not found in current directory or any parent directory; see 'go help modules'
	// 아래 코드를 같이 실행헀을 때 위와 같은 오류가 나면 go mod init을 통해 go.mod파일을 초기화해주면 된다.
	something.SayHello()
}