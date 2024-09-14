# Mutex 

**Mutex** (Mutual Exclusion - qarşılıqlı istisna) Go proqramlaşdırma dilində çoxlu işçi qovşaqların (goroutine) eyni zamanda bir resursa girişini idarə etmək üçün istifadə olunan sinxronizasiya obyektidir. Go dilində **sync** paketində mövcud olan **Mutex** strukturu, bir neçə işçi qovşağın eyni zamanda eyni resursa girməsinin qarşısını almaq üçün istifadə olunur. Bu, məlumatların təhlükəsizliyini və düzgünlüyünü təmin edir.

## Mutex-in Məntiqi

Goroutinlər eyni resurs üzərində eyni zamanda əməliyyatlar yerinə yetirməyə çalışarsa, bu zaman **race condition** (rəqabət vəziyyəti) yaranır. Bu halda məlumatların ardıcıllığı pozula bilər və gözlənilməz nəticələr ortaya çıxa bilər. **Mutex** istifadə edərək, hər hansı bir goroutine-in digərinə mane olmadan resursa daxil olmasını təmin edə bilərsiniz.

Mutex iki əsas funksiyaya malikdir:
- **Lock()**: Mutex kilidini açır. Bir goroutine bu funksiyanı çağırdıqdan sonra Mutex kilidlənir və digər goroutinlər həmin resursa giriş əldə edə bilməzlər.
- **Unlock()**: Mutex kilidini sərbəst buraxır. Mutex kilidindən sonra digər goroutinlər resursa giriş əldə edə bilər.

## Mutex Nümunəsi

Aşağıdakı nümunədə, Mutex istifadə edərək iki goroutine-in eyni resursa girişini sinxronizasiya edirik:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    count int
    mutex sync.Mutex
)

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    
    mutex.Lock() // Resursu kilidləyirik
    count++
    fmt.Println("Current Count:", count)
    mutex.Unlock() // Resursu sərbəst buraxırıq
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println("Final Count:", count)
}
```

### Nümunənin İzahı:

- **mutex.Lock()** funksiyası ilə resursa giriş kilidlənir. Yəni, hər hansı bir goroutine bu nöqtəyə çatdıqda, resurs kilidlənir və digər goroutinlər resursa giriş əldə edə bilmirlər.
- **mutex.Unlock()** funksiyası ilə kilid açılır və digər goroutinlər resursa giriş imkanı əldə edir.
- `count++` əməliyyatı Mutex ilə qorunduğu üçün birdən çox goroutine eyni anda bu əməliyyatı yerinə yetirə bilməz. Bu, **race condition**-dan qorunma təmin edir.

Bu nümunədə 5 fərqli goroutine `increment` funksiyasını çağırır və hər biri `count` dəyişənini bir artırır. Mutex istifadə edərək hər goroutine yalnız kilid açıldıqda əməliyyatı yerinə yetirə bilir. Son nəticə olaraq, `count` dəyişəninin dəyəri ardıcıl və düzgün şəkildə artırılır.

## Mutex-dən İstifadənin Vacibliyi

Mutex istifadə edilmədikdə, çoxlu goroutinlərin eyni resursa eyni anda giriş edə biləcəyi üçün `race condition` yaranır və bu vəziyyət gözlənilməz nəticələrə gətirib çıxara bilər. Buna görə də, resursun təhlükəsizliyi təmin edilməli və goroutinlər arasında düzgün sinxronizasiya aparılmalıdır.

Mutex-lərdən başqa, Go-da daha kompleks sinxronizasiya vasitələri olan **RWMutex**, **WaitGroup**, **Cond** və **Once** kimi obyektlər də mövcuddur. Hər biri xüsusi hallar üçün nəzərdə tutulmuşdur və daha effektiv sinxronizasiya və resurs idarəsi təmin edir.

## Nəticə

**Mutex**, çoxlu goroutinlərin eyni resursa eyni anda daxil olmalarının qarşısını almaq üçün Go-da çox vacib sinxronizasiya vasitəsidir. Kilidləmə və kilidi açma əməliyyatları vasitəsilə **race condition**-ların qarşısını almaq və məlumatların təhlükəsizliyini təmin etmək mümkündür.