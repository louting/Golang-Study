## 可取代性例子

```golang
package main

import (
    "fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) { // 其实这个返回值根本没用到
    *c += ByteCounter(len(p)) // 转换int为ByteCounter类型
    return len(p), nil
}

func main() {

    var c ByteCounter
    c.Write([]byte("hello"))
    fmt.Println(c) // "5" = len("hello")

    c = 0 // 重置计数器
    var name = "Dolly"
    fmt.Fprintf(&c, "hello, %s", name)
    fmt.Println(c)

}

```