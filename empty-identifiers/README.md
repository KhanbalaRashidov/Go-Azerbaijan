# Boş Identifikatorlar (Empty Identifiers) Go dilində

Go proqramlaşdırma dilində boş identifikator, `_` işarəsi ilə təmsil olunur və bu, dəyəri göz ardı etmək üçün istifadə olunur. Boş identifikator, müəyyən bir kontekstdə istifadə edilməyəcək dəyərləri nəzərə almadan, sadəcə yer tutucu kimi xidmət edir. Bu xüsusiyyət müxtəlif hallarda, məsələn, çoxlu geri dönüş dəyərləri ilə işləmək, dəyişənləri göz ardı etmək və daha çox istifadə olunur.

## Boş Identifikatorların Əsas İstifadə Sahələri

### 1. İstifadəsiz Geri Dönüş Dəyərlərini Gözardı Etmək

Go dilində funksiyalar bir neçə dəyər qaytara bilər və əgər siz bu dəyərlərin yalnız bəzilərini istifadə etmək istəyirsinizsə, `_` ilə istifadə etmədiyiniz dəyərləri gözardı edə bilərsiniz.

**Misal:**

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Bir neçə geri dönüş dəyəri olan funksiya
    value, _ := strconv.Atoi("123") // Xəta dəyərini gözardı edirik

    fmt.Println(value) // Çıxış: 123
}
```

Bu misalda, `strconv.Atoi` funksiyası tam ədəd və xəta qaytarır. Xəta bu xüsusi istifadədə vacib deyil, ona görə `_` ilə gözardı edilir.

### 2. İstifadəsiz Dəyişənləri Gözardı Etmək

Dəyişənləri elan edib amma istifadə etmədikdə, onları `_`-ə təyin edərək, məqsədli olaraq istifadə etmədiyinizi göstərə bilərsiniz.

**Misal:**

```go
package main

import "fmt"

func main() {
    // İstifadəsiz dəyişən
    _, b := 10, 20

    fmt.Println(b) // Çıxış: 20
}
```

Burada, `a` dəyişəni `_`-ə təyin edilib və yalnız `b` istifadə olunur.

### 3. `range` Dövrlərində Dəyərləri Gözardı Etmək

`range` açar sözünü istifadə edərək dilimlər, xəritələr və ya kanallarla iterasiya edərkən, yalnız dəyərləri istifadə etmək və indeksləri və ya açarları gözardı etmək istəyə bilərsiniz. Bu halda `_` istifadə edərək lazımsız hissələri gözardı edə bilərsiniz.

**Misal:**

```go
package main

import "fmt"

func main() {
    meyvələr := []string{"Alma", "Banan", "Çiyələk"}

    for _, meyvə := range meyvələr {
        fmt.Println(meyvə)
    }
}
```

Bu döngüdə, indeks `_` ilə gözardı edilir və yalnız dəyər (`meyvə`) istifadə olunur.