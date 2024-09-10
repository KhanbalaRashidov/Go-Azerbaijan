# Worker pools

Go dilində, worker pool'lar müəyyən bir işi yerinə yetirmək üçün təyin olunmuş işçi qrupudur. İşlər `channel`-a göndərilir və bu işlər işçilər tərəfindən emal olunur. Bu metod, paralel emal və işlərin işçilər arasında paylanması ilə iş yükünü balanslaşdırmağa kömək edir.

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
```

Bu nümunədə, `worker` adlı bir funksiya yaradılır. Bu funksiya işçilər tərəfindən emal olunacaq işləri alır və nəticələri `results` kanalına geri qaytarır.

* `jobs` adlı bir kanal yaradılır və bu kanal vasitəsilə 100-ə qədər iş yükləri göndərilə bilər.
* `results` adlı başqa bir kanal yaradılır və 100 nəticə bu kanal vasitəsilə geri alına bilər.

İlk `for` dövrü ilə, 3 işçi üçün `goroutine` yaradılır. Hər işçi `worker` funksiyasını icra edərək, ona verilmiş işləri emal edir.

Növbəti `for` dövrü ilə, 5 iş göndərilir və bu işlər `jobs` kanalına daxil edilir. `close` funksiyası vasitəsilə `jobs` kanalı bağlanır.

Sonuncu `for` dövrü ilə, emal olunmuş nəticələr `results` kanalından alınır.

**Output:**

```go
Worker 3 started job 1
Worker 2 started job 2
Worker 1 started job 3
Worker 3 finished job 1
Worker 3 started job 4
Worker 2 finished job 2
Worker 1 finished job 3
Worker 1 started job 5
Worker 3 finished job 4
Worker 1 finished job 5
```

Bu nümunə, işlərin işçilər arasında necə paylanmasını və nəticələrin necə geri alındığını göstərir.
