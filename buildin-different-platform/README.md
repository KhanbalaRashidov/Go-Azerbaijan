### Building Go Applications on Different Platforms

Go proqramlaşdırma dili, çox platformlu tətbiqlər yaratmaq üçün geniş imkanlar təqdim edir. Bir Go proqramını tək bir əməliyyat sistemi üçün deyil, müxtəlif əməliyyat sistemləri və arxitekturalar üçün də tərtib etmək mümkündür. Bu xüsusiyyət Go-nun daxilində olan "cross-compilation" (çarpaz tərtib) dəstəyi sayəsində həyata keçirilir.

#### Əsas Əmrlər və Müddəalar

Go-da çarpaz tərtibat üçün `GOOS` (operation system) və `GOARCH` (architecture) adlı iki mühüm mühit dəyişəni var. Bu dəyişənlər istifadə olunaraq, bir tətbiq istənilən əməliyyat sistemi və arxitektura üçün tərtib edilə bilər.

##### `GOOS` Dəyərləri:
- **windows**: Windows əməliyyat sistemi
- **linux**: Linux əməliyyat sistemi
- **darwin**: macOS (Apple) əməliyyat sistemi
- **freebsd**: FreeBSD əməliyyat sistemi

##### `GOARCH` Dəyərləri:
- **amd64**: 64-bit x86 prosessorlar üçün
- **386**: 32-bit x86 prosessorlar üçün
- **arm**: ARM arxitekturası üçün
- **arm64**: 64-bit ARM arxitekturası üçün

#### Tərtibat Prosesi

Fərqli platformalar üçün Go tətbiqi qurmaq olduqca sadədir. Siz yalnız uyğun mühit dəyişənlərini təyin etməklə tətbiqinizi tərtib edə bilərsiniz. Aşağıda bəzi nümunələrə nəzər salaq.

#### Linux üçün Windows-da Tərtibat:
Bir Windows maşınında Linux üçün Go tətbiqi qurmaq üçün aşağıdakı əmrdən istifadə edə bilərsiniz:

```bash
GOOS=linux GOARCH=amd64 go build -o myapp-linux
```

#### macOS üçün Tərtibat:
Linux və ya Windows sistemində macOS üçün Go tətbiqi qurmaq üçün:

```bash
GOOS=darwin GOARCH=amd64 go build -o myapp-macos
```

#### Windows üçün Tərtibat:
Linux və ya macOS sistemində Windows üçün Go tətbiqi qurmaq üçün:

```bash
GOOS=windows GOARCH=amd64 go build -o myapp-windows.exe
```

#### ARM Arxitekturası üçün Tərtibat:
Əgər ARM cihazları üçün tərtib etmək istəyirsinizsə, məsələn, Raspberry Pi üçün:

```bash
GOOS=linux GOARCH=arm go build -o myapp-arm
```

#### Faydaları
- **Daha sürətli inkişaf**: Bir platformda işləyərək bir çox fərqli platformalar üçün tətbiq tərtib etmək mümkündür.
- **Daha az konfiqurasiya**: Əlavə alət və ya mühitə ehtiyac olmadan tək bir maşından çoxlu platforma üçün tətbiqlər tərtib etmək mümkündür.
- **Sadə tərtibat prosesi**: Yalnız `GOOS` və `GOARCH` dəyişənlərini təyin etməklə fərqli platformalar üçün tərtib həyata keçirmək mümkündür.

Bu yanaşma Go dilini xüsusilə platformadan asılı olmayan tətbiqlər və mikroservislər üçün çox uyğun edir.