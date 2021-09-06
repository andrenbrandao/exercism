package lasagna

func OvenTime() int {
	expectedTime := 40
	return expectedTime
}

func RemainingOvenTime(currentTime int) int {
	return OvenTime() - currentTime
}

func PreparationTime(layers int) int {
	return layers * 2
}

func ElapsedTime(layers, minutesInOven int) int {
	return PreparationTime(layers) + minutesInOven
}
