# golang with env

## use https://github.com/joho/godotenv  package to read env

***NOTICE***

可以利用 godotenv 的init function來做auto load

所以可以使用
```golang===
import (
    _ "github.com/joho/godotent"
)
```
注意這種寫法
是指load init function 其他 function都不被loading 進來

因此須除了 特殊auto load的需求 大部分不建議這樣做
[godotenv repo](https://github.com/joho/godotenv)