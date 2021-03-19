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

// 练习：rot13Reader
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		b = (b - 'A' + 13) % 26 + 'A'
	case 'a' <= b && b <= 'z':
		b = (b - 'a' + 13) % 26 + 'a'
	}
	return b
}

func (rot rot13Reader) Read(b []byte) (int,error){
	n, err := rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
   return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// 练习：图像
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	w int
	h int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.w,i.h)
}

func (i Image) ColorModel() color.Model {
   return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x),uint8(y),uint8(255),uint8(255)}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}

// 练习：等价二叉查找树
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	add(t, ch)
	close(ch)
}

func add(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		add(t.Left, ch)
	}
	if t.Right != nil {
		add(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	i := 0
	for n := range ch1 {
		fmt.Println(i)
		i = i ^ n // 按位异或
	}
	fmt.Println()
	for n := range ch2 {
		fmt.Println(i)
		i = i ^ n // 按位异或
	}
	return i == 0
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}