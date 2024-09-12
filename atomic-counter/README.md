## Atomic Counters

Atomic Counters mövzusu, Go dilində paralelliyi idarə etmək üçün istifadə olunan bir mövzudur. Çoxsaylı Go proqramları bir neçə goroutine tərəfindən paylaşılan verilərlə işlədiyinə görə, bu verilərin eyni anda bir neçə goroutine tərəfindən dəyişdirilməsi labüddür. Bu dəyişikliklər düzgün idarə edilmədikdə, proqram gözlənilməz şəkildə işləyə bilər.

Go dilində Atomic Counter-lər `sync/atomic` paketində mövcuddur. Bu paketdə `AddInt64`, `CompareAndSwapInt64`, `SwapInt64`, `LoadInt64`, və `StoreInt64` kimi metodlar mövcuddur. Bu metodlar, bir neçə goroutine tərəfindən eyni anda daxil olunan dəyişkənin təhlükəsiz şəkildə dəyişdirilməsini təmin edir.

Aşağıdakı nümunə, Atomic Counter istifadə edərək, 10 goroutine tərəfindən paylaşılan bir sayaç dəyişkəninin təhlükəsiz şəkildə artırılmasını göstərir:

```golang
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64 = 0
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

Bu nümunədə, `counter` adlı dəyişkən `int64` tipində bir Atomic Counter olaraq təyin olunur və başlanğıcda 0 dəyəri ilə müəyyən edilir. Daha sonra, 10 goroutine yaradılır və hər biri `atomic.AddInt64` metodundan istifadə edərək `counter` dəyişkənini artırır. Bu metod `&counter` ünvanını alır və `counter` dəyişkəninə atomik şəkildə 1 əlavə edir.

Nəticədə, Atomic Counter veri quruluşu, çoxsaylı goroutine-lər tərəfindən paylaşılan dəyişkənlərin təhlükəsiz şəkildə dəyişdirilməsini təmin edir və paralellik idarəsində mühüm rol oynayır.