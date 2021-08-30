// #1.1 package and import

package main

// fmt는 포맷팅을 위한 go가 기본적으로 가지고 있는 패키지 중 하나로 import를 통해 가져온 것이다.
// import "github.com/githubgw/learngo/something"는 go폴더 아래에 있는 해당 폴더 경로에서 import해온다는 의미이다.
import (
	"fmt"

	"github.com/githubgw/learngo/something"
)


func main() {
	// 자바스크립트에서는 해당 함수를 내보내려 할 때 export를 사용하지만 go에서는 함수를 대문자로 시작해준다. 
	// 즉 Println()함수는 이미 만들어져서 export되었기 때문에 사용할 수 있는 것이다. 
	fmt.Println("Hello GO!")

	// something패키지 안에 있는 SayHello함수를 실행한다. 
	// go.mod file not found in current directory or any parent directory; see 'go help modules'
	// 아래 코드를 같이 실행헀을 때 위와 같은 오류가 나면 go mod init을 통해 go.mod파일을 초기화해주면 된다.
	something.SayHello()

}