# Range

range ifadəsi, Go proqramlaşdırma dilində müəyyən bir array, dilim və ya map üzərində dövr yaratmaq üçün istifadə olunur. range ifadəsi, məlumat strukturundakı bütün elementləri tək-tək götürmək üçün istifadə olunur.

```golang
numbers := []int{1, 2, 3, 4, 5}
for i, num := range numbers {
    fmt.Println("index:", i, "number:", num)
}
```

Bu nümunədə, numbers adlı bir dilim təyin edilir və elementlərinə {1, 2, 3, 4, 5} dəyərləri verilir. range ifadəsi istifadə edilərək, numbers dilimindəki bütün elementlər tək-tək götürülür və for dövrü içində istifadə olunur. i dəyişəni dövrün sırasındakı indeks dəyərini, num dəyişəni isə indeksdəki elementi təmsil edir. Bu nümunədə, indekslər və elementlər birgə yazdırılır.

```golang
colors := map[string]string{
    "red":   "#FF0000",
    "green": "#00FF00",
    "blue":  "#0000FF",
}
for color, code := range colors {
    fmt.Println("color:", color, "code:", code)
}
```

Bu nümunədə, colors adlı bir map təyin edilir və red, green və blue açarlarına sahib rəng kodları verilir. range ifadəsi istifadə edilərək, colors xəritəsindəki bütün açar-dəyər cütləri tək-tək götürülür və for dövrü içində istifadə olunur. color dəyişəni açarı, code dəyişəni isə açara uyğun dəyəri təmsil edir. Bu nümunədə, rənglər və onların kodları birgə yazdırılır.


<br><br>

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

