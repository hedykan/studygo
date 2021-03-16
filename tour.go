// 练习：循环与函数
package main

import (
	"fmt",
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.00
	temp := 0.00
	for {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
		if math.Abs(z-temp) < 0.001 {
			break
		} else {
			temp = z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

// 练习：切片
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		swap := make([]uint8, dx)
		for x, _ := range swap {
			swap[x] = uint8((x * y) / 2)
		}
		s[y] = swap
	}
	return s
}

func main() {
	pic.Show(Pic)
}

// 练习：映射
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = map[string]int{}
	sArr := strings.Fields(s)
	for _, value := range sArr {
		_, ok := m[value]
		if ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

// 练习：斐波纳契闭包
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		sum3 := sum1 + sum2
		sum1 = sum2
		sum2 = sum3
		return sum3
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 练习：Stringer
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string{  
   return fmt.Sprintf("%v.%v.%v.%v", ip[0],ip[1],ip[2],ip[3])  
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// 练习：错误
package main

import (
	"fmt"
	"math"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
   return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// 练习：Reader
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read(b []byte) (int,error){
    b[0] = 'A'
    return 1,nil
}

func main() {
	reader.Validate(MyReader{})
}
