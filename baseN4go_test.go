package baseN4go

import (
	"testing"
	convey "github.com/smartystreets/goconvey/convey"

	"strconv"
)

func test(name string, radix int8, testMinNum int64, testMaxNum int64, t *testing.T) {
	convey.Convey(name, t, func() {
			_, baseN := NewBaseN(radix)
			//convey.So(err, convey.ShouldBeNil)
			var i int64
			for i = testMinNum; i < testMaxNum; i++ {
				err, encodeResult := baseN.Encode(i)
				//convey.So(err, convey.ShouldBeNil)
				err, decodeResult := baseN.Decode(encodeResult)
				convey.So(err, convey.ShouldBeNil)
				convey.So(decodeResult, convey.ShouldEqual, int64(i))
			}
		})
}

func TestRadix2(t *testing.T) {
	test("test radix 2", 2, 0, 1<<10, t)
}

func TestRadix8(t *testing.T) {
	test("test radix 8", 8, 0, 1<<10, t)
}

func TestRadix10(t *testing.T) {
	test("test radix 10", 10, 1<<16-100, 1<<16, t)
}

func TestRadix16(t *testing.T) {
	test("test radix 16", 16, 1<<16-100, 1<<16, t)
}

func TestRadix62(t *testing.T) {
	test("test radix 62", 62, 1<<16-100, 1<<16, t)
}


func TestRadix16_Stand(t *testing.T) {

	//warning...注意ParseUint与ParseInt、uint64和int64对测试的影响
	//	v1, _ := strconv.ParseUint("200", 16, 10)
	//	fmt.Println(v1)
	//
	//	v2, _ := strconv.ParseUint("1ff", 16, 10)
	//	fmt.Println(v2)
	//
	//	v3, _ := strconv.ParseUint("1fe", 16, 10)
	//	fmt.Println(v3)
	//
	//	v4, _ := strconv.ParseInt("200", 16, 10)
	//	fmt.Println(v4)
	//
	//	v5, _ := strconv.ParseInt("1ff", 16, 10)
	//	fmt.Println(v5)
	//
	//	v6, _ := strconv.ParseInt("1fe", 16, 10)
	//	fmt.Println(v6)

	convey.Convey("should equal with srtconv.ParseUnit result", t, func() {
			err, baseN := NewBaseN(int8(16))
			convey.So(err, convey.ShouldBeNil)
			var i int64
			for i = 0; i < 1024; i++ {
				_, encodeResult := baseN.Encode(i)
				value, e := strconv.ParseUint(encodeResult, 16, 10)
				convey.So(e, convey.ShouldBeNil)
				convey.So(value, convey.ShouldEqual, i)
			}
		})
}

func TestBasics(t *testing.T) {
	convey.Convey("test basics", t, func() {
			convey.So(len(defaultBase), convey.ShouldEqual, 62)
			convey.So(maxNum, convey.ShouldEqual, 1<<63-1)
		})
}

func TestConstructor(t *testing.T) {
	convey.Convey("constructor should be right",t, func(){
			err, _ := NewBaseN([]string{"a", "b", "c", "d"})
			convey.So(err, convey.ShouldBeNil)
			err, _ = NewBaseN(4)
			convey.So(err, convey.ShouldBeNil)
			err, _ = NewBaseN(int(4))
			convey.So(err, convey.ShouldBeNil)
			err, _ = NewBaseN(int8(4))
			convey.So(err, convey.ShouldBeNil)

			err, _ = NewBaseN(int32(4))
			convey.So(err, convey.ShouldBeNil)
			err, _ = NewBaseN(int64(4))
			convey.So(err, convey.ShouldBeNil)

			err, baseN := NewBaseN(int(10))
			convey.So(err, convey.ShouldBeNil)
			convey.So(baseN,convey.ShouldNotBeNil)
			convey.So(baseN.radix,convey.ShouldEqual,10)

			convey.So(baseN,convey.ShouldHaveSameTypeAs,new(BaseN))
		})

	convey.Convey("constructor should be wrong",t, func(){
			err, _ := NewBaseN([]string{
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
			"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
			"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
			"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
			"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
			"-","+"})
			convey.So(err, convey.ShouldNotBeNil)


			err, _ = NewBaseN(63)
			convey.So(err, convey.ShouldNotBeNil)
			err, _ = NewBaseN(1)
			convey.So(err, convey.ShouldNotBeNil)

			err, _ = NewBaseN(int8(63))
			convey.So(err, convey.ShouldNotBeNil)
			err, _ = NewBaseN(int32(1))
			convey.So(err, convey.ShouldNotBeNil)

			err, _ = NewBaseN(int(63))
			convey.So(err, convey.ShouldNotBeNil)
			err, _ = NewBaseN(int(1))
			convey.So(err, convey.ShouldNotBeNil)
			err, _ = NewBaseN(int64(1))
			convey.So(err, convey.ShouldNotBeNil)


			err, baseN := NewBaseN(0)
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(baseN,convey.ShouldBeNil)

			err, baseN = NewBaseN("abc")
			convey.So(err, convey.ShouldNotBeNil)
			convey.So(baseN,convey.ShouldBeNil)


		})
}


func TestNegative(t *testing.T) {
	convey.Convey("test negative", t, func() {
			err,baseN := NewBaseN(62)
			err,encode := baseN.Encode(-10)
			convey.So(err, convey.ShouldBeNil)
			convey.So(encode, convey.ShouldEqual, "-a")

			err,decode := baseN.Decode("-a")
			convey.So(decode,convey.ShouldEqual,-10)

			//			err,encode = baseN.Encode(1<<64)
			//			convey.So(err,convey.ShouldNotBeNil)
		})
}

