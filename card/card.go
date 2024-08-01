package card

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SuitType int

const (
	Clubs SuitType = iota
	Diamonds
	Hearts
	Spades
)

func (s SuitType) ToString() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return "尚未定義"
	}
}

type Card struct {
	Idx    int
	Suit   SuitType
	Number int
}

func NewCard(suit SuitType, number int) *Card {
	return &Card{
		Suit:   suit,
		Number: number,
		Idx:    int(suit)*13 + number,
	}
}

func (c *Card) ToString() string {
	return fmt.Sprintf("%s%d", c.Suit.ToString(), c.Number)
}

type HandType int

const (
	HighCard HandType = iota
	Pair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

func (h HandType) ToString() string {
	switch h {
	case HighCard:
		return "高牌"
	case Pair:
		return "對子"
	case ThreeOfAKind:
		return "三條"
	case Straight:
		return "順子"
	case Flush:
		return "同花"
	case FullHouse:
		return "葫蘆"
	case FourOfAKind:
		return "四條"
	case StraightFlush:
		return "同花順"
	default:
		return "尚未定義"
	}
}

func (h HandType) GetOdds() int {
	switch h {
	case HighCard:
		return 0
	case Pair:
		return 2
	case ThreeOfAKind:
		return 10
	case Straight:
		return 20
	case Flush:
		return 30
	case FullHouse:
		return 50
	case FourOfAKind:
		return 250
	case StraightFlush:
		return 1000
	default:
		log.Errorf("尚未定義的HandType牌型賠率: %d", h)
		return 0
	}
}

func ShowCards(cards []*Card) {
	var str strings.Builder
	for i, card := range cards {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(card.ToString())
	}
	log.Println(str.String())
}

func GetHandType(cards []*Card) HandType {
	if IsStraightFlush(cards) {
		return StraightFlush
	} else if IsFourOfAKind(cards) {
		return FourOfAKind
	} else if IsFullHouse(cards) {
		return FullHouse
	} else if IsThreeOfAKind(cards) {
		return ThreeOfAKind
	} else if IsStraight(cards) {
		return Straight
	} else if IsFlush(cards) {
		return Flush
	} else if IsPair(cards) {
		return Pair
	} else {
		return HighCard
	}
}

func IsFlush(cards []*Card) bool {
	suitDics := make(map[SuitType]int)
	for _, card := range cards {
		suitDics[card.Suit]++
		if suitDics[card.Suit] >= 5 {
			return true
		}
	}
	return false
}

func GetFlushIndices(cards []*Card) []int {
	suitIndices := make(map[SuitType][]int)
	for i, card := range cards {
		suitIndices[card.Suit] = append(suitIndices[card.Suit], i)
		if len(suitIndices[card.Suit]) >= 5 {
			return suitIndices[card.Suit][:5]
		}
	}
	return nil
}

func IsThreeOfAKind(cards []*Card) bool {
	numberDic := make(map[int]int)
	for _, card := range cards {
		numberDic[card.Number]++
		if numberDic[card.Number] >= 3 {
			return true
		}
	}
	return false
}

func GetThreeOfAKindIndices(cards []*Card) []int {
	numberIndices := make(map[int][]int)
	for i, card := range cards {
		numberIndices[card.Number] = append(numberIndices[card.Number], i)
		if len(numberIndices[card.Number]) >= 3 {
			return numberIndices[card.Number][:3]
		}
	}
	return nil
}

func IsPair(cards []*Card) bool {
	numberDic := make(map[int]int)
	for _, card := range cards {
		numberDic[card.Number]++
		if numberDic[card.Number] >= 2 {
			return true
		}
	}
	return false
}

func GetPairIndices(cards []*Card) []int {
	numberIndices := make(map[int][]int)
	for i, card := range cards {
		numberIndices[card.Number] = append(numberIndices[card.Number], i)
		if len(numberIndices[card.Number]) >= 2 {
			return numberIndices[card.Number][:2]
		}
	}
	return nil
}

func IsFullHouse(cards []*Card) bool {
	numberDic := make(map[int]int)
	for _, card := range cards {
		numberDic[card.Number]++
	}

	threeOfAKindCount := 0
	pairCount := 0

	for _, count := range numberDic {
		if count >= 3 {
			threeOfAKindCount++
		}
		if count >= 2 {
			pairCount++
		}
	}

	return (threeOfAKindCount >= 1 && pairCount >= 2) || (threeOfAKindCount > 1)
}

func GetFullHouseIndices(cards []*Card) []int {
	numberIndices := make(map[int][]int)
	for i, card := range cards {
		numberIndices[card.Number] = append(numberIndices[card.Number], i)
	}

	var threeOfAKindIndices, pairIndices []int
	for _, indices := range numberIndices {
		if len(indices) >= 3 && threeOfAKindIndices == nil {
			threeOfAKindIndices = indices[:3]
		} else if len(indices) >= 2 && pairIndices == nil {
			pairIndices = indices[:2]
		}
	}

	if threeOfAKindIndices != nil && pairIndices != nil {
		return append(threeOfAKindIndices, pairIndices...)
	}

	for _, indices := range numberIndices {
		if len(indices) >= 3 {
			if threeOfAKindIndices == nil {
				threeOfAKindIndices = indices[:3]
			} else {
				return append(threeOfAKindIndices, indices[:2]...)
			}
		}
	}

	return nil
}

func IsFourOfAKind(cards []*Card) bool {
	rankCount := make(map[int]int)
	for _, card := range cards {
		rankCount[card.Number]++
		if rankCount[card.Number] == 4 {
			return true
		}
	}
	return false
}

func GetFourOfAKindIndices(cards []*Card) []int {
	rankCount := make(map[int][]int)
	for i, card := range cards {
		rankCount[card.Number] = append(rankCount[card.Number], i)
		if len(rankCount[card.Number]) == 4 {
			return rankCount[card.Number]
		}
	}
	return nil
}

func IsStraightFlush(cards []*Card) bool {
	suitCards := make(map[SuitType][]*Card)
	for _, card := range cards {
		suitCards[card.Suit] = append(suitCards[card.Suit], card)
	}

	for _, suit := range suitCards {
		if len(suit) < 5 {
			continue
		}
		sortedCards := suit
		sort.Slice(sortedCards, func(i, j int) bool {
			return sortedCards[i].Number < sortedCards[j].Number
		})
		if IsStraight(sortedCards) {
			return true
		}
	}
	return false
}

func IsStraight(cards []*Card) bool {
	values := make(map[int]bool)
	for _, card := range cards {
		values[card.Number] = true
	}

	var nums []int
	for k := range values {
		nums = append(nums, k)
	}
	sort.Ints(nums)

	for i := 0; i <= len(nums)-5; i++ {
		if nums[i+4]-nums[i] == 4 {
			return true
		}
	}

	if values[1] && values[10] && values[11] && values[12] && values[13] {
		return true
	}

	return false
}

func GetStraightFlushIndices(cards []*Card) []int {
	suitIndices := make(map[SuitType][]*Card)
	for _, card := range cards {
		suitIndices[card.Suit] = append(suitIndices[card.Suit], card)
	}

	for _, suitCards := range suitIndices {
		if len(suitCards) < 5 {
			continue
		}
		sort.Slice(suitCards, func(i, j int) bool {
			return suitCards[i].Number < suitCards[j].Number
		})
		straightIndices := GetStraightIndices(suitCards)
		if len(straightIndices) > 0 {
			return straightIndices
		}
	}
	return nil
}

func GetStraightIndices(cards []*Card) []int {
	values := make(map[int]int)
	for i, card := range cards {
		values[card.Number] = i
	}

	var nums []int
	for num := range values {
		nums = append(nums, num)
	}
	sort.Ints(nums)

	for i := 0; i <= len(nums)-5; i++ {
		if nums[i+4]-nums[i] == 4 {
			return []int{values[nums[i]], values[nums[i+1]], values[nums[i+2]], values[nums[i+3]], values[nums[i+4]]}
		}
	}

	if values[1] != 0 && values[10] != 0 && values[11] != 0 && values[12] != 0 && values[13] != 0 {
		return []int{values[1], values[10], values[11], values[12], values[13]}
	}

	return nil
}
