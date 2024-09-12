## Rate Limiting

Rate Limiting, bir API və ya xidmətin istifadəçinin müəyyən bir zaman ərzində edə biləcəyi sorğu sayını məhdudlaşdırmaq üçün istifadə edilən bir metoddur. Bu metod, tətbiqin və ya xidmətin həddindən artıq yüklənməsinin qarşısını almaq üçün istifadə oluna bilər.

Go dilində, Rate Limiting etmək üçün `time` paketi istifadə oluna bilər. Bu paket, müəyyən bir zaman çərçivəsində müəyyən bir əməliyyatın yerinə yetirilməsi üçün nə qədər gözləmək lazım olduğunu hesablamaq üçün istifadə edilir.

Aşağıdakı nümunədə, `time.Ticker` və `time.Sleep` istifadə edilərək Rate Limiting nümunəsi göstərilir:

```golang
package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
}
```

Bu nümunədə, `requests` adlı bir kanal yaradılır və içinə 5 sorğu əlavə edilir. Daha sonra `limiter` adlı bir `time.Ticker` yaradılır və 200 millisaniyəlik bir müddətlə məhdudlaşdırılır.

Daha sonra, `requests` kanalındakı hər bir sorğu məhdudlaşdırıcıya uyğun olaraq işlədir. Hər sorğu arasında 200 millisaniyəlik fasilə ilə işlənməsi təmin edilir.

```
request 1 2023-05-23 15:56:01.46705 +0300 +03 m=+0.201167209
request 2 2023-05-23 15:56:01.667008 +0300 +03 m=+0.401132584
request 3 2023-05-23 15:56:01.867014 +0300 +03 m=+0.601146918
request 4 2023-05-23 15:56:02.067052 +0300 +03 m=+0.801193043
request 5 2023-05-23 15:56:02.267025 +0300 +03 m=+1.001173793
```

Bu şəkildə, Rate Limiting istifadə edərək sorğular müəyyən bir sürətlə işlənir.