package models

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CookingAppliance int

const (
	None CookingAppliance = iota
	Stove
	Oven
	Microwave
	Toaster
	ToasterOven
	Blender
	FoodProcessor
	SlowCooker
	PressureCooker
	RiceCooker
	DeepFryer
	AirFryer
	OutdoorGrill
	Griddle
	SandwichMaker
	WaffleMaker
	CoffeeMaker
	EspressoMachine
	ElectricKettle
	SousVide
	InductionCooktop
	BreadMaker
	IceCreamMaker
	StandMixer
	HandMixer
	Juicer
	HotPlate
	Rotisserie
	Smokers
	DutchOven
	InstantPot
	Steamer
	FonduePot
	PopcornMaker
	EggCooker
	Unknown
)

type CookingAction int

const (
	Other CookingAction = iota
	Stir
	Mix
	Whisk
	Beat
	Fold
	Knead
	Chop
	Slice
	Dice
	Mince
	Grate
	Shred
	Peel
	Julienne
	Sear
	Saute
	Fry
	DeepFry
	Boil
	Simmer
	Blanch
	Poach
	Steam
	Roast
	Bake
	Broil
	Grill
	Braise
	Stew
	Marinate
	Glaze
	Baste
	Caramelize
	Deglaze
	Puree
	Blend
	Chill
	Freeze
	Thaw
	Garnish
	Measure
	Weigh
	Sift
	Strain
	Season
	Brine
	Smoke
	Cure
	Ferment
	Proof
	Skim
	Toast
	Crush
	Grind
	Melt
	Reheat
	Rest
	Trim
)

type RecipeStepStage int

const (
	Preparation RecipeStepStage = iota
	Cooking
	Resting
)

type FlavorProfile struct {
	Sweet  int `json:"sweet" dynamodbav:"sweet"`
	Salty  int `json:"salty" dynamodbav:"salty"`
	Spicy  int `json:"spicy" dynamodbav:"spicy"`
	Sour   int `json:"sour" dynamodbav:"sour"`
	Savory int `json:"savory" dynamodbav:"savory"`
}

func (fp *FlavorProfile) ValidateValue(value int) int {
	if value < 0 {
		return 0
	} else if value > 5 {
		return 5
	} else {
		return value
	}
}

type Ingredient struct {
	Id   string `json:"id" dynamodbav:"id"`
	Name string `json:"name" dynamodbav:"name"`
}

func (i *Ingredient) CapitalizeWords(value string) string {
	caser := cases.Title(language.English)
	return caser.String(value)
}

type RecipeStep struct {
	Id              string           `json:"id" dynamodbav:"id"`
	Name            string           `json:"name" dynamodbav:"name"`
	ApplianceUsed   CookingAppliance `json:"applianceUsed" dynamodbav:"applianceUsed"`
	CookingAction   CookingAction    `json:"cookingAction" dynamodbav:"cookingAction"`
	TimeSeconds     int              `json:"timeSeconds" dynamodbav:"timeSeconds"`
	PrepTimeSeconds int              `json:"prepTimeSeconds" dynamodbav:"prepTimeSeconds"`
	Ingredients     []Ingredient     `json:"ingredients" dynamodbav:"ingredients"`
	Description     string           `json:"description" dynamodbav:"description"`
	RecipeStage     RecipeStepStage  `json:"recipeStage" dynamodbav:"recipeStage"`
}

type Recipe struct {
	Id                   string        `json:"id" dynamodbav:"id"`
	Name                 string        `json:"name" dynamodbav:"name"`
	Steps                []RecipeStep  `json:"steps" dynamodbav:"steps"`
	Ingredients          []Ingredient  `json:"ingredients" dynamodbav:"ingredients"`
	TotalTimeSeconds     int           `json:"totalTimeSeconds" dynamodbav:"totalTimeSeconds"`
	TotalTimePrepSeconds int           `json:"totalTimePrepSeconds" dynamodbav:"totalTimePrepSeconds"`
	FlavorProfile        FlavorProfile `json:"flavorProfile" dynamodbav:"flavorProfile"`
	Country              string        `json:"country" dynamodbav:"country"`
	Note                 string        `json:"note,omitempty" dynamodbav:"note,omitempty"`
}
