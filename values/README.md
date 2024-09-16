# Values

Go proqramlaşdırma dilində **dəyərlər** dəyişənlər tərəfindən daşınan məlumatlardır. Bu dəyərlər sabit və ya dəyişən ola bilər və müxtəlif məlumat növlərini təmsil edir. Bu məlumat növlərini başa düşmək və düzgün istifadə etmək Go dilində proqramlaşdırmanın əsasını təşkil edir.

#### Əsas Məlumat Növləri

Go dilində bir neçə əsas məlumat növü mövcuddur:

1. **Ədədi Məlumat Növləri**:
    - **Tam Ədədlər**: Tam ədədləri təmsil etmək üçün istifadə olunur.
        - **İmzalı Tam Ədədlər**:
            - `int`: Platformaya görə dəyişə bilən ölçü (adətən 32 və ya 64 bit).
            - `int8`, `int16`, `int32`, `int64`: 8, 16, 32 və 64 bit ölçüləri ilə təyin edilir.
        - **İmzalı Olmayan Tam Ədədlər**:
            - `uint`: Platformaya görə dəyişə bilən ölçü (adətən 32 və ya 64 bit).
            - `uint8`, `uint16`, `uint32`, `uint64`: 8, 16, 32 və 64 bit ölçüləri ilə təyin edilir.
        - **`uintptr`**: Pointer dəyərlərini tam ədəd kimi saxlamaq üçün istifadə olunur.
    - **Onluq Ədədlər**: Onluq nöqtə olan ədədləri təmsil edir.
        - `float32`: Tək dəqiqlikli onluq ədəd.
        - `float64`: İkili dəqiqlikli onluq ədəd.
    - **Mürəkkəb Ədədlər**: Real və təsirli hissələri olan ədədləri təmsil edir.
        - `complex64`: `float32` real və təsirli hissələri olan mürəkkəb ədəd.
        - `complex128`: `float64` real və təsirli hissələri olan mürəkkəb ədəd.
    - **`byte`**: `uint8` üçün bir aliasdır və mətn simvollarını təmsil edir.
    - **`rune`**: `int32` üçün bir aliasdır və Unicode kod nöqtələrini təmsil edir.

2. **Boolean Məlumat Növü**:
    - **`bool`**: Yalnız iki dəyəri qəbul edə bilər: `true` (doğru) və `false` (yanlış). Məntiqi əməliyyatlar və proqramın idarə edilməsi üçün istifadə olunur.

3. **String Məlumat Növü**:
    - **`string`**: Mətn və ya simvol sətirlərini təmsil edir. Mətn əməliyyatları üçün istifadə olunur.

4. **Mürəkkəb Məlumat Növləri**:
    - **`struct`**: Müxtəlif növ məlumatları bir yerdə saxlamaq üçün istifadə olunan məlumat strukturu.
    - **`array`**: Eyni növ elementlərin sabit ölçüdə olan ardıcıllığını təmsil edir.
    - **`slice`**: Dinamik uzunluqda olan və array-ə çevik baxış təqdim edən növdür.
    - **`map`**: Açar-dəyər cütlərini saxlamaq üçün istifadə olunur.
    - **`channel`**: Go rutininin digər Go rutini ilə məlumat mübadiləsi üçün istifadə olunur.

#### Dəyişənlər və Tip Müəyyənləşdirmə

Go dilində bir dəyişən elan edildikdə, onun tipi müəyyən edilir və bu tip daha sonra dəyişdirilə bilmir. Bu, proqramın tip təhlükəsizliyini təmin edir və səhvlərin qarşısını alır. Aşağıdakı nümunə bu prinsipi göstərir:

```go
var age int
age = 32

age = "thirty-two" // compile error: cannot use "thirty-two" (type string) as type int in assignment
```

Burada `age` adlı dəyişən `int` tipi ilə elan edilir. İlk olaraq, `age` dəyişəninə `32` dəyəri təyin olunur, bu da düzgündür. Lakin, `"thirty-two"` adlı sətir tipli dəyəri `age` dəyişəninə təyin etməyə çalışdığınızda, bu, tərtib xətasına səbəb olur. Çünki `age` dəyişəni `int` tipi ilə təyin edildiyi üçün yalnız tam ədəd dəyərləri qəbul edə bilər.

#### Nəticə

Dəyərlər proqramlaşdırmada müxtəlif məqsədlər üçün istifadə olunur: məsələn, `int` dəyərini rəqəmsal əməliyyatlar üçün və `string` dəyərini mətn əməliyyatları üçün istifadə edə bilərik. Dəyərlərin doğru növlərinin seçilməsi proqramların düzgün işləməsini təmin edir.