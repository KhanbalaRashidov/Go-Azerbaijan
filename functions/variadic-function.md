# Variadic function

## Variadic Functions

Variadic funksiyalar, Go proqramlaşdırma dilində dəyişən sayda arqument qəbul edən funksiyalardır. Bu funksiyalar bir və ya daha çox arqument qəbul edə bilər və arqumentlərin sayını dəyişən olaraq təyin etməyə imkan verir.

```golang
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

```golang
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

\
\


## Closures

Closures, Go proqramlaşdırma dilində, bir funksiyanın başqa bir funksiyanın daxilində yaradılması ilə əmələ gələn bir strukturdur. Bu struktur, bir funksiyanın daxilində olan başqa bir funksiyaya istinad edərək, funksiyanın işlədiyi mühitin xaricindəki dəyişənlərə giriş imkanı verir.

```golang
func outer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

increment := outer()
fmt.Println(increment())
fmt.Println(increment())
fmt.Println(increment())
```

Bu örnekte, outer adlı bir fonksiyon tanımlanır. Fonksiyon, bir iç fonksiyon döndürür ve iç fonksiyon, count adlı bir değişkene erişim sağlar. increment adlı bir değişkene outer() fonksiyonu atanır ve bu değişken ile iç fonksiyon çalıştırılır. count değişkeni, increment() çağrıldıkça artar ve her seferinde artışı ekrana yazdırılır.

```golang
func adder(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}

addFive := adder(5)
fmt.Println(addFive(2))
fmt.Println(addFive(3))
```

Bu nümunədə, adder adlı bir funksiya təyin edilir. Funksiya, a adlı bir int tipində parametr qəbul edir və bir daxili funksiya qaytarır. Daxili funksiya, b adlı bir int tipində parametr qəbul edir və a ilə b parametrlərinin cəmini geri qaytarır. addFive adlı bir dəyişənə adder(5) funksiyası təyin edilir və bu dəyişən ilə daxili funksiya işlədilir. addFive(2) və addFive(3) çağırıldıqca nəticələr ekrana yazdırılır.

\
\
