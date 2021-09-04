package main

import (
	"fmt"

	"github.com/githubgw/Go/bank/accounts"
)

func main(){
	// accounts폴더 안에 Account struct을 가져와서 변수 account에 담는다.
	// 그리고 필드에 값을 전달해줘서 하나의 account를 생성할 수 있다. 
	// account:=accounts.Account{Owner:"GW", Balance:1000}

	account:=accounts.NewAccount("GW")

	fmt.Println("account",account)
}