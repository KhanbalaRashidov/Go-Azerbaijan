Bu repository Go dilini tez öyrənmək istəyənlər üçün hazırlanmışdır. Ümumilikdə mövzuları nümunələrlə izah edərək dilin əsas strukturlarını əhatə edir. Eyni zamanda,  Go dilini öyrənmək istəyənlər üçün müntəzəm resurs təmin etmək məqsədi daşıyır. Hər bir mövzu təsviri və başa düşülən şəkildə əhatə olunur ki, oxucular Go dilini asanlıqla və tez öyrənə bilsinlər.

Repository-ni bəyənirsinizsə, ulduz qoyub sosial media hesablarınızda paylaşa bilərsiniz ki, daha çox insana çatsın ⭐️.

<br>

# Bu Go nədir?

Golang (digər adı Go) 2007-ci ildən Google tərəfindən hazırlanmış açıq mənbəli proqramlaşdırma dilidir. O, əsasən alt sistem proqramlaşdırması üçün nəzərdə tutulmuşdur və tərtib edilə bilən və statik olaraq yazılmış dildir. İlk versiya 2009-cu ilin noyabrında buraxıldı. Onun tərtibçisi "gc" (Go Compiler) bir çox əməliyyat sistemi üçün açıq mənbə kimi işlənib hazırlanmışdır.

Golang ilk dəfə Google mühəndisləri Robert Griesemer, Rob Pike və Ken Thompson tərəfindən eksperimental məqsədlər üçün yaradılmışdır. Tənqid olunan problemlərin digər dillərdə həllini təmin etmək və onların yaxşı tərəflərini qorumaq üçün yaradılmışdır.

Go sadə, sürətli və etibarlı proqram təminatı hazırlamaq üçün nəzərdə tutulmuş açıq mənbəli proqramlaşdırma dilidir. Go dili sürətli tərtib prosesi, yüngül sintaksis strukturu və effektiv garbage collection sistemi ilə seçilir.

Go dilinin istifadəsinə müxtəlif inkişaf alətləri, paketlər və modullar daxildir. Tərtibatçılar Go dili ilə veb proqramlar, API-lər, verilənlər bazası sistemləri, şəbəkə proqram təminatı və sair kimi çoxlu müxtəlif növ proqramlar yarada bilərlər.

Go dilinin əsas xüsusiyyətlərindən biri sürətli tərtib prosesidir. Go C dilinə bənzər sintaksis strukturundan və C++ kimi dillərdə obyekt yönümlü xüsusiyyətlərdən istifadə edir. Zibil toplama, yüngül iplər, kapsullaşdırılmış tip sistemlər və dinamik yaddaşın idarə edilməsi kimi xüsusiyyətlər də Go dilində mövcuddur.

Go dilinin digər mühüm xüsusiyyəti onun effektiv paket idarəetmə sistemidir. Go dilində modullar və paketlər proqram tərtibatçılarına kodlarını nizamlı şəkildə təşkil etməyə və təkmilləşdirməyə imkan verir. Beləliklə, tərtibatçılar asanlıqla təkrar istifadə edilə bilən kodları yaza bilərlər.

Nəticədə, Go dilinin əsas xüsusiyyətlərinə sürətli tərtib prosesi, yüngül sintaksis strukturu, effektiv paket idarəetməsi, zibil toplama və kapsullaşdırılmış tip sistemlər kimi xüsusiyyətlər daxildir. Go-da inkişaf etdirmək bir çox proqram növləri üçün əlverişli seçimdir və proqram təminatının hazırlanması prosesini sürətləndirməyə kömək edə bilər.

<br>

# Məzmun:

- [Values](#values)
- [Variables](#variables)
- [Constants](#constants)
- [If-Else](#if-else)
- [For](#for)
- [Switch](#switch)
- [Arrays](#arrays)
- [Slices](#slices)
- [Maps](#maps)
- [Range](#range)
- [Functions](#functions)
- [Variadic Functions](#variadic-functions)
- [Closures](#closures)
- [Recursion](#recursion)
- [Pointers](#pointers)
- [String Functions](#string-functions)
- [Structs](#structs)
- [Methods](#methods)
- [Interfaces](#interfaces)
- [Struct Embedding](#struct-embedding)
- [Errors](#errors)
- [Goroutines](#goroutines)
- [Channel](#channel)
- [Select](#select)
- [Timeouts](#timeouts)
- [Non-Blocking Channel](#non-blocking-channel)
- [Closing Channels](#closing-channels)
- [Range over Channels](#range-over-channels)
- [Timers](#timers)
- [Tickers](#tickers)
- [Worker Pools](#worker-pools)
- [Wait Groups](#wait-groups)
- [Rate Limiting](#rate-limiting)
- [Atomic Counters](#atomic-counters)
- [Sorting](#sorting)
- [Panic](#panic)
- [Defer](#defer)
- [Recover](#recover)
- [Strings and Runes](#strings-and-runes)
- [Text Templates](#text-templates)
- [JSON](#json)

---

<br>

# Values

Dəyərlər Go proqramlaşdırma dilində dəyişənlər tərəfindən daşınan məlumatlardır. Dəyərlər sabit və ya dəyişən ola bilər və müxtəlif məlumat növlərini təmsil edir.

Go dilində əsas məlumat növləri bunlardır:

- Ədədi məlumat növləri: Ədədi dəyərləri ifadə etmək üçün int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr, float32, float64, complex64, complex128, byte, rune kimi məlumat növlərindən istifadə olunur.
- Boolean məlumat növü: bool məlumat növü yalnız doğru və ya yanlış qiymətlər qəbul edə bilən məlumat növüdür.
- String data type: string data type mətn və ya simvol sətirinin qiymətlərini ifadə etmək üçün istifadə olunur.
- Mürəkkəb məlumat növləri: struct, array, slice, map, channel kimi məlumat tipləri çoxsaylı məlumat elementlərini bir yerdə saxlamaq və müəyyən bir məqsədə xidmət etmək üçün istifadə olunur.

Dəyərlərin növləri Go dilində statik olaraq müəyyən edilir. Dəyişənə qiymət təyin etdikdə bu dəyişənin tipi müəyyən edilir və biz bu tipi sonra dəyişə bilmərik.

```golang
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu nümunədə yaş dəyişəni ilk olaraq int kimi elan edilir və sonra dəyişən yaş üçün 32 rəqəmi təyin edilir. Sonra "thirty-two" sətri yaş dəyişəninə təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipli qiymət təyin edilə bilməz.
Dəyərlər proqramları müxtəlif məqsədlər üçün istifadə etməyə imkan verir. Məsələn, rəqəmsal əməliyyatı yerinə yetirmək üçün int dəyərindən istifadə edə bilərik və ya mətn əməliyyatını yerinə yetirmək üçün string dəyərindən istifadə edə bilərik.

<br>

---

<br>

<br><br>

# Variables

Dəyişənlər Go proqramlaşdırma dilində məlumatların saxlanması üçün istifadə olunan əsaslardan biridir. Dəyişənə qiymət təyin etməklə biz proqramda həmin dəyişənin saxladığı qiymətdən istifadə edə bilərik.

Dəyişənlərin elan edilməsi var açar sözü ilə həyata keçirilir. Müəyyən ediləcək dəyişənin adı və məlumat tipi müəyyən edilir. Heç bir ilkin dəyər verilmədikdə, Go dilində dəyişənlər standart dəyərə malikdirlər.

```golang
var name string
name = "John Doe"
```

Bu nümunədə ad adlı dəyişən müəyyən edilir və bu dəyişənin tipi sətir kimi müəyyən edilir. Sonra, ad dəyişəninə "John Doe" sətri təyin edilir.
Dəyişənlərin dəyəri dəyişdirilə bilər və onlar müxtəlif məlumat tiplərində ola bilər.

```golang
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
````

Bu misalda əvvəlcə yaş adlı dəyişən müəyyən edilir və bu dəyişənin növü int kimi müəyyən edilir. Sonra, dəyişən yaşa 32 nömrəsi verilir.

Daha sonra yaş dəyişəninə "thirty-two" sətri təyin edilməyə çalışılır və proqram xəta verir. Yaş dəyişəni int növü kimi təyin olunduğu üçün ona sətir tipi dəyəri təyin edilə bilməz.

Dəyişənlər proqramlarda müəyyən bir məqsədə xidmət etmək üçün istifadə olunur. Məsələn, istifadəçinin adı, yaşı və ya bir sıra ədədi dəyərlər saxlanıla bilər.


<br><br>

# Constants

Sabitlər (constants) Go proqramlaşdırma dilində proqramın heç bir yerində dəyişdirilə bilməyən sabit dəyərlərdir. Sabitlər bir dəfə müəyyən edilir və sonra proqramın istənilən yerində istifadə edilə bilər.

Sabitlər const açar sözü ilə müəyyən edilir və onların məlumat növləri göstərilir. Sabitlər işə salınmalıdır və sonra dəyişdirilə bilməz.

```golang
const pi = 3.14159
const welcomeMessage = "Welcome to Go programming"
````


Bu nümunədə, sabit pi float64 növü kimi müəyyən edilir və 3.14159-a mənimsədilir. Sonra, welcomeMessage sabiti sətir kimi müəyyən edilir və "Welcome to Go programming" olaraq işə salınır.

Sabitlər müəyyən bir məqsəd üçün istifadə ediləcək proqramların sabit dəyərlərini saxlamaq üçün istifadə olunur. Məsələn, pi sayı və ya xüsusi mesaj və ya səhv kodu. Sabitlər proqramları daha oxunaqlı və davamlı etməyə kömək edir və proqramın müxtəlif yerlərində eyni sabitdən dəfələrlə istifadə etməyi asanlaşdırır.

Sabitlər proqramın istənilən yerində istifadə oluna bildiyi üçün onlar müxtəlif fayllar arasında da paylaşıla bilər. Bundan əlavə, sabitlərin dəyişdirilə bilməməsi proqram səhvlərini azaldır və təhlükəsizliyi artırır.


<br><br>

# if/else

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

Bu nümunədə, if ifadəsi, x dəyişənin 10 ve 20 arasında olması halında "x is between 10 and 20" mətnini yazdıracaqdır. Əgər x dəyişəni 10 ve 20 arasında deyilsə, else if ifadesi yoxlanılacaq  və x dəyişənin 20 və 30 arasında olması halında "x is between 20 and 30" mətnini yazdıracaqdır. Əgər x dəyişəni 10 və 30 arasında deyilsə, else bloku işələyəcəkdir və "x is not between 10 and 30" mətnini yazdıracaqdır.



