package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Game_model struct {
	deck  []string
	hand1 [6]string
	hand2 [6]string
	hand3 [6]string
	table [4]string
}

func min_card(cards [6]string) int {
	var index int
	min_num := 20
	for i := 0; i < 6; i++ {
		num, err := strconv.Atoi(cards[i][:len(cards[i])-1])
		if err != nil {
			fmt.Println("smt goes wrong|min_card", num)
			return 0
		}
		if num < min_num {
			min_num = num
			index = i
		}

	}
	return index
}

func comp_cards(card1 string, card2 string) bool {
	if card1[:len(card1)-1] == card2[:len(card2)-1] {
		return true
	}
	return false
}

func (gm Game_model) print_deck() {
	for i, s := range gm.deck {
		fmt.Println(i, s)
	}
}
func (gm Game_model) print_hands() {
	for i := 0; i < 6; i++ {
		fmt.Println(gm.hand1[i], gm.hand2[i], gm.hand3[i])
	}
}

func (gm *Game_model) begin() {
	for i := 0; i < 6*3; i += 3 {
		gm.hand1[i/3] = gm.deck[i]
		gm.hand2[i/3] = gm.deck[i+1]
		gm.hand3[i/3] = gm.deck[i+2]
	}
}

func (gm *Game_model) Init(isLargeDeck bool) {
	var min_card int
	if isLargeDeck {
		min_card = 2
	} else {
		min_card = 6
	}

	for i := min_card; i < 15; i++ {
		gm.deck = append(gm.deck, strconv.Itoa(i)+"a")
		gm.deck = append(gm.deck, strconv.Itoa(i)+"b")
		gm.deck = append(gm.deck, strconv.Itoa(i)+"c")
		gm.deck = append(gm.deck, strconv.Itoa(i)+"d")
	}
	rand.Shuffle(len(gm.deck), func(i, j int) { gm.deck[i], gm.deck[j] = gm.deck[j], gm.deck[i] })

}

func (gm *Game_model) Game_sim() int {
	ans := 0
	gm.table[0] = gm.hand1[min_card(gm.hand1)]
	gm.hand1[min_card(gm.hand1)] = "18n"
	for i := 0; i < 6; i++ {
		if comp_cards(gm.hand2[i], gm.table[0]) {
			ans++
			gm.table[1] = gm.hand2[i]
			gm.hand2[i] = "18n"
			break
		}
	}
	if ans == 0 {
		return ans
	}

	for i := 0; i < 6; i++ {
		if comp_cards(gm.hand3[i], gm.table[0]) {
			ans++
			gm.table[2] = gm.hand3[i]
			gm.hand3[i] = "18n"
			break
		}
	}
	if ans == 1 {
		return ans
	}

	for i := 0; i < 6; i++ {
		if comp_cards(gm.hand1[i], gm.table[0]) {
			ans++
			gm.table[3] = gm.hand1[i]
			gm.hand1[i] = "18n"
			break
		}
	}

	return ans
}

func sim(num int, ans *[]int, is52 bool) {
	var gm Game_model
	gm.Init(is52)
	gm.begin()
	(*ans)[num] = gm.Game_sim()

}

func main() {
	n := 1000000
	res := make([]int, n)
	var ans []float64 = []float64{0, 0, 0, 0}
	is52 := true

	for i := 0; i < n; i++ {
		go sim(i, &res, is52)
	}

	for i := 0; i < n; i++ {
		switch res[i] {
		case 0:
			ans[0] += 1
		case 1:
			ans[1] += 1
		case 2:
			ans[2] += 1
		case 3:
			ans[3] += 1
		}
	}
	fmt.Println("Проведено игр: ", n)
	fmt.Println("52 карты:", is52)
	fmt.Printf("0:\t%f\n1:\t%f\n2:\t%f\n3:\t%f\n", ans[0] / float64(n), ans[1] / float64(n), ans[2] / float64(n), ans[3] / float64(n))

}
