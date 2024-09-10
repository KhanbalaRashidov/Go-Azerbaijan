# Range over Channels

## Range over Channels

Go dilində, `range` açar sözü channel-lar üzərində dövr etməyə imkan verir. Bu üsul ilə channel bağlanana qədər bütün mesajlar `for` dövrü ilə oxuna bilər.

```go
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
```

Bu nümunədə:

* `queue` adlı bir channel yaradılır və bu channel-ə iki mesaj göndərilir.
* `close` funksiyası ilə channel bağlanır.
* Daha sonra, `range` açar sözü ilə for dövrü istifadə edərək, channel-dakı bütün mesajlar oxunur.

#### Output:

```go
one
two
```

Bu nümunədə `range` açar sözü ilə channel-dakı bütün mesajlar oxundu və `close` funksiyası ilə channel bağlandıqdan sonra belə, göndərilmiş bütün mesajlar alındı. `range` dövrü channel bağlanana qədər davam edir və channel-dan yeni mesajlar gəldikcə onları oxuyur.
