package main

import (
	"fmt"
	"github.com/ceub/avaliacao-1/bonus"
	"github.com/ceub/avaliacao-1/q1"
	"github.com/ceub/avaliacao-1/q2"
	"github.com/ceub/avaliacao-1/q3"
	"github.com/ceub/avaliacao-1/q4"
	"github.com/ceub/avaliacao-1/q5"
)

func main() {
	possible, err := q1.DivideWatermelon(100)
	fmt.Printf("Q1:\tpossible: %v,\terr: %v\n", possible, err)

	solved := q2.ProblemsSolved([][3]bool{{true, true, true}, {false, false, false}, {true, false, true}})
	fmt.Printf("Q2:\tsolved: %v\n", solved)

	pieces, err := q3.DominoPieces(5, 5)
	fmt.Printf("Q3:\tpieces: %v,\terr: %v\n", pieces, err)

	order, err := q4.ClassifyPrices([]int{5, 4, 3, 2, 1})
	fmt.Printf("Q4:\torder: %v,\terr: %v\n", order, err)

	processString := q5.ProcessString("CEUB")
	fmt.Printf("Q5:\tprocessString: %v\n", processString)

	maxHeight, towers := bonus.CalculateTowers([]int{112, 277, 170, 247, 252, 115, 157, 293, 256, 143, 196, 90, 12, 164, 164, 42, 8, 223, 167, 109, 175, 232, 239, 111, 148, 51, 9, 254, 93, 32, 268, 162, 231, 91, 47, 162, 161, 191, 195, 145, 247, 292, 129, 199, 230, 94, 144, 217, 18, 205, 176, 20, 143, 198, 121, 243, 211, 262, 230, 277, 195, 255, 108, 290, 220, 275, 158, 2, 286, 200, 60, 267, 278, 207, 123, 150, 123, 116, 131, 13, 12, 226, 33, 244, 30, 275, 263, 45, 158, 192, 254, 149, 242, 176, 62, 224, 221, 288, 250, 160, 155, 225, 132, 143, 276, 293, 218, 145, 197, 175, 33, 129, 79, 206, 210, 192, 222, 262, 190, 52, 274, 243, 233})
	fmt.Printf("Bonus:\tmaxHeight: %v,\ttowers: %v", maxHeight, towers)
}
