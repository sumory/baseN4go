package baseN4go

import "math"
import "errors"
import "strings"
import "fmt"

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
	radix int64
}

func NewBaseN(param interface {}) (error, *BaseN) {
	var baseN *BaseN

	switch paramType:= param.(type){
	case int:
		a := int64(param.(int))
		if a > 62 || a < 2 {
			return errors.New("error param: if the param is numeric, it must be between 2 and 62."), nil
		}
		baseN = &BaseN{
			base:  defaultBase[0:a],
			radix: a,
		}
	case int32:
		a := int64(param.(int32))
		if a > 62 || a < 2 {
			return errors.New("error param: if the param is numeric, it must be between 2 and 62."), nil
		}
		baseN = &BaseN{
			base:  defaultBase[0:a],
			radix: a,
		}
	case int64:
		a := param.(int64)
		if a > 62 || a < 2 {
			return errors.New("error param: if the param is numeric, it must be between 2 and 62."), nil
		}
		baseN = &BaseN{
			base:  defaultBase[0:a],
			radix: a,
		}
	case int8:
		a := int64(param.(int8))
		if a > 62 || a < 2 {
			return errors.New("error param: if the param is numeric, it must be between 2 and 62."), nil
		}
		baseN = &BaseN{
			base:  defaultBase[0:a],
			radix: a,
		}

	case []string:
		b := param.([]string)
		radix := len(b)
		if radix > 62 || radix < 2 {
			return errors.New("error param: if the type of param is []string, it's length must be between 2 and 62."), nil
		}
		baseN = &BaseN{
			base:  b,
			radix: int64(radix),
		}
	default:
		fmt.Println("illegel param type: ", paramType)
		return errors.New("error param type: param must be int8(2~62) or []string(length:2~62)"), nil
	}

	return nil, baseN
}

func (this *BaseN) Encode(num int64) (error, string) {

	// type int64 is not bigger than maxNum(math.MaxInt64), needn't check again
	//	if num > maxNum {
	//		return errors.New("input param is bigger than maxNum(1<<63-1)"), ""
	//	}

	var tmp string
	var result string
	var negative bool = false

	if num == 0 {
		return nil, this.base[0]
	}
	if num < 0 {
		negative = true
		num = num*(-1)
	}

	for num > 0 {
		tmp += this.base[int8(num % int64(this.radix))]
		num = int64(num / int64(this.radix))
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		result += string(tmp[i])
	}

	if negative {
		result = "-"+result
	}
	return nil, result
}

func (this *BaseN) Decode(str string) (error, int64) {
	var result int64
	var negative int64 = 1

	if strings.HasPrefix(str, "-") {
		negative = -1
		str = str[1:]
	}

	for index := 0; index < len(str); index++ {
		c := string(str[index])
		var tableIndex,i int64
		for i=0; i < this.radix; i++ {
			if string(this.base[i]) == c {
				tableIndex = i
				break
			}
		}

		var tmp int64 = 1
		for j := len(str) - index - 1; j > 0; j-- {
			tmp = tmp*this.radix
		}
		result += tableIndex*tmp
	}

	return nil, result*negative
}

