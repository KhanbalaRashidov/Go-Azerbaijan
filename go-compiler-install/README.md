# Go Compiler Quraşdırılması

Go proqramlaşdırma dilini istifadə etməyə başlamaq üçün Go compiler (dərleyicisi) quraşdırmaq lazımdır. Aşağıda Go compiler-in müxtəlif əməliyyat sistemlərində necə quraşdırılacağına dair addım-addım təlimatları təqdim edirəm.

## 1. Go Compiler-in Yüklənməsi

**1.1. Windows**

1. **Go-nun Rəsmi Veb Saytına Gedin**:
    - [Go-nun rəsmi yükləmə səhifəsinə](https://golang.org/dl/) daxil olun.

2. **Yükləmə Faylını Seçin**:
    - Windows üçün uyğun `.msi` faylını seçin və yükləyin (64-bit və ya 32-bit versiya).

3. **Quraşdırmanı Başladın**:
    - Yüklədiyiniz `.msi` faylını açın və quraşdırma prosesini izləyin.

4. **Sistem Dəyişənlərini Konfiqurasiya Edin**:
    - Quraşdırma başa çatdıqdan sonra, `GOPATH` və `GOROOT` mühit dəyişənlərini konfiqurasiya etməyə ehtiyac ola bilər. Adətən quraşdırma avtomatik olaraq bu dəyişənləri əlavə edir, amma əl ilə əlavə etmək lazım ola bilər.

5. **Təsdiqləmə**:
    - Komanda xəttinə (`cmd`) daxil olun və aşağıdakı əmri işlədərək quraşdırmanı təsdiqləyin:
      ```sh
      go version
      ```

**1.2. macOS**

1. **Go-nun Rəsmi Veb Saytına Gedin**:
    - [Go-nun rəsmi yükləmə səhifəsinə](https://golang.org/dl/) daxil olun.

2. **Yükləmə Faylını Seçin**:
    - macOS üçün uyğun `.pkg` faylını seçin və yükləyin.

3. **Quraşdırmanı Başladın**:
    - Yüklədiyiniz `.pkg` faylını açın və quraşdırma prosesini izləyin.

4. **Sistem Dəyişənlərini Konfiqurasiya Edin**:
    - Quraşdırma başa çatdıqdan sonra, `GOPATH` və `GOROOT` mühit dəyişənlərini `~/.bash_profile` və ya `~/.zshrc` faylınıza əlavə edə bilərsiniz:
      ```sh
      export GOROOT=/usr/local/go
      export GOPATH=$HOME/go
      export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
      ```

5. **Təsdiqləmə**:
    - Terminala daxil olun və aşağıdakı əmri işlədərək quraşdırmanı təsdiqləyin:
      ```sh
      go version
      ```

**1.3. Linux**

1. **Go-nun Rəsmi Veb Saytına Gedin**:
    - [Go-nun rəsmi yükləmə səhifəsinə](https://golang.org/dl/) daxil olun.

2. **Yükləmə Faylını Seçin**:
    - Linux üçün uyğun `.tar.gz` faylını seçin və yükləyin.

3. **Quraşdırma Faylını Açın**:
    - Terminala daxil olun və yüklədiyiniz `.tar.gz` faylını açın:
      ```sh
      tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz
      ```

4. **Sistem Dəyişənlərini Konfiqurasiya Edin**:
    - `~/.bashrc` və ya `~/.zshrc` faylınıza aşağıdakı sətirləri əlavə edin:
      ```sh
      export GOROOT=/usr/local/go
      export GOPATH=$HOME/go
      export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
      ```

5. **Təsdiqləmə**:
    - Terminala daxil olun və aşağıdakı əmri işlədərək quraşdırmanı təsdiqləyin:
      ```sh
      go version
      ```

## 2. Quraşdırma Təsdiqlənməsi

Quraşdırma tamamlandıqdan sonra, `go version` əmrini işlədərək Go compiler-in düzgün quraşdırıldığını və versiyasını yoxlayın. Bu əmr Go-nun versiyasını göstərməlidir, məsələn:

```sh
go version
```

Çıxış:

```
go version go1.22.5 linux/amd64
```

Bu təlimatlarla Go compiler-i sisteminizə uğurla quraşdıra bilərsiniz.