# Type Conversion in Go

Type conversion (tip çevirmə) Go dilində müxtəlif məlumat tipləri arasında çevirmələr etmək üçün istifadə olunur. Go, statik tipli bir dildir, bu səbəbdən bir tipdən digərinə çevirmə etmək üçün müvafiq çevirmə operatorlarından istifadə etməlisiniz.

## Tip Çevirmənin Əsasları

Go dilində tip çevirməsi bir dəyişənin bir tipdən digər tipə çevrilməsi üçün istifadə olunur. Çevirmə yalnız uyumlu tiplər arasında mümkündür. Tip çevirmə prosesində, konkret bir məlumat tipi digər bir tipə çevrilir, bu da müəyyən bir funksiyanı və ya əməliyyatı həyata keçirmək üçün lazım ola bilər.

## Tip Çevirmə Nümunələri

### Ən Sadə Tip Çevirmə

Aşağıdakı misalda, `int` tipində bir dəyişəni `float64` tipinə çevirmək göstərilir:

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i) // int tipindən float64 tipinə çevirmə
    fmt.Println(f) // 42.0
}
```

Burada, `i` dəyişəni `int` tipindədir və onu `float64` tipinə çevirmək üçün `float64(i)` ifadəsi istifadə olunur.

### Stringə Çevirmə

Bir dəyişəni `string` tipinə çevirmək üçün `string()` funksiyasından istifadə edilir:

```go
package main

import "fmt"

func main() {
    var i int = 123
    var s string = string(i) // int tipindən string tipinə çevirmə
    fmt.Println(s) // Əslində bu doğru nəticə verməyəcək, çünki bu ASCII kodunu alacaq.
}
```

Bu misalda, `int` tipindəki dəyişəni `string` tipinə çevirmək `string(i)` istifadə edilir, amma burada `i` dəyişəni ASCII kodu kimi təfsir olunur, bu səbəbdən nəticə gözlənilməz ola bilər.

### Numeric Tiplər Arasında Çevirmə

Fərqli numeric tiplər arasında çevirmə (məsələn, `int` və `float64` arasında) `float64` kimi daha geniş tipdən daha dar tiplərə çevrilərkən məlumat itkisi ola bilər.

```go
package main

import "fmt"

func main() {
    var f float64 = 3.14159
    var i int = int(f) // float64 tipindən int tipinə çevirmə
    fmt.Println(i) // 3
}
```

Bu misalda, `float64` tipindəki dəyişəni `int` tipinə çevirmək `int(f)` istifadə edilir. Çevirmə prosesində ondalıklı hissə itirilir.

### Type Conversion with Functions

Dəyişənlərin və funksiyaların parametr tipləri arasında çevirmələr də edə bilərsiniz:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var s string = "123"
    var i int
    var err error

    i, err = strconv.Atoi(s) // string tipindən int tipinə çevirmə
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(i) // 123
}
```

Burada, `strconv.Atoi` funksiyası string tipindəki dəyəri `int` tipinə çevirmək üçün istifadə olunur.

## Nəticə

Go dilində tip çevirmə, fərqli məlumat tipləri arasında çevirmələr edərək müxtəlif əməliyyatları həyata keçirmək üçün əhəmiyyətlidir. Çevirmə prosesində məlumat itkisini və potensial səhvləri nəzərə almaq vacibdir. Tip çevirməni düzgün istifadə edərək, proqramınızı daha effektiv və çevik hala gətirə bilərsiniz.