# Range

range ifadəsi, Go proqramlaşdırma dilində müəyyən bir array, dilim və ya map üzərində dövr yaratmaq üçün istifadə olunur. range ifadəsi, məlumat strukturundakı bütün elementləri tək-tək götürmək üçün istifadə olunur.

```go
numbers := []int{1, 2, 3, 4, 5}
for i, num := range numbers {
    fmt.Println("index:", i, "number:", num)
}
```

Bu nümunədə, numbers adlı bir dilim təyin edilir və elementlərinə {1, 2, 3, 4, 5} dəyərləri verilir. range ifadəsi istifadə edilərək, numbers dilimindəki bütün elementlər tək-tək götürülür və for dövrü içində istifadə olunur. i dəyişəni dövrün sırasındakı indeks dəyərini, num dəyişəni isə indeksdəki elementi təmsil edir. Bu nümunədə, indekslər və elementlər birgə yazdırılır.

```go
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
