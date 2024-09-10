# Bu Go nədir?

## Bu Go nədir?

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

#### Məzmun:

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
* [String Functions](./#strings-and-runes)
* [Structs](./#structs)
* [Methods](./#methods)
* [Interfaces](./#interfaces)
* [Struct Embedding](./#struct-embedding)
* [Errors](./#errors)
* [Goroutines](./#goroutines)
* [Channel](./#channel)
* [Select](./#select)
* [Timeouts](./#timeouts)
* [Non-Blocking Channel](README%20\(1\).md#non-blocking-channel)
* [Closing Channels](./#closing-channels)
* [Range over Channels](./#range-over-channels)
* [Timers](./#timers)
* [Tickers](README%20\(1\).md#tickers)
* [Worker Pools](README%20\(1\).md#worker-pools)
* [Wait Groups](README%20\(1\).md#wait-groups)
* [Rate Limiting](README%20\(1\).md#rate-limiting)
* [Atomic Counters](README%20\(1\).md#atomic-counters)
* [Sorting](README%20\(1\).md#sorting)
* [Panic](README%20\(1\).md#panic)
* [Defer](README%20\(1\).md#defer)
* [Recover](README%20\(1\).md#recover)
* [Strings and Runes](README%20\(1\).md#strings-and-runes)
* [Text Templates](README%20\(1\).md#text-templates)
* [JSON](README%20\(1\).md#json)

***

## Values

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

## Variables

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

## Constants

Sabitlər (constants) Go proqramlaşdırma dilində proqramın heç bir yerində dəyişdirilə bilməyən sabit dəyərlərdir. Sabitlər bir dəfə müəyyən edilir və sonra proqramın istənilən yerində istifadə edilə bilər.

Sabitlər const açar sözü ilə müəyyən edilir və onların məlumat növləri göstərilir. Sabitlər işə salınmalıdır və sonra dəyişdirilə bilməz.

```go
const pi = 3.14159
const welcomeMessage = "Welcome to Go programming"
```

Bu nümunədə, sabit pi float64 növü kimi müəyyən edilir və 3.14159-a mənimsədilir. Sonra, welcomeMessage sabiti sətir kimi müəyyən edilir və "Welcome to Go programming" olaraq işə salınır.

Sabitlər müəyyən bir məqsəd üçün istifadə ediləcək proqramların sabit dəyərlərini saxlamaq üçün istifadə olunur. Məsələn, pi sayı və ya xüsusi mesaj və ya səhv kodu. Sabitlər proqramları daha oxunaqlı və davamlı etməyə kömək edir və proqramın müxtəlif yerlərində eyni sabitdən dəfələrlə istifadə etməyi asanlaşdırır.

Sabitlər proqramın istənilən yerində istifadə oluna bildiyi üçün onlar müxtəlif fayllar arasında da paylaşıla bilər. Bundan əlavə, sabitlərin dəyişdirilə bilməməsi proqram səhvlərini azaldır və təhlükəsizliyi artırır.

## if/else

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

## For

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

## Switch

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

## Arrays

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

## Slices

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

## Maps

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

## Range

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

## Functions

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

## Variadic Functions

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

## Closures

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

## Recursion

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

## Pointers

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

## Strings and Runes

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

## Structs

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

## Methods

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

## Interfaces

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

## Struct Embedding

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

## Errors

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

## Goroutines

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

## Channel

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

\\

## Select

Go dilində, select açar sözü bir neçə kanalı (channel) dinləyərək hansı kanalın mesaj göndərdiyini müəyyənləşdirə bilər. Bu xüsusiyyət kanalların sinxronizasiyasını asanlaşdırır və fərqli goroutin-lər arasında mesajlaşmanı idarə edir.

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

Bu nümunədə, c1 və c2 adlı iki kanal yaradılır və mesaj göndərmə prosesi üçün goroutin-lər yaradılır.

main funksiyasında, select açar sözü istifadə edilərək, c1 və c2 kanalları dinlənilir. İlk olaraq, goroutin-lər arasındakı gözləmə müddətinə görə, c1 kanalından bir mesaj alınır və ekrana yazdırılır. Daha sonra, c2 kanalından bir mesaj alınır və ekrana yazdırılır.

Output:

```go
received one
received two
```

Bu nümunədə, select açar sözü istifadə edilərək, c1 və c2 kanallarını dinləyən bir for döngüsü yaradıldı. Bu, mesaj alım müddətinə əsaslanaraq fərqli kanalların dinlənilməsinə imkan verir. Nəticədə, goroutin-lər arasındakı mesajlaşma müəyyən bir qaydada həyata keçirilir və select açar sözü istifadə edilərək sinxronizasiya təmin edilir.

## Timeouts

Go dilində, **timeout** əməliyyatları xüsusilə şəbəkə əməliyyatları zamanı əhəmiyyətlidir. `time` paketi istifadə edilərək, müəyyən bir müddət gözləmək təmin edilə bilər. Əgər bu müddət keçərsə, timeout səhvi yaranır.

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

Bu nümunədə, `c1` və `c2` adlı iki kanal yaradılır və mesaj göndərmə əməliyyatları üçün goroutin-lər yaradılır.

İlk olaraq, `select` açar sözü istifadə edilərək, `c1` kanalından mesaj gözlənilir. Amma `time.After` istifadə edilərək, `c1` kanalından bir mesaj alınmazdan əvvəl 1 saniyə gözlənilir. Əgər 1 saniyədən çox müddət keçərsə, timeout səhvi yaranır və "timeout 1" mesajı ekrana yazdırılır.

Daha sonra, `select` açar sözü ilə `c2` kanalından mesaj gözlənilir. Bu dəfə `time.After` ilə 3 saniyə gözlənilir və nəticə mesajı alınır.

#### Çıxış:

```go
timeout 1
result 2
```

Bu nümunədə `time.After` istifadə edərək timeout əməliyyatları icra edilir. İlk nümunədə, `c1` kanalına 1 saniyədən əvvəl mesaj göndərilmədiyi üçün timeout səhvi yaranır. İkinci nümunədə isə, `c2` kanalına mesaj vaxtında göndərildiyi üçün mesaj qəbul edilir və ekrana yazdırılır.



## Non-Blocking Channel Operations

Go dilində, channel-lar adətən bloklayan xüsusiyyətə malikdir, yəni bir goroutine bir channel-a mesaj göndərmək və ya almaq istədikdə həmin əməliyyat tamamlanana qədər gözləyir. Amma `select` açar sözü ilə channel-ların bloklanmadan işləməsi təmin edilə bilər.

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

Bu nümunədə iki channel (`messages` və `signals`) yaradılır və onlardan bloklanmadan necə məlumat göndərmək və qəbul etmək göstərilir.

#### Output:

```go
no message received
sent message hello
no message received
no signal received
sent signal true
received signal true
```

Bu nümunədə:

* İlk öncə `messages` channel-dan mesaj alınmazdan əvvəl `no message received` mesajı göstərilir.
* Daha sonra `hello` mesajı `messages` channel-ına göndərilir və "sent message hello" yazısı ekrana çıxır.
* İkinci dəfə `messages` channel-dan mesaj almağa çalışılsa da, yeni mesaj olmadığından "no message received" göstərilir.
* Eyni üsulla `signals` channel-ı üçün sinyal göndərilməzdən əvvəl "no signal received" göstərilir, sinyal göndəriləndən sonra isə qəbul edildiyi göstərilir.

Bloklamadan əməliyyatlar `select` ilə kanal vasitəsilə sinxronlaşma və qarşılıqlı əlaqənin daha çevik formada həyata keçirilməsinə imkan yaradır.



## Closing Channels

Go dilində, channel-ların bağlanması, məlumatların göndərilməsi və alınması prosesində sinxronizasiyanı təmin edir. `close` funksiyası vasitəsilə channel bağlanır. Bağlanmış bir channel-a mesaj göndərmək mümkün deyil və həmin channel-dan daha çox mesaj alınmaz.

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

Bu nümunədə:

* `jobs` adlı bir channel yaradılır və işlər bu kanala göndərilir.
* `done` adlı bir channel yaradılır və işlərin tamamlandığını bildirmək üçün istifadə olunur.
* Bir `goroutine` yaradılır və bu `goroutine`, `jobs` channel-ından mesajları alır. Alınan mesajların sayı azaldıqca hər dəfə "received job" yazısı çap edilir.
* `for` döngüsü ilə `jobs` channel-ına 3 iş göndərilir və hər göndərmə zamanı "sent job" mesajı göstərilir.
* `jobs` channel-ı bağlandıqdan sonra `goroutine`, channel-dan bütün mesajları alır və son olaraq "received all jobs" mesajı göstərilir.

#### Output:

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

Bu nümunə, Go dilində bir channel-a mesaj göndərdikdən sonra onu necə bağlayacağınızı və bağlanmış channel-dan daha çox mesaj alınmadığını göstərir. Bağlı bir channel-dan oxumağa davam edə bilərsiniz, lakin artıq yeni mesajlar göndərilə bilməz.



## Range over Channels

Go dilində, `range` açar sözü channel-lar üzərində dövr etməyə imkan verir. Bu üsul ilə channel bağlanana qədər bütün mesajlar `for` dövrü ilə oxuna bilər.

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

Bu nümunədə:

* `queue` adlı bir channel yaradılır və bu channel-ə iki mesaj göndərilir.
* `close` funksiyası ilə channel bağlanır.
* Daha sonra, `range` açar sözü ilə for dövrü istifadə edərək, channel-dakı bütün mesajlar oxunur.

#### Output:

```go
one
two
```

Bu nümunədə `range` açar sözü ilə channel-dakı bütün mesajlar oxundu və `close` funksiyası ilə channel bağlandıqdan sonra belə, göndərilmiş bütün mesajlar alındı. `range` dövrü channel bağlanana qədər davam edir və channel-dan yeni mesajlar gəldikcə onları oxuyur.



## Timers

Go dilində timer-lər, müəyyən bir müddət keçdikdən sonra bir əməliyyatın yerinə yetirilməsini təmin etmək üçün istifadə olunur. `time` paketi daxilindəki `NewTimer` funksiyası ilə timer yaradıla bilər.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
```

Bu nümunədə:

* `NewTimer` funksiyası ilə iki timer yaradılır.
* Birinci timer (`timer1`) 2 saniyədən sonra bitəcək şəkildə təyin olunur. Timer bitdikdə `<-timer1.C` ilə gözlənilir və "Timer 1 expired" mesajı ekrana yazdırılır.
* İkinci timer (`timer2`) 1 saniyədən sonra sona çatacaq. `goroutine` ilə bu timer izlənir və əgər timer vaxtı bitərsə, "Timer 2 expired" mesajı çıxar. Lakin `Stop` funksiyası ilə bu timer vaxtı dolmadan dayandırılır və "Timer 2 stopped" mesajı ekrana yazdırılır.

#### Output:

```go
Timer 1 expired
Timer 2 stopped
```

Bu nümunədə, timer-lərdən biri müəyyən müddətdən sonra bitir və bir əməliyyat yerinə yetirilir, digəri isə vaxtı dolmadan əvvəl `Stop` funksiyası ilə dayandırılır. Timer-lərin bu cür idarə olunması zamanlama əməliyyatlarının nəzarətində faydalıdır.
