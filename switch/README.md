# Switch

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
