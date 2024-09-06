# Strings and Runes

Strings, Go proqramlaşdırma dilində, Unicode simvollarının birləşdirilməsi ilə yaradılan bir simvol sətiridir. Hər bir simvol 1-4 bayt arasında dəyişən ölçülərdə ola bilər.

```go
str := "hello"
fmt.Println(str)
fmt.Println(str[0])
fmt.Println(str[1:3])
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "hello" dəyəri verilir. str dəyişəni ekrana çap olunur, sonra isə str\[0] ifadəsi istifadə edilərək sətirin ilk simvolu olan "h" ekrana çap edilir. str\[1:3] ifadəsi istifadə edilərək sətirin ikinci və üçüncü simvolları olan "el" ekrana çap edilir.

```go
for i, r := range "hello" {
    fmt.Printf("%d: %s\n", i, string(r))
}
```

Bu nümunədə, for döngüsü istifadə edilərək range funksiyası ilə "hello" stringindəki hər bir simvola giriş edilir. Hər simvolun mövqeyi və dəyəri ekrana çap olunur. Runes, Go proqramlaşdırma dilində, bir Unicode simvolunun birləşməsini ifadə edən bir məlumat tipidir. Runes, 1-4 bayt arasında dəyişən ölçülərdə ola bilən simvolları təmsil etmək üçün istifadə olunur.

```go
str := "こんにちは"
for i, r := range str {
    fmt.Printf("%d: %c\n", i, r)
}
```

Bu nümunədə, str adlı bir string dəyişəni təyin edilir və "こんにちは" dəyəri verilir. for döngüsü istifadə edilərək range funksiyası ilə hər bir simvola giriş edilir və simvolun mövqeyi və dəyəri ekrana çap olunur.
