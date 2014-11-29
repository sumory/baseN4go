package baseN4go

import (
	"fmt"
	"testing"
	"strconv"
)

func test(radix int8, testMinNum int64, testMaxNum int64, t *testing.T) {
	//runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Test:", radix, testMinNum, testMaxNum)
	err, base := NewBaseN(radix)

	if err != nil {
		fmt.Println("can not initialize a new BaseN4go instance")
	} else {
		for i := testMinNum; i < testMaxNum; i++ {
			//fmt.Print("origin is ", i)
			err, encodeResult := base.encode(int64(i))
			if err != nil {
				//fmt.Println("error when encoding..."
				t.FailNow()
			} else {
				//fmt.Print("    encode is ", encodeResult)
				err2, decodeResult := base.decode(encodeResult)
				if err2 != nil {
					//fmt.Println("error when decoding...")

				}else {
					if decodeResult == int64(i) {
						//fmt.Print("    decode is ", decodeResult)
					}else {
						fmt.Print("decode is equal encode ", i, encodeResult, decodeResult)
						t.Fail()
					}
				}
			}
			//fmt.Println()
		}
	}
}


func TestRadix8(t *testing.T) {
	test(8, 0, 1<<10, t)
}


func TestRadix16(t *testing.T) {
	test(16, 1<<10, 1<<20, t)
}

func TestRadix62(t *testing.T) {
	test(62, 1<<30-100000, 1<<30, t)
}


func TestRadix2(t *testing.T) {
	test(2, 1<<60-100000, 1<<60, t)
}

func TestRadix10(t *testing.T) {
	test(10, 1<<63-1<<10, 1<<63-1, t)
}

/*
 * 注意ParseUint与ParseInt、uint64和int64对测试的影响
 */
func TestRadix16_Stand(t *testing.T) {
	fmt.Println("TestRadix16WithHex...")
	err, base := NewBaseN(16)

	//warning...
	v1, _ := strconv.ParseUint("200", 16, 10)
	fmt.Println(v1)

	v2, _ := strconv.ParseUint("1ff", 16, 10)
	fmt.Println(v2)

	v3, _ := strconv.ParseUint("1fe", 16, 10)
	fmt.Println(v3)

	v4, _ := strconv.ParseInt("200", 16, 10)
	fmt.Println(v4)

	v5, _ := strconv.ParseInt("1ff", 16, 10)
	fmt.Println(v5)

	v6, _ := strconv.ParseInt("1fe", 16, 10)
	fmt.Println(v6)

	if err != nil {
		fmt.Println("can not initialize a new BaseN4go instance")
	} else {
		for i := 0; i < 1<<10; i++ {
			err, encodeResult := base.encode(int64(i))
			if err != nil {
				t.FailNow()
			} else {
				value, e := strconv.ParseUint(encodeResult, 16, 10)

				if e != nil {
					fmt.Println(i, encodeResult, value)
					t.FailNow()
				}
				if value != uint64(i) {
					fmt.Println(i, encodeResult, value)
					t.FailNow()
				}
			}
		}
	}
}
