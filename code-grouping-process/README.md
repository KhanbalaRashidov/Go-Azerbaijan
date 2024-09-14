# Code Grouping Process

Kodların qruplaşdırılması prosesi çox sadədir. Bu proses sayəsində eyni dəyişənlər müxtəlif kod bloklarında fərqli şəkildə işləyə bilər. Go dilində kodları qruplaşdırmaq üçün süslü mötərizələrdən istifadə edilir. Aşağıda bunun necə həyata keçirildiyini nümunə ilə göstərək:

```go
package main

import "fmt"

func main() {
    variable := "one"
    {
		variable := "two"
        fmt.Println(variable)
    }
    fmt.Println(variable)
}
```

### Output:
```
two
one
```

Yuxarıda `variable` adlı bir dəyişən yaratdıq. Dəyişənin içindəki qiymət `"one"`dir. Daha sonra isə süslü mötərizələr arasında yenidən `variable` adlı bir dəyişən elan etdik və ona `"two"` qiymətini verdik. Bu iki dəyişən eyni ada malikdir, lakin fərqli kod bloklarında olduqları üçün bir-biri ilə əlaqəsi yoxdur. Hər biri öz blokunda müxtəlif dəyərlər daşıyır.

Bu prosesin daha yaxşı başa düşülməsi üçün hər dəyişənin yaddaş ünvanına baxaq. Pointerlərdən istifadə edərək, hər dəyişənin yaddaşdakı yerini öyrənək.

```go
package main

import "fmt"

func main() {
	variable := "one"
    {
		variable := "two"
        fmt.Println(variable)
        fmt.Println(&variable)  // Dəyişənin yaddaş ünvanı
    }
    fmt.Println(variable)
    fmt.Println(&variable)      // Yenə yaddaş ünvanı
}
```

### Output:
```
two
0xc00008c1d0
one
0xc00008c1c0
```

Gördüyünüz kimi, `variable` adlı dəyişən iki fərqli yaddaş ünvanında saxlanılır. Bu, hər iki dəyişənin eyni ada malik olmasına baxmayaraq, müxtəlif kod bloklarında fərqli yaddaş yerlərində saxlanıldığını göstərir.