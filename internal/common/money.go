package common

type Money struct {
	Amount   int64
	Currency string
}

func NewMoney(amount int64, currency string) Result[*Money] {
	if amount < 0 {
		return Failure[*Money](NewError("money_negative_amount", "amount must be >= 0"))
	}

	if currency == "" {
		return Failure[*Money](NewError("money_currency_required", "currency is required"))
	}

	return Success(&Money{Amount: amount, Currency: currency})
}

func NewUsd(amount int64) Result[*Money] {
	return NewMoney(amount, "USD")
}
