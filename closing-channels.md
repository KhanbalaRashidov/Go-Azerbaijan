# Closing Channels

Go dilində, channel-ların bağlanması, məlumatların göndərilməsi və alınması prosesində sinxronizasiyanı təmin edir. `close` funksiyası vasitəsilə channel bağlanır. Bağlanmış bir channel-a mesaj göndərmək mümkün deyil və həmin channel-dan daha çox mesaj alınmaz.

```go
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
```

Bu nümunədə:

* `jobs` adlı bir channel yaradılır və işlər bu kanala göndərilir.
* `done` adlı bir channel yaradılır və işlərin tamamlandığını bildirmək üçün istifadə olunur.
* Bir `goroutine` yaradılır və bu `goroutine`, `jobs` channel-ından mesajları alır. Alınan mesajların sayı azaldıqca hər dəfə "received job" yazısı çap edilir.
* `for` döngüsü ilə `jobs` channel-ına 3 iş göndərilir və hər göndərmə zamanı "sent job" mesajı göstərilir.
* `jobs` channel-ı bağlandıqdan sonra `goroutine`, channel-dan bütün mesajları alır və son olaraq "received all jobs" mesajı göstərilir.

#### Output:

```go
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
```

Bu nümunə, Go dilində bir channel-a mesaj göndərdikdən sonra onu necə bağlayacağınızı və bağlanmış channel-dan daha çox mesaj alınmadığını göstərir. Bağlı bir channel-dan oxumağa davam edə bilərsiniz, lakin artıq yeni mesajlar göndərilə bilməz.
