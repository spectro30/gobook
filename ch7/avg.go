package main
import "fmt"

func average (arr []float64) float64{
	sum := 0.0
	for _, i := range arr {
		sum += i
	}
	return sum/float64(len(arr))

}
func main() {
	arr := []float64{1, 5, 7, 4, 2}
	fmt.Println(average(arr))


}
