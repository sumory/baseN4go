package main

import "fmt"
import "math"
import "errors"

func Test(str string) {
	fmt.Println("just test...", str)
}

//utf8
var defaultBase = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

const (
	maxNum = math.MaxInt64 // int64(1<<63 - 1)
)

type BaseN struct {
	base  []string
	radix int8
}

func New(radix int8) (error, *BaseN) {
	if radix > 62 || radix < 2 {
		return errors.New("rror param: if the param is numeric, it must be between 2 and 62."), nil
	}

	baseN := &BaseN{
		base:  defaultBase[0:radix],
		radix: radix,
	}

	return nil, baseN
}

func (this *BaseN) encode(num int64) string {
	if num > maxNum {
		return "-"
	}

	var tmp string
	var result string
	var negative = false

	if num == 0 {
		return this.base[0]
	}
	if num < 0 {
		negative = true
		num = num * (-1)
	}

	for num > 0 {
		tmp += this.base[int8(num%int64(this.radix))]
		num = int64(num / int64(this.radix))
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		result += string(tmp[i])
	}

	if negative {
		result = "-" + result
	}
	return result
}

func main() {
	fmt.Println(maxNum)
	fmt.Println(defaultBase)
	_, base := New(2)
	for i := 0; i < 20; i++ {
		fmt.Println(base.encode(int64(i)))
		fmt.Println(base.encode(int64(-i)))
	}

}
