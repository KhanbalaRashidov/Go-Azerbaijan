# Recursion

Recursion, Go proqramlaşdırma dilində, bir funksiyanın özünü çağırmasıdır. Bu struktur, müəyyən bir şərt yerinə yetirilənə qədər funksiyanın təkrarlanaraq işləməsini təmin edir.

```golang
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

fmt.Println(factorial(5))
```

Bu nümunədə, factorial adlı bir funksiya təyin edilir. Funksiya, n adlı bir int tipində parametr qəbul edir və faktorialı hesablayır. Funksiya daxilində, if şərti istifadə edilərək n dəyərinin 0 olub-olmadığı yoxlanılır. Əgər n 0-dırsa, 1 qaytarılır. Əgər n 0 deyil, başqa bir dəyərdirsə, funksiya özünü yenidən çağıraraq faktorialı hesablayır. factorial(5) çağırıldıqda nəticə ekrana yazdırılır.

```golang
func fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

fmt.Println(fibonacci(10))
```

Bu nümunədə, fibonacci adlı bir funksiya təyin edilir. Funksiya, n adlı bir int tipində parametr qəbul edir və Fibonacci sayını hesablayır. Funksiya daxilində, if şərti istifadə edilərək n dəyərinin 2-dən kiçik olub-olmadığı yoxlanılır. Əgər n 2-dən kiçikdirsə, n dəyəri geri qaytarılır. Əks halda, funksiya özünü yenidən çağıraraq Fibonacci sayını hesablayır. fibonacci(10) çağırıldıqda nəticələr ekrana yazdırılır.

\
\
