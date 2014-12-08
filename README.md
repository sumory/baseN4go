## baseN4go

[![Build Status](https://travis-ci.org/sumory/baseN4go.svg?branch=master)](https://travis-ci.org/sumory/baseN4go) [![](http://gocover.io/_badge/github.com/sumory/baseN4go)](http://gocover.io/github.com/sumory/baseN4go) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/sumory/baseN4go/master/LICENSE)


#### 介绍

 - 查看[baseN][1]，[baseN4j][2]，这是go版本.



#### 使用

go get github.com/sumory/baseN4go

默认数组:

```go
[]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
```

取默认数组前N项作为基数

```go
var N int64 = 16
err, baseN := baseN4go.NewBaseN(N) //16进制
_, encodeResult := baseN.Encode(123456)
_, decodeResult := baseN.Decode(encodeResult)
```

自定义基数

```go
err, baseN := baseN4go.NewBaseN([]string{"a", "b", "c", "d"}) //四进制
_, encodeResult := baseN.Encode(123456)
_, decodeResult := baseN.Decode(encodeResult)
```


#### TODO

性能较低，待优化`Encode`和`Decode`方法


[1]: https://github.com/sumory/baseN
[2]: https://github.com/sumory/baseN4j