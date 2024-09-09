# Non-Blocking Channel Operations

Go dilində, channel-lar adətən bloklayan xüsusiyyətə malikdir, yəni bir goroutine bir channel-a mesaj göndərmək və ya almaq istədikdə həmin əməliyyat tamamlanana qədər gözləyir. Amma `select` açar sözü ilə channel-ların bloklanmadan işləməsi təmin edilə bilər.

```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no signal received")
	}

	signal := true
	select {
	case signals <- signal:
		fmt.Println("sent signal", signal)
	default:
		fmt.Println("no signal sent")
	}

	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no signal received")
	}
}
```

Bu nümunədə iki channel (`messages` və `signals`) yaradılır və onlardan bloklanmadan necə məlumat göndərmək və qəbul etmək göstərilir.

#### Output:

```go
no message received
sent message hello
no message received
no signal received
sent signal true
received signal true
```

Bu nümunədə:

* İlk öncə `messages` channel-dan mesaj alınmazdan əvvəl `no message received` mesajı göstərilir.
* Daha sonra `hello` mesajı `messages` channel-ına göndərilir və "sent message hello" yazısı ekrana çıxır.
* İkinci dəfə `messages` channel-dan mesaj almağa çalışılsa da, yeni mesaj olmadığından "no message received" göstərilir.
* Eyni üsulla `signals` channel-ı üçün sinyal göndərilməzdən əvvəl "no signal received" göstərilir, sinyal göndəriləndən sonra isə qəbul edildiyi göstərilir.

Bloklamadan əməliyyatlar `select` ilə kanal vasitəsilə sinxronlaşma və qarşılıqlı əlaqənin daha çevik formada həyata keçirilməsinə imkan yaradır.
