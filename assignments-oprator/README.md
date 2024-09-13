# Assignment Operators in Go

Assignment operators (təyin etmə operatorları) proqramlaşdırma dilində bir dəyişənə dəyər təyin etmək üçün istifadə olunur. Go dilində müxtəlif assignment operatorları mövcuddur və aşağıda onların əsas növləri göstərilmişdir:

1. **Assignment (`=`)**  
   Bu operator bir dəyişənə dəyər təyin edir. Əgər dəyişən daha əvvəl təyin edilməyibsə, əvvəlcə onu elan etməlisiniz.

   ```go
   var a int
   a = 10
   ```

   Burada, `a` dəyişəninə 10 dəyəri təyin edilir.

2. **Add and Assign (`+=`)**  
   Bu operator mövcud dəyişənin dəyərinə müəyyən bir dəyəri əlavə edir və nəticəni həmin dəyişənə təyin edir.

   ```go
   x := 5
   x += 3 // x = x + 3
   fmt.Println(x) // 8
   ```

   Burada, `x` dəyişəninə 3 əlavə edilir və nəticə `x`-ə təyin olunur.

3. **Subtract and Assign (`-=`)**  
   Bu operator mövcud dəyişənin dəyərindən müəyyən bir dəyəri çıxarır və nəticəni həmin dəyişənə təyin edir.

   ```go
   x := 10
   x -= 4 // x = x - 4
   fmt.Println(x) // 6
   ```

   Burada, `x` dəyişənindən 4 çıxarılır və nəticə `x`-ə təyin olunur.

4. **Multiply and Assign (`*=`)**  
   Bu operator mövcud dəyişənin dəyərini müəyyən bir dəyərlə vurur və nəticəni həmin dəyişənə təyin edir.

   ```go
   x := 6
   x *= 7 // x = x * 7
   fmt.Println(x) // 42
   ```

   Burada, `x` dəyişəni 7 ilə vurulur və nəticə `x`-ə təyin olunur.

5. **Divide and Assign (`/=`)**  
   Bu operator mövcud dəyişənin dəyərini müəyyən bir dəyərlə bölür və nəticəni həmin dəyişənə təyin edir.

   ```go
   x := 20
   x /= 4 // x = x / 4
   fmt.Println(x) // 5
   ```

   Burada, `x` dəyişəni 4-ə bölünür və nəticə `x`-ə təyin olunur.

6. **Modulus and Assign (`%=`)**  
   Bu operator mövcud dəyişənin dəyərini müəyyən bir dəyərlə modulyasiya edir (qalanı tapır) və nəticəni həmin dəyişənə təyin edir.

   ```go
   x := 17
   x %= 5 // x = x % 5
   fmt.Println(x) // 2
   ```

   Burada, `x` dəyişəninin 5-ə bölünməsi ilə qalan 2 olur və bu nəticə `x`-ə təyin olunur.

## Misal

Aşağıdakı kod parçaları müxtəlif assignment operatorlarının istifadəsini göstərir:

```go
package main

import "fmt"

func main() {
    a := 10
    fmt.Println("Initial value of a:", a)

    a += 5
    fmt.Println("After a += 5:", a) // 15

    a -= 3
    fmt.Println("After a -= 3:", a) // 12

    a *= 2
    fmt.Println("After a *= 2:", a) // 24

    a /= 4
    fmt.Println("After a /= 4:", a) // 6

    a %= 3
    fmt.Println("After a %= 3:", a) // 0
}
```

## Nəticə

Assignment operators Go proqramlaşdırma dilində dəyişənlərə dəyər təyin etməyə və bu dəyərləri dəyişməyə imkan verir. Bu operatorlar müxtəlif hesablama əməliyyatlarını sadə və qısa şəkildə yerinə yetirmək üçün istifadə olunur. Onların düzgün istifadəsi proqramın səmərəliliyini artırır və kodu daha oxunaqlı edir.