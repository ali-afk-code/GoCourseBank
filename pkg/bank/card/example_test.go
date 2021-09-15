package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// func ExampleWithdraw_positive() {
// 	res := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 1000000
// }

// func ExampleWithdraw_noMoney() {
// 	res := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 0
// }

// func ExampleWithdraw_inactive() {
// 	res := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)

// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000
// }

// func ExampleWithdraw_limit() {
// 	res := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 200000000)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000
// }
// func ExampleDeposit_inactive() {
// 	res := types.Card{Balance: 20_000_00, Active: false}
// 	Deposit(&res, 10_000_00)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000

// }

// func ExampleDeposit_limit() {
// 	res := types.Card{Balance: 20_000_00, Active: true}
// 	Deposit(&res, 51_000_00)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000
// }
// func ExampleDeposit() {
// 	res := types.Card{Balance: 20_000_00, Active: true}
// 	Deposit(&res, 10_000_00)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 3000000
// }

// func ExampleAddBonus() {
// 	res := types.Card{Balance: 20_000_00, Active: true, MinBalance: 10_000}
// 	AddBonus(&res, 3, 30, 365)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000024
// }

// func ExampleAddBonus_inactive() {
// 	res := types.Card{Balance: 20_000_00, Active: false, MinBalance: 10_000}
// 	AddBonus(&res, 3, 30, 365)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000
// }

// func ExampleAddBonus_negativeBalance() {
// 	res := types.Card{Balance: -50, Active: true, MinBalance: 10_000}
// 	AddBonus(&res, 3, 30, 365)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// -50
// }

// func ExampleAddBonus_limitExceeded() {
// 	res := types.Card{Balance: 20_000_00, Active: true, MinBalance: 200000000}
// 	AddBonus(&res, 3, 30, 365)
// 	fmt.Println(res.Balance)
// 	// Output:
// 	// 2000000
// }
// func ExampleTotal() {
// 	cards := []types.Card{
// 		{
// 			Balance: 10_000_00,
// 			Active:  true,
// 		},
// 		{
// 			ID:         0,
// 			PAN:        "",
// 			Balance:    10_000_00,
// 			Currency:   "",
// 			Color:      "",
// 			Name:       "",
// 			Active:     true,
// 			MinBalance: 0,
// 		},
// 	}
// 	fmt.Println(Total(cards))
// 	// Output:
// 	// 2000000

// }

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
