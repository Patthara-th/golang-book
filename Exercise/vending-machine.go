package main

import (
	"fmt"
	"github.com/Patthara-th/vendingMachine"
)

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
	items["DW"] = 7

}

func main() {

	vm := vendingMachine.NewVendingMachine(coins, items)

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
	vm.InsertCoin("TW")
	vm.InsertCoin("F")
	vm.InsertCoin("O")
	vm.InsertCoin("T")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	change := vm.CoinReturn()
	fmt.Println(change) // T, T, F

}
