package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 45,
		},
		{
			ID:     99,
			Amount: 54,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	// Output:
	// 99
}

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			ID:         0,
			PAN:        "D",
			Balance:    1_000_00,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     true,
			MinBalance: 0,
		},
		{

			ID:         0,
			PAN:        "A",
			Balance:    1_000_00,
			Currency:   "",
			Color:      "",
			Name:       "",
			Active:     true,
			MinBalance: 0,
		},
	}
	results := PaymentSources(cards)
	for _, source := range results {
		fmt.Println(source.Number)
	}
	//Output:
	//D
	//A

}
