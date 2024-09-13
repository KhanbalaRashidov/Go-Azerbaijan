# Go Proyekti Yaratma

Go proqramlaşdırma dilində yeni bir layihə başlatmaq, müəyyən addımları izləməklə həyata keçirilir. Aşağıda Go-da yeni bir layihənin yaradılması üçün əsas addımlar və təlimatlar təqdim olunur.

## 1. Go-nun Quraşdırılmasını Təsdiqləmə

Əvvəlcə Go-nun sisteminizdə düzgün quraşdırıldığını təsdiqləməlisiniz. Terminala aşağıdakı əmri yazaraq Go-nun versiyasını yoxlayın:

```sh
go version
```

Çıxışda Go-nun versiyası göstərilməlidir, məsələn:

```
go version go1.20.2 linux/amd64
```

## 2. Proyekt Qovluğunun Yaradılması

Yeni bir Go layihəsi yaratmaq üçün ilk olaraq layihə qovluğunu yaradın:

```sh
mkdir my-go-project
cd my-go-project
```

Bu əmr yeni bir qovluq yaradacaq və həmin qovluğa daxil edəcəkdir.

## 3. Go Modulu İlə Proyekti İdarə Etmək

Go modulları layihənin asılılıqlarını idarə etmək üçün istifadə olunur. Proyektinizin kök qovluğunda `go.mod` faylı yaradaraq layihəni bir Go modulu kimi başlatmaq üçün aşağıdakı əmri işlədin:

```sh
go mod init my-go-project
```

Bu əmr `go.mod` faylını yaradacaq və layihənizin adını `my-go-project` olaraq təyin edəcək.

## 4. Go Kodunu Yazmaq

Layihə qovluğunda əsas Go kodunu yazmaq üçün bir `main.go` faylı yaradın. Məsələn:

```sh
touch main.go
```

`main.go` faylını bir mətn redaktoru ilə açın və aşağıdakı kodu əlavə edin:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Bu sadə proqram ekrana "Salam, Dünya!" mesajını çap edəcək.

## 5. Proyekti İcra Etmək

Layihənizi icra etmək üçün terminalda aşağıdakı əmri yazın:

```sh
go run main.go
```

Çıxışda proqramın çap etdiyi mesajı görməlisiniz:

```
Hello, World!!
```

## 6. Proyekti Build Etmək

Proyekti binary fayl olaraq build etmək üçün aşağıdakı əmri istifadə edin:

```sh
go build
```

Bu əmr, `main.go` faylını oxuyacaq və `my-go-project` adında bir icra edilə bilən fayl yaradacaq. Binary faylı işə salmaq üçün aşağıdakı əmri yazın:

```sh
./my-go-project
```

Bu əmr eyni "SHello, World!" mesajını çap edəcək.

## 7. Proyektin Asılılıqlarını İdarə Etmək

Əgər layihənizdə xarici asılılıqlardan istifadə edirsinizsə, `go get` əmrini istifadə edərək bu asılılıqları əlavə edə bilərsiniz. Məsələn, bir kitabxana əlavə etmək üçün:

```sh
go get github.com/some/library
```

Bu əmr, `go.mod` faylınıza lazımlı asılılığı əlavə edəcək və `go.sum` faylını müvafiq şəkildə yeniləyəcək.

## 8. Proyekti Test Etmək

Layihənizdə testlər yazmışsınızsa, testləri icra etmək üçün aşağıdakı əmrdən istifadə edin:

```sh
go test
```

Bu əmr layihə qovluğunda olan bütün test fayllarını icra edəcək və nəticələri göstərəcək.

## 9. Proyekti Formatlamaq

Go kodunun standart formatlaşdırılmasını təmin etmək üçün `go fmt` əmrindən istifadə edin:

```sh
go fmt ./...
```

Bu əmr layihə qovluğundakı bütün Go fayllarını standart formatda formatlayacaq.

## Nəticə

Bu təlimatlarla siz Go-da yeni bir layihə yaratmağı, kod yazmağı, build etməyi və idarə etməyi öyrəndiniz. Go-nun modul idarəetməsi, test yazma və kod formatlama kimi xüsusiyyətlərindən istifadə edərək layihənizi effektiv bir şəkildə inkişaf etdirə bilərsiniz.