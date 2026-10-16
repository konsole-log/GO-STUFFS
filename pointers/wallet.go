package pointers

import "fmt"

type Wallet struct{
	balance int
}
//By default if there is just (w Wallet) then it is just the copy of the actual parameter from the method calling it but if we use the pointer (w *Wallet) then it is the pointer to the actual value of the parameter calling the method.
func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p\n",&w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int{
	return w.balance
}
