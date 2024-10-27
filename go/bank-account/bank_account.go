package account

import "sync"

// Account holds information about bank accounts.
type Account struct {
	mutex   *sync.Mutex
	balance int64
	closed  bool
}

// Open handles creating accounts.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		mutex:   &sync.Mutex{},
		balance: initialDeposit,
	}
}

// Close handles closing accounts.
func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return 0, false
	}

	payout = a.balance

	a.balance = 0
	a.closed = true

	return payout, true
}

// Balance handles returning account balances.
func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit handles account deposits and withdrawals.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if a.closed {
		return 0, false
	}

	a.balance += amount
	if a.balance < 0 {
		a.balance -= amount
		return 0, false
	}

	return a.balance, true
}
