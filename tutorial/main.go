// 主包，可执行文件所在包
package main

// 导入包
import (
	"fmt"
	"math"
	"math/rand"
	// "test01/tutorial"
	//"test01/morestrings"
)

type binOp func(int, int) int

// main函数为入口函数
func main() {
	fmt.Println("This is in test.go main func.")
	// strings.Min()
	// tutorial.Logical()
	// tutorial.Integer()
	// tutorial.String()
	// tutorial.IF()
	// tutorial.For()
	// tutorial.Array()
	// tutorial.Slice()
	// tutorial.Map()
	// tutorial.Func()
	// tutorial.Pointer()
	// tutorial.Type()
	//tutorial.Struct()
	//typeMixing()
	//fmt.Println(Uint8FromInt(223))
	//PositiveRand()
	//println(strings.HasPrefix("This is an example", "This"))
	//println(strings.HasSuffix("This is an example", "ple"))
	//println(strings.Contains("This is an example", "example"))
	//PrintG()
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 { break }
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
}

func typeMixing() {
	var a int16 = 21
	var b int32
	b = int32(a)
	fmt.Println(b)
	fmt.Println(a)
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n < math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func PositiveRand() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(8))
	}
}

func PrintG() {
	//for str, i := "", 0; i < 25; i++ {
	//	str += "G"
	//	println(str)
	//}

	//for t, err := p.Token(); err == nil; t, err = p.Token() {
	//}
	string := "Go is a beautiful language!"
	for pos, val := range string {
		fmt.Printf("%d -- %c\n", pos, val)
	}
}

// 编译执行
// go run test.go
