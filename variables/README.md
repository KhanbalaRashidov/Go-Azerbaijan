# Variables

Dəyişənlər Go proqramlaşdırma dilində məlumatları saxlamaq üçün istifadə olunan əsas komponentlərdən biridir. Dəyişənlərə qiymət təyin etməklə, proqramın müxtəlif yerlərində həmin qiymətləri istifadə edə bilirik. Bu, proqramın dinamikliyini və çevikliyini artırır.

#### Dəyişənlərin Elan Edilməsi

Dəyişənlərin elan edilməsi üçün `var` açar sözündən istifadə olunur. Dəyişənin adı və məlumat tipi müəyyən edilir. Əgər dəyişənə ilkin dəyər verilməzsə, Go dilində dəyişənlər standart (default) dəyərlərlə təyin edilir. Məsələn, `int` tipli dəyişənlər üçün standart dəyər `0`, `string` tipli dəyişənlər üçün isə boş sətirdir.

**Nümunə:**

```go
var name string
name = "John Doe"
```

Bu nümunədə `name` adlı dəyişən `string` tipi ilə elan edilir. Daha sonra bu dəyişənə `"John Doe"` dəyəri təyin edilir. Bu, `name` dəyişəninin sətir tipli olduğunu göstərir.

#### Dəyişənlərin Dəyərinin Dəyişdirilməsi

Dəyişənlər proqramın icrası zamanı dəyişdirilə bilər. Lakin, dəyişənin tipi müəyyən edildikdən sonra, ona uyğun olmayan tipdəki dəyəri təyin etmək mümkün deyil.

**Nümunə:**

```go
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Bu nümunədə `age` adlı dəyişən `int` tipi ilə elan edilir. İlk olaraq, `age` dəyişəninə `32` dəyəri verilir, bu, düzgündür. Lakin, daha sonra `"thirty-two"` adlı sətir tipi dəyəri `age` dəyişəninə təyin edilməyə çalışılır və bu, tərtib xətasına səbəb olur. Çünki `age` dəyişəni `int` tipinə malikdir və yalnız tam ədəd dəyərləri qəbul edə bilər.

#### Dəyişənlərin İstifadə Məqsədləri

Dəyişənlər proqramlarda müxtəlif məqsədlər üçün istifadə olunur. Məsələn, istifadəçinin adı, yaşı və digər məlumatları saxlamaq üçün dəyişənlərdən istifadə edə bilərik. Dəyişənlər proqramın müxtəlif hissələrində məlumatları saxlamaq və idarə etmək üçün əvəzsizdir.

**Misal:**

```go
var userName string
var userAge int

userName = "Alice"
userAge = 25
```

Bu nümunədə `userName` və `userAge` adlı dəyişənlər elan edilir. `userName` dəyişəninə `Alice` dəyəri, `userAge` dəyişəninə isə `25` dəyəri təyin edilir. Bu, istifadəçinin adını və yaşını saxlamaq üçün istifadə olunan dəyişənlərdir.

#### Nəticə

Dəyişənlər proqramlaşdırmanın əsasını təşkil edir. Onlar məlumatları saxlamaq, idarə etmək və proqramın müxtəlif hissələrində istifadə etmək üçün vacibdir. Dəyişənlərin tipi müəyyən edildikdən sonra, bu tipə uyğun olan dəyərlər ilə işləmək, proqramın düzgün işləməsini təmin edir.