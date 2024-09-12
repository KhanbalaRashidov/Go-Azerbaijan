# Info

Bu repository Go dilini tez öyrənmək istəyənlər üçün hazırlanmışdır. Ümumilikdə mövzuları nümunələrlə izah edərək dilin əsas strukturlarını əhatə edir. Eyni zamanda, Go dilini öyrənmək istəyənlər üçün müntəzəm resurs təmin etmək məqsədi daşıyır. Hər bir mövzu təsviri və başa düşülən şəkildə əhatə olunur ki, oxucular Go dilini asanlıqla və tez öyrənə bilsinlər.

Repository-ni bəyənirsinizsə, ulduz qoyub sosial media hesablarınızda paylaşa bilərsiniz ki, daha çox insana çatsın ⭐️.

## **Bu Go nədir?**

Golang (digər adı Go) 2007-ci ildən Google tərəfindən hazırlanmış açıq mənbəli proqramlaşdırma dilidir. O, əsasən alt sistem proqramlaşdırması üçün nəzərdə tutulmuşdur və tərtib edilə bilən və statik olaraq yazılmış dildir. İlk versiya 2009-cu ilin noyabrında buraxıldı. Onun tərtibçisi "gc" (Go Compiler) bir çox əməliyyat sistemi üçün açıq mənbə kimi işlənib hazırlanmışdır.

Golang ilk dəfə Google mühəndisləri Robert Griesemer, Rob Pike və Ken Thompson tərəfindən eksperimental məqsədlər üçün yaradılmışdır. Tənqid olunan problemlərin digər dillərdə həllini təmin etmək və onların yaxşı tərəflərini qorumaq üçün yaradılmışdır.

Go sadə, sürətli və etibarlı proqram təminatı hazırlamaq üçün nəzərdə tutulmuş açıq mənbəli proqramlaşdırma dilidir. Go dili sürətli tərtib prosesi, yüngül sintaksis strukturu və effektiv garbage collection sistemi ilə seçilir.

Go dilinin istifadəsinə müxtəlif inkişaf alətləri, paketlər və modullar daxildir. Tərtibatçılar Go dili ilə veb proqramlar, API-lər, verilənlər bazası sistemləri, şəbəkə proqram təminatı və sair kimi çoxlu müxtəlif növ proqramlar yarada bilərlər.

Go dilinin əsas xüsusiyyətlərindən biri sürətli tərtib prosesidir. Go C dilinə bənzər sintaksis strukturundan və C++ kimi dillərdə obyekt yönümlü xüsusiyyətlərdən istifadə edir. Zibil toplama, yüngül iplər, kapsullaşdırılmış tip sistemlər və dinamik yaddaşın idarə edilməsi kimi xüsusiyyətlər də Go dilində mövcuddur.

Go dilinin digər mühüm xüsusiyyəti onun effektiv paket idarəetmə sistemidir. Go dilində modullar və paketlər proqram tərtibatçılarına kodlarını nizamlı şəkildə təşkil etməyə və təkmilləşdirməyə imkan verir. Beləliklə, tərtibatçılar asanlıqla təkrar istifadə edilə bilən kodları yaza bilərlər.

Nəticədə, Go dilinin əsas xüsusiyyətlərinə sürətli tərtib prosesi, yüngül sintaksis strukturu, effektiv paket idarəetməsi, zibil toplama və kapsullaşdırılmış tip sistemlər kimi xüsusiyyətlər daxildir. Go-da inkişaf etdirmək bir çox proqram növləri üçün əlverişli seçimdir və proqram təminatının hazırlanması prosesini sürətləndirməyə kömək edə bilər.

**Məzmun:**

