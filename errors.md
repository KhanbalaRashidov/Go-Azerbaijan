# Errors

Go dilində səhv idarəetməsi `error` adlı bir məlumat tipi ilə həyata keçirilir. `error` tipi, səhv mesajı və ya `nil` dəyəri qaytarmaq üçün istifadə olunur.

```go
package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Sıfıra bölmək mümkün deyil")
	}
	return x / y, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

Bu nümunədə, `divide` adlı bir funksiya təyin edilir və iki ədədi bölür. Lakin, əgər ikinci ədəd `0` olarsa, bir səhv qaytarır.

`main` funksiyasında, `divide` funksiyası iki fərqli parametr ilə çağırılır. `err` adlı bir səhv dəyişəni istifadə edilərək hər bir çağırışın nəticəsi yoxlanılır. Əgər səhv varsa, səhv mesajı ekrana yazdırılır. Əks halda, nəticə ekrana yazdırılır.

#### Output:

```go
5
Sıfıra bölmək mümkün deyil
```
