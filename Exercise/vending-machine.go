package main

import "fmt"

type VendingMachine struct {
	current int
	coin    map[string]int
	items   map[string]int
}

func NewVendingMachine(x, y map[string]int) *VendingMachine {

	return &VendingMachine{0, x, y}

}

func (v *VendingMachine) InsertCoin(c string) {

	v.current += v.coin[c]

}

func (v *VendingMachine) GetInsertedMoney() int {

	return v.current

}

func (v *VendingMachine) SelectSD() string {

	if v.current >= v.items["SD"] {
		v.current -= v.items["SD"]
		var change string
		if v.current > 0 {
			change = GetChange(v.current)
		}
		if change != "" {
			change = ", " + change
		}
		return "SD" + change
	}
	return "Not Enough Coins"

}

func (v *VendingMachine) SelectCC() string {

	if v.current >= v.items["CC"] {
		v.current -= v.items["CC"]
		var change string
		if v.current > 0 {
			change = GetChange(v.current)
			v.current = 0
		}
		if change != "" {
			change = ", " + change
		}
		return "CC" + change
	}

	return "Not Enough Coins"

}

func (v *VendingMachine) CoinReturn() string {

	result := GetChange(v.current)
	v.current = 0

	return result

}

func GetChange(x int) string {

	value := [4]int{10, 5, 2, 1}
	c := [4]string{"T", "F", "TW", "O"}
	var change string

	for x > 0 {

		for i, v := range value {
			if x >= v {
				if change != "" {
					change += ", "
				}
				change += c[i]
				x -= v
				break
			}
		}

	}

	return change
}

var coins map[string]int
var items map[string]int

func init() {
	coins = make(map[string]int)
	coins["T"] = 10
	coins["F"] = 5
	coins["TW"] = 2
	coins["O"] = 1

	items = make(map[string]int)
	items["SD"] = 18
	items["CC"] = 12

}

func main() {
	vm := NewVendingMachine(coins, items)
	// Buy SD(soft drink) with exact change
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	can := vm.SelectSD()
	fmt.Println(can) // SD

	// Buy CC(canned coffee) without exact change
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	// Start adding change but hit coin return
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F
}
