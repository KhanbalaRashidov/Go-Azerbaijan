# Məntiqi Operatorlar

Go proqramlaşdırma dilində məntiqi operatorlar şərtlər üzərində məntiqi əməliyyatlar aparmaq üçün istifadə olunur. Bu operatorlar əsasən şərti ifadələrin qiymətləndirilməsi zamanı istifadə edilir və nəticə olaraq `true` və ya `false` (boolean) dəyərlər qaytarır.

Go dilində üç əsas məntiqi operator mövcuddur:

- **Və (AND) operatoru** `&&`
- **Və ya (OR) operatoru** `||`
- **Deyil (NOT) operatoru** `!`

Aşağıda bu operatorların istifadəsi haqqında məlumat verilib.

## Və (AND) Operatoru `&&`

Və operatoru iki şərti qiymətləndirir və hər iki şərt `true` olduqda nəticə olaraq `true` qaytarır. Əgər şərtlərdən biri belə `false` olarsa, nəticə `false` olur.

```go
package main

import "fmt"

func main() {
    a := true
    b := false

    // Və operatoru hər iki şərt doğru olduqda nəticə doğru olur
    result := a && b // Nəticə: false
    fmt.Println("a və b:", result)

    c := 10
    d := 20

    // Şərtlərdən hər ikisi doğru olduğu üçün nəticə true olur
    result = (c < d) && (c != d) // Nəticə: true
    fmt.Println("c < d və c != d:", result)
}
```

Yuxarıdakı nümunədə, `a && b` ifadəsi qiymətləndirildikdə nəticə `false` olur, çünki `b` dəyişəni `false` dəyərinə malikdir. İkinci nümunədə isə hər iki şərt doğru olduğundan nəticə `true` olur.

## Və ya (OR) Operatoru `||`

Və ya operatoru iki şərti qiymətləndirir və hər hansı bir şərt `true` olduqda nəticə olaraq `true` qaytarır. Hər iki şərt `false` olduqda isə nəticə `false` olur.

```go
package main

import "fmt"

func main() {
    a := true
    b := false

    // Və ya operatoru hər hansı bir şərt doğru olduqda nəticə doğru olur
    result := a || b // Nəticə: true
    fmt.Println("a və ya b:", result)

    c := 10
    d := 20

    // Şərtlərdən biri doğru olduğu üçün nəticə true olur
    result = (c == d) || (c < d) // Nəticə: true
    fmt.Println("c == d və ya c < d:", result)
}
```

Bu nümunədə `a || b` ifadəsi qiymətləndirildikdə, `a` dəyişəni `true` olduğu üçün nəticə `true` olur. İkinci nümunədə isə `c < d` şərti doğru olduğu üçün nəticə yenə `true` qaytarılır.

## Deyil (NOT) Operatoru `!`

Deyil operatoru tək bir şərti tərsinə çevirir. Əgər şərt `true` olarsa, `!` operatoru nəticəni `false` edəcək və əksinə, şərt `false` olarsa, nəticə `true` olacaq.

```go
package main

import "fmt"

func main() {
    a := true
    b := false

    // Deyil operatoru şərti tərsinə çevirir
    result := !a // Nəticə: false
    fmt.Println("!a:", result)

    result = !b // Nəticə: true
    fmt.Println("!b:", result)
}
```

Bu nümunədə `!a` ifadəsi qiymətləndirildikdə, `a` dəyişəni `true` olduğu üçün nəticə `false` olur. Eyni qaydada `!b` ifadəsi qiymətləndirildikdə, `b` dəyişəni `false` olduğu üçün nəticə `true` olur.

## Nəticə

Məntiqi operatorlar Go proqramlaşdırma dilində şərti ifadələri qiymətləndirmək üçün çox faydalıdır. `&&` (Və), `||` (Və ya), və `!` (Deyil) operatorları müxtəlif şərtlərin kombinasiya edilməsi və qiymətləndirilməsi üçün geniş istifadə edilir. Bu operatorlar if-else, dövrlər və digər məntiqi quruluşlar ilə sıx bağlıdır.