# Closures

Closures, Go proqramlaşdırma dilində, bir funksiyanın başqa bir funksiyanın daxilində yaradılması ilə əmələ gələn bir strukturdur. Bu struktur, bir funksiyanın daxilində olan başqa bir funksiyaya istinad edərək, funksiyanın işlədiyi mühitin xaricindəki dəyişənlərə giriş imkanı verir.

```go
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

Bu nümunədə, `outer` adlı bir funksiya təyin edilir. Funksiya, bir daxili funksiya qaytarır və daxili funksiya, `count` adlı bir dəyişənə giriş imkanı əldə edir. `increment` adlı bir dəyişənə `outer()` funksiyası təyin edilir və bu dəyişən vasitəsilə daxili funksiya işlədilir. `count` dəyişəni hər dəfə `increment()` çağırıldıqda artır və hər çağırışda artım ekrana yazdırılır.

```go
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
