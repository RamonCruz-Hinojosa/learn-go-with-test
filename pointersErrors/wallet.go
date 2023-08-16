package pointersErrors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

//pointers to structs (*Wallet is the pointer here and without it Wallet is just a copy and not changing the real values) are automatically dereferenced
func (w *Wallet) Balance() int {
	return w.balance
}
