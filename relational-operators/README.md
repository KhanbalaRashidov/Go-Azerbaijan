# Relational Operators in Go

Relational operators (qarşılaşdırma operatorları) proqramlaşdırmada iki dəyəri müqayisə etmək üçün istifadə olunur. Go proqramlaşdırma dilində də relational operators mövcuddur və aşağıdakıları əhatə edir:

1. **Equal to (bərabərdir) `==`**  
   İki dəyərin bərabər olub-olmadığını yoxlayır. Əgər dəyərlər bərabərdirsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 20
   fmt.Println(a == b) // false
   ```

2. **Not equal to (bərabər deyil) `!=`**  
   İki dəyərin bərabər olmadığını yoxlayır. Əgər dəyərlər bərabər deyilsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 20
   fmt.Println(a != b) // true
   ```

3. **Greater than (böyükdür) `>`**  
   Bir dəyərin digərindən böyük olub-olmadığını yoxlayır. Əgər birinci dəyər ikinci dəyərdən böyükdürsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 20
   fmt.Println(a > b) // false
   ```

4. **Less than (kiçikdir) `<`**  
   Bir dəyərin digərindən kiçik olub-olmadığını yoxlayır. Əgər birinci dəyər ikinci dəyərdən kiçikdirsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 20
   fmt.Println(a < b) // true
   ```

5. **Greater than or equal to (böyük və ya bərabərdir) `>=`**  
   Bir dəyərin digərindən böyük və ya bərabər olub-olmadığını yoxlayır. Əgər birinci dəyər ikinci dəyərdən böyükdür və ya bərabərdirsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 10
   fmt.Println(a >= b) // true
   ```

6. **Less than or equal to (kiçik və ya bərabərdir) `<=`**  
   Bir dəyərin digərindən kiçik və ya bərabər olub-olmadığını yoxlayır. Əgər birinci dəyər ikinci dəyərdən kiçikdir və ya bərabərdirsə, `true`, əks halda `false` qaytarır.

   ```go
   a := 10
   b := 20
   fmt.Println(a <= b) // true
   ```

## Misal

Aşağıdakı kod parçaları relational operatorlardan istifadəni göstərir:

```go
package main

import "fmt"

func main() {
    x := 15
    y := 20

    fmt.Println("x == y:", x == y) // false
    fmt.Println("x != y:", x != y) // true
    fmt.Println("x > y:", x > y)   // false
    fmt.Println("x < y:", x < y)   // true
    fmt.Println("x >= y:", x >= y) // false
    fmt.Println("x <= y:", x <= y) // true
}
```

## Nəticə

Relational operators proqramlaşdırma dilində dəyərlərin müqayisə edilməsi üçün istifadə olunur. Bu operatorlar müxtəlif şərtləri yoxlamağa kömək edir və şərtli ifadələrdə geniş istifadə olunur. Go dilində bu operatorların doğru istifadə edilməsi proqramların daha düzgün və effektiv işləməsini təmin edir.