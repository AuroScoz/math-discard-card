package main

import (
	"bufio"
	"fmt"
	"math-discard-card/game"
	"os"
	"strconv"
	"strings"
)

func main() {

	test()
	return

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("============指令清單============ \n1. reset(重置遊戲), \n2. play(開始遊戲), \n3. d-0,2(換第1與第3張手牌)")
	fmt.Println()
	resetGame()
	for {

		fmt.Println()
		fmt.Print("請輸入指令: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, "-")
		if len(parts) < 1 {
			fmt.Println("輸入錯誤")
			continue
		}

		switch parts[0] {
		case "reset":
			resetGame()
		case "play":
			game.MyGame.Settlement()
			game.MyGame.NewGame()
		case "d":
			if len(parts) < 2 {
				fmt.Println("要輸入想替換的手牌索引")
				continue
			}
			idxStrs := strings.Split(parts[1], ",")
			idxs := []int{}
			for _, idxStr := range idxStrs {
				idx, err := strconv.Atoi(idxStr)
				if err != nil {
					fmt.Println("索引輸入錯誤:", idxStr)
					continue
				}
				idxs = append(idxs, idx)
			}
			game.MyGame.DiscardCard(idxs...)
			game.MyGame.ShowCards()
		default:
			fmt.Println("輸入錯誤")
		}
	}

}
func resetGame() {
	fmt.Println("重置遊戲")
	game.NewPlayer(100)
	game.InitCardGame(10, 1, 1)
	game.MyGame.NewGame()
	fmt.Println()
}
