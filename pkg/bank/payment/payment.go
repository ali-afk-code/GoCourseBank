package payment

import "bank/pkg/bank/types"

func Max(payments []types.Payment) types.Payment {
	// var max = (payments[0].Amount)
	maxPayment := payments[0]
	for _, payment := range payments {
		if maxPayment.Amount < payment.Amount {
			maxPayment = payment
		}
	}
	return maxPayment
}
