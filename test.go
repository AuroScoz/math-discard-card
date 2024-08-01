package main

import (
	"github.com/sirupsen/logrus"
	"math-discard-card/card"
	"math-discard-card/game"
)

func getHandCombinations(cards []*card.Card, n int) [][]*card.Card {
	var combinations [][]*card.Card
	var helper func([]*card.Card, int, int)
	temp := make([]*card.Card, n)
	helper = func(arr []*card.Card, start int, k int) {
		if k == 0 {
			combo := make([]*card.Card, n)
			copy(combo, temp)
			combinations = append(combinations, combo)
			return
		}
		for i := start; i <= len(arr)-k; i++ {
			temp[len(temp)-k] = arr[i]
			helper(arr, i+1, k-1)
		}
	}
	helper(cards, 0, n)
	return combinations
}

func test() {
	game.NewPlayer(100)
	game.InitCardGame(10, 1, 1)

	card1 := card.NewCard(card.Clubs, 1)
	card2 := card.NewCard(card.Diamonds, 1)
	card3 := card.NewCard(card.Hearts, 1)
	card4 := card.NewCard(card.Diamonds, 4)
	card5 := card.NewCard(card.Hearts, 10)

	cardIdxs := []int{card1.Idx, card2.Idx, card3.Idx, card4.Idx, card5.Idx}
	game.MyGame.NewGame(cardIdxs...)

	logrus.Infof("total: %v", len(game.MyGame.Deck))

	combinations := getHandCombinations(game.MyGame.Deck, 2)

	var fullHouseCount int
	for _, combo := range combinations {
		hand := []*card.Card{
			card1,
			card2,
			card3,
		}
		hand = append(hand, combo...)
		if card.IsFullHouse(hand) {
			fullHouseCount++
		}
	}

	totalCombinations := len(combinations)
	probability := float64(fullHouseCount) / float64(totalCombinations)

	logrus.Printf("Total Combinations: %d", totalCombinations)
	logrus.Printf("Combinations: %d", fullHouseCount)
	logrus.Printf("Probability : %.4f", probability)
}
