package main
import "fmt"
func main() {
	arr := []float64{1, 5, 7, 4, 2}
	sum := 0.0
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum/float64(len(arr)))


}
