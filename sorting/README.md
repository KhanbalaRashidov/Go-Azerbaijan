# Sorting

Golang-də sıralama əməliyyatı üçün `sort` paketi istifadə olunur. Bu paket daxilində müxtəlif məlumat tiplərinə aid sıralama əməliyyatları aparmaq üçün funksiyalar mövcuddur.

Ən çox istifadə olunan iki sıralama funksiyası `sort.Ints()` və `sort.Strings()` funksiyalarıdır. `sort.Ints()` funksiyası `int` tipində slice-i kiçikdən böyüyə doğru sıralayır. `sort.Strings()` funksiyası isə `string` tipində slice-i əlifba sırası ilə sıralayır.

Bundan əlavə, `sort.Float64s()` funksiyası da mövcuddur və `float64` tipində slice-i kiçikdən böyüyə doğru sıralayır.

Bundan başqa, `sort.Slice()` funksiyası ilə xüsusi sıralama əməliyyatları da həyata keçirilə bilər. Bu üsul, bir slice və bir `Less` funksiyası alır. `Less` funksiyası iki slice elementini müqayisə etmək üçün istifadə olunur və nəticəsinə əsasən elementlərin sıralanmasına qərar verilir. Bu funksiya ilə slice elementlərinə görə xüsusi sıralama əməliyyatları həyata keçirilə bilər.

Məsələn, aşağıdakı nümunədə `sort.Slice()` funksiyası istifadə edilərək xüsusi bir sıralama əməliyyatı həyata keçirilir:

```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 20},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}
```

Bu nümunədə, `Person` tipində bir slice təyin edilir və `sort.Slice()` funksiyası istifadə edilərək `Age` sahəsinə görə kiçikdən böyüyə doğru sıralama aparılır.

Output:

```
[{Charlie 20} {Alice 25} {Bob 30}]
```

İlk nümunədə, `sort.Ints()` funksiyası istifadə edilərək `ints` adlı bir slice kiçikdən böyüyə doğru sıralanır və nəticə ekrana yazdırılır:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{5, 2, 7, 1, 9, 3, 8, 6, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	if sort.IntsAreSorted(ints) {
		fmt.Println("Slice is sorted.")
	} else {
		fmt.Println("Slice is not sorted.")
	}
}
```

Output:

```go
[1 2 3 4 5 6 7 8 9]
Slice is sorted.
```

İkinci nümunədə isə `sort.Strings()` funksiyası istifadə edilərək `strings` adlı bir slice əlifba sırasına görə sıralanır və nəticə ekrana yazdırılır:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"apple", "banana", "cherry", "date", "elderberry", "fig"}
	sort.Strings(strings)
	fmt.Println(strings)

	if sort.StringsAreSorted(strings) {
		fmt.Println("Slice is sorted.")
	} else {
		fmt.Println("Slice is not sorted.")
	}
}
```

Output:

```go
[apple banana cherry date elderberry fig]
Slice is sorted.
```