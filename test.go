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

	card1 := card.NewCard(card.Clubs, 5)
	card2 := card.NewCard(card.Clubs, 6)
	card3 := card.NewCard(card.Hearts, 12)
	card4 := card.NewCard(card.Diamonds, 4)
	card5 := card.NewCard(card.Spades, 10)

	cardIdxs := []int{card1.Idx, card2.Idx, card3.Idx, card4.Idx, card5.Idx}
	game.MyGame.NewGame(cardIdxs...)

	discardCount := 3
	combinations := getHandCombinations(game.MyGame.Deck, discardCount)

	targetType := card.Pair

	var targetTypeCount int
	for _, combo := range combinations {
		hand := []*card.Card{
			card1,
			card2,
		}
		hand = append(hand, combo...)
		// if card.IsOnlyHighCard(hand) {
		// 	targetTypeCount++
		// }
		if card.IsHandType(hand, targetType) && !card.IsHandType(hand, card.FullHouse) && !card.IsHandType(hand, card.ThreeOfAKind) {
			targetTypeCount++
		}
	}

	totalCombinations := len(combinations)
	probability := float64(targetTypeCount) / float64(totalCombinations)

	logrus.Infof("換%v張", discardCount)
	logrus.Printf("總組合數: %d", totalCombinations)
	logrus.Printf("%s有幾總組合: %d", targetType.ToString(), targetTypeCount)
	logrus.Printf("%s出線機率: %.10f", targetType.ToString(), probability)
}
