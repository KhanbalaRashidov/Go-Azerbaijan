# variables

## Variables

Dəyişənlər Go proqramlaşdırma dilində məlumatların saxlanması üçün istifadə olunan əsaslardan biridir. Dəyişənə qiymət təyin etməklə biz proqramda həmin dəyişənin saxladığı qiymətdən istifadə edə bilərik.

Dəyişənlərin elan edilməsi var açar sözü ilə həyata keçirilir. Müəyyən ediləcək dəyişənin adı və məlumat tipi müəyyən edilir. Heç bir ilkin dəyər verilmədikdə, Go dilində dəyişənlər standart dəyərə malikdirlər.

```golang
var name string
name = "John Doe"
```

Bu nümunədə ad adlı dəyişən müəyyən edilir və bu dəyişənin tipi sətir kimi müəyyən edilir. Sonra, ad dəyişəninə "John Doe" sətri təyin edilir. Dəyişənlərin dəyəri dəyişdirilə bilər və onlar müxtəlif məlumat tiplərində ola bilər.

```golang
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu misalda əvvəlcə yaş adlı dəyişən müəyyən edilir və bu dəyişənin növü int kimi müəyyən edilir. Sonra, dəyişən yaşa 32 nömrəsi verilir.

Daha sonra yaş dəyişəninə "thirty-two" sətri təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipi dəyəri təyin edilə bilməz.

Dəyişənlər proqramlarda müəyyən bir məqsədə xidmət etmək üçün istifadə olunur. Məsələn, istifadəçinin adı, yaşı və ya bir sıra ədədi dəyərlər saxlanıla bilər.

\
\


## Range

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

\
\
