package game

import (
	"fmt"
)

var (
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// DebugMode makes log output enable
	DebugMode = true
)

type (
	// Question returns hits, blow by guess []int
	Question func(guess []int) (hits, blow int)
	// Estimate returns a next guess
	Estimate func(q Question) (guess []int)
	// Game means one game field
	Game struct {
		difficulty int
		answer     []int
	}
)

// NewGame returns a new game field
func NewGame(d int) *Game {
	if d < 1 || d > 9 {
		fmt.Println(d, "is invalid moo digit, difficulty set to", d)
		d = 4
	}
	return &Game{
		difficulty: d,
		answer:     GetMooNum(d),
	}
}

// GetDifficulty returns digits
func (x *Game) GetDifficulty() int {
	return x.difficulty
}

// GetAnswer returns answer
func (x *Game) GetAnswer() []int {
	return x.answer
}

// GetQuestion returns a question func
func (x *Game) GetQuestion(count *int) Question {
	*count = 0
	return func(g []int) (h, b int) {
		*count++
		h = x.GetHit(g)
		b = x.GetBlow(g)
		if DebugMode {
			fmt.Println(g, ": hits:", h, "blow:", b)
		}
		return
	}
}

// GetHit returns hit count in this game
func (x *Game) GetHit(g []int) int {
	return GetHit(g, x.answer)
}

// GetBlow returns blow count in this game
func (x *Game) GetBlow(g []int) int {
	return GetBlow(g, x.answer)
}

// Equals returns bool which guess = answer
func (x *Game) Equals(g []int) bool {
	return Equals(g, x.answer)
}

// GetHit returns hit
func GetHit(guess []int, answer []int) int {
	count := 0
	if len(guess) != len(answer) {
		return 0
	}
	for i, v := range answer {
		if guess[i] == v {
			count++
		}
	}
	return count
}

// GetBlow returns blow
func GetBlow(guess []int, answer []int) int {
	count := 0
	if len(guess) != len(answer) {
		return 0
	}
	for i, g := range guess {
		for j, a := range answer {
			if g == a && i != j {
				count++
			}
		}
	}
	return count
}

func Parmutation(n, r int) int {
	limit := n - r
	var max_candidates func(int) int
	max_candidates = func(n int) int {
		if n == limit {
			return 1
		}
		return n * max_candidates(n-1)
	}
	return max_candidates(n)
}

func GetCandidates(difficulty int) [][]int {
	var index int
	var flag bool
	var CreateCstimate func(int)
	used_num := make([]int, difficulty-1)
	max_candidates := Parmutation(10, difficulty)

	candidates := make([][]int, max_candidates)
	for i := 0; i < max_candidates; i++ {
		candidates[i] = make([]int, difficulty)
	}

	CreateCstimate = func(current_dif int) {
		for i := 0; i < 10; i++ {
			// 上の位に同じ数値が使われていないか確認
			for j := 0; j < difficulty-current_dif; j++ {
				if i == used_num[j] {
					flag = true
					break
				}
			}
			if flag {
				flag = false
				continue
			}
			if current_dif > 1 {
				candidates[index][difficulty-current_dif] = i
				used_num[difficulty-current_dif] = i
				CreateCstimate(current_dif - 1)
			} else if current_dif == 1 {
				copy(candidates[index], used_num)
				candidates[index][difficulty-1] = i
				index++
			}
		}
	}
	CreateCstimate(difficulty)
	return candidates
}
