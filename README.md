# guid
[![goproxy.cn](https://goproxy.cn/stats/github.com/ego-component/eguid/badges/download-count.svg)](https://goproxy.cn/stats/github.com/ego-component/eguid)
[![Release](https://img.shields.io/github/v/release/ego-component/eguid.svg?style=flat-square)](https://github.com/ego-component/eguid)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Example](https://img.shields.io/badge/Examples-2ca5e0?style=flat&logo=appveyor)](https://github.com/ego-component/eguid/tree/master/examples)

## 参数
```bash
Salt     string  #必须设置
Length   int     #如果在8位，加随机数到不了1亿，建议必须大于等于10
Alphabet string  #hashid默认的字母表
```

## 使用方法
```go
g := guid.Load("guid").Build()
g.EncodeRandomInt64(45) // 使用这种比较安全，因为有随机数
```
