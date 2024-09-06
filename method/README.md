# Method

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

Bu nümunədə, `Rectangle` adlı bir `struct` yaradılır və `width` adlı bir `float64` tipində və `height` adlı bir `float64` tipində iki xüsusiyyət təyin olunur. `area` adlı funksiya, `Rectangle` tipində parametr qəbul edir və düzbucaqlının sahəsini hesablayır. Funksiyanın qaytardığı nəticə, düzbucaqlının sahəsidir. `perimeter` adlı funksiya da `Rectangle` tipində parametr alır və düzbucaqlının perimetrini hesablayır. Funksiyanın qaytardığı nəticə, düzbucaqlının perimetridir. `r` adlı bir `Rectangle` dəyişəni yaradılır və `width` və `height` xüsusiyyətlərinə 3.0 və 4.0 dəyərləri təyin olunur. `r.area()` və `r.perimeter()` ifadələrindən istifadə edərək düzbucaqlının sahəsi və perimetri hesablanır və nəticələr ekrana yazdırılır.

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

Bu nümunədə, `Person` adlı bir `struct` yaradılır və `Name` adlı bir `string` tipində və `Age` adlı bir `int` tipində iki xüsusiyyət təyin olunur. `setName` adlı funksiya, `Person` tipində pointer (`*Person`) parametr qəbul edir və şəxsin adını dəyişir. `setAge` adlı funksiya, `Person` tipində pointer (`*Person`) parametr qəbul edir və şəxsin yaşını dəyişir. `getName` adlı funksiya, `Person` tipində parametr qəbul edir və şəxsin adını qaytarır. `getAge` adlı funksiya isə şəxsin yaşını qaytarır.

`p` adlı bir `Person` dəyişəni yaradılır və `Name` və `Age` xüsusiyyətlərinə "John Doe" və 42 dəyərləri təyin olunur. Daha sonra, `p.setName("Jane Doe")` və `p.setAge(35)` ifadələrindən istifadə edərək şəxsin adı və yaşı dəyişdirilir.
