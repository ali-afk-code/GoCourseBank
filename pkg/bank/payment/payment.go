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

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			paySource := types.PaymentSource{
				Type:    "card",
				Number:  string(card.PAN),
				Balance: card.Balance,
			}
			paymentSources = append(paymentSources, paySource)
		}
	}
	return paymentSources
}
