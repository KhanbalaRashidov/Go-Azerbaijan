# For

for dövür operatırı, Go programlasdırma dilində, müəyyən bir şərt doğru olduğu halda təkrarlanan kod bloklarını ifadə etmək üçün istifadə olunur. for dövrü, bir başlanğıc halı, bir şərt və bir addım dəyərinə sahibdir.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Bu nümunədə, for dövrü, i dəyişəninin 0 dəyərindən başlayaraq, i dəyişəni < 5 şərtini ödədiyi müddətcə təkrarlanan bir dövrdür. Dövrdəki hər adımdda i dəyişəni 1 artırılır və nəticədə i dəyəri 4 olur.

for dövrü, şərt hissəsi doğru olana qədər təkrarlanır. Şərt hissəsi doğru olmadığında, dövr sonlanır. Ayrıca, for dövrü break və ya continue ifadələri ilə də idarə edilə bilir.

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        break
    }
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

Bu nümunədə, for dövrü, i dəyişəninin 0 dəyərindən başlayaraq, i dəyişəni < 10 şərtini ödəyənədək təkrarlanan bir dövrdür. Dövrdəki hər adımdda, i dəyişəni 1 artırılır.

Dövr, if şərti ilə kəsilir ve break ifadesi işə düşür. Ayrıca, i dəyəri 3 olduğunda da dövr sonlanır.

Dövr içindəki digər if şərti isə, əgər i dəyəri cüt ədəddirsə bir sonrakı addıma keçərək i dəyərini yazdırmaz. Bu şərtin səbəbi, sadəcə tək ədədləri yazdırmaq istəməmizdir.

Sonsuz dövr yaratmaq üçün aşağıdakı nümunəyə baxaq.

```go
i:=0
for  {
    if i == 100 {
        break
    }
    fmt.Println(i)
	
	i++
}
```

Bu nümunədə i dəyişənimizə 0 dəyərini mənimsədirik. Daha sonra dövr daxilində i əgər 100-ə bərabər olarsa break ifadəsi ilə dövrü sonlandırırıq, əks halda dövr davam edəcəkdir
