# Interface

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
