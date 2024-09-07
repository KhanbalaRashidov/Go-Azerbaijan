# Channel

Go dilində kanal (channel), goroutine-lər arasında məlumat ötürmək üçün istifadə olunan bir məlumat strukturudur. Kanal, make açar sözü ilə yaradılır və <- operatoru ilə məlumat göndərmə və qəbul etmə əməliyyatları həyata keçirilir.

```golang
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "Hello" }()

	msg := <-messages
	fmt.Println(msg)
}
```

Bu nümunədə, messages adlı bir kanal yaradılır.

go açar sözü ilə bir goroutine yaradılır və bu goroutine messages kanalına "Hello" mesajını göndərir.

main funksiyasında, msg adlı bir dəyişkənə messages kanalından bir mesaj alınır və ekrana yazdırılır.

Output:
```
Hello
```

Bu örnekte, channel kullanarak, bir goroutine'dan ana iş parçacığına bir mesaj gönderildi. Bu nedenle, goroutine işlemi tamamlandıktan sonra ana iş parçacığı channel'dan mesajı alır ve sonucu ekrana yazdırır.

Kanallar, Golang'de birçok durumda kullanılabilir, örneğin:

Bu nümunədə, kanal istifadə edərək bir goroutine-dən ana iş parçacığına bir mesaj göndərildi. Bu səbəbdən, goroutine tamamlandıqdan sonra ana iş parçacığı kanaldan mesajı alır və nəticəni ekrana yazdırır.

Kanallar Go dilində bir çox vəziyyətdə istifadə oluna bilər, məsələn:

1. Goroutine-lər arasında məlumat mübadiləsi üçün
2. Sinxronizasiya əməliyyatları üçün
3. Tətbiqinizin performansını artırmaq üçün (paralel emal etmək)
4. Goroutine-lər arasında məlumat yarışlarını qarşısını almaq üçün
5. Tapşırıqların koordinasiyası və sinxronizasiyası üçün

<br><br>




