# Aritmetik Operatorlar

Go proqramlaşdırma dilində aritmetik operatorlar ədədi dəyərlər üzərində riyazi əməliyyatlar aparmaq üçün istifadə olunur. Bu operatorlar cəmləmə, çıxma, vurma, bölmə və qalıq kimi əməliyyatları həyata keçirir. Aritmetik operatorlar əsasən integer (tam ədədlər), float (ondalıq ədədlər) kimi növlər üzərində işləyir.

Aşağıda Go dilində mövcud olan əsas aritmetik operatorlar və onların istifadəsi haqqında məlumat verilib.

## Cəmləmə operatoru (`+`)

Cəmləmə operatoru iki ədədin cəmini hesablayır.

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    c := a + b // Cəmləmə əməliyyatı: 10 + 5 = 15
    fmt.Println("Cəmi:", c)
}
```

Yuxarıdakı nümunədə, `a + b` əməliyyatı iki tam ədədin cəmini hesablayır və nəticə `c` dəyişəninə təyin olunur.

## Çıxma operatoru (`-`)

Çıxma operatoru iki ədədin fərqini hesablayır.

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    c := a - b // Çıxma əməliyyatı: 10 - 5 = 5
    fmt.Println("Fərq:", c)
}
```

Bu nümunədə, `a - b` əməliyyatı ilə iki tam ədədin fərqi hesablanır.

## Vurma operatoru (`*`)

Vurma operatoru iki ədədin hasilini hesablayır.

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    c := a * b // Vurma əməliyyatı: 10 * 5 = 50
    fmt.Println("Hasil:", c)
}
```

Bu nümunədə, `a * b` iki ədədin hasilini hesablayır.

## Bölmə operatoru (`/`)

Bölmə operatoru iki ədədin bölünməsini həyata keçirir. Lakin, Go dilində tam ədədlərlə bölmə zamanı nəticə tam ədəd olur. Əgər dəqiq nəticə istənilirsə, ondalıq növlərdən (`float32`, `float64`) istifadə edilməlidir.

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5
    c := a / b // Bölmə əməliyyatı: 10 / 5 = 2
    fmt.Println("Nəticə:", c)

    x := 10.0
    y := 4.0
    z := x / y // Ondalıq ədədlər üzərində bölmə: 10.0 / 4.0 = 2.5
    fmt.Println("Ondalıq nəticə:", z)
}
```

Bu nümunədə, tam ədədlərlə bölmə nəticəsi tam ədəd kimi göstərilir, ondalıq növlərlə bölmə isə dəqiq nəticə verir.

## Qalıq operatoru (`%`)

Qalıq operatoru iki ədədin bölünməsindən qalan qalıq dəyərini qaytarır.

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    c := a % b // Qalıq əməliyyatı: 10 % 3 = 1
    fmt.Println("Qalıq:", c)
}
```

Bu nümunədə, `a % b` əməliyyatı 10-un 3-ə bölünməsindən sonra qalan qalıq olan `1` dəyərini qaytarır.

## Nəticə

Aritmetik operatorlar Go proqramlaşdırma dilində ədədi dəyərlər üzərində əsas riyazi əməliyyatları aparmaq üçün istifadə olunur. Bu operatorlar proqramın müxtəlif hesablamalarında geniş istifadə edilir.