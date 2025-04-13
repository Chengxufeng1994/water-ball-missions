package main

import (
	"fmt"

	"github.com/Chengxufeng1994/water-ball-missions/chapter05/computation_model/pkg/computationmodel"
)

func main() {
	models := computationmodel.NewModels()
	model, _ := models.CreateModel("Reflection")
	vector := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		vector[i] = float64(i)
	}
	result := model.Transform(computationmodel.Vector(vector))
	fmt.Println(result)
}
