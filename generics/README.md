# Generics 

**Generics**, Go 1.18 versiyasından etibarən dilə daxil edilmiş çox güclü bir xüsusiyyətdir. Generics, müxtəlif məlumat tipləri ilə işləyən funksiyalar və strukturlar yazmağa imkan verir, bu isə kodun yenidən istifadə olunabilirliyini artırır. Generics olmadan, biz adətən eyni funksiyanı müxtəlif tiplərə uyğun olaraq çoxaldırdıq, lakin generics sayəsində bu ehtiyac aradan qalxır.

## Generics-in Əsas Məntiqi

Generics bizə bir `function`ın və ya `struct`ın tip parametrlərindən asılı olmadan işləməsinə imkan verir. Go-da generics istifadə edərkən, biz **type parameters** (tip parametrləri) təyin edirik. Bu parametrlər `[]` kvadrat mötərizələrin içərisində müəyyən edilir və həmin tip parametrlər `function` və ya `struct` daxilində istifadə edilə bilər.

## Generics ilə Function Nümunəsi

Tutaq ki, biz iki elementi müqayisə edən bir funksiya yazmaq istəyirik. Bu funksiya həm integer, həm float, həm də string tiplərində işləyə bilməlidir.

```go
package main

import "fmt"

func Compare[T comparable](a, b T) bool {
    return a == b
}

func main() {
    fmt.Println(Compare(10, 10))       // True, both are int
    fmt.Println(Compare("Go", "Go"))   // True, both are string
    fmt.Println(Compare(3.14, 2.71))   // False, both are float64
}
```

### Bu Nümunədə:

- `Compare` funksiyası generik bir funksiyadır və `T` tip parametri ilə işləyir.
- `T` tipi `comparable` məhdudiyyəti ilə təyin olunmuşdur. Bu məhdudiyyət, `T` tipinin müqayisə edilə bilən (`==` və `!=` operatorları ilə) olmasını tələb edir.
- Eyni funksiyadan fərqli tiplərlə istifadə edə bilərik, məsələn, `int`, `string`, `float64` və s.

## Generics ilə Struct Nümunəsi

Generics sadəcə funksiyalarda deyil, həmçinin strukturlarda da istifadə edilə bilər. Bu, strukturların içərisində müxtəlif tiplərlə işləməyi mümkün edir.

```go
package main

import "fmt"

type Box[T any] struct {
    value T
}

func (b Box[T]) GetValue() T {
    return b.value
}

func main() {
    intBox := Box[int]{value: 100}
    fmt.Println("Integer Box:", intBox.GetValue())

    strBox := Box[string]{value: "Hello"}
    fmt.Println("String Box:", strBox.GetValue())
}
```

### Bu Nümunədə:

- `Box` adlı bir generik struktur təyin edirik. `T` tip parametri struktur daxilindəki `value` sahəsinin tipini təyin edir.
- `GetValue` metodu `T` tipini qaytarır, yəni bu metod dinamik olaraq fərqli tipləri geri qaytara bilər.
- Biz `Box` strukturundan həm integer, həm də string tipləri ilə istifadə edirik.

## Generics ilə Məhdudiyyətlər (Constraints)

Generics-də tiplər üçün məhdudiyyətlər təyin edə bilərik. Məhdudiyyətlər vasitəsilə, bir generik funksiyanın və ya strukturun hansı tip parametrləri qəbul edəcəyini idarə edə bilərik. Məsələn, `comparable`, `any`, və ya öz xüsusi interfeyslərimizdən istifadə edə bilərik.

```go
package main

import "fmt"

// Generic function with constraint
func Add[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(5, 10))       // Integer addition
    fmt.Println(Add(3.14, 2.71))  // Float addition
}
```

### Bu Nümunədə:

- `Add` funksiyası yalnız `int` və `float64` tiplərində işləyə bilər, çünki biz məhdudiyyət kimi `T int | float64` təyin etmişik.
- Funksiya iki parametri əlavə edir və onların tipinə uyğun olaraq qaytarır.

## Nəticə

Generics Go dilində kodu daha çevik və yenidən istifadə edilə bilən edir. Funksiyalar və strukturlar bir neçə məlumat tipi ilə işləyə bilər və generics ilə yazılan kodun daha az çoxaldılması ehtiyacı olur. Generics-dən istifadə edərkən tip parametrlərini və məhdudiyyətləri düzgün təyin etməklə kodun daha oxunaqlı və stabil olmasını təmin edə bilərik.