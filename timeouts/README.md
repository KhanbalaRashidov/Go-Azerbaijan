# Timeouts

Go dilində, **timeout** əməliyyatları xüsusilə şəbəkə əməliyyatları zamanı əhəmiyyətlidir. `time` paketi istifadə edilərək, müəyyən bir müddət gözləmək təmin edilə bilər. Əgər bu müddət keçərsə, timeout səhvi yaranır.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
```

Bu nümunədə, `c1` və `c2` adlı iki kanal yaradılır və mesaj göndərmə əməliyyatları üçün goroutin-lər yaradılır.

İlk olaraq, `select` açar sözü istifadə edilərək, `c1` kanalından mesaj gözlənilir. Amma `time.After` istifadə edilərək, `c1` kanalından bir mesaj alınmazdan əvvəl 1 saniyə gözlənilir. Əgər 1 saniyədən çox müddət keçərsə, timeout səhvi yaranır və "timeout 1" mesajı ekrana yazdırılır.

Daha sonra, `select` açar sözü ilə `c2` kanalından mesaj gözlənilir. Bu dəfə `time.After` ilə 3 saniyə gözlənilir və nəticə mesajı alınır.

#### Output:

```go
timeout 1
result 2
```

Bu nümunədə `time.After` istifadə edərək timeout əməliyyatları icra edilir. İlk nümunədə, `c1` kanalına 1 saniyədən əvvəl mesaj göndərilmədiyi üçün timeout səhvi yaranır. İkinci nümunədə isə, `c2` kanalına mesaj vaxtında göndərildiyi üçün mesaj qəbul edilir və ekrana yazdırılır.
