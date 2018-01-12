# syncmap

Golang官方在 1.9 加入了sync.map协程安全的map, 性能和安全得以保证，就是没有Length方法. 当然，官方说 map 本来就不应该有length的实现. 

### How

在sync.map结构体里加了计数器，在触发Store和Delete时，Atomic.AddInit64 +1 -1就可以了.
