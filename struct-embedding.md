# Struct embedding

Go dilində struct embedding, bir struct-ın başqa bir struct daxilində yerləşdirilməsi ilə digər struct-ın sahələrinə və metodlarına birbaşa giriş imkanı verir. Bu, təkcə kod təkrarını azaltmır, həm də kompozisiyanı təşviq edir, beləliklə, bir struct-ın başqa bir struct-a miras buraxmadan onun funksionallığını təmin etməyə imkan verir.

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d years old\n", p.Name, p.Age)
}

type Employee struct {
	Person
	Company string
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "John",
			Age:  30,
		},
		Company: "ABC Inc.",
	}

	emp.SayHello()
}
```

Bu nümunədə, Person adlı bir struct təyin edilir və Name və Age adlı iki sahəsi var. Eyni zamanda SayHello adlı bir funksiyası mövcuddur.

Employee adlı bir struct təyin edilir və onun içərisində Person struct-ı yerləşdirilir. Employee-nin əlavə olaraq Company adlı bir sahəsi də var. main funksiyasında, emp adlı bir Employee dəyişəni yaradılır. Bu dəyişənin Person sahəsinə Person tipində bir dəyər atanır. SayHello funksiyası emp dəyişəni üzərindən çağırılır və nəticə ekrana yazdırılır.

Output:

```go
Hello, my name is John and I'm 30 years old
```

Bu nümunədə, Employee adlı struktur, Person strukturunu daxili olaraq yerləşdirir. Bu, Employee strukturuna Person-un bütün sahələrinə və funksiyalarına birbaşa çıxış imkanı verir. Beləliklə, Employee strukturunu istifadə edərək həm Employee, həm də Person məlumatlarına asanlıqla giriş əldə etmək mümkündür. Bu xüsusiyyət, kod təkrarını azaldır və strukturları daha modul halına gətirir.
