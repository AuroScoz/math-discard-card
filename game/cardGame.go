package game

import (
	"fmt"
	"math-discard-card/card"
	"math/rand"
)

var MyGame *CardGame

type CardGame struct {
	Deck               []*card.Card
	HandCards          []*card.Card
	DeckAvailableDic   map[int]bool
	GameCost           int
	DefaultDiscardCost int
	DiscardAddCost     int
	CurDiscardCount    int
}

func InitCardGame(gameCost, defaultDiscardCost, discardAddCost int) {
	MyGame = &CardGame{
		GameCost:           gameCost,
		DefaultDiscardCost: defaultDiscardCost,
		DiscardAddCost:     discardAddCost,
	}
	MyGame.initDeck()
}

func (g *CardGame) curDiscardCost() int {
	return g.DefaultDiscardCost + (g.CurDiscardCount * g.DiscardAddCost)
}

func (g *CardGame) initDeck() {
	g.Deck = []*card.Card{}
	g.DeckAvailableDic = make(map[int]bool)
	for suit := 0; suit < 4; suit++ {
		for number := 1; number <= 13; number++ {
			card := card.NewCard(card.SuitType(suit), number)
			g.Deck = append(g.Deck, card)
			g.DeckAvailableDic[card.Idx] = true
		}
	}
}

func (g *CardGame) resetDeckAvailableDic() {
	for key := range g.DeckAvailableDic {
		g.DeckAvailableDic[key] = true
	}
}

func (g *CardGame) NewGame(handIdxs ...int) {
	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
	g.CurDiscardCount = 0
	g.resetDeckAvailableDic()
	if len(handIdxs) == 0 {
		g.drawInitialHand()
	} else {
		g.HandCards = []*card.Card{}
		for i := 0; i < 5; i++ {
			if i < len(handIdxs) {
				g.drawCard(handIdxs[i])
			} else {
				g.drawCard(0)
			}
		}
	}
	MyPlayer.AddPt(-g.GameCost)
	log := fmt.Sprintf("新的一局遊戲 花費%v點遊玩 玩家點數: %v", g.GameCost, MyPlayer.Pt)
	println(log)
	g.ShowCards()
}

func (g *CardGame) firstDrawInitialHand() {
	g.HandCards = []*card.Card{}
	// 四條
	g.drawCard(1)
	g.drawCard(14)
	g.drawCard(27)
	g.drawCard(40)
	g.drawCard(15)

}

func (g *CardGame) drawInitialHand() {
	g.HandCards = []*card.Card{}
	for i := 0; i < 7; i++ {
		g.drawCard(0)
	}
}

func (g *CardGame) drawCard(idx int) *card.Card {
	if idx != 0 {
		for i, card := range g.Deck {
			if card.Idx == idx {
				g.HandCards = append(g.HandCards, card)
				g.DeckAvailableDic[card.Idx] = false
				g.Deck = append(g.Deck[:i], g.Deck[i+1:]...)
				return card
			}
		}
		fmt.Printf("牌池無此idx的牌: %d\n", idx)
		return nil
	} else {
		if len(g.Deck) > 0 {
			card := g.Deck[0]
			g.HandCards = append(g.HandCards, card)
			g.DeckAvailableDic[card.Idx] = false
			g.Deck = g.Deck[1:]
			return card
		}
		return nil
	}
}

func (g *CardGame) Settlement() {
	handType := g.GetHandType()
	gainPT := handType.GetOdds()
	MyPlayer.AddPt(gainPT)
	log := fmt.Sprintf("結算牌型: %v  獲得點數: %v   玩家點數: %v", handType.ToString(), gainPT, MyPlayer.Pt)
	println(log)
}

func (g *CardGame) DiscardCard(handIdxs ...int) {
	if len(handIdxs) == 0 {
		fmt.Println("傳入參數錯誤")
		return
	}
	if MyPlayer.Pt < g.curDiscardCost() {
		fmt.Println("點數不夠")
		return
	}
	MyPlayer.AddPt(-g.curDiscardCost())
	costStr := fmt.Sprintf("重抽花費點數%v  玩家點數: %v", g.curDiscardCost(), MyPlayer.Pt)
	fmt.Println(costStr)

	newCards := []*card.Card{}
	for _, handIdx := range handIdxs {
		if handIdx >= 0 && handIdx < len(g.HandCards) {
			if len(g.Deck) > 0 {
				log := fmt.Sprintf("丟棄: %v", g.HandCards[handIdx].ToString())
				newCard := g.Deck[0]
				g.Deck = g.Deck[1:]
				newCards = append(newCards, newCard)
				g.HandCards[handIdx] = newCard
				g.DeckAvailableDic[newCard.Idx] = false
				log += fmt.Sprintf("  抽到: %v", newCard.ToString())
				fmt.Println(log)
			}
		}
	}

	g.CurDiscardCount++
}

func (g *CardGame) GetHandType() card.HandType {
	return card.GetHandType(g.HandCards)
}

func (g *CardGame) ShowCards() {
	cardStr := "手牌: "
	for _, card := range MyGame.HandCards {
		cardStr += fmt.Sprintf("[%v] ", card.ToString())
	}
	cardStr += "   目前牌型: " + g.GetHandType().ToString()
	fmt.Println(cardStr)
}