* [Values](./#values)
* [Variables](./#variables)
* [Constants](./#constants)
* [If-Else](./#if-else)
* [For](./#for)
* [Switch](./#switch)
* [Arrays](./#arrays)
* [Slices](./#slices)
* [Maps](./#maps)
* [Range](./#range)
* [Functions](./#functions)
* [Variadic Functions](./#variadic-functions)
* [Closures](./#closures)
* [Recursion](./#recursion)
* [Pointers](./#pointers)
* [String and Runes](./#strings-and-runes)
* [Structs](./#structs)
* [Methods](./#methods)
* [Interfaces](./#interfaces)
* [Struct Embedding](./#struct-embedding)
* [Errors](./#errors)
* [Goroutines](./#goroutines)
* [Channel](./#channel)
* [Select](./#select)
* [Timeouts](./#timeouts)
* [Non-Blocking Channel](./#non-blocking-channel)
* [Closing Channels](./#closing-channels)
* [Range over Channels](./#range-over-channels)
* [Timers](./#timers)
* [Tickers](./#ticker)
* [Worker Pools](./#worker-pools)
* [Wait Groups](./#waitgroups)
* [Rate Limiting](./#rate-limiting)
* [Atomic Counters](./#atomic-counters)
* [Sorting](README%20\(1\).md#sorting)
* [Panic](./#panic)
* [Defer](./md#defer)
* [Recover](README%20\(1\).md#recover)
* [String Functions](./#string-functions)
* [Text Templates](README%20\(1\).md#text-templates)
* [JSON](README%20\(1\).md#json)

***

### Values

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

***

### Variables

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

### Constants

Sabitlər (constants) Go proqramlaşdırma dilində proqramın heç bir yerində dəyişdirilə bilməyən sabit dəyərlərdir. Sabitlər bir dəfə müəyyən edilir və sonra proqramın istənilən yerində istifadə edilə bilər.

Sabitlər const açar sözü ilə müəyyən edilir və onların məlumat növləri göstərilir. Sabitlər işə salınmalıdır və sonra dəyişdirilə bilməz.

```go
const pi = 3.14159
const welcomeMessage = "Welcome to Go programming"
```

Bu nümunədə, sabit pi float64 növü kimi müəyyən edilir və 3.14159-a mənimsədilir. Sonra, welcomeMessage sabiti sətir kimi müəyyən edilir və "Welcome to Go programming" olaraq işə salınır.

Sabitlər müəyyən bir məqsəd üçün istifadə ediləcək proqramların sabit dəyərlərini saxlamaq üçün istifadə olunur. Məsələn, pi sayı və ya xüsusi mesaj və ya səhv kodu. Sabitlər proqramları daha oxunaqlı və davamlı etməyə kömək edir və proqramın müxtəlif yerlərində eyni sabitdən dəfələrlə istifadə etməyi asanlaşdırır.

Sabitlər proqramın istənilən yerində istifadə oluna bildiyi üçün onlar müxtəlif fayllar arasında da paylaşıla bilər. Bundan əlavə, sabitlərin dəyişdirilə bilməməsi proqram səhvlərini azaldır və təhlükəsizliyi artırır.

### if/else

if və else ifadələri, Go programlaşdırma dilində, müəyyən şərtlərin doğru və ya yanlış olduğu hallarda fərqli kod bloklarının işləməsini həyata keçirir.

```go
if x > 0 {
    fmt.Println("Positive number")
} else if x < 0 {
    fmt.Println("Negative number")
} else {
    fmt.Println("Zero")
}
```

Bu nümunədə, if ifadəsi, x dəyişkənin 0 dan böyük olması halında "Positive number" mətnini yazdıracaq. Əgər x dəyişəni 0 dan böyük deyilsə, else if ifadəsi yoxlanılacaq ve x dəyişəninin 0 dan kiçik olması halında "Negative number" mətnini yazdıracaq. Əgər x dəyişəni 0 dan böyük və ya kiçik deyilsə, else bloku işləyəcək ve "Zero" mətnini yazdıracaqdır.

```go
if və else ifadələri, qarışıq şərtləri idarə etmək üçün də istifadə oluna bilər.
if x > 10 && x < 20 {
    fmt.Println("x is between 10 and 20")
} else if x > 20 && x < 30 {
    fmt.Println("x is between 20 and 30")
} else {
    fmt.Println("x is not between 10 and 30")
}
```

Bu nümunədə, if ifadəsi, x dəyişənin 10 ve 20 arasında olması halında "x is between 10 and 20" mətnini yazdıracaqdır. Əgər x dəyişəni 10 ve 20 arasında deyilsə, else if ifadesi yoxlanılacaq və x dəyişənin 20 və 30 arasında olması halında "x is between 20 and 30" mətnini yazdıracaqdır. Əgər x dəyişəni 10 və 30 arasında deyilsə, else bloku işələyəcəkdir və "x is not between 10 and 30" mətnini yazdıracaqdır.

### For

for dövür operatırı, Go programlasdırma dilində, müəyyən bir şərt doğru olduğu halda təkrarlanan kod bloklarını ifadə etmək üçün istifadə olunur. for dövrü, bir başlanğıc halı, bir şərt və bir addım dəyərinə sahibdir.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Bu nümunədə, for dövrü, i dəyişəninin 0 dəyərindən başlayaraq, i dəyişəni < 5 şərtini ödədiyi müddətcə təkrarlanan bir dövrdür. Dövrdəki hər adımdda i dəyişəni 1 artırılır və nəticədə i dəyəri 4 olur.

for dövrü, şərt hissəsi doğru olana qədər təkrarlanır. Şərt hissəsi doğru olmadığında, dövr sonlanır. Ayrıca, for dövrü break və ya continue ifadələri ilə də idarə edilə bilir.

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

Bu nümunədə, for dövrü, i dəyişəninin 0 dəyərindən başlayaraq, i dəyişəni < 10 şərtini ödəyənədək təkrarlanan bir dövrdür. Dövrdəki hər adımdda, i dəyişəni 1 artırılır.

Dövr, if şərti ilə kəsilir ve break ifadesi işə düşür. Ayrıca, i dəyəri 3 olduğunda da dövr sonlanır.

Dövr içindəki digər if şərti isə, əgər i dəyəri cüt ədəddirsə bir sonrakı addıma keçərək i dəyərini yazdırmaz. Bu şərtin səbəbi, sadəcə tək ədədləri yazdırmaq istəməmizdir.

Sonsuz dövr yaratmaq üçün aşağıdakı nümunəyə baxaq.

```go
i:=0
for  {
    if i == 100 {
        break
    }
    fmt.Println(i)
	
	i++
}
```

Bu nümunədə i dəyişənimizə 0 dəyərini mənimsədirik. Daha sonra dövr daxilində i əgər 100-ə bərabər olarsa break ifadəsi ilə dövrü sonlandırırıq, əks halda dövr davam edəcəkdir

### Switch

Switch ifadəsi müəyyən şərtlər əsasında müxtəlif əməliyyatları yerinə yetirmək üçün Go proqramlaşdırma dilində istifadə olunur:

```go
day := "sunday"

switch day {
case "monday":
    fmt.Println("Today is Monday")
case "tuesday":
    fmt.Println("Today is Tuesday")
case "wednesday":
    fmt.Println("Today is Wednesday")
case "thursday":
    fmt.Println("Today is Thursday")
case "friday":
    fmt.Println("Today is Friday")
case "saturday":
    fmt.Println("Today is Saturday")
case "sunday":
    fmt.Println("Today is Sunday")
default:
    fmt.Println("Invalid day")
}
```

Bu misalda switch ifadəsi gün dəyişəninin dəyərindən asılı olaraq müxtəlif əməliyyatlar yerinə yetirir. Gün dəyişəni "monday"dirsə, "Today is Monday" mətnini çap edir. Gün dəyişəni "tuesday"dırsa, "Today is Tuesday" mətnini çap edir. Bu şəkildə, gün dəyişəninin dəyərindən asılı olaraq müxtəlif kod blokları icra edilə bilər.

default ifadəsi bütün şərtlər doğru olmadığı təqdirdə işləyəcək kod blokuna aiddir.

```go
switch x {
case 1:
    fmt.Println("x is 1")
case 2:
    fmt.Println("x is 2")
case 3:
    fmt.Println("x is 3")
default:
    fmt.Println("x is not 1, 2 or 3")
}
```

Bu misalda switch ifadəsi x dəyişəninin qiymətindən asılı olaraq müxtəlif əməliyyatlar yerinə yetirir. Əgər x dəyişəni 1-dirsə, o, "x-1-dir" mətnini çap edir. Əgər x dəyişəni 2-dirsə, o, "x-2-dir" mətnini çap edir. Əgər x dəyişəni 3-dürsə, o, “x-3-dür” mətnini çap edir. Əgər x dəyişəni 1, 2 və ya 3 deyilsə, standart blok işləyəcək və "x is not 1, 2 or 3" mətnini çap edəcək.

### Arrays

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

### Slices

Dilimin əsas məlumat strukturlarından biri olan slices, Go proqramlaşdırma dilində ölçüsü dəyişdirilə bilən, elastik bir array məlumat strukturdur. Slices, array-lərin bir alt hissəsi kimi qəbul edilə bilər və Go dilində tez-tez istifadə olunur.

```go
a := []int{1, 2, 3, 4, 5}
fmt.Println(a)
```

Bu nümunədə, a adlı bir dilim təyin edilir və elementləri {1, 2, 3, 4, 5} olaraq təyin edilir. Array-lərdən fərqli olaraq, dilimlərin ölçüsü təyin edilmə mərhələsində göstərilmir. Slices boş da təyin edilə bilər.

```go
var a []int
a = append(a, 1)
a = append(a, 2, 3, 4, 5)
fmt.Println(a)
```

Bu nümunədə, a adlı bir dilim təyin edilir və append() funksiyası ilə elementlərinə müvafiq olaraq 1, 2, 3, 4 və 5 dəyərləri əlavə edilir. append() funksiyası, bir dilimə bir və ya bir neçə element əlavə etmək üçün istifadə olunur.

Dilimlər, array-lərdən fərqli olaraq, bir alt hissə kimi də təyin edilə bilər.

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:4]
fmt.Println(b)
```

Bu nümunədə, a adlı bir dilim təyin edilir və elementlərinə {1, 2, 3, 4, 5} dəyərləri verilir. b adlı bir dilim isə a diliminin 1-ci indeksindən 4-cü indeksinə qədər olan alt hissəsini əhatə edir. b dilimindəki elementlər {2, 3, 4} olaraq göstərilir.

### Maps

Map, Go proqramlaşdırma dilində bir açar-dəyər cütləri kolleksiyasıdır. Map məlumat strukturu, digər proqramlaşdırma dillərindəki dictionary, hash table və ya associative array kimi məlumat strukturlarına bənzəyir. Bir Map məlumat strukturu, müəyyən bir açar üçün bir dəyər saxlayır.

```go
var colors map[string]string
colors = make(map[string]string)
colors["red"] = "#FF0000"
colors["green"] = "#00FF00"
colors["blue"] = "#0000FF"
fmt.Println(colors)
```

Bu nümunədə, colors adlı bir Map təyin edilir və make() funksiyası ilə yaradılır. colors məlumat strukturu red, green və blue açarlarına sahib rəng kodları ilə doldurulur və fmt.Println(colors) ifadəsi istifadə edilərək Map məlumat strukturundakı bütün açar-dəyər cütləri ekrana çıxarılır.

Map məlumat strukturu, digər proqramlaşdırma dillərindəki məlumat strukturlarından fərqli olaraq, açarlar və dəyərlər üçün müəyyən bir məlumat tipini göstərmək məcburiyyətində deyil. Açarlar və dəyərlər fərqli məlumat tiplərində ola bilər.

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 35,
}
fmt.Println(ages)
```

Bu nümunədə, ages adlı bir Map təyin edilir və string tipində açarlar və int tipində dəyərlərlə əlaqəli cütlər təyin olunur. Map məlumat strukturunun yaradılması və elementlərin əlavə edilməsi make() funksiyası ilə birləşdirilərək də həyata keçirilə bilər.

### Range

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

### Functions

Funksiya, Go proqramlaşdırma dilində, müəyyən bir vəzifəni yerinə yetirən kod bloklarını ifadə edir. Bir funksiya, bir və ya bir neçə parametr ala bilər, bir əməliyyatı həyata keçirə bilər və bir və ya bir neçə nəticə qaytara bilər.

```go
func add(a int, b int) int {
return a + b
}

result := add(5, 10)
fmt.Println(result)
```

Bu nümunədə, add adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki int tipində parametr qəbul edir və bu parametrlər toplanaraq nəticə int tipində geri qaytarılır. add funksiyası, 5 və 10 parametrləri ilə çağırılır və nəticə fmt.Println(result) ifadəsi ilə yazdırılır.

```go
func swap(a, b string) (string, string) {
    return b, a
}

x, y := swap("hello", "world")
fmt.Println(x, y)
```

Bu nümunədə, swap adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki string tipində parametr qəbul edir və bu parametrləri bir-biri ilə dəyişdirərək geri qaytarır. swap funksiyası, "hello" və "world" parametrləri ilə çağırılır və geri qaytarılan nəticələr x və y dəyişənləri tərəfindən qəbul edilir. Nəticələr fmt.Println(x, y) ifadəsi ilə yazdırılır.

### Variadic Functions

Variadic funksiyalar, Go proqramlaşdırma dilində dəyişən sayda arqument qəbul edən funksiyalardır. Bu funksiyalar bir və ya daha çox arqument qəbul edə bilər və arqumentlərin sayını dəyişən olaraq təyin etməyə imkan verir.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

fmt.Println(sum(1, 2, 3, 4, 5))
fmt.Println(sum(2, 4, 6))
```

Bu nümunədə, sum adlı bir funksiya təyin edilir. Funksiya, nums adlı int tipində dəyişən sayda parametr qəbul edir və onların cəmini qaytarır. Funksiyanın daxilində for dövrü istifadə edərək, nums dilimindəki bütün elementlərin cəmi hesablanır. Funksiya sum(1, 2, 3, 4, 5) və sum(2, 4, 6) şəklində çağırılır və nəticələr ekrana yazdırılır.

```go
func concatenate(sep string, strs ...string) string {
    result := ""
    for i, str := range strs {
        if i > 0 {
            result += sep
        }
        result += str
    }
    return result
}

fmt.Println(concatenate(", ", "foo", "bar", "baz"))
fmt.Println(concatenate("-", "hello", "world"))
```

Bu nümunədə, concatenate adlı bir funksiya təyin edilir. Funksiya, sep adlı string tipində bir parametr və strs adlı dəyişən sayda string tipində parametrlər qəbul edir. Funksiyanın daxilində for dövrü istifadə edilərək, strs dilimindəki bütün elementlər birləşdirilir və sep ayracı ilə birgə yazılır. Funksiya concatenate(", ", "foo", "bar", "baz") və concatenate("-", "hello", "world") şəklində çağırılır və nəticələr ekrana yazdırılır.

### Closures

Closures, Go proqramlaşdırma dilində, bir funksiyanın başqa bir funksiyanın daxilində yaradılması ilə əmələ gələn bir strukturdur. Bu struktur, bir funksiyanın daxilində olan başqa bir funksiyaya istinad edərək, funksiyanın işlədiyi mühitin xaricindəki dəyişənlərə giriş imkanı verir.

```go
func outer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

increment := outer()
fmt.Println(increment())
fmt.Println(increment())
fmt.Println(increment())
```

Bu örnekte, outer adlı bir fonksiyon tanımlanır. Fonksiyon, bir iç fonksiyon döndürür ve iç fonksiyon, count adlı bir değişkene erişim sağlar. increment adlı bir değişkene outer() fonksiyonu atanır ve bu değişken ile iç fonksiyon çalıştırılır. count değişkeni, increment() çağrıldıkça artar ve her seferinde artışı ekrana yazdırılır.

```go
func adder(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}

addFive := adder(5)
fmt.Println(addFive(2))
fmt.Println(addFive(3))
```

Bu nümunədə, adder adlı bir funksiya təyin edilir. Funksiya, a adlı bir int tipində parametr qəbul edir və bir daxili funksiya qaytarır. Daxili funksiya, b adlı bir int tipində parametr qəbul edir və a ilə b parametrlərinin cəmini geri qaytarır. addFive adlı bir dəyişənə adder(5) funksiyası təyin edilir və bu dəyişən ilə daxili funksiya işlədilir. addFive(2) və addFive(3) çağırıldıqca nəticələr ekrana yazdırılır.

### Recursion

Recursion, Go proqramlaşdırma dilində, bir funksiyanın özünü çağırmasıdır. Bu struktur, müəyyən bir şərt yerinə yetirilənə qədər funksiyanın təkrarlanaraq işləməsini təmin edir.

```go
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5))
```

Bu nümunədə, factorial adlı bir funksiya təyin edilir. Funksiya, n adlı bir int tipində parametr qəbul edir və faktorialı hesablayır. Funksiya daxilində, if şərti istifadə edilərək n dəyərinin 0 olub-olmadığı yoxlanılır. Əgər n 0-dırsa, 1 qaytarılır. Əgər n 0 deyil, başqa bir dəyərdirsə, funksiya özünü yenidən çağıraraq faktorialı hesablayır. factorial(5) çağırıldıqda nəticə ekrana yazdırılır.

```go
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

fmt.Println(fibonacci(10))
```

Bu nümunədə, fibonacci adlı bir funksiya təyin edilir. Funksiya, n adlı bir int tipində parametr qəbul edir və Fibonacci sayını hesablayır. Funksiya daxilində, if şərti istifadə edilərək n dəyərinin 2-dən kiçik olub-olmadığı yoxlanılır. Əgər n 2-dən kiçikdirsə, n dəyəri geri qaytarılır. Əks halda, funksiya özünü yenidən çağıraraq Fibonacci sayını hesablayır. fibonacci(10) çağırıldıqda nəticələr ekrana yazdırılır.

### Pointers

Pointers, Go proqramlaşdırma dilində, bir dəyişənin yaddaş ünvanını saxlayan bir məlumat tipidir. Yaddaş ünvanı, dəyişənin yaddaşda yerləşdiyi yerdir.

```go
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

```go
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

### Strings and Runes

Strings, Go proqramlaşdırma dilində, Unicode simvollarının birləşdirilməsi ilə yaradılan bir simvol sətiridir. Hər bir simvol 1-4 byte arasında dəyişən ölçülərdə ola bilər.

```go
str := "hello"
fmt.Println(str)
fmt.Println(str[0])
fmt.Println(str[1:3])
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "hello" dəyəri verilir. str dəyişəni ekrana çap olunur, sonra isə str\[0] ifadəsi istifadə edilərək sətirin ilk simvolu olan "h" ekrana çap edilir. str\[1:3] ifadəsi istifadə edilərək sətirin ikinci və üçüncü simvolları olan "el" ekrana çap edilir.

```go
for i, r := range "hello" {
    fmt.Printf("%d: %s\n", i, string(r))
}
```

Bu nümunədə, for döngüsü istifadə edilərək range funksiyası ilə "hello" stringindəki hər bir simvola giriş edilir. Hər simvolun mövqeyi və dəyəri ekrana çap olunur. Runes, Go proqramlaşdırma dilində, bir Unicode simvolunun birləşməsini ifadə edən bir məlumat tipidir. Runes, 1-4 bayt arasında dəyişən ölçülərdə ola bilən simvolları təmsil etmək üçün istifadə olunur.

```go
str := "こんにちは"
for i, r := range str {
    fmt.Printf("%d: %c\n", i, r)
}
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "こんにちは" dəyəri verilir. for döngüsü istifadə edilərək range funksiyası ilə hər bir simvola giriş edilir və simvolun mövqeyi və dəyəri ekrana çap olunur.

### Structs

Structlar, Go proqramlaşdırma dilində, fərqli məlumat tiplərini özündə birləşdirən bir məlumat quruluşudur. Bu quruluşda, müxtəlif məlumat tiplərinə sahib məlumatları bir arada saxlaya və bu məlumatlar üzərində əməliyyatlar apara bilərsiniz.

```go
type Person struct {
    Name string
    Age  int
}

var p Person
p.Name = "John Doe"
p.Age = 42
fmt.Println(p)
fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
```

Bu nümunədə, Person adlı bir struct yaradılır və Name adlı bir string tipində və Age adlı bir int tipində iki xüsusiyyət təyin edilir. var açar sözü ilə p adlı bir Person tipində dəyişən təyin edilir. p dəyişəninin Name və Age xüsusiyyətləri, p.Name və p.Age ifadələri istifadə edilərək təyin edilir və fmt.Println() funksiyası ilə p dəyişəni ekrana çap edilir. Əlavə olaraq, p.Name və p.Age ifadələri ilə, Name və Age xüsusiyyətləri ayrıca ekrana çap edilir.

```go
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

r := Rectangle{width: 3.0, height: 4.0}
fmt.Println(r.area())
```

Bu nümunədə, Rectangle adlı bir struct yaradılır və width adlı bir float64 tipində və height adlı bir float64 tipində iki xüsusiyyət təyin edilir. area adlı bir funksiya, Rectangle tipində bir parametr qəbul edir və düzbucaqlının sahəsini hesablayır. Funksiyanın qaytardığı nəticə düzbucaqlının sahəsidir. r adlı bir Rectangle dəyişəni yaradılır və width və height xüsusiyyətlərinə 3.0 və 4.0 dəyərləri verilir. r.area() ifadəsi istifadə edilərək, düzbucaqlının sahəsi hesablanır və nəticə ekrana çap edilir.

### Methods

Metodlar, Go proqramlaşdırma dilində, bir məlumat quruluşuna məxsus əməliyyatları yerinə yetirmək üçün istifadə edilən funksiya növüdür. Bu əməliyyatlar, məlumat quruluşunun xüsusiyyətləri üzərində işləyərək nəticə yaradır.

```go
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
    return 2 * (r.width + r.height)
}

r := Rectangle{width: 3.0, height: 4.0}
fmt.Println(r.area())
fmt.Println(r.perimeter())
```

Bu nümunədə, Rectangle adlı bir struct yaradılır və width adlı bir float64 tipində və height adlı bir float64 tipində iki xüsusiyyət təyin olunur. area adlı funksiya, Rectangle tipində parametr qəbul edir və düzbucaqlının sahəsini hesablayır. Funksiyanın qaytardığı nəticə, düzbucaqlının sahəsidir. perimeter adlı funksiya da Rectangle tipində parametr alır və düzbucaqlının perimetrini hesablayır. Funksiyanın qaytardığı nəticə, düzbucaqlının perimetridir. r adlı bir Rectangle dəyişəni yaradılır və width və height xüsusiyyətlərinə 3.0 və 4.0 dəyərləri təyin olunur. r.area() və r.perimeter() ifadələrindən istifadə edərək düzbucaqlının sahəsi və perimetri hesablanır və nəticələr ekrana yazdırılır.

```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) setName(name string) {
    p.Name = name
}

func (p *Person) setAge(age int) {
    p.Age = age
}

func (p Person) getName() string {
    return p.Name
}

func (p Person) getAge() int {
    return p.Age
}

p := Person{Name: "John Doe", Age: 42}
p.setName("Jane Doe")
p.setAge(35)
fmt.Printf("Name: %s, Age: %d\n", p.getName(), p.getAge())
```

Bu nümunədə, Person adlı bir struct yaradılır və Name adlı bir string tipində və Age adlı bir int tipində iki xüsusiyyət təyin olunur. setName adlı funksiya, Person tipində pointer (\*Person) parametr qəbul edir və şəxsin adını dəyişir. setAge adlı funksiya, Person tipində pointer (\*Person) parametr qəbul edir və şəxsin yaşını dəyişir. getName adlı funksiya, Person tipində parametr qəbul edir və şəxsin adını qaytarır. getAge adlı funksiya isə şəxsin yaşını qaytarır. p adlı bir Person dəyişəni yaradılır və Name və Age xüsusiyyətlərinə "John Doe" və 42 dəyərləri təyin olunur. Daha sonra, p.setName("Jane Doe") və p.setAge(35) ifadələrindən istifadə edərək şəxsin adı və yaşı dəyişdirilir.

### Interfaces

Go dilində interfeys, bir və ya bir neçə metodun müəyyən imza dəstini təyin edən bir məlumat tipidir. Bu imza dəsti, bir məlumat tipinin hansı metodları tətbiq etməli olduğunu göstərir. Bu səbəbdən, interfeyslər, bir məlumat tipinin hansı xüsusiyyətlərə malik olduğunu təyin etmək üçün istifadə olunur.

Məsələn, bir fiqurun sahəsini hesablamaq üçün area adlı bir funksiya təyin etdiyinizi düşünün. Kvadrat, dairə və düzbucaqlı kimi müxtəlif fiqurlar area funksiyasını fərqli şəkildə tətbiq edirlər. Bu halda, hər bir fiqurun sahəsini hesablamaq üçün ayrı-ayrı funksiyalar təyin etmək əvəzinə, interfeys istifadə edərək hamısını eyni tipdə bir məlumatda toplaya bilərsiniz.

```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

type Square struct {
	side float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (s Square) area() float64 {
	return s.side * s.side
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 5, height: 10}
	square := Square{side: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
	fmt.Printf("Square area: %f\n", getArea(square))
}
```

Bu nümunədə, Shape adlı bir interfeys təyin olunur və area adlı bir funksiya imzası göstərilir. Circle, Rectangle və Square adlı üç fərqli struktura malik məlumat tipləri təyin olunur və hər biri area adlı funksiyanı həyata keçirir. getArea adlı bir funksiya təyin olunur və parametri Shape tipindədir. Bu funksiya, sahəsi hesablanacaq fiqur məlumatını alır və area funksiyasını çağıraraq fiqurun sahəsini hesablamağa imkan verir. main funksiyasında isə, nümunə olaraq circle, rectangle və square adlı üç fərqli dəyişən yaradılır. Bu dəyişənlərin hər biri getArea funksiyasına parametr olaraq verilir və hər bir fiqurun sahəsi hesablanaraq ekrana yazdırılır.

Output:

```go
Circle area: 78.539816
Rectangle area: 50.000000
Square area: 25.000000
```

### Struct Embedding

Go dilində struct embedding, bir struct-ın başqa bir struct daxilində yerləşdirilməsi ilə digər struct-ın sahələrinə və metodlarına birbaşa giriş imkanı verir. Bu, təkcə kod təkrarını azaltmır, həm də kompozisiyanı təşviq edir, beləliklə, bir struct-ın başqa bir struct-a miras buraxmadan onun funksionallığını təmin etməyə imkan verir.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old\n", p.Name, p.Age)
}

type Employee struct {
	Person
	Company string
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Company: "ABC Inc.",
	}

	emp.SayHello()
}
```

Bu nümunədə, Person adlı bir struct təyin edilir və Name və Age adlı iki sahəsi var. Eyni zamanda SayHello adlı bir funksiyası mövcuddur.

Employee adlı bir struct təyin edilir və onun içərisində Person struct-ı yerləşdirilir. Employee-nin əlavə olaraq Company adlı bir sahəsi də var. main funksiyasında, emp adlı bir Employee dəyişəni yaradılır. Bu dəyişənin Person sahəsinə Person tipində bir dəyər atanır. SayHello funksiyası emp dəyişəni üzərindən çağırılır və nəticə ekrana yazdırılır.

Output:

```go
Hello, my name is John and I'm 30 years old
```

Bu nümunədə, Employee adlı struktur, Person strukturunu daxili olaraq yerləşdirir. Bu, Employee strukturuna Person-un bütün sahələrinə və funksiyalarına birbaşa çıxış imkanı verir. Beləliklə, Employee strukturunu istifadə edərək həm Employee, həm də Person məlumatlarına asanlıqla giriş əldə etmək mümkündür. Bu xüsusiyyət, kod təkrarını azaldır və strukturları daha modul halına gətirir.

### Errors

Go dilində səhv (error) idarəetməsi, error adlı bir məlumat tipi vasitəsilə həyata keçirilir. Bu tip, bir funksiyanın nəticəsi olaraq ya bir səhv mesajı, ya da nil dəyərini qaytarmaq üçün istifadə olunur. Bu, proqramın gözlənilməyən vəziyyətlərdə necə davranmalı olduğunu idarə etməyə kömək edir.

```go
package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
```

Bu nümunədə, divide adlı bir funksiya təyin edilir və iki ədədi bölür. Lakin, əgər ikinci ədəd 0 olarsa, bir səhv qaytarır.

main funksiyasında, divide funksiyası iki fərqli parametr ilə çağırılır. err adlı bir səhv dəyişəni istifadə edilərək hər bir çağırışın nəticəsi yoxlanılır. Əgər səhv varsa, səhv mesajı ekrana yazdırılır. Əks halda, nəticə ekrana yazdırılır.

Output:

```go
5 <nil>
Cannot divide by zero
```

Bu nümunədə divide funksiyası sıfıra bölmə vəziyyətini idarə etmək üçün səhv idarəetməsindən istifadə edir. Əgər y sıfır olarsa, errors.New ilə yaradılmış səhv qaytarılır. Proqram err dəyişəni ilə səhv olub-olmadığını yoxlayır və əgər səhv varsa, onu ekrana yazdırır. Bu yanaşma Go dilində səhv idarəetməsini daha etibarlı və idarə olunan edir.

### Goroutines

Go dilində goroutine-lər, eyni anda çalışan əməliyyatlardır. Goroutine-lər go açar sözü istifadə edilərək yaradılır və fərqli əməliyyatları eyni vaxtda həyata keçirmək üçün istifadə olunur.

```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
}

func main() {
	go sayHello() // goroutine
	time.Sleep(time.Second)
	fmt.Println("World")
}
```

Bu nümunədə, sayHello adlı bir funksiya təyin edilir və "Hello" mesajını ekrana yazdırır.

main funksiyasında, sayHello funksiyası bir goroutine olaraq çağırılır. Bu səbəbdən, sayHello funksiyasının icrası digər əməliyyatlardan müstəqil olaraq baş verir. time.Sleep funksiyası bir saniyəlik gözləmə müddəti əlavə edir. Nəticədə, "World" mesajı ekrana yazdırılır.

Output:

```go
Hello
World
```

Bu nümunədə, goroutine istifadə edərək sayHello funksiyası eyni anda çalışdırıldı. sayHello funksiyası goroutine olaraq çağırıldığı üçün digər əməliyyatlardan müstəqil işləd və nəticədə ekrana "Hello" mesajı yazdırıldıktan sonra "World" mesajı yazdırıldı.

### Channel

Go dilində kanal (channel), goroutine-lər arasında məlumat ötürmək üçün istifadə olunan bir məlumat strukturudur. Kanal, make açar sözü ilə yaradılır və <- operatoru ilə məlumat göndərmə və qəbul etmə əməliyyatları həyata keçirilir.

```go
package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "Hello" }()

	msg := <-messages
	fmt.Println(msg)
}
```

Bu nümunədə, messages adlı bir kanal yaradılır.

go açar sözü ilə bir goroutine yaradılır və bu goroutine messages kanalına "Hello" mesajını göndərir.

main funksiyasında, msg adlı bir dəyişkənə messages kanalından bir mesaj alınır və ekrana yazdırılır.

Output:

```go
Hello
```

Bu örnekte, channel kullanarak, bir goroutine'dan ana iş parçacığına bir mesaj gönderildi. Bu nedenle, goroutine işlemi tamamlandıktan sonra ana iş parçacığı channel'dan mesajı alır ve sonucu ekrana yazdırır.

Kanallar, Golang'de birçok durumda kullanılabilir, örneğin:

Bu nümunədə, kanal istifadə edərək bir goroutine-dən ana iş parçacığına bir mesaj göndərildi. Bu səbəbdən, goroutine tamamlandıqdan sonra ana iş parçacığı kanaldan mesajı alır və nəticəni ekrana yazdırır.

Kanallar Go dilində bir çox vəziyyətdə istifadə oluna bilər, məsələn:

1. Goroutine-lər arasında məlumat mübadiləsi üçün
2. Sinxronizasiya əməliyyatları üçün
3. Tətbiqinizin performansını artırmaq üçün (paralel emal etmək)
4. Goroutine-lər arasında məlumat yarışlarını qarşısını almaq üçün
5. Tapşırıqların koordinasiyası və sinxronizasiyası üçün

### Select

Go dilində, select açar sözü bir neçə kanalı `(channel)` dinləyərək hansı kanalın mesaj göndərdiyini müəyyənləşdirə bilər. Bu xüsusiyyət kanalların sinxronizasiyasını asanlaşdırır və fərqli goroutin-lər arasında mesajlaşmanı idarə edir.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```

Bu nümunədə, `c1` və `c2` adlı iki kanal yaradılır və mesaj göndərmə prosesi üçün goroutin-lər yaradılır.

main funksiyasında, select açar sözü istifadə edilərək, c1 və c2 kanalları dinlənilir. İlk olaraq, goroutin-lər arasındakı gözləmə müddətinə görə, `c1` kanalından bir mesaj alınır və ekrana yazdırılır. Daha sonra, `c2` kanalından bir mesaj alınır və ekrana yazdırılır.

Output:

```golang
received one
received two
```

Bu nümunədə, select açar sözü istifadə edilərək, `c1` və `c2` kanallarını dinləyən bir for döngüsü yaradıldı. Bu, mesaj alım müddətinə əsaslanaraq fərqli kanalların dinlənilməsinə imkan verir. Nəticədə, goroutin-lər arasındakı mesajlaşma müəyyən bir qaydada həyata keçirilir və select açar sözü istifadə edilərək sinxronizasiya təmin edilir.

### Timeouts

Go dilində, timeout əməliyyatları, xüsusilə şəbəkə əməliyyatları zamanı vacibdir. `time` paketindən istifadə edərək, müəyyən bir müddət gözləmək mümkündür. Əgər müddət aşılırsa, timeout xətası baş verir.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
```

Bu nümunədə `c1` və `c2` adlı iki kanal yaradılır və mesaj göndərmə əməliyyatları üçün goroutine-lər istifadə olunur. İlk olaraq, `select` açar sözü istifadə edərək, `c1` kanalından bir mesaj gözlənilir. Amma `time.After` ilə bir saniyə vaxt limiti təyin edilir. Əgər bu vaxt aşılırsa, "timeout 1" mesajı çap olunur.

Sonra, eyni əməliyyat `c2` kanalı üçün edilir və bu dəfə 3 saniyə gözlənilir.

**Output**:

```go
timeout 1
result 2
```

### Non-Blocking Channel Operations

Go dilində, kanallar adətən bloklama xüsusiyyətinə malikdir. Yəni, bir goroutine bir kanala mesaj göndərmək və ya mesaj almaq istəyirsə, həmin əməliyyat tamamlanana qədər dayanar. Lakin, `select` açar sözü istifadə edilərək, non-blocking əməliyyatlar da həyata keçirmək mümkündür.

```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hello"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no signal received")
	}

	signal := true
	select {
	case signals <- signal:
		fmt.Println("sent signal", signal)
	default:
		fmt.Println("no signal sent")
	}

	select {
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no signal received")
	}
}
```

Bu nümunədə, `messages` adlı bir kanal yaradılır və göndərilməmiş mesajın yoxlanılması üçün `select` istifadə edilir. Ardından, bir mesaj yaradılır və kanala göndərilməyə çalışılır. Əgər kanal boşdursa, "no message sent" çap olunur.

**Output**:

```go
no message received
sent message hello
no message received
no signal received
sent signal true
received signal true
```

### Closing Channels

Go dilində, kanalların bağlanması, mesajların göndərilməsi və alınması arasında sinxronizasiya yaradır. `close` funksiyası ilə bir kanal bağlana bilər. Bağlanan kanala daha artıq mesaj göndərilə bilməz və bu kanaldan mesajlar alınmaz.

```go
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
```

Bu nümunədə `jobs` adlı bir kanal yaradılır və goroutine bu kanaldan mesajlar qəbul edir. `close` funksiyası ilə kanal bağlanır və bütün işlərin tamamlandığı siqnalı `done` kanalına göndərilir.

**Output**:

```go
sent job 1
sent job 2
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
```

### Range over Channels

Go-da `range` açar sözü istifadə edərək kanaldan mesajlar almaq mümkündür. Bu zaman kanal açıq qaldığı müddətdə mesajlar qəbul edilir. Kanal bağlandıqda isə `range` döngüsü dayanır.

```go
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
```

Bu nümunədə, `queue` adlı bir kanal yaradılır və iki mesaj kanala yerləşdirilir. Kanal bağlandıqdan sonra, `range` döngüsü ilə kanalın bütün elementləri çap edilir.

**Output**:

```
one
two
```

### Timers

Go dilində `time` paketindən istifadə edərək timerlər yaratmaq mümkündür. Timerlər müəyyən bir müddətdən sonra bir siqnal göndərir. `time.NewTimer` funksiyası müəyyən edilmiş vaxtdan sonra bir siqnal göndərən bir timer yaradır.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(time.Second * 2)
}
```

Bu nümunədə, iki timer yaradılır. Birincisi 2 saniyə sonra işə düşür və "Timer 1 fired" çap edir. İkincisi isə 1 saniyə sonra işə düşmədən dayandırılır, buna görə də "Timer 2 stopped" çap olunur.

**Output**:

```go
Timer 1 fired
Timer 2 stopped
```

### Tickers

Go-da `Ticker` periodik olaraq müəyyən bir intervalla siqnal göndərən bir mexanizmdir. `time.NewTicker` funksiyası ilə ticker yaradılır və hər dəfə təyin edilmiş intervalla siqnallar alır.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```

Bu nümunədə `ticker` hər 500 millisekundda bir "tick" siqnalı göndərir və bu siqnallar çap edilir. `time.Sleep` ilə müəyyən müddətdən sonra ticker dayandırılır və proqram sonlanır.

**Output**:

```go
Tick at 2023-09-10 12:34:56.123456789 +0000 UTC m=+0.500123456
Tick at 2023-09-10 12:34:56.623456789 +0000 UTC m=+1.000123456
Tick at 2023-09-10 12:34:57.123456789 +0000 UTC m=+1.500123456
Ticker stopped
```

### Worker Pools

Worker pool-ları Go-da paralel hesablama üçün istifadə olunur. Bir neçə goroutine bir-birindən müstəqil işləri eyni zamanda yerinə yetirə bilər. Bu şəkildə bir neçə işçinin bir pool-da işləməsi təşkil edilə bilər.

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
```

Bu nümunədə 3 işçi yaradılır və 5 iş onları icra etməsi üçün göndərilir. Hər işçinin hansı işi icra etməyə başladığı və tamamladığı çap olunur.

**Output**:

```go
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
```

### WaitGroups

WaitGroup quruluşu Go dilində goroutinlər arasında sinxronizasiya təmin etmək üçün istifadə edilən bir mexanizmdir. WaitGroup quruluşu Go-nun `sync` paketində yerləşir.

WaitGroup proqramçılara işləyəcək funksiyaların sayını əvvəlcədən müəyyən etməyə və həmin funksiyaların tamamlanmasını gözləməyə imkan verir. Hər goroutine işləmə bitdikdə, WaitGroup quruluşundakı `Done()` funksiyasını çağırır. WaitGroup-dakı `Wait()` funksiyası isə bütün funksiyaların tamamlanmasını gözləyir.

Aşağıdakı nümunə, WaitGroup quruluşunun istifadəsini göstərir:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()

    fmt.Println("All workers done")
}
```

Bu nümunədə, 5 işçi funksiyası `worker()` funksiyası çağırılır. WaitGroup quruluşu, hər işçi funksiyası başlamazdan əvvəl `Add()` funksiyası ilə gözlənilən işlərin sayını artırır. Hər işçi funksiyası tamamlandıqda, `Done()` funksiyası ilə bir işin başa çatdığı bildirilir. `Wait()` funksiyası isə bütün işlərin bitməsini gözləmək üçün istifadə olunur. Nəticədə, bütün işçi funksiyaları bitdikdən sonra `main()` funksiyası "All workers done" mesajını çap edir.

**Ouput**:

```go

Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 5 done
Worker 3 done
Worker 2 done
Worker 1 done
Worker 4 done
All workers done
```

### Rate Limiting

Rate Limiting, bir API və ya xidmətin istifadəçinin müəyyən bir zaman ərzində edə biləcəyi sorğu sayını məhdudlaşdırmaq üçün istifadə edilən bir metoddur. Bu metod, tətbiqin və ya xidmətin həddindən artıq yüklənməsinin qarşısını almaq üçün istifadə oluna bilər.

Go dilində, Rate Limiting etmək üçün `time` paketi istifadə oluna bilər. Bu paket, müəyyən bir zaman çərçivəsində müəyyən bir əməliyyatın yerinə yetirilməsi üçün nə qədər gözləmək lazım olduğunu hesablamaq üçün istifadə edilir.

Aşağıdakı nümunədə, `time.Ticker` və `time.Sleep` istifadə edilərək Rate Limiting nümunəsi göstərilir:

```golang
package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }
}
```

Bu nümunədə, `requests` adlı bir kanal yaradılır və içinə 5 sorğu əlavə edilir. Daha sonra `limiter` adlı bir `time.Ticker` yaradılır və 200 millisaniyəlik bir müddətlə məhdudlaşdırılır.

Daha sonra, `requests` kanalındakı hər bir sorğu məhdudlaşdırıcıya uyğun olaraq işlədir. Hər sorğu arasında 200 millisaniyəlik fasilə ilə işlənməsi təmin edilir.

```
request 1 2023-05-23 15:56:01.46705 +0300 +03 m=+0.201167209
request 2 2023-05-23 15:56:01.667008 +0300 +03 m=+0.401132584
request 3 2023-05-23 15:56:01.867014 +0300 +03 m=+0.601146918
request 4 2023-05-23 15:56:02.067052 +0300 +03 m=+0.801193043
request 5 2023-05-23 15:56:02.267025 +0300 +03 m=+1.001173793
```

Bu şəkildə, Rate Limiting istifadə edərək sorğular müəyyən bir sürətlə işlənir.

### Atomic Counters

Atomic Counters mövzusu, Go dilində paralelliyi idarə etmək üçün istifadə olunan bir mövzudur. Çoxsaylı Go proqramları bir neçə goroutine tərəfindən paylaşılan verilərlə işlədiyinə görə, bu verilərin eyni anda bir neçə goroutine tərəfindən dəyişdirilməsi labüddür. Bu dəyişikliklər düzgün idarə edilmədikdə, proqram gözlənilməz şəkildə işləyə bilər.

Go dilində Atomic Counter-lər `sync/atomic` paketində mövcuddur. Bu paketdə `AddInt64`, `CompareAndSwapInt64`, `SwapInt64`, `LoadInt64`, və `StoreInt64` kimi metodlar mövcuddur. Bu metodlar, bir neçə goroutine tərəfindən eyni anda daxil olunan dəyişkənin təhlükəsiz şəkildə dəyişdirilməsini təmin edir.

Aşağıdakı nümunə, Atomic Counter istifadə edərək, 10 goroutine tərəfindən paylaşılan bir sayaç dəyişkəninin təhlükəsiz şəkildə artırılmasını göstərir:

```golang
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64 = 0
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    wg.Wait()
    fmt.Println("Counter:", counter)
}
```

Bu nümunədə, `counter` adlı dəyişkən `int64` tipində bir Atomic Counter olaraq təyin olunur və başlanğıcda 0 dəyəri ilə müəyyən edilir. Daha sonra, 10 goroutine yaradılır və hər biri `atomic.AddInt64` metodundan istifadə edərək `counter` dəyişkənini artırır. Bu metod `&counter` ünvanını alır və `counter` dəyişkəninə atomik şəkildə 1 əlavə edir.

Nəticədə, Atomic Counter veri quruluşu, çoxsaylı goroutine-lər tərəfindən paylaşılan dəyişkənlərin təhlükəsiz şəkildə dəyişdirilməsini təmin edir və paralellik idarəsində mühüm rol oynayır.

## Sorting

Golang-də sıralama əməliyyatı üçün `sort` paketi istifadə olunur. Bu paket daxilində müxtəlif məlumat tiplərinə aid sıralama əməliyyatları aparmaq üçün funksiyalar mövcuddur.

Ən çox istifadə olunan iki sıralama funksiyası `sort.Ints()` və `sort.Strings()` funksiyalarıdır. `sort.Ints()` funksiyası `int` tipində slice-i kiçikdən böyüyə doğru sıralayır. `sort.Strings()` funksiyası isə `string` tipində slice-i əlifba sırası ilə sıralayır.

Bundan əlavə, `sort.Float64s()` funksiyası da mövcuddur və `float64` tipində slice-i kiçikdən böyüyə doğru sıralayır.

Bundan başqa, `sort.Slice()` funksiyası ilə xüsusi sıralama əməliyyatları da həyata keçirilə bilər. Bu üsul, bir slice və bir `Less` funksiyası alır. `Less` funksiyası iki slice elementini müqayisə etmək üçün istifadə olunur və nəticəsinə əsasən elementlərin sıralanmasına qərar verilir. Bu funksiya ilə slice elementlərinə görə xüsusi sıralama əməliyyatları həyata keçirilə bilər.

Məsələn, aşağıdakı nümunədə `sort.Slice()` funksiyası istifadə edilərək xüsusi bir sıralama əməliyyatı həyata keçirilir:

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
```

Bu nümunədə, `Person` tipində bir slice təyin edilir və `sort.Slice()` funksiyası istifadə edilərək `Age` sahəsinə görə kiçikdən böyüyə doğru sıralama aparılır.

Output:

```
[{Charlie 20} {Alice 25} {Bob 30}]
```

İlk nümunədə, `sort.Ints()` funksiyası istifadə edilərək `ints` adlı bir slice kiçikdən böyüyə doğru sıralanır və nəticə ekrana yazdırılır:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{5, 2, 7, 1, 9, 3, 8, 6, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	if sort.IntsAreSorted(ints) {
		fmt.Println("Slice is sorted.")
	} else {
		fmt.Println("Slice is not sorted.")
	}
}
```

Output:

```go
[1 2 3 4 5 6 7 8 9]
Slice is sorted.
```

İkinci nümunədə isə `sort.Strings()` funksiyası istifadə edilərək `strings` adlı bir slice əlifba sırasına görə sıralanır və nəticə ekrana yazdırılır:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}
	sort.Strings(strings)
	fmt.Println(strings)

	if sort.StringsAreSorted(strings) {
		fmt.Println("Slice is sorted.")
	} else {
		fmt.Println("Slice is not sorted.")
	}
}
```

Output:

```go
[apple banana cherry date elderberry fig]
Slice is sorted.
```

## Panic

`panic`, Go dilində bir proqram işləmə zamanı gözlənilməz bir vəziyyətlə qarşılaşdıqda, proqramın icrasını dayandıraraq bir xəta mesajı göstərməsini təmin edən xüsusiyyətdir.

`panic` ciddi səhvləri idarə etmək üçün istifadə edilməlidir və mümkün qədər qarşısını almağa çalışılmalıdır. Lakin, kritik bir xəta baş verdikdə, proqramın dərhal dayandırılması lazım olan vəziyyətlərdə `panic` faydalı ola bilər.

#### `panic` istifadə nümunəsi:

```golang
package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic yakalandı:", err)
		}
	}()

	divideByZero()
	fmt.Println("Bu kod çalışacak mı?")
}

func divideByZero() {
	numerator := 10
	denominator := 0
	result := numerator / denominator
	fmt.Println("Sonuç:", result)
}
```

Bu nümunədə, `divideByZero` adlı bir funksiya var. Bu funksiya daxilində `numerator` 10 və `denominator` 0 olaraq təyin edilir, daha sonra `result` dəyişəninə `numerator`u `denominator`ə bölürük. Lakin, `denominator`un 0 olması vəziyyətində, sıfıra bölmə xətası baş verir və bu, `panic` vəziyyətini yaradır.

#### **Açıqlama**:

1. **Panic və recover istifadəçi:**\
   `defer` istifadə edərək bir funksiyanı gecikdiririk, burada `recover()` funksiyası `panic` vəziyyətlərini ələ keçirir. Əgər `panic` vəziyyəti yaranarsa, `recover()` həmin `panic` səbəbini tapır və işləyir.
2. **Sıfıra bölmə xətası:**\
   `divideByZero()` funksiyasında sıfıra bölmə cəhdi olur, bu isə `panic` vəziyyətinə səbəb olur. Lakin, `recover()` bu `panic`i ələ keçirir və proqramın dayanmadan davam etməsinə imkan yaradır.
3. **Nəticə:**\
   Proqram bu vəziyyəti tutaraq "`Panic yakalandı: runtime error: integer divide by zero`" kimi bir mesaj göstərir və proqram sonrakı kodu icra etməyə davam edir.

Bu yanaşma proqramların gözlənilməz xətaları idarə etməsinə və stabil qalmasına kömək edir.


## Defer

`defer`, Go dilindəki xüsusi bir açar sözdür və müəyyən bir funksiyanın sonunda icra ediləcək funksiyaları və ya ifadələri təyin etmək üçün istifadə olunur. Defer ifadələri, funksiyanın sonunda nə olursa olsun, yəni funksiyanın hər hansı bir səbəbdən sona çatması halında belə icra ediləcəkdir.

Defer ifadələrini istifadə edərək, bir funksiyanın sonunda açılmış faylları, bağlanmamış verilənlər bazası bağlantılarını, şəbəkə bağlantılarını və s. tez və təhlükəsiz bir şəkildə bağlamaq mümkündür.

Nümunə olaraq, bir fayl açılır və `defer` ilə funksiyanın sonunda fayl bağlanılır:

```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	
	f, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer f.Close()

	// faylı oxuma
	b := make([]byte, 1024)
	n, err := f.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(b[:n]))
}
```

Bu nümunədə, `os.Open` funksiyası istifadə edilərək "example.txt" faylı açılır. `defer` ifadəsi istifadə edilərək fayl bağlama əməliyyatı funksiyanın sonunda icra ediləcək şəkildə planlaşdırılır. Daha sonra, `Read` funksiyası istifadə edilərək faylın məzmunu oxunur və ekrana çap edilir.

Burada diqqət edilməli mühüm bir məqam, `defer` ifadəsinin ən sona yazılmamasıdır. `defer` ifadəsi, bağlanacaq olan faylı açan ifadə ilə eyni blokda olmalıdır. Əks halda, `defer` ifadəsi funksiyanın sonuna qədər gözləməyə davam edəcəkdir.



Aşağıda `Recover` mövzusunda kod və şərhlərin azərbaycanca tərcüməsi, çıxış hissəsi isə ingiliscədir:

## Recover

`recover` funksiyası, bir Go proqramının işləmə müddətində meydana gələn panic-ləri idarə etmək üçün istifadə olunur. Əgər bir panic baş verərsə, `recover` funksiyası panic-in baş verdiyi funksiyanın içində çağırılarsa, panic səbəbindən dayandırılmış əməliyyatın idarəsini geri alır və proqramın normal şəkildə davam etməsini təmin edir.

`recover` funksiyası adətən `defer` ifadələri ilə birlikdə istifadə olunur. Bu sayədə, mümkün bir panic vəziyyətində proqramın idarəsi `defer` ifadələrindəki əməliyyatlardan ən sonuncusu olan `recover` funksiyasına ötürülür və beləliklə proqramın davam etməsi təmin edilir.

Nümunə olaraq, bir panic baş verdikdə proqramın işləməsini dayandırmaq əvəzinə, `recover` funksiyası vasitəsilə proqramın normal şəkildə davam etməsi təmin edilə bilər:

```golang
package main

import "fmt"

func main() {
    fmt.Println("Program starting...")

    // Defer ifadəsi istifadə edərək panic vəziyyətlərində işlədiləcək funksiyanı müəyyən edirik.
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic occurred: This is a panic!:", r)
        }
    }()

    fmt.Println("Program continues...")

    // Burada bilərəkdən panic yaradırıq.
    panic("This is a panic")
}
```

Bu nümunədə, `panic` ifadəsi istifadə edilərək bilərəkdən bir panic yaradılır. `defer` ifadəsi istifadə edərək panic vəziyyətlərində işlədiləcək bir funksiya müəyyən edilir. `recover` funksiyası, funksiyanın içində çağırıldığı zaman, panic vəziyyətindən geri dönən dəyəri alır və `if` strukturu istifadə edilərək bu dəyər ilə işləyir.

```
Program starting...
Program continues...
Panic occurred: This is a panic!
This is a panic
```


# String Functions

Golang, string məlumatları üzərində əməliyyatlar aparmaq üçün bir çox hazır funksiyalar təqdim edir. Bu funksiyalar, string məlumatlarını parçalamaq, birləşdirmək, müqayisə etmək, axtarmaq, dəyişdirmək və daha bir çox əməliyyat aparmaq üçün istifadə edilə bilər.

Bəzi çox istifadə olunan string funksiyaları bunlardır:

- `len(s string) int`: Verilmiş string'in uzunluğunu qaytarır.
- `s[i] byte`: Verilmiş string'dəki göstərilən indeks nömrəsindəki simvolu qaytarır.
- `s + t`: İki string'i birləşdirir.
- `strings.Split(s, sep string) []string`: Verilmiş string'i, göstərilən ayırıcı simvollara görə parçalayır və bir array olaraq return edir.
- `strings.Join(a []string, sep string) string`: Verilmiş string arrayini, göstərilən ayırıcı simvol ilə birləşdirir və tək bir string olaraq return edir.
- `strings.Contains(s, substr string) bool`: Verilmiş string daxilində, göstərilən alt string'in olub olmadığını yoxlayır.
- `strings.HasPrefix(s, prefix string) bool`: Verilmiş string'in, göstərilən prefix ilə başlayıb başlamadığını yoxlayır.
- `strings.HasSuffix(s, suffix string) bool`: Verilmiş string'in, göstərilən suffix ilə bitib-bitmədiyini yoxlayır.
- `strings.Replace(s, old, new string, n int) string`: Verilmiş string'də, göstərilən köhnə string'i, göstərilən yeni string ilə dəyişdirir. İstəyə bağlı olaraq, göstərilən sayda dəyişdirmə edir.

Nümunə olaraq, bir string daxilində müəyyən bir xarakter dizisinin neçə dəfə keçdiyini tapmaq üçün `strings.Count` funksiyası istifadə edilə bilər:

```golang
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
}
```

Output:

```go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST
```


## Text Templates

Template (şablon) faylları, Go dilində mətn əsaslı məlumatları formatlamaq üçün istifadə olunur. Bir şablon faylı, bir mətn sənədindəki məlumatların bir qismini yer tutucular və ya ifadələrlə dəyişdirərək yaradılır. Bu əməliyyat, bir şablon faylını bir məlumat strukturu ilə birləşdirərək həyata keçirilir.

Məsələn, bir e-poçt göndərimi zamanı, bir şablon faylı istifadə edərək mesajın məzmununu müəyyən bir formatda yaratmaq olar. Şablon faylında, dəyişənlər və funksiyalar istifadə edərək yer tutucular və ifadələr təyin edilə bilər. Daha sonra, bu dəyişənlər və funksiyalar real məlumat strukturu ilə birləşdirilərək nəticədə tamamlanmış mətn sənədi yaradılır.

Go dilində, şablon faylları `text/template` və `html/template` paketləri ilə yaradılır. Bu paketlər, şablon fayllarının yaradılmasını və işlənməsini təmin edən bir çox funksiya təqdim edir.

```golang
package main

import (
    "fmt"
    "os"
    "text/template"
)

type Person struct {
    Name    string
    Age     int
    Country string
}

func main() {
    person := Person{
        Name:    "John",
        Age:     30,
        Country: "USA",
    }

    tpl := "My name is {{.Name}} and I am {{.Age}} years old. I live in {{.Country}}.\n"
    t := template.Must(template.New("person").Parse(tpl))

    err := t.Execute(os.Stdout, person)
    if err != nil {
        fmt.Println("Error executing template:", err)
    }
}
```

Bu nümunədə, `Person` adlı bir məlumat strukturu təyin edilmişdir. Sonra, `tpl` dəyişənində, şablonun bir mətn parçası mövcuddur. Şablon içərisində istifadə olunacaq dəyişənlər `{{.}}` işarələri arasına yazılır. Bu dəyişənlərə məlumat strukturundakı sahələr təyin ediləcəkdir.

Daha sonra, `template.Must()` funksiyası istifadə edilərək yeni bir `Template` obyekti yaradılır və şablon mətnini `Parse()` funksiyası ilə bu obyektə təyin edilir. Sonra, `Execute()` funksiyası istifadə edilərək, `Template` obyekti, məlumat strukturu və `os.Stdout` (standart çıxış) printeri istifadə edilərək işlənir və nəticə ekrana çıxarılır.

Proqramın çıxışı belə olacaq:

```go
My name is John and I am 30 years old. I live in USA.
```

