package main
import "fmt"
import "learning-go/custom_package"
func main() {
xs := []float64{1,2,3,4}
avg := custom_package.Average(xs)
fmt.Println(avg)
}