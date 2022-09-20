package main

import "fmt"

type Money struct {
	Amount int
	Currency string
}

func (m Money) Add(money Money) (*Money, error) {
	if money.Currency != m.Currency {
		return nil, fmt.Errorf("通貨単位が違います。")
	}

	added := m.Amount + money.Amount
	return NewMoney(added, money.Currency)
}

func NewMoney(amount int, currency string) (*Money, error) {
	if amount < 0 {
		return nil, fmt.Errorf("金額には0以上を指定してください。")
	}
	if currency == "" {
		return nil, fmt.Errorf("通貨単位を指定してください。")
	}

	return &Money{
		Amount: amount,
		Currency: currency,
	}, nil
}
