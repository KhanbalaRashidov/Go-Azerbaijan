# Select

Go dilində, select açar sözü bir neçə kanalı (channel) dinləyərək hansı kanalın mesaj göndərdiyini müəyyənləşdirə bilər. Bu xüsusiyyət kanalların sinxronizasiyasını asanlaşdırır və fərqli goroutin-lər arasında mesajlaşmanı idarə edir.

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```

Bu nümunədə, c1 və c2 adlı iki kanal yaradılır və mesaj göndərmə prosesi üçün goroutin-lər yaradılır.

main funksiyasında, select açar sözü istifadə edilərək, c1 və c2 kanalları dinlənilir. İlk olaraq, goroutin-lər arasındakı gözləmə müddətinə görə, c1 kanalından bir mesaj alınır və ekrana yazdırılır. Daha sonra, c2 kanalından bir mesaj alınır və ekrana yazdırılır.

Output:
```
received one
received two
```

Bu nümunədə, select açar sözü istifadə edilərək, c1 və c2 kanallarını dinləyən bir for döngüsü yaradıldı. Bu, mesaj alım müddətinə əsaslanaraq fərqli kanalların dinlənilməsinə imkan verir. Nəticədə, goroutin-lər arasındakı mesajlaşma müəyyən bir qaydada həyata keçirilir və select açar sözü istifadə edilərək sinxronizasiya təmin edilir.





