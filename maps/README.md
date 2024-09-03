# Maps

Map, Go proqramlaşdırma dilində bir açar-dəyər cütləri kolleksiyasıdır. Map məlumat strukturu, digər proqramlaşdırma dillərindəki dictionary, hash table və ya associative array kimi məlumat strukturlarına bənzəyir. Bir Map məlumat strukturu, müəyyən bir açar üçün bir dəyər saxlayır.

```golang
var colors map[string]string
colors = make(map[string]string)
colors["red"] = "#FF0000"
colors["green"] = "#00FF00"
colors["blue"] = "#0000FF"
fmt.Println(colors)
```

Bu nümunədə, colors adlı bir Map təyin edilir və make() funksiyası ilə yaradılır. colors məlumat strukturu red, green və blue açarlarına sahib rəng kodları ilə doldurulur və fmt.Println(colors) ifadəsi istifadə edilərək Map məlumat strukturundakı bütün açar-dəyər cütləri ekrana çıxarılır.

Map məlumat strukturu, digər proqramlaşdırma dillərindəki məlumat strukturlarından fərqli olaraq, açarlar və dəyərlər üçün müəyyən bir məlumat tipini göstərmək məcburiyyətində deyil. Açarlar və dəyərlər fərqli məlumat tiplərində ola bilər.

```golang
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
fmt.Println(ages)
```

Bu nümunədə, ages adlı bir Map təyin edilir və string tipində açarlar və int tipində dəyərlərlə əlaqəli cütlər təyin olunur. Map məlumat strukturunun yaradılması və elementlərin əlavə edilməsi make() funksiyası ilə birləşdirilərək də həyata keçirilə bilər.


<br><br>