为了实现一个辅助处理 Golang 中 map 类型的方法包，我们可以创建一个名为 `maps` 的包，其中包含添加、删除、查找、求并集、交集，以及 map 和 reduce 操作的功能。

以下是具体实现以及相应的测试文件：

### add.go

```go
package maps

// Add 向 map 中添加一个键值对，并返回新的 map
// 参数:
//   m map[K]V: 需要添加键值对的 map
//   key K: 要添加的键
//   value V: 要添加的值
// 返回值:
//   map[K]V: 添加键值对后的新 map
func Add[K comparable, V any](m map[K]V, key K, value V) map[K]V {
    m[key] = value
    return m
}
```

### add_test.go

```go
package maps

import (
    "reflect"
    "testing"
)

func TestAdd(t *testing.T) {
    m := map[int]string{1: "a", 2: "b"}
    key := 3
    value := "c"
    result := Add(m, key, value)
    expected := map[int]string{1: "a", 2: "b", 3: "c"}

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Add(%v, %d, %s) = %v; want %v", m, key, value, result, expected)
    }
}
```

### remove.go

```go
package maps

// Remove 从 map 中删除指定的键，并返回新的 map
// 参数:
//   m map[K]V: 需要删除键值对的 map
//   key K: 要删除的键
// 返回值:
//   map[K]V: 删除键值对后的新 map
func Remove[K comparable, V any](m map[K]V, key K) map[K]V {
    delete(m, key)
    return m
}
```

### remove_test.go

```go
package maps

import (
    "reflect"
    "testing"
)

func TestRemove(t *testing.T) {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    key := 2
    result := Remove(m, key)
    expected := map[int]string{1: "a", 3: "c"}

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Remove(%v, %d) = %v; want %v", m, key, result, expected)
    }
}
```

### contains.go

```go
package maps

// Contains 检查 map 中是否包含指定的键
// 参数:
//   m map[K]V: 需要检查的 map
//   key K: 要检查的键
// 返回值:
//   bool: 如果 map 中包含该键则返回 true，否则返回 false
func Contains[K comparable, V any](m map[K]V, key K) bool {
    _, exists := m[key]
    return exists
}
```

### contains_test.go

```go
package maps

import (
    "testing"
)

func TestContains(t *testing.T) {
    m := map[int]string{1: "a", 2: "b", 3: "c"}
    key := 2
    result := Contains(m, key)

    if !result {
        t.Errorf("Contains(%v, %d) = %v; want true", m, key, result)
    }

    key = 4
    result = Contains(m, key)

    if result {
        t.Errorf("Contains(%v, %d) = %v; want false", m, key, result)
    }
}
```

### union.go

```go
package maps

// Union 返回两个 map 的并集
// 参数:
//   m1 map[K]V: 第一个 map
//   m2 map[K]V: 第二个 map
// 返回值:
//   map[K]V: 两个 map 的并集
func Union[K comparable, V any](m1, m2 map[K]V) map[K]V {
    result := make(map[K]V)
    for k, v := range m1 {
        result[k] = v
    }
    for k, v := range m2 {
        result[k] = v
    }
    return result
}
```

### union_test.go

```go
package maps

import (
    "reflect"
    "testing"
)

func TestUnion(t *testing.T) {
    m1 := map[int]string{1: "a", 2: "b"}
    m2 := map[int]string{2: "B", 3: "c"}
    result := Union(m1, m2)
    expected := map[int]string{1: "a", 2: "B", 3: "c"}

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Union(%v, %v) = %v; want %v", m1, m2, result, expected)
    }
}
```

### intersection.go

```go
package maps

// Intersection 返回两个 map 的交集
// 参数:
//   m1 map[K]V: 第一个 map
//   m2 map[K]V: 第二个 map
// 返回值:
//   map[K]V: 两个 map 的交集
func Intersection[K comparable, V any](m1, m2 map[K]V) map[K]V {
    result := make(map[K]V)
    for k, v := range m1 {
        if val, exists := m2[k]; exists && val == v {
            result[k] = v
        }
    }
    return result
}
```

### intersection_test.go

```go
package maps

import (
    "reflect"
    "testing"
)

func TestIntersection(t *testing.T) {
    m1 := map[int]string{1: "a", 2: "b", 3: "c"}
    m2 := map[int]string{2: "b", 3: "C", 4: "d"}
    result := Intersection(m1, m2)
    expected := map[int]string{2: "b"}

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Intersection(%v, %v) = %v; want %v", m1, m2, result, expected)
    }
}
```

### map.go

```go
package maps

// Map 对 map 中的每个键值对应用函数，并返回一个包含结果的新 map
// 参数:
//   m map[K]V: 要处理的 map
//   f func(K, V) (K, U): 应用到每个键值对的函数
// 返回值:
//   map[K]U: 包含函数应用结果的新 map
func Map[K comparable, V any, U any](m map[K]V, f func(K, V) (K, U)) map[K]U {
    result := make(map[K]U)
    for k, v := range m {
        nk, nv := f(k, v)
        result[nk] = nv
    }
    return result
}
```

### map_test.go

```go
package maps

import (
    "reflect"
    "testing"
)

func TestMap(t *testing.T) {
    m := map[int]string{1: "a", 2: "b"}
    result := Map(m, func(k int, v string) (int, string) {
        return k, v + "!"
    })
    expected := map[int]string{1: "a!", 2: "b!"}

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Map(%v, func) = %v; want %v", m, result, expected)
    }
}
```

### reduce.go

```go
package maps

// Reduce 对 map 中的每个键值对应用函数，累计结果并返回
// 参数:
//   m map[K]V: 要处理的 map
//   initial U: 初始值
//   f func(U, K, V) U: 应用到每个键值对的函数
// 返回值:
//   U: 累计结果
func Reduce[K comparable, V any, U any](m map[K]V, initial U, f func(U, K, V) U) U {
    result := initial
    for k, v := range m {
        result = f(result, k, v)
    }
    return result
}
```

### reduce_test.go

```go
package maps

import (
    "testing"
)

func TestReduce(t *testing.T) {
    m := map[int]int{1: 1, 2: 2, 3: 3}
    result := Reduce(m, 0, func(acc, k, v int) int {
        return acc + v
    })
    expected := 6

    if result != expected {
        t.Errorf("Reduce(%v, 0, func) = %d; want %d", m, result, expected)
    }
}
```

通过上述代码，我们可以创建一个名为 `maps` 的包，包含所有所需的 map 操作方法。每个文件都有详细的中文注释，并且每个方法都有对应的测试文件。

你可以在终端使用以下命令来创建这些文件：

```sh
mkdir -p maps
touch maps/add

.go maps/add_test.go
touch maps/remove.go maps/remove_test.go
touch maps/contains.go maps/contains_test.go
touch maps/union.go maps/union_test.go
touch maps/intersection.go maps/intersection_test.go
touch maps/map.go maps/map_test.go
touch maps/reduce.go maps/reduce_test.go
```