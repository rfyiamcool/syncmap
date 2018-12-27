# syncmap

Golang官方在 1.9 加入了sync.map协程安全的map, 性能和安全得以保证，就是没有Length方法. 自己丰衣足食加了个补丁.

至于官方为什么不加Length方法原因，有兴趣的可以看看issue. 简单说官方认为 map 本来就不应该有length的实现.

![](why.jpeg)

### How

在sync.map结构体里加了计数器，在触发Store和Delete时，Atomic.AddInit64 +1 -1就可以了.

### To Do List

* 加入 test 测试单元
