package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	numberOfLayers := len(layers)
	if timePerLayer == 0 {
		return 2 * numberOfLayers
	}

	return numberOfLayers * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	gramsPerNoodle := 50
	literPerSauce := 0.2
	noodleCount := 0
	sauceCount := 0

	for _, layer := range layers {
		if layer == "noodles" {
			noodleCount += 1
		} else if layer == "sauce" {
			sauceCount += 1
		}
	}

	return noodleCount * gramsPerNoodle, float64(sauceCount) * literPerSauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	lengthFriendsList := len(friendsList)
	lengthMyList := len(myList)

	myList[lengthMyList-1] = friendsList[lengthFriendsList-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	originalPortions := 2
	multiplier := float64(portions) / float64(originalPortions)

	scaledQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		scaledQuantities[i] = multiplier * quantity
	}

	return scaledQuantities
}
