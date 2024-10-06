package concurrency

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeadlock(t *testing.T) {
	account1 := &BankAccount{}
	account2 := &BankAccount{}
	wg := sync.WaitGroup{}

	account1.Deposit(100)
	account2.Deposit(100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		account1.TransferWithDeadlock(account2, 50)
	}()

	go func() {
		defer wg.Done()
		account2.TransferWithDeadlock(account1, 50)
	}()

	// This test will likely deadlock and never complete
	wg.Wait()
	assert.Equal(t, 100, account1.Balance())
	assert.Equal(t, 100, account2.Balance())
}

func TestDeadlockWarn(t *testing.T) {
	account1 := &BankAccount{}
	account2 := &BankAccount{}
	wg := sync.WaitGroup{}

	account1.Deposit(100)
	account2.Deposit(100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		account1.Transfer(account2, 50)
	}()

	go func() {
		defer wg.Done()
		account2.Transfer(account1, 50)
	}()

	wg.Wait()
	assert.Equal(t, 100, account1.Balance())
	assert.Equal(t, 100, account2.Balance())
}
