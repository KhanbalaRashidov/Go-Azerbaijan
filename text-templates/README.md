## Text Templates

Template (şablon) faylları, Go dilində mətn əsaslı məlumatları formatlamaq üçün istifadə olunur. Bir şablon faylı, bir mətn sənədindəki məlumatların bir qismini yer tutucular və ya ifadələrlə dəyişdirərək yaradılır. Bu əməliyyat, bir şablon faylını bir məlumat strukturu ilə birləşdirərək həyata keçirilir.

Məsələn, bir e-poçt göndərimi zamanı, bir şablon faylı istifadə edərək mesajın məzmununu müəyyən bir formatda yaratmaq olar. Şablon faylında, dəyişənlər və funksiyalar istifadə edərək yer tutucular və ifadələr təyin edilə bilər. Daha sonra, bu dəyişənlər və funksiyalar real məlumat strukturu ilə birləşdirilərək nəticədə tamamlanmış mətn sənədi yaradılır.

Go dilində, şablon faylları `text/template` və `html/template` paketləri ilə yaradılır. Bu paketlər, şablon fayllarının yaradılmasını və işlənməsini təmin edən bir çox funksiya təqdim edir.

```golang
package main

import (
    "fmt"
    "os"
    "text/template"
)

type Person struct {
    Name    string
    Age     int
    Country string
}

func main() {
    person := Person{
        Name:    "John",
        Age:     30,
        Country: "USA",
    }

    tpl := "My name is {{.Name}} and I am {{.Age}} years old. I live in {{.Country}}.\n"
    t := template.Must(template.New("person").Parse(tpl))

    err := t.Execute(os.Stdout, person)
    if err != nil {
        fmt.Println("Error executing template:", err)
    }
}
```

Bu nümunədə, `Person` adlı bir məlumat strukturu təyin edilmişdir. Sonra, `tpl` dəyişənində, şablonun bir mətn parçası mövcuddur. Şablon içərisində istifadə olunacaq dəyişənlər `{{.}}` işarələri arasına yazılır. Bu dəyişənlərə məlumat strukturundakı sahələr təyin ediləcəkdir.

Daha sonra, `template.Must()` funksiyası istifadə edilərək yeni bir `Template` obyekti yaradılır və şablon mətnini `Parse()` funksiyası ilə bu obyektə təyin edilir. Sonra, `Execute()` funksiyası istifadə edilərək, `Template` obyekti, məlumat strukturu və `os.Stdout` (standart çıxış) printeri istifadə edilərək işlənir və nəticə ekrana çıxarılır.

Proqramın çıxışı belə olacaq:

```go
My name is John and I am 30 years old. I live in USA.
```

