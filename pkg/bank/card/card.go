package card

import "bank/pkg/bank/types"

const (
	maxWithdraw = 20_000_00
	maxDeposit  = 50_000_00
	minBonus    = 5_000
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	// TODO: произвести операøии с картой
	if card.Active && card.Balance >= amount && amount >= 0 && amount < maxWithdraw {
		card.Balance -= amount
	}
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	// TODO: произвести операøии с картой
	if card.Active && amount <= maxDeposit && amount > 0 {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if card.Active && card.Balance > 0 && card.MinBalance > 0 {
		bonus := card.MinBalance * types.Money(percent*daysInMonth) / types.Money(100*daysInYear)
		if bonus < minBonus {
			card.Balance += bonus
		}
	}
}
func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _, card := range cards {
		sum += (card.Balance)
	}
	return sum
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
