# Anonymous Functions in Go

Anonymous functions (anonim funksiyalar) Go proqramlaşdırma dilində adlandırılmamış funksiyalardır. Bu funksiyalar ad vermədən yaradılır və lazım olduqda birbaşa istifadə olunur. Anonymous funksiyalar adətən bir dəfəlik əməliyyatlar üçün və ya daha çevik funksiyaların yaradılması üçün istifadə olunur.

## Anonymous Functionların İstifadəsi

Anonymous funksiyalar Go-da çox məqsədli istifadə üçün uyğundur. Onlar kodun daha səliqəli və funksional olmasını təmin edir, həmçinin funksiya içində lokal olaraq müəyyənləşdirilən dəyişənlərin və məntiqin istifadəsinə imkan tanıyır.

### Basic Syntax

Anonymous funksiyanın yaradılması və çağırılması sintaksisi aşağıdakı kimidir:

```go
package main

import "fmt"

func main() {
    // Anonymous funksiyanın yaradılması və birbaşa icrası
    func() {
        fmt.Println("Hello from anonymous function!")
    }()
}
```

Bu nümunədə, anonim funksiya dərhal çağırılır və "Hello from anonymous function!" mesajını ekrana yazdırır.

### Parameters və Return Values ilə Anonim Funksiyalar

Anonymous funksiyalar həmçinin parametr qəbul edə və dəyər qaytara bilər:

```go
package main

import "fmt"

func main() {
    // Anonim funksiya ilə toplama əməliyyatı
    sum := func(a int, b int) int {
        return a + b
    }
    
    result := sum(3, 4)
    fmt.Println("Sum:", result) // Sum: 7
}
```

Bu nümunədə, anonim funksiya `a` və `b` parametrlərini qəbul edir və onların cəmini qaytarır. Bu funksiya `sum` adlı dəyişəndə saxlanılır və sonra çağırılır.

### Closures ilə Anonim Funksiyalar

Anonymous funksiyalar closures (bağlanma) kimi də istifadə edilə bilər. Bu, funksiyanın yaradıldığı mühitdəki dəyişənlərə daxil olmaq imkanı verir:

```go
package main

import "fmt"

func main() {
    // Bağlanma nümunəsi
    increment := func(x int) func() int {
        return func() int {
            x++
            return x
        }
    }
    
    counter := increment(0) // counter dəyişəni anonim funksiya alır
    
    fmt.Println(counter()) // 1
    fmt.Println(counter()) // 2
    fmt.Println(counter()) // 3
}
```

Bu nümunədə, `increment` anonim funksiyası digər anonim funksiya qaytarır. Bu iç funksiyanın `x` dəyişəninə bağlandığına diqqət yetirin, beləliklə, `counter` funksiyası hər çağırıldığında `x` dəyişənini artırır.

### Anonim Funksiyaların Parametr Kimi İstifadəsi

Anonymous funksiyalar çox vaxt digər funksiyalara parametr kimi ötürülür:

```go
package main

import "fmt"

// Funksiya anonim funksiyanı parametr kimi qəbul edir
func execute(fn func(int) int, value int) {
    result := fn(value)
    fmt.Println("Result:", result)
}

func main() {
    // Anonim funksiyanın digər funksiya üçün parametr kimi istifadə olunması
    execute(func(x int) int {
        return x * x
    }, 5) // Result: 25
}
```

Bu nümunədə, `execute` funksiyası anonim funksiyanı `fn` parametri kimi qəbul edir və sonra həmin anonim funksiyanı çağırır.

## Nəticə

Anonymous funksiyalar Go dilində kodun daha modul və çevik olmasına kömək edir. Onlar qısa müddətli funksiyaları daha sadə şəkildə təyin etməyə imkan tanıyır və bağlanma, parametr kimi ötürmə və digər müasir proqramlaşdırma texnikalarını dəstəkləyir. Anonim funksiyalar Go proqramlarının daha funksional və səliqəli olmasına yardımçı olur.