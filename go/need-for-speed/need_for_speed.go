package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	driveTimes := car.battery / car.batteryDrain
	distanceCapacity := driveTimes*car.speed + car.distance

	distanceLeft := track.distance - distanceCapacity

	if distanceLeft <= 0 {
		return true
	}

	return false
}
