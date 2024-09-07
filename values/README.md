# Values

Dəyərlər Go proqramlaşdırma dilində dəyişənlər tərəfindən daşınan məlumatlardır. Dəyərlər sabit və ya dəyişən ola bilər və müxtəlif məlumat növlərini təmsil edir.

Go dilində əsas məlumat növləri bunlardır:

* Ədədi məlumat növləri: Ədədi dəyərləri ifadə etmək üçün int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr, float32, float64, complex64, complex128, byte, rune kimi məlumat növlərindən istifadə olunur.
* Boolean məlumat növü: bool məlumat növü yalnız doğru və ya yanlış qiymətlər qəbul edə bilən məlumat növüdür.
* String data type: string data type mətn və ya simvol sətirinin qiymətlərini ifadə etmək üçün istifadə olunur.
* Mürəkkəb məlumat növləri: struct, array, slice, map, channel kimi məlumat tipləri çoxsaylı məlumat elementlərini bir yerdə saxlamaq və müəyyən bir məqsədə xidmət etmək üçün istifadə olunur.

Dəyərlərin növləri Go dilində statik olaraq müəyyən edilir. Dəyişənə qiymət təyin etdikdə bu dəyişənin tipi müəyyən edilir və biz bu tipi sonra dəyişə bilmərik.

```go
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu nümunədə yaş dəyişəni ilk olaraq int kimi elan edilir və sonra dəyişən yaş üçün 32 rəqəmi təyin edilir. Sonra "thirty-two" sətri yaş dəyişəninə təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipli qiymət təyin edilə bilməz. Dəyərlər proqramları müxtəlif məqsədlər üçün istifadə etməyə imkan verir. Məsələn, rəqəmsal əməliyyatı yerinə yetirmək üçün int dəyərindən istifadə edə bilərik və ya mətn əməliyyatını yerinə yetirmək üçün string dəyərindən istifadə edə bilərik.
