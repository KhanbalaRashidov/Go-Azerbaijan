# Ticker

Go dilində, ticker'lər müəyyən bir müddət ərzində müəyyən aralıqlarla bir əməliyyatı yerinə yetirmək üçün istifadə olunur. `time` paketi daxilində yer alan `NewTicker` funksiyası istifadə edilərək, bir ticker yaradılır.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```

Bu nümunədə, `NewTicker` funksiyası istifadə edilərək, bir ticker yaradılır və hər 500 millisekundda bir əməliyyat yerinə yetirilir.

`goroutine` istifadə edilərək, `select` açar sözü ilə, ticker'ın müəyyən edilmiş zaman aralığına uyğun olaraq müəyyən aralıqlarla əməliyyat həyata keçirilir. `done` adlı kanal vasitəsilə mesaj göndərilərək, `goroutine` dayandırılır və "Ticker stopped" mesajı ekrana yazdırılır.

Output:

```go
Tick at 2022-11-27 16:37:14.750915 +0300 MSK m=+0.501356162
Tick at 2022-11-27 16:37:15.249568 +0300 MSK m=+1.000018689
Tick at 2022-11-27 16:37:15.751064 +0300 MSK m=+1.501505306
Ticker stopped
```
