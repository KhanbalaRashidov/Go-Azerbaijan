# Arrays

Array , Go proqramlaşdırma dilində eyni tipdə bir neçə dəyişəni tək bir dəyişkəndə saxlamağa imkan verən bir məlumat strukturudur. Bir array əvvəlcədən müəyyən edilmiş ölçüyə və müəyyən bir məlumat tipinə malikdir.

```go
var a [5]int
a[0] = 1
a[1] = 2
a[2] = 3
a[3] = 4
a[4] = 5
fmt.Println(a)
```

Bu nümunədə, a adlı bir array müəyyən edilir və ölçüsü 5 olaraq təyin edilir. Array-in elementlərinə a\[index] sintaksisi ilə daxil olmaq olar və elementlər tək-tək təyin edilə bilər. Bu nümunədə, array elementlərinə müvafiq olaraq 1, 2, 3, 4 və 5 dəyərləri təyin edilir və fmt.Println(a) ifadəsi istifadə edilərək bütün array-in elementləri ekrana çıxarılır.

Array-lərin ölçüsü bir dəfə təyin edildikdən sonra dəyişdirilə bilməz. Ancaq Go dilindəki dil xüsusiyyətləri ilə array ölçüləri dəyişdirilə bilər. Bu xüsusiyyət, Go dilində slice adlanan bir məlumat strukturu ilə həyata keçirilir.

```go
a := [5]int{1, 2, 3, 4, 5}
fmt.Println(a)
```

Bu nümunədə, a adlı bir array müəyyən edilir və elementləri {1, 2, 3, 4, 5} olaraq təyin edilir. Array-in ölçüsü elementlərin sayına bərabər olmalıdır. Bu nümunədə, a adlı array {1, 2, 3, 4, 5} dəyərləri ilə başlatılır və bütün elementləri ekrana çıxarılır.
