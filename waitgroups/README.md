# WaitGroups

`WaitGroup` strukturu, Go dilində goroutine-lərin sinxronizasiyası üçün istifadə olunan bir mexanizmdir. Go-nun `sync` paketində tapılan bu struktur, birdən çox goroutine-in icrasını izləmək və onların hamısı tamamlandıqda müəyyən bir əməliyyatı yerinə yetirmək üçün imkan yaradır.

`WaitGroup`, birdən çox goroutine-i izləməyə, onların bitməsini gözləməyə və hər bir goroutine tamamlandıqda xəbərdar olmağa imkan verir. Bu məqsədlə, `Add()`, `Done()` və `Wait()` metodlarından istifadə edilir:

* `Add(n)`: `n` sayda goroutine əlavə edir.
* `Done()`: Goroutine bitdikdən sonra çağırılır və gözləmə sayını azaldır.
* `Wait()`: Bütün goroutine-lər tamamlanana qədər gözləyir.

Aşağıdakı nümunədə, `WaitGroup` istifadə edərək goroutine-lərin idarə olunması göstərilir:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Goroutine tamamlandıqda Done() çağırılır.

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)  // İşçinin işləməsi üçün bir saniyə gözləmə.
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup  // WaitGroup dəyişəni yaradılır.

    for i := 1; i <= 5; i++ {
        wg.Add(1)  // Hər goroutine üçün Add(1) edilir.
        go worker(i, &wg)  // Goroutine işə salınır.
    }

    wg.Wait()  // Bütün goroutine-lər bitənə qədər gözləyir.

    fmt.Println("All workers done")
}
```

**Nümunə açıqlaması:**

* `worker` funksiyası bir işçini təmsil edir. Hər işçi bir saniyə işləyir və tamamlandıqda `wg.Done()` çağırılır.
* `main` funksiyasında, 5 goroutine işə salınır və hər biri `worker` funksiyasını icra edir.
* `wg.Add(1)` hər işçi başlamazdan əvvəl sayğacı bir vahid artırır.
* `wg.Wait()` bütün goroutine-lər bitmədən `main` funksiyasının bitməsini gecikdirir.

&#x20;**Output:**

```go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 5 done
Worker 3 done
Worker 2 done
Worker 1 done
Worker 4 done
All workers done
```

Bu nümunədə, 5 goroutine paralel olaraq işləyir, hər biri tamamlandıqda `WaitGroup` onları izləyir və bütün işlər tamamlandıqda proqram "All workers done" mesajını yazdırır.
