package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int 
	Positive  int 
	Negative  int 
	ScoreType ScoreType
}


var scoreToLetter = []string{"A", "B", "C", "D", "E"}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcids float64

type SodiumMilligram float64

type FibreGram float64

type ProteinGram float64

type FruitPercent float64



type Nutritionaldata struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Protein             ProteinGram
	Fibre               FibreGram
	Fruits              FruitPercent
	isWater			    bool
}

var energyLevels = []float64{3350, 3015, 2680,2345,2010,1675,1340,1005,670,335}
var sugarLevels = []float64{45, 60, 36,31,27, 22.5, 18, 13.5, 9, 4.5}
var saturatedFattyAcidsLevels = []float64{10,9,8,7,6,5,4,3,2,1}
var sodiumLevels = []float64{900,810,720,630,540,450,360,270,180,90}
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8,6.4,4.8,3.2,1.6}

var energyLevelBeverage = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelBeverage = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Food {
		return getPointsFromRange(float64(e), energyLevels)
	} else if st == Beverage {
		return getPointsFromRange(float64(e), energyLevelBeverage)
	} else {
		return 0
	}
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if st == Food {
		return getPointsFromRange(float64(s), sugarLevels)
	} else if st == Beverage {
		return getPointsFromRange(float64(s), sugarLevelBeverage)
	} else {
		return 0
	}
}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidsLevels)
}

func (smg SodiumMilligram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(smg), sodiumLevels)
}

func (f FibreGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(f), fibreLevels)
}

func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p), proteinLevels)
}

func (f FruitPercent) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80  {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		} 
		return 0
	}

	if f > 80  {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

func EnergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SugarFromGram(g float64) SodiumMilligram {
	return SodiumMilligram(g / 2.5)
}


func GetNutritionalScore( n Nutritionaldata, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fibrePoints := n.Fibre.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints(st) + n.Sodium.GetPoints(st)
		positive = n.Protein.GetPoints(st) + fruitPoints + fibrePoints

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	}

	return NutritionalScore{
		Value: value,
		Positive: positive,
		Negative: negative,
		ScoreType: st,
	}
}


func (ns NutritionalScore) GetNutriScore() string {
	if ns.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{2, 10, 18, -1})]
	}

	if ns.ScoreType == Water {
		return scoreToLetter[0]
	}

	return scoreToLetter[getPointsFromRange(float64(ns.Value), []float64{1, 5, 9, -2})]
}


func getPointsFromRange(v float64, steps []float64) int {
	lenSteps := len(steps)

	for i, l := range steps {
		if v >= l {
			return lenSteps - i
		}
	}
	return 0
}


