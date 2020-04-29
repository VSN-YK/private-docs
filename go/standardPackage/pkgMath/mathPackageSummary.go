package pkgMath

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MathPackageSummary() {
	/*constant*/
	//π
	fmt.Printf("π: %f\n", math.Pi)
	//root
	fmt.Printf("√2 : %f\n", math.Sqrt2)

	//absolute vlaue
	fmt.Printf("Distance of (-)Point : %f\n", math.Abs(-245.567))
	fmt.Printf("Distance of (+)Point : %f\n", math.Abs(+245.567))

	//pow
	fmt.Printf("%d\n", int(math.Pow(2, 0))) //actual result = 1.0000

	//Sqrt and Cbrt 1 To 100
	for i := 1.0; i <= 100.0; i++ {
		fmt.Printf("[Sqrt]idx:%d\tvalue:%f\n", int(i), math.Sqrt(i))
		fmt.Printf("[Cbrt]idx:%d\tvalue:%f\n", int(i), math.Cbrt(i))
	}
	// max and min
	fmt.Println(math.Max(3.1412, 3.14))
	fmt.Println(math.Min(3.1412, 3.14))

	for j := 0.0; j <= 180.0; j++ {
		fmt.Printf("[idx]:%d\t[Sin]:%f\n", int(j), math.Sin(j))
	}

	//Common logarithm
	for k := 0.0; k <= 100.0; k++ {
		fmt.Printf("[idx]%d\t[LogValue]:%f\n", int(k), math.Log10(k))
	}

	// fix seed value
	rand.Seed(256)
	fmt.Println(rand.Int())
	fmt.Println(rand.Float64())

	// generate 1 ~ 99
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float64())

	//create new rand generator
	seedSrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(seedSrc)
	fmt.Println(rnd.Intn(100))
}
