# Şərhlər (Comments)

Şərhlər (comments) Go proqramlaşdırma dilində kodun oxunmasını asanlaşdırmaq, izah etmək və ya müəyyən qeydlər əlavə etmək üçün istifadə olunur. Şərhlər kompilyator tərəfindən nəzərə alınmır, yəni onlar proqramın icrasına təsir etmir. Go dilində iki növ şərh mövcuddur: tək sətrli şərhlər və çox sətrli şərhlər.

## Tək sətrli şərhlər

Tək sətrli şərhlər `//` simvolları ilə başlayır və sətirin qalan hissəsi şərh olaraq qəbul edilir. Bu tip şərhlər adətən kodun müəyyən hissəsini qısaca izah etmək üçün istifadə olunur.

```go
package main

import "fmt"

func main() {
    // Bu tək sətrli şərhdir və "Hello, World!" mesajını çap edəcək.
    fmt.Println("Hello, World!")
}
```

Yuxarıdakı nümunədə `//` simvollarından sonra gələn hissə şərh olaraq qəbul edilir və proqramın işləmə prosesinə təsir etmir.

## Çox sətrli şərhlər

Çox sətrli şərhlər `/*` ilə başlayır və `*/` ilə bitir. Bu şərhlər birdən çox sətirdə yazıla bilər və daha ətraflı izahlar üçün istifadə olunur.

```go
package main

import "fmt"

func main() {
    /*
       Bu çox sətrli şərhdir.
       Aşağıdakı kod konsolda "Hello, World!" mesajını çap edəcək.
    */
    fmt.Println("Hello, World!")
}
```

Bu nümunədə, `/*` və `*/` arasında olan bütün mətn çox sətrli şərh kimi qəbul edilir.

## Şərhlərin istifadə məqsədləri

1. **Kodun izah edilməsi:** Kodun funksionallığını, mürəkkəb hissələrini izah etmək üçün istifadə olunur.
2. **Gələcəkdə işarələmə:** Şərhlər gələcəkdə kodda hansı dəyişikliklər və ya təkmilləşdirmələrin edilməsi lazım olduğunu qeyd etmək üçün istifadə edilə bilər.
3. **Debugging:** Debugging zamanı müəyyən bir kod hissəsinin müvəqqəti olaraq şərh edilməsi və proqramın o hissəsinin icra edilməməsi təmin edilə bilər.

### Yaxşı şərh yazma təcrübələri

- **Anlayışlı və qısa olun:** Şərhlər lazımsız detallara girmədən kodu izah etməlidir.
- **Yenilənmiş saxlayın:** Kodun funksionallığı dəyişdikdə, şərhləri də ona uyğun şəkildə yeniləmək vacibdir. Əks halda, şərhlər köhnəlmiş ola bilər və çaşqınlığa səbəb ola bilər.
- **Çoxlu şərhlərdən çəkinin:** Şərhlər faydalı olmalı, lakin kodu şərhlərlə doldurmaq kodun oxunmasını çətinləşdirə bilər.

### Nəticə

Go dilində şərhlər kodun təmiz və başa düşülən olmasına kömək edir. Hər iki şərh növü, yəni tək sətrli və çox sətrli şərhlər, proqramlaşdırma zamanı daha səliqəli və başadüşülən kod yazmağa imkan verir.