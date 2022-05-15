package sample

import (
	"fmt"
	// "math/rand"
	"os"
	"strings"

	"github.com/speecan/moo/game"
)

// EstimateHuman is played by human
func EstimateHuman(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		var input string
		fmt.Print("?: ")
		fmt.Fscanln(os.Stdin, &input)
		guess := game.Str2Int(strings.Split(input, ""))
		fn(guess)
		return guess
	}
}

// EstimateWithRandom is idiot algo.
// returns estimate number with simply random
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}

// EstimateWithRandom2 is idiot algo.
// exclude duplicate queries
func EstimateWithRandom2(difficulty int) game.Estimate {
	query := make([][]int, 0)
	isDuplicated := func(i []int) bool {
		for _, v := range query {
			if game.Equals(v, i) {
				return true
			}
		}
		return false
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if !isDuplicated(r) {
				break
			}
		}
		fn(r)
		query = append(query, r)
		return r
	}
}

func EstimateAuto(difficulty int) game.Estimate {
	index := 0
	count := 0
	pre_hit := 0
	pre_blow := 0
	candidates := game.GetCandidates(difficulty)

	r := make([]int, difficulty)

	max_candidates := game.Parmutation(10, difficulty)
	new_candidates := make([][]int, max_candidates)
	for i := 0; i < max_candidates; i++ {
		new_candidates[i] = make([]int, difficulty)
	}

	return func(fn game.Question) (res []int) {
		// 初回は[0 1 2 3]と入力
		if count == 0 {
			r = candidates[0]
			candidates = candidates[1:]
		} else if count > 0 {
			index = 0
			// 前の解答と同じhitとblowとなる候補を見つけ出す
			for i := 0; i < len(candidates); i++ {
				if game.GetHit(candidates[i], r) == pre_hit && game.GetBlow(candidates[i], r) == pre_blow {
					new_candidates[index] = candidates[i]
					index++
				}
			}
			copy(candidates, new_candidates)
			r = candidates[0]
			candidates = candidates[1:]
		} else {
			fmt.Printf("error: count is %d!", count)
			os.Exit(1)
		}
		pre_hit, pre_blow = fn(r)
		count++
		return r
	}
}
