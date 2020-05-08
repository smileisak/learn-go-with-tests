package pointers

import "fmt"

// Stringer is a interface
type Stringer interface {
	String() string
}

// Bitcoin is a new revolutionnary type
type Bitcoin int

// Wallet struct to store wallet infromations
type Wallet struct {
	balance Bitcoin
}

// Deposit deposit to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance return the balance of the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw withdraw from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
