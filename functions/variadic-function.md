# Variadic function

## Variadic Functions

Variadic funksiyalar, Go proqramlaşdırma dilində dəyişən sayda arqument qəbul edən funksiyalardır. Bu funksiyalar bir və ya daha çox arqument qəbul edə bilər və arqumentlərin sayını dəyişən olaraq təyin etməyə imkan verir.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4, 5))
fmt.Println(sum(2, 4, 6))
```

Bu nümunədə, sum adlı bir funksiya təyin edilir. Funksiya, nums adlı int tipində dəyişən sayda parametr qəbul edir və onların cəmini qaytarır. Funksiyanın daxilində for dövrü istifadə edərək, nums dilimindəki bütün elementlərin cəmi hesablanır. Funksiya sum(1, 2, 3, 4, 5) və sum(2, 4, 6) şəklində çağırılır və nəticələr ekrana yazdırılır.

```go
func concatenate(sep string, strs ...string) string {
    result := ""
    for i, str := range strs {
        if i > 0 {
            result += sep
        }
        result += str
    }
    return result
}

fmt.Println(concatenate(", ", "foo", "bar", "baz"))
fmt.Println(concatenate("-", "hello", "world"))
```

Bu nümunədə, concatenate adlı bir funksiya təyin edilir. Funksiya, sep adlı string tipində bir parametr və strs adlı dəyişən sayda string tipində parametrlər qəbul edir. Funksiyanın daxilində for dövrü istifadə edilərək, strs dilimindəki bütün elementlər birləşdirilir və sep ayracı ilə birgə yazılır. Funksiya concatenate(", ", "foo", "bar", "baz") və concatenate("-", "hello", "world") şəklində çağırılır və nəticələr ekrana yazdırılır.
