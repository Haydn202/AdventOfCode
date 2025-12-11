package main

func main() {

}

func readData(fileName string) {

}

type Point struct {
	x int
	y int
	z int
}

func CalculateDistance(firstPoint Point, secondPoint Point) int {
	return sq(secondPoint.x-firstPoint.x) + sq(secondPoint.y-firstPoint.y) + sq(secondPoint.z-firstPoint.z)
}

func sq(x int) int {
	return x * x
}
