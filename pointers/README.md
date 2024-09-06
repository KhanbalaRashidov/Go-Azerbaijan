# Pointers

Pointers, Go proqramlaşdırma dilində, bir dəyişənin yaddaş ünvanını saxlayan bir məlumat tipidir. Yaddaş ünvanı, dəyişənin yaddaşda yerləşdiyi yerdir.

```golang
func zeroVal(val int) {
    val = 0
}

func zeroPtr(ptr *int) {
    *ptr = 0
}

x := 5
zeroVal(x)
fmt.Println(x)

y := 5
zeroPtr(&y)
fmt.Println(y)
```

Bu nümunədə, zeroVal adlı bir funksiya təyin edilir. Funksiya int tipində bir parametr qəbul edir və val dəyişəninin dəyərini 0 olaraq dəyişir, amma bu dəyişiklik orijinal dəyişənə təsir etmir. Digər tərəfdən, zeroPtr adlı bir funksiya isə bir pointer (\*int) qəbul edir və bu göstərici ilə göstərilən dəyişənin dəyərini 0 olaraq dəyişdirir. x dəyişəninə 5 dəyəri atanır və zeroVal(x) çağırılır, amma x dəyişməz. y dəyişəninə isə 5 dəyəri atanır və zeroPtr(\&y) çağırıldıqda, y 0 olur.

```golang
func swap(a, b *int) {
    temp := *a
    *a = *b
    *b = temp
}

x := 5
y := 10
swap(&x, &y)
fmt.Println(x, y)
```

Bu nümunədə, swap adlı funksiya iki pointer (\*int) qəbul edir və həmin dəyişənlərin dəyərlərini bir-biri ilə dəyişir. x-ə 5 və y-ə 10 dəyəri verilir. swap(\&x, \&y) çağırıldıqdan sonra, x 10 olur, y isə 5 olur.

\
\
