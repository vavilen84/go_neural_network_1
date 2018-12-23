package main

import (
	"fmt"
	"github.com/fxsjy/gonn/gonn"
)

func CreateNN() {
	nn := gonn.DefaultNetwork(3, 16, 4, false)
	input := [][]float64{
		[]float64{0.3, 1, 1}, []float64{0.6, 1, 1}, []float64{0.9, 1, 1},
	}
	target := [][]float64{
		[]float64{1, 0, 0, 0}, []float64{0, 1, 0, 0}, []float64{0, 0, 1, 0},
	}
	nn.Train(input, target, 100000)
	gonn.DumpNN("gonn", nn)
}

func GetResult(output []float64) string {
	var max float64 = -99999
	pos := -1
	for i, value := range output {
		if value > max {
			max = value
			pos = i
		}
	}
	switch pos {
	case 0:
		return "Attack"
	case 1:
		return "Steal"
	case 2:
		return "Run away"
	case 3:
		return "Do nothing"
	}

	return ""
}

func main() {
	CreateNN()
	nn := gonn.LoadNN("gonn")
	var hp float64 = 0.76
	var weapon float64 = 1.0
	var enemyCount float64 = 1.0
	out := nn.Forward([]float64{hp, weapon, enemyCount})
	fmt.Println(GetResult(out))
}
