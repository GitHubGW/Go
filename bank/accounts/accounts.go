package accounts

// Account struct: Account struct을 생성하고 그 안에 타입을 지정해준다.
// 단, 주의할점은 이 struct을 export해주기 위해서 이름 시작을 대문자로 지정해준다.
// struct내부의 필드 또한 대문자로 지정해줘야 private하지 않게 되서 접근가능하다.
// 대문자=public, 소문자=private
type Account struct {
	// Owner string
	// Balance int
	
	owner string
	balance int
}

// NewAccount함수를 생성하고 export하기위해서 대문자로 만들어 줬다.
// 그리고 NewAccount함수가 리턴하는 값은 *Account이다.
// 즉 &account가 리턴하는 것의 타입이 *Account(Account내부를 들여다 봄)인 것이다.
// NewAccount함수를 통해 객체를 리턴시킨 것이다. 실제 메모리 주소를 리턴한 것이다. 
func NewAccount(owner string) *Account {
	// Account struct을 가져와서 각각의 필드에 값을 넣은 결과값을 account변수에 받는다.
	// 그리고 &account을 통해 account의 메모리주소를 리턴해준다. 
	// account가 아닌 &를 붙여서 메모리 주소로 리턴해준 것이다.
	account:=Account{owner: owner, balance:0}
	return &account;
}