package main

import (
	"fmt"
	"sort"
)

type VendingMachine struct {
	current     int
	coin        map[string]int
	items       map[string]int
	changecoins map[int]string
	value       []int
}

func NewVendingMachine(x, y map[string]int) *VendingMachine {

	value := make([]int, 0)
	var changecoins map[int]string
	changecoins = make(map[int]string)
	for k, v := range coins {
		changecoins[v] = k
		value = append(value, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(value)))
	return &VendingMachine{0, x, y, changecoins, value}

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
		if v.current > 0 {
			return "SD" + ", " + v.CoinReturn()
		}
		return "SD"
	}
	return "Not Enough Coins"
}

func (v *VendingMachine) SelectCC() string {

	if v.current >= v.items["CC"] {
		v.current -= v.items["CC"]
		if v.current > 0 {
			return "CC" + ", " + v.CoinReturn()
		}
		return "CC"
	}
	return "Not Enough Coins"

}

func (v *VendingMachine) CoinReturn() string {

	var change string
	for _, x := range v.value {
		for v.current >= x {
			change += ", " + v.changecoins[x]
			v.current -= x
		}
	}
	return change[2:len(change)]

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