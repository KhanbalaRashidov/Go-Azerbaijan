# Golang-da Dinamik Dəyişənlər və `any` Tipi

Go 1.18 versiyasından etibarən, Go dilinə yeni bir tip sinonimi olan `any` əlavə edildi. `any`, əslində, `interface{}` boş interfeysinin sinonimidir və proqramın daha oxunaqlı olmasını təmin etmək məqsədini daşıyır. Yəni, `any` istifadə edərək, `interface{}`-nin gördüyü işləri daha sadə şəkildə ifadə edə bilərik.

## Boş Interfeys (`interface{}`) və `any`

Go-da dinamik dəyişənlər digər dinamik dillərdə olduğu kimi istifadə edilməsə də, `interface{}` və `any` tipləri ilə dinamikliyə nail ola bilirik. Bu tiplər vasitəsilə, bir dəyişənə müxtəlif tiplərdə dəyərlər təyin edə bilərik.

### `any` ilə Nümunə:

```go
package main

import "fmt"

func main() {
    var dynamicVar any
	
    dynamicVar = 42
    fmt.Println("Integer value:", dynamicVar)

    dynamicVar = "Hello, Go!"
    fmt.Println("String value:", dynamicVar)

    dynamicVar = 3.14
    fmt.Println("Float value:", dynamicVar)
}
```

### Output:
```
Integer value: 42
String value: Hello, Go!
Float value: 3.14
```

Bu nümunədə biz `any` tipindən istifadə edərək, bir dəyişənə müxtəlif tiplərdə qiymətlər təyin edirik. `any` əslində `interface{}`-nin sinonimidir və Go 1.18 ilə daha oxunaqlı şəkildə yazmaq üçün tətbiq olunmuşdur.

### Məlumatın Tipinin Yoxlanılması

`any` tipi ilə işləyərkən, `interface{}` ilə olduğu kimi, məlumatın tipini yoxlamaq üçün **type assertion** (tip təsdiqi) və **type switch** (tip çevirmə) üsullarından istifadə edə bilərik.

```go
package main

import "fmt"

func main() {
    var dynamicVar any = "Hello, Go!"

    // Type assertion with `any`
    str, ok := dynamicVar.(string)
    if ok {
        fmt.Println("String value:", str)
    } else {
        fmt.Println("Not a string")
    }

    // Type switch with `any`
    switch v := dynamicVar.(type) {
    case string:
        fmt.Println("String value:", v)
    case int:
        fmt.Println("Integer value:", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

Bu nümunədə, `any` tipi ilə təyin edilmiş bir dəyişənin tipi yoxlanılır və uyğun əməliyyatlar həyata keçirilir.

## Nəticə

Golang-da dinamik dəyişənlər birbaşa mövcud olmasa da, `interface{}` və onun sinonimi olan `any` tipləri vasitəsilə dinamik tiplərlə işləmək mümkündür. Bu tiplər, dəyişənlərə müxtəlif məlumat tipləri təyin etməyə və daha çevik proqramlar yazmağa imkan verir. `any` tipi ilə Go dilində kod yazmaq daha oxunaqlı olur və Go 1.18-dən etibarən daha geniş istifadə olunur.