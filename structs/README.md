# Structs

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
