# Functions

Funksiya, Go proqramlaşdırma dilində, müəyyən bir vəzifəni yerinə yetirən kod bloklarını ifadə edir. Bir funksiya, bir və ya bir neçə parametr ala bilər, bir əməliyyatı həyata keçirə bilər və bir və ya bir neçə nəticə qaytara bilər.

```golang
func add(a int, b int) int {
return a + b
}
​
result := add(5, 10)
fmt.Println(result)
```

Bu nümunədə, add adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki int tipində parametr qəbul edir və bu parametrlər toplanaraq nəticə int tipində geri qaytarılır. add funksiyası, 5 və 10 parametrləri ilə çağırılır və nəticə fmt.Println(result) ifadəsi ilə yazdırılır.

```golang
func swap(a, b string) (string, string) {
    return b, a
}
​
x, y := swap("hello", "world")
fmt.Println(x, y)
```

Bu nümunədə, swap adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki string tipində parametr qəbul edir və bu parametrləri bir-biri ilə dəyişdirərək geri qaytarır. swap funksiyası, "hello" və "world" parametrləri ilə çağırılır və geri qaytarılan nəticələr x və y dəyişənləri tərəfindən qəbul edilir. Nəticələr fmt.Println(x, y) ifadəsi ilə yazdırılır.

<br><br>

