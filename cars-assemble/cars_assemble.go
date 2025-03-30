package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRate = successRate / 100
	var workingCarsPerHour float64 = float64(productionRate) * successRate
	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successRate = successRate / 100
	var workingCarsPerMinute int = int(float64(productionRate) * successRate / 60)
	return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var singleCarCost uint = 10000
	var tenCarCost uint = 95000
	var totalCost uint = uint(carsCount/10) * tenCarCost + uint(carsCount%10) * singleCarCost
	return totalCost
}
