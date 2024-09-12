# Defer

`defer`, Go dilindəki xüsusi bir açar sözdür və müəyyən bir funksiyanın sonunda icra ediləcək funksiyaları və ya ifadələri təyin etmək üçün istifadə olunur. Defer ifadələri, funksiyanın sonunda nə olursa olsun, yəni funksiyanın hər hansı bir səbəbdən sona çatması halında belə icra ediləcəkdir.

Defer ifadələrini istifadə edərək, bir funksiyanın sonunda açılmış faylları, bağlanmamış verilənlər bazası bağlantılarını, şəbəkə bağlantılarını və s. tez və təhlükəsiz bir şəkildə bağlamaq mümkündür.

Nümunə olaraq, bir fayl açılır və `defer` ilə funksiyanın sonunda fayl bağlanılır:

```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer f.Close()

	// faylı oxuma
	b := make([]byte, 1024)
	n, err := f.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(b[:n]))
}
```

Bu nümunədə, `os.Open` funksiyası istifadə edilərək "example.txt" faylı açılır. `defer` ifadəsi istifadə edilərək fayl bağlama əməliyyatı funksiyanın sonunda icra ediləcək şəkildə planlaşdırılır. Daha sonra, `Read` funksiyası istifadə edilərək faylın məzmunu oxunur və ekrana çap edilir.

Burada diqqət edilməli mühüm bir məqam, `defer` ifadəsinin ən sona yazılmamasıdır. `defer` ifadəsi, bağlanacaq olan faylı açan ifadə ilə eyni blokda olmalıdır. Əks halda, `defer` ifadəsi funksiyanın sonuna qədər gözləməyə davam edəcəkdir.