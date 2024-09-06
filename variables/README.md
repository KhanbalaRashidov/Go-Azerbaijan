# Variables

Dəyişənlər Go proqramlaşdırma dilində məlumatların saxlanması üçün istifadə olunan əsaslardan biridir. Dəyişənə qiymət təyin etməklə biz proqramda həmin dəyişənin saxladığı qiymətdən istifadə edə bilərik.

Dəyişənlərin elan edilməsi var açar sözü ilə həyata keçirilir. Müəyyən ediləcək dəyişənin adı və məlumat tipi müəyyən edilir. Heç bir ilkin dəyər verilmədikdə, Go dilində dəyişənlər standart dəyərə malikdirlər.

```go
var name string
name = "John Doe"
```

Bu nümunədə ad adlı dəyişən müəyyən edilir və bu dəyişənin tipi sətir kimi müəyyən edilir. Sonra, ad dəyişəninə "John Doe" sətri təyin edilir. Dəyişənlərin dəyəri dəyişdirilə bilər və onlar müxtəlif məlumat tiplərində ola bilər.

```go
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu misalda əvvəlcə yaş adlı dəyişən müəyyən edilir və bu dəyişənin növü int kimi müəyyən edilir. Sonra, dəyişən yaşa 32 nömrəsi verilir.

Daha sonra yaş dəyişəninə "thirty-two" sətri təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipi dəyəri təyin edilə bilməz.

Dəyişənlər proqramlarda müəyyən bir məqsədə xidmət etmək üçün istifadə olunur. Məsələn, istifadəçinin adı, yaşı və ya bir sıra ədədi dəyərlər saxlanıla bilər.
