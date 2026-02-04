package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	RemainTime := OvenTime - actualMinutesInOven
	if RemainTime > 0 {
		return RemainTime
	}
	panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	PreparTime := numberOfLayers * 2
	if PreparTime > 0 {
		return PreparTime
	}
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	Elapsed := PreparationTime(numberOfLayers) + actualMinutesInOven
	if Elapsed > 0 {
		return Elapsed
	}
	panic("ElapsedTime not implemented")
}
