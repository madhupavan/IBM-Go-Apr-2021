package shapes

import "fmt"

func PrintDimentions(s StandardShape) {
	fmt.Println("Area : ", s.Area())
	fmt.Println("Perimeter : ", s.Perimeter())
}
