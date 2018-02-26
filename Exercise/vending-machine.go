package main

import (
	"fmt"
	"sort"
)

type vendingMachine struct {
	current     int
	coin        map[string]int
	items       map[string]int
	changecoins map[int]string
	value       []int
	insertedcoin		[]string
}

func newVendingMachine(x, y map[string]int) *vendingMachine {

	value := make([]int, 0)
	var changecoins map[int]string
	changecoins = make(map[int]string)
	for k, v := range coins {
		changecoins[v] = k
		value = append(value, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(value)))
	return &vendingMachine{0, x, y, changecoins, value,make([]string,0)}

}

func (v *vendingMachine) InsertCoin(c string) {

	v.current += v.coin[c]
	v.insertedcoin = append(v.insertedcoin,c)

}

func (v *vendingMachine) GetInsertedMoney() int {

	return v.current

}

func (v *vendingMachine) SelectSD() string {

	if v.current >= v.items["SD"] {
		v.insertedcoin = make([]string,0)
		v.current -= v.items["SD"]
		if v.current > 0 {
			return "SD" + v.getchange()
		}
		return "SD"
	}	
	return "Not Enough Coins"
}

func (v *vendingMachine) SelectCC() string {

	if v.current >= v.items["CC"] {
		v.insertedcoin = make([]string,0)
		v.current -= v.items["CC"]
		if v.current > 0 {
			return "CC" + v.getchange()
		}
		return "CC"
	}	
	return "Not Enough Coins"

}

func (v *vendingMachine) getchange() string {
		var change string
	for _, x := range v.value {
		for v.current >= x {
			change += ", " + v.changecoins[x]
			v.current -= x
		}
	}
	return change
}

func (v *vendingMachine) CoinReturn() string {

	var change string
	
	for _, x := range v.insertedcoin {
			change += ", " + x			
	}
	v.insertedcoin = make([]string,0)
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
	items["DW"] = 7

}

func main() {

	vm := newVendingMachine(coins, items)

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