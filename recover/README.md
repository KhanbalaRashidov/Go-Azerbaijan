## Recover

`recover` funksiyası, bir Go proqramının işləmə müddətində meydana gələn panic-ləri idarə etmək üçün istifadə olunur. Əgər bir panic baş verərsə, `recover` funksiyası panic-in baş verdiyi funksiyanın içində çağırılarsa, panic səbəbindən dayandırılmış əməliyyatın idarəsini geri alır və proqramın normal şəkildə davam etməsini təmin edir.

`recover` funksiyası adətən `defer` ifadələri ilə birlikdə istifadə olunur. Bu sayədə, mümkün bir panic vəziyyətində proqramın idarəsi `defer` ifadələrindəki əməliyyatlardan ən sonuncusu olan `recover` funksiyasına ötürülür və beləliklə proqramın davam etməsi təmin edilir.

Nümunə olaraq, bir panic baş verdikdə proqramın işləməsini dayandırmaq əvəzinə, `recover` funksiyası vasitəsilə proqramın normal şəkildə davam etməsi təmin edilə bilər:

```golang
package main

import "fmt"

func main() {
    fmt.Println("Program starting...")

    // Defer ifadəsi istifadə edərək panic vəziyyətlərində işlədiləcək funksiyanı müəyyən edirik.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic occurred: This is a panic!:", r)
        }
    }()

    fmt.Println("Program continues...")

    // Burada bilərəkdən panic yaradırıq.
    panic("This is a panic")
}
```

Bu nümunədə, `panic` ifadəsi istifadə edilərək bilərəkdən bir panic yaradılır. `defer` ifadəsi istifadə edərək panic vəziyyətlərində işlədiləcək bir funksiya müəyyən edilir. `recover` funksiyası, funksiyanın içində çağırıldığı zaman, panic vəziyyətindən geri dönən dəyəri alır və `if` strukturu istifadə edilərək bu dəyər ilə işləyir.

```
Program starting...
Program continues...
Panic occurred: This is a panic!
This is a panic
```


# String Functions

Golang, string məlumatları üzərində əməliyyatlar aparmaq üçün bir çox hazır funksiyalar təqdim edir. Bu funksiyalar, string məlumatlarını parçalamaq, birləşdirmək, müqayisə etmək, axtarmaq, dəyişdirmək və daha bir çox əməliyyat aparmaq üçün istifadə edilə bilər.

Bəzi çox istifadə olunan string funksiyaları bunlardır:

- `len(s string) int`: Verilmiş string'in uzunluğunu qaytarır.
- `s[i] byte`: Verilmiş string'dəki göstərilən indeks nömrəsindəki simvolu qaytarır.
- `s + t`: İki string'i birləşdirir.
- `strings.Split(s, sep string) []string`: Verilmiş string'i, göstərilən ayırıcı simvollara görə parçalayır və bir array olaraq return edir.
- `strings.Join(a []string, sep string) string`: Verilmiş string arrayini, göstərilən ayırıcı simvol ilə birləşdirir və tək bir string olaraq return edir.
- `strings.Contains(s, substr string) bool`: Verilmiş string daxilində, göstərilən alt string'in olub olmadığını yoxlayır.
- `strings.HasPrefix(s, prefix string) bool`: Verilmiş string'in, göstərilən prefix ilə başlayıb başlamadığını yoxlayır.
- `strings.HasSuffix(s, suffix string) bool`: Verilmiş string'in, göstərilən suffix ilə bitib-bitmədiyini yoxlayır.
- `strings.Replace(s, old, new string, n int) string`: Verilmiş string'də, göstərilən köhnə string'i, göstərilən yeni string ilə dəyişdirir. İstəyə bağlı olaraq, göstərilən sayda dəyişdirmə edir.

Nümunə olaraq, bir string daxilində müəyyən bir xarakter dizisinin neçə dəfə keçdiyini tapmaq üçün `strings.Count` funksiyası istifadə edilə bilər:

```golang
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
}
```

Output:

```go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST
```