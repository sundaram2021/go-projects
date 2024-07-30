package main

import (
	"fmt"
)

func main() {
	ns := GetNutritionalScore(Nutritionaldata{
		Energy: EnergyFromKcal(10),
		Sugars: SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium: SodiumMilligram(50),
		Protein: ProteinGram(60),
		Fibre: FibreGram(4),
		Fruits: FruitPercent(2),
	}, Food)


	fmt.Printf("Nutritional Score : %v\n",ns.Value)
	fmt.Printf("NutriScore : %s",ns.GetNutriScore())
}