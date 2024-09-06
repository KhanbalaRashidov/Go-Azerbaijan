
Bu repository Go dilini tez öyrənmək istəyənlər üçün hazırlanmışdır. Ümumilikdə mövzuları nümunələrlə izah edərək dilin əsas strukturlarını əhatə edir. Eyni zamanda, Go dilini öyrənmək istəyənlər üçün müntəzəm resurs təmin etmək məqsədi daşıyır. Hər bir mövzu təsviri və başa düşülən şəkildə əhatə olunur ki, oxucular Go dilini asanlıqla və tez öyrənə bilsinlər.

Repository-ni bəyənirsinizsə, ulduz qoyub sosial media hesablarınızda paylaşa bilərsiniz ki, daha çox insana çatsın ⭐️.


## Bu Go nədir?

Golang (digər adı Go) 2007-ci ildən Google tərəfindən hazırlanmış açıq mənbəli proqramlaşdırma dilidir. O, əsasən alt sistem proqramlaşdırması üçün nəzərdə tutulmuşdur və tərtib edilə bilən və statik olaraq yazılmış dildir. İlk versiya 2009-cu ilin noyabrında buraxıldı. Onun tərtibçisi "gc" (Go Compiler) bir çox əməliyyat sistemi üçün açıq mənbə kimi işlənib hazırlanmışdır.

Golang ilk dəfə Google mühəndisləri Robert Griesemer, Rob Pike və Ken Thompson tərəfindən eksperimental məqsədlər üçün yaradılmışdır. Tənqid olunan problemlərin digər dillərdə həllini təmin etmək və onların yaxşı tərəflərini qorumaq üçün yaradılmışdır.

Go sadə, sürətli və etibarlı proqram təminatı hazırlamaq üçün nəzərdə tutulmuş açıq mənbəli proqramlaşdırma dilidir. Go dili sürətli tərtib prosesi, yüngül sintaksis strukturu və effektiv garbage collection sistemi ilə seçilir.

Go dilinin istifadəsinə müxtəlif inkişaf alətləri, paketlər və modullar daxildir. Tərtibatçılar Go dili ilə veb proqramlar, API-lər, verilənlər bazası sistemləri, şəbəkə proqram təminatı və sair kimi çoxlu müxtəlif növ proqramlar yarada bilərlər.

Go dilinin əsas xüsusiyyətlərindən biri sürətli tərtib prosesidir. Go C dilinə bənzər sintaksis strukturundan və C++ kimi dillərdə obyekt yönümlü xüsusiyyətlərdən istifadə edir. Zibil toplama, yüngül iplər, kapsullaşdırılmış tip sistemlər və dinamik yaddaşın idarə edilməsi kimi xüsusiyyətlər də Go dilində mövcuddur.

Go dilinin digər mühüm xüsusiyyəti onun effektiv paket idarəetmə sistemidir. Go dilində modullar və paketlər proqram tərtibatçılarına kodlarını nizamlı şəkildə təşkil etməyə və təkmilləşdirməyə imkan verir. Beləliklə, tərtibatçılar asanlıqla təkrar istifadə edilə bilən kodları yaza bilərlər.

Nəticədə, Go dilinin əsas xüsusiyyətlərinə sürətli tərtib prosesi, yüngül sintaksis strukturu, effektiv paket idarəetməsi, zibil toplama və kapsullaşdırılmış tip sistemlər kimi xüsusiyyətlər daxildir. Go-da inkişaf etdirmək bir çox proqram növləri üçün əlverişli seçimdir və proqram təminatının hazırlanması prosesini sürətləndirməyə kömək edə bilər.


## Məzmun:

* [Values](<README (1).md#values>)
* [Variables](<README (1).md#variables>)
* [Constants](<README (1).md#constants>)
* [If-Else](<README (1).md#if-else>)
* [For](<README (1).md#for>)
* [Switch](<README (1).md#switch>)
* [Arrays](<README (1).md#arrays>)
* [Slices](<README (1).md#slices>)
* [Maps](<README (1).md#maps>)
* [Range](<README (1).md#range>)
* [Functions](<README (1).md#functions>)
* [Variadic Functions](<README (1).md#variadic-functions>)
* [Closures](<README (1).md#closures>)
* [Recursion](<README (1).md#recursion>)
* [Pointers](<README (1).md#pointers>)
* [String Functions](<README (1).md#string-functions>)
* [Structs](<README (1).md#structs>)
* [Methods](<README (1).md#methods>)
* [Interfaces](<README (1).md#interfaces>)
* [Struct Embedding](<README (1).md#struct-embedding>)
* [Errors](<README (1).md#errors>)
* [Goroutines](<README (1).md#goroutines>)
* [Channel](<README (1).md#channel>)
* [Select](<README (1).md#select>)
* [Timeouts](<README (1).md#timeouts>)
* [Non-Blocking Channel](<README (1).md#non-blocking-channel>)
* [Closing Channels](<README (1).md#closing-channels>)
* [Range over Channels](<README (1).md#range-over-channels>)
* [Timers](<README (1).md#timers>)
* [Tickers](<README (1).md#tickers>)
* [Worker Pools](<README (1).md#worker-pools>)
* [Wait Groups](<README (1).md#wait-groups>)
* [Rate Limiting](<README (1).md#rate-limiting>)
* [Atomic Counters](<README (1).md#atomic-counters>)
* [Sorting](<README (1).md#sorting>)
* [Panic](<README (1).md#panic>)
* [Defer](<README (1).md#defer>)
* [Recover](<README (1).md#recover>)
* [Strings and Runes](<README (1).md#strings-and-runes>)
* [Text Templates](<README (1).md#text-templates>)
* [JSON](<README (1).md#json>)

***


## Values

Dəyərlər Go proqramlaşdırma dilində dəyişənlər tərəfindən daşınan məlumatlardır. Dəyərlər sabit və ya dəyişən ola bilər və müxtəlif məlumat növlərini təmsil edir.

Go dilində əsas məlumat növləri bunlardır:

* Ədədi məlumat növləri: Ədədi dəyərləri ifadə etmək üçün int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr, float32, float64, complex64, complex128, byte, rune kimi məlumat növlərindən istifadə olunur.
* Boolean məlumat növü: bool məlumat növü yalnız doğru və ya yanlış qiymətlər qəbul edə bilən məlumat növüdür.
* String data type: string data type mətn və ya simvol sətirinin qiymətlərini ifadə etmək üçün istifadə olunur.
* Mürəkkəb məlumat növləri: struct, array, slice, map, channel kimi məlumat tipləri çoxsaylı məlumat elementlərini bir yerdə saxlamaq və müəyyən bir məqsədə xidmət etmək üçün istifadə olunur.

Dəyərlərin növləri Go dilində statik olaraq müəyyən edilir. Dəyişənə qiymət təyin etdikdə bu dəyişənin tipi müəyyən edilir və biz bu tipi sonra dəyişə bilmərik.

```golang
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu nümunədə yaş dəyişəni ilk olaraq int kimi elan edilir və sonra dəyişən yaş üçün 32 rəqəmi təyin edilir. Sonra "thirty-two" sətri yaş dəyişəninə təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipli qiymət təyin edilə bilməz. Dəyərlər proqramları müxtəlif məqsədlər üçün istifadə etməyə imkan verir. Məsələn, rəqəmsal əməliyyatı yerinə yetirmək üçün int dəyərindən istifadə edə bilərik və ya mətn əməliyyatını yerinə yetirmək üçün string dəyərindən istifadə edə bilərik.

***

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


## Constants

Sabitlər (constants) Go proqramlaşdırma dilində proqramın heç bir yerində dəyişdirilə bilməyən sabit dəyərlərdir. Sabitlər bir dəfə müəyyən edilir və sonra proqramın istənilən yerində istifadə edilə bilər.

Sabitlər const açar sözü ilə müəyyən edilir və onların məlumat növləri göstərilir. Sabitlər işə salınmalıdır və sonra dəyişdirilə bilməz.

```golang
const pi = 3.14159
const welcomeMessage = "Welcome to Go programming"
```

Bu nümunədə, sabit pi float64 növü kimi müəyyən edilir və 3.14159-a mənimsədilir. Sonra, welcomeMessage sabiti sətir kimi müəyyən edilir və "Welcome to Go programming" olaraq işə salınır.

Sabitlər müəyyən bir məqsəd üçün istifadə ediləcək proqramların sabit dəyərlərini saxlamaq üçün istifadə olunur. Məsələn, pi sayı və ya xüsusi mesaj və ya səhv kodu. Sabitlər proqramları daha oxunaqlı və davamlı etməyə kömək edir və proqramın müxtəlif yerlərində eyni sabitdən dəfələrlə istifadə etməyi asanlaşdırır.

Sabitlər proqramın istənilən yerində istifadə oluna bildiyi üçün onlar müxtəlif fayllar arasında da paylaşıla bilər. Bundan əlavə, sabitlərin dəyişdirilə bilməməsi proqram səhvlərini azaldır və təhlükəsizliyi artırır.

## if/else

if və else ifadələri, Go programlaşdırma dilində, müəyyən şərtlərin doğru və ya yanlış olduğu hallarda fərqli kod bloklarının işləməsini həyata keçirir.

```golang
if x > 0 {
    fmt.Println("Positive number")
} else if x < 0 {
    fmt.Println("Negative number")
} else {
    fmt.Println("Zero")
}
```

Bu nümunədə, if ifadəsi, x dəyişkənin 0 dan böyük olması halında "Positive number" mətnini yazdıracaq. Əgər x dəyişəni 0 dan böyük deyilsə, else if ifadəsi yoxlanılacaq ve x dəyişəninin 0 dan kiçik olması halında "Negative number" mətnini yazdıracaq. Əgər x dəyişəni 0 dan böyük və ya kiçik deyilsə, else bloku işləyəcək ve "Zero" mətnini yazdıracaqdır.

```golang
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

```golang
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Bu nümunədə, for dövrü, i dəyişəninin 0 dəyərindən başlayaraq, i dəyişəni < 5 şərtini ödədiyi müddətcə təkrarlanan bir dövrdür. Dövrdəki hər adımdda i dəyişəni 1 artırılır və nəticədə i dəyəri 4 olur.

for dövrü, şərt hissəsi doğru olana qədər təkrarlanır. Şərt hissəsi doğru olmadığında, dövr sonlanır. Ayrıca, for dövrü break və ya continue ifadələri ilə də idarə edilə bilir.

```golang
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

```golang
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

```golang
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

```golang
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

```golang
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

```golang
a := [5]int{1, 2, 3, 4, 5}
fmt.Println(a)
```

Bu nümunədə, a adlı bir array müəyyən edilir və elementləri {1, 2, 3, 4, 5} olaraq təyin edilir. Array-in ölçüsü elementlərin sayına bərabər olmalıdır. Bu nümunədə, a adlı array {1, 2, 3, 4, 5} dəyərləri ilə başlatılır və bütün elementləri ekrana çıxarılır.


## Slices

Dilimin əsas məlumat strukturlarından biri olan slices, Go proqramlaşdırma dilində ölçüsü dəyişdirilə bilən, elastik bir array məlumat strukturdur. Slices, array-lərin bir alt hissəsi kimi qəbul edilə bilər və Go dilində tez-tez istifadə olunur.

```golang
a := []int{1, 2, 3, 4, 5}
fmt.Println(a)
```

Bu nümunədə, a adlı bir dilim təyin edilir və elementləri {1, 2, 3, 4, 5} olaraq təyin edilir. Array-lərdən fərqli olaraq, dilimlərin ölçüsü təyin edilmə mərhələsində göstərilmir. Slices boş da təyin edilə bilər.

```golang
var a []int
a = append(a, 1)
a = append(a, 2, 3, 4, 5)
fmt.Println(a)
```

Bu nümunədə, a adlı bir dilim təyin edilir və append() funksiyası ilə elementlərinə müvafiq olaraq 1, 2, 3, 4 və 5 dəyərləri əlavə edilir. append() funksiyası, bir dilimə bir və ya bir neçə element əlavə etmək üçün istifadə olunur.

Dilimlər, array-lərdən fərqli olaraq, bir alt hissə kimi də təyin edilə bilər.

```golang
a := []int{1, 2, 3, 4, 5}
b := a[1:4]
fmt.Println(b)
```

Bu nümunədə, a adlı bir dilim təyin edilir və elementlərinə {1, 2, 3, 4, 5} dəyərləri verilir. b adlı bir dilim isə a diliminin 1-ci indeksindən 4-cü indeksinə qədər olan alt hissəsini əhatə edir. b dilimindəki elementlər {2, 3, 4} olaraq göstərilir.

## Maps

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

## Functions

Funksiya, Go proqramlaşdırma dilində, müəyyən bir vəzifəni yerinə yetirən kod bloklarını ifadə edir. Bir funksiya, bir və ya bir neçə parametr ala bilər, bir əməliyyatı həyata keçirə bilər və bir və ya bir neçə nəticə qaytara bilər.

```golang
func add(a int, b int) int {
return a + b
}

result := add(5, 10)
fmt.Println(result)
```

Bu nümunədə, add adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki int tipində parametr qəbul edir və bu parametrlər toplanaraq nəticə int tipində geri qaytarılır. add funksiyası, 5 və 10 parametrləri ilə çağırılır və nəticə fmt.Println(result) ifadəsi ilə yazdırılır.

```golang
func swap(a, b string) (string, string) {
    return b, a
}

x, y := swap("hello", "world")
fmt.Println(x, y)
```

Bu nümunədə, swap adlı bir funksiya təyin edilir. Funksiya, a və b adlı iki string tipində parametr qəbul edir və bu parametrləri bir-biri ilə dəyişdirərək geri qaytarır. swap funksiyası, "hello" və "world" parametrləri ilə çağırılır və geri qaytarılan nəticələr x və y dəyişənləri tərəfindən qəbul edilir. Nəticələr fmt.Println(x, y) ifadəsi ilə yazdırılır.

## Variadic Functions

Variadic funksiyalar, Go proqramlaşdırma dilində dəyişən sayda arqument qəbul edən funksiyalardır. Bu funksiyalar bir və ya daha çox arqument qəbul edə bilər və arqumentlərin sayını dəyişən olaraq təyin etməyə imkan verir.

```golang
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

```golang
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

## Pointers

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


## Strings and Runes

Strings, Go proqramlaşdırma dilində, Unicode simvollarının birləşdirilməsi ilə yaradılan bir simvol sətiridir. Hər bir simvol 1-4 byte arasında dəyişən ölçülərdə ola bilər.

```golang
str := "hello"
fmt.Println(str)
fmt.Println(str[0])
fmt.Println(str[1:3])
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "hello" dəyəri verilir. str dəyişəni ekrana çap olunur, sonra isə str\[0] ifadəsi istifadə edilərək sətirin ilk simvolu olan "h" ekrana çap edilir. str\[1:3] ifadəsi istifadə edilərək sətirin ikinci və üçüncü simvolları olan "el" ekrana çap edilir.

```golang
for i, r := range "hello" {
    fmt.Printf("%d: %s\n", i, string(r))
}
```

Bu nümunədə, for döngüsü istifadə edilərək range funksiyası ilə "hello" stringindəki hər bir simvola giriş edilir. Hər simvolun mövqeyi və dəyəri ekrana çap olunur. Runes, Go proqramlaşdırma dilində, bir Unicode simvolunun birləşməsini ifadə edən bir məlumat tipidir. Runes, 1-4 bayt arasında dəyişən ölçülərdə ola bilən simvolları təmsil etmək üçün istifadə olunur.

```golang
str := "こんにちは"
for i, r := range str {
    fmt.Printf("%d: %c\n", i, r)
}
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "こんにちは" dəyəri verilir. for döngüsü istifadə edilərək range funksiyası ilə hər bir simvola giriş edilir və simvolun mövqeyi və dəyəri ekrana çap olunur.

## Structs

Structlar, Go proqramlaşdırma dilində, fərqli məlumat tiplərini özündə birləşdirən bir məlumat quruluşudur. Bu quruluşda, müxtəlif məlumat tiplərinə sahib məlumatları bir arada saxlaya və bu məlumatlar üzərində əməliyyatlar apara bilərsiniz.

```golang
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

```golang
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
