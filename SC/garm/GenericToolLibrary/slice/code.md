# Slice

### slice/add.go

```go
package slice

/** 
 * Add 将一个元素添加到切片并返回新的切片
 * @param slice []T 需要添加元素的切片
 * @param element T 需要添加的元素
 * @return []T 添加元素后的新切片
 */
func Add[T any](slice []T, element T) []T {
    return append(slice, element)
}
```

### slice/add_test.go

```go
package slice

import (
    "testing"
)

func TestAdd(t *testing.T) {
    slice := []int{1, 2, 3}
    newSlice := Add(slice, 4)
    expected := []int{1, 2, 3, 4}

    for i, v := range newSlice {
        if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
```

### slice/remove.go

```go
package slice

/** 
 * Remove 从切片中删除第一次出现的元素
 * @param slice []T 需要删除元素的切片
 * @param element T 需要删除的元素
 * @return []T 删除元素后的新切片
 */
func Remove[T comparable](slice []T, element T) []T {
    for i, v := range slice {
        if v == element {
            return append(slice[:i], slice[i+1:]...)
        }
    }
    return slice
}
```

### slice/remove_test.go

```go
package slice

import (
    "testing"
)

func TestRemove(t *testing.T) {
    slice := []int{1, 2, 3, 4}
    newSlice := Remove(slice, 3)
    expected := []int{1, 2, 4}

    if len(newSlice) != len(expected) {
        t.Fatalf("Expected length %d, got %d", len(expected), len(newSlice))
    }

    for i, v := range newSlice {
        if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
```

### slice/contains.go

```go
package slice

/** 
 * Contains 检查切片中是否存在某个元素
 * @param slice []T 需要检查的切片
 * @param element T 需要查找的元素
 * @return bool 如果元素存在返回 true，否则返回 false
 */
func Contains[T comparable](slice []T, element T) bool {
    for _, v := range slice {
        if v == element {
            return true
        }
    }
    return false
}
```

### slice/contains_test.go

```go
package slice

import (
    "testing"
)

func TestContains(t *testing.T) {
    slice := []int{1, 2, 3, 4}
    if !Contains(slice, 3) {
        t.Errorf("Expected slice to contain 3")
    }
    if Contains(slice, 5) {
        t.Errorf("Expected slice not to contain 5")
    }
}
```

### slice/union.go

```go
package slice

/** 
 * Union 返回两个切片的并集
 * @param slice1 []T 第一个切片
 * @param slice2 []T 第二个切片
 * @return []T 两个切片的并集
 */
func Union[T comparable](slice1, slice2 []T) []T {
    result := make([]T, len(slice1))
    copy(result, slice1)
    for _, v := range slice2 {
        if !Contains(result, v) {
            result = append(result, v)
        }
    }
    return result
}
```

### slice/union_test.go

```go
package slice

import (
    "testing"
)

func TestUnion(t *testing.T) {
    slice1 := []int{1, 2, 3}
    slice2 := []int{3, 4, 5}
    result := Union(slice1, slice2)
    expected := []int{1, 2, 3, 4, 5}

    if len(result) != len(expected) {
        t.Fatalf("Expected length %d, got %d", len(expected), len(result))
    }

    for i, v := range result {
        if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
```

### slice/intersection.go

```go
package slice

/** 
 * Intersection 返回两个切片的交集
 * @param slice1 []T 第一个切片
 * @param slice2 []T 第二个切片
 * @return []T 两个切片的交集
 */
func Intersection[T comparable](slice1, slice2 []T) []T {
    var result []T
    for _, v := range slice1 {
        if Contains(slice2, v) {
            result = append(result, v)
        }
    }
    return result
}
```

### slice/intersection_test.go

```go
package slice

import (
    "testing"
)

func TestIntersection(t *testing.T) {
    slice1 := []int{1, 2, 3}
    slice2 := []int{3, 4, 5}
    result := Intersection(slice1, slice2)
    expected := []int{3}

    if len(result) != len(expected) {
        t.Fatalf("Expected length %d, got %d", len(expected), len(result))
    }

    for i, v := range result {
        if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
```

### slice/map.go

```go
package slice

/** 
 * Map 对切片中的每个元素应用一个函数，并返回一个包含结果的新切片
 * @param slice []T 输入切片
 * @param f func(T) U 应用于每个元素的函数
 * @return []U 应用函数后的新切片
 */
func Map[T any, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}
```

### slice/map_test.go

```go
package slice

import (
    "testing"
)

func TestMap(t *testing.T) {
    slice := []int{1, 2, 3}
    result := Map(slice, func(x int) int {
        return x * 2
    })
    expected := []int{2, 4, 6}

    if len(result) != len(expected) {
        t.Fatalf("Expected length %d, got %d", len(expected), len(result))
    }

    for i, v := range result {
        if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
```

### slice/reduce.go

```go
package slice

/** 
 * Reduce 对切片中的每个元素应用一个函数，并将结果累积
 * @param slice []T 输入切片
 * @param initial U 累积的初始值
 * @param f func(U, T) U 应用于每个元素的函数
 * @return U 累积的结果
 */
func Reduce[T any, U any](slice []T, initial U, f func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = f(result, v)
    }
    return result
}
```

### slice/reduce_test.go

```go
package slice

import (
    "testing"
)

func TestReduce(t *testing.T) {
    slice := []int{1, 2, 3, 4}
    result := Reduce(slice, 0, func(acc, x int) int {
        return acc + x
    })
    expected := 10

    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}
```

以上是每个功能的详细实现及其测试文件。你可以将这些文件分别放到你的项目目录中，并运行 `go test ./slice/...` 来测试这些功能。
