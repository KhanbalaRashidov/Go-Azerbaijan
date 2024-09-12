## JSON

JSON (JavaScript Object Notation), yüngül bir veri mübadiləsi formatıdır və insanlar tərəfindən oxunması və yazılması asandır. Go dilində, `encoding/json` paketi JSON məlumatları ilə işləmək üçün istifadə edilə bilər.

JSON məlumatları, açar/dəyər cütləri və ya siyahılardan ibarət ola bilər. Açarlar bir string, dəyərlər isə JSON məlumatı ola bilər. JSON məlumatları `{}` (qıvrım mötərizələr) ilə göstərilən obyektlər şəklində də ola bilər.

Misal JSON məlumatı belə görünə bilər:

```json
{
    "name": "John Doe",
    "age": 30,
    "city": "New York"
}
```

Bu məlumat, `name`, `age` və `city` adlı üç açar saxlayır. `name` açarı bir string dəyəri, `age` açarı bir sayısal dəyər və `city` açarı da bir string dəyəridir.

Go dilində, JSON məlumatları üçün `struct` quruluşu istifadə edilə bilər. `struct` içərisindəki sahələr JSON məlumatındakı açarlarla uyğunlaşdırıla bilər. Daha sonra, `encoding/json` paketində olan `Marshal()` və `Unmarshal()` funksiyaları istifadə edilərək, Go obyekti və JSON məlumatı arasında dönüşüm edilə bilər.

Misal üçün, bir `Person` struct'ını tanıdaq və bu struct-ı JSON məlumatı ilə uyğunlaşdıraq:

```golang
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}
```

Burada, `Person` struct-ının sahələri JSON məlumatındakı açarlarla uyğunlaşdırılmışdır. `Name` sahəsi üçün `name`, `Age` sahəsi üçün `age`, `City` sahəsi üçün isə `city` açarları istifadə olunmuşdur.

Bu struct-ı JSON məlumatına çevirmək üçün, `json.Marshal()` funksiyası istifadə edilə bilər:

```go
person := Person{Name: "John Doe", Age: 30, City: "New York"}
jsonData, err := json.Marshal(person)
if err != nil {
    fmt.Println("Error encoding JSON:", err)
}
fmt.Println(string(jsonData)) // {"name":"John Doe","age":30,"city":"New York"}
```

Burada, `Person` struct-ı bir `person` dəyişəni olaraq təyin edilir və `json.Marshal()` funksiyası istifadə edilərək `person` obyekti JSON məlumatına çevrilir. Əldə edilən JSON məlumatı `string()` funksiyası istifadə edilərək string tipinə çevrilir və ekrana çap olunur.

Başqa bir misal isə, JSON məlumatlarını Go məlumat strukturlarına çevirmək və əksinə etmək üçün `json.Marshal()` və `json.Unmarshal()` funksiyalarını istifadə etməkdir. Məsələn:

```golang
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p1 := Person{"Alice", 30}
    jsonStr, err := json.Marshal(p1)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(jsonStr))

  
    var p2 Person
    jsonStr = []byte(`{"Name":"Bob","Age":40}`)
    err = json.Unmarshal(jsonStr, &p2)
    if err != nil {
        panic(err)
    }
    fmt.Println(p2)
}
```

Bu nümunədə, `Person` adlı bir struktur təyin edilir və `json.Marshal()` və `json.Unmarshal()` funksiyaları istifadə edilərək bu struktur JSON məlumatı olaraq kodlanır və təhlil edilir.

#### **Output**

```go
{"Name":"Alice","Age":30}
{Bob 40}
```