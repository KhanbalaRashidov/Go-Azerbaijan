# Errors

Go dilində səhv (error) idarəetməsi, error adlı bir məlumat tipi vasitəsilə həyata keçirilir. Bu tip, bir funksiyanın nəticəsi olaraq ya bir səhv mesajı, ya da nil dəyərini qaytarmaq üçün istifadə olunur. Bu, proqramın gözlənilməyən vəziyyətlərdə necə davranmalı olduğunu idarə etməyə kömək edir.

```golang
package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
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

Bu nümunədə, divide adlı bir funksiya təyin edilir və iki ədədi bölür. Lakin, əgər ikinci ədəd 0 olarsa, bir səhv qaytarır.

main funksiyasında, divide funksiyası iki fərqli parametr ilə çağırılır. err adlı bir səhv dəyişəni istifadə edilərək hər bir çağırışın nəticəsi yoxlanılır. Əgər səhv varsa, səhv mesajı ekrana yazdırılır. Əks halda, nəticə ekrana yazdırılır.

Output:
```
5 <nil>
Cannot divide by zero
```

Bu nümunədə divide funksiyası sıfıra bölmə vəziyyətini idarə etmək üçün səhv idarəetməsindən istifadə edir. Əgər y sıfır olarsa, errors.New ilə yaradılmış səhv qaytarılır. Proqram err dəyişəni ilə səhv olub-olmadığını yoxlayır və əgər səhv varsa, onu ekrana yazdırır. Bu yanaşma Go dilində səhv idarəetməsini daha etibarlı və idarə olunan edir.

<br><br>




