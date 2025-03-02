Aşağıda Golang ilə PostgreSQL verilənlər bazasında necə işləmək barədə təcrübəmi paylaşacağam. Addım-addım izah edəcəyəm ki, PostgreSQL serverinə necə qoşulursunuz, CRUD (yaratma, oxuma, yeniləmə, silmə) əməliyyatlarını necə yerinə yetirirsiniz, həmçinin hazırlanan ifadələr və transaction idarəetməsini necə tətbiq edirsiniz.

---

## 1. Başlamazdan Əvvəl: Tələblər və Quraşdırma

### Nə Lazımdır:
- **PostgreSQL Serveri:** İşlək bir PostgreSQL verilənlər bazası (lokal və ya uzaq server).
- **PostgreSQL Go Driver:** [lib/pq](https://github.com/lib/pq) paketini istifadə edəcəyik.

### Quraşdırma:
Terminalınızda aşağıdakı əmri işə salın:
```bash
go get -u github.com/lib/pq
```
Bu driver Golang-ın `database/sql` paketinin PostgreSQL ilə işləməsini təmin edir.

---

## 2. Verilənlər Bazası ilə Əlaqə

İlk addım PostgreSQL serverinə qoşulmaqdır. Bunun üçün DSN (Data Source Name) formatında məlumatları təyin etmək lazımdır. DSN-də istifadəçi adı, parol, serverin ünvanı, port və verilənlər bazasının adı yer alır.

### Məsələn:
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driverini əlavə edirik
)

func main() {
	// DSN formatı: "postgres://user_name:password@localhost:5432/database_name?sslmode=disable"
	dsn := "postgres://user_name:password@localhost:5432/database_name?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Bağlantı xətası:", err)
	}
	defer db.Close()

	// Əlaqəni sınayırıq
	if err := db.Ping(); err != nil {
		log.Fatal("Verilənlər bazasına qoşularkən problem:", err)
	}

	fmt.Println("PostgreSQL verilənlər bazasına uğurla qoşuldunuz!")
}
```

> **Qeyd:** `sql.Open` funksiyası əlaqəni dərhal yoxlamır, ona görə də `db.Ping()` istifadə edərək əlaqənin aktiv olduğunu təsdiq etmək vacibdir.

---

## 3. CRUD Əməliyyatları

### a) Məlumat Əlavə Etmək (Create)
Yeni məlumat əlavə etmək üçün `INSERT` sorğusundan istifadə edirik.

```go
// Məlumat əlavə etmək nümunəsi
query := "INSERT INTO users (name, email) VALUES ($1, $2)"
result, err := db.Exec(query, "Ahmet", "ahmet@example.com")
if err != nil {
	log.Fatal("Məlumat əlavə edilərkən xəta baş verdi:", err)
}

// Əlavə olunan son ID-ni əldə edirik (PostgreSQL-də qaytarılması üçün RETURNING klauzulası lazımdır)
fmt.Println("Məlumat əlavə olundu!")
```

> **Qeyd:** PostgreSQL-də son əlavə olunan ID-ni əldə etmək üçün sorğunu `RETURNING id` klauzulası ilə yazmaq tövsiyə olunur. Məsələn:
> 
> ```go
> query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
> var lastInsertID int
> err = db.QueryRow(query, "Ahmet", "ahmet@example.com").Scan(&lastInsertID)
> if err != nil {
>     log.Fatal("ID əldə edilərkən xəta:", err)
> }
> fmt.Println("Əlavə olunan məlumatın ID-si:", lastInsertID)
> ```

### b) Məlumat Oxumaq (Read)
Mövcud məlumatları oxumaq üçün `SELECT` sorğusunu istifadə edirik.

```go
// Məlumat oxumaq nümunəsi
rows, err := db.Query("SELECT id, name, email FROM users")
if err != nil {
	log.Fatal("Sorğu zamanı xəta baş verdi:", err)
}
defer rows.Close()

for rows.Next() {
	var id int
	var name, email string
	if err := rows.Scan(&id, &name, &email); err != nil {
		log.Fatal("Məlumat oxunarkən xəta:", err)
	}
	fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
}

// Satırların oxunması zamanı yaranan xətaları yoxlayırıq
if err := rows.Err(); err != nil {
	log.Fatal("Oxuma zamanı xəta baş verdi:", err)
}
```

### c) Məlumat Yeniləmək (Update)
Mövcud məlumat üzərində dəyişiklik etmək üçün `UPDATE` sorğusundan istifadə edirik.

```go
// Məlumat yeniləmə nümunəsi
updateQuery := "UPDATE users SET email = $1 WHERE id = $2"
res, err := db.Exec(updateQuery, "yeni_email@example.com", 1)
if err != nil {
	log.Fatal("Yeniləmə zamanı xəta baş verdi:", err)
}

// Neçə sətirin yeniləndiyini öyrənirik
affectedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Əlavə edilən əldə sayını əldə edərkən xəta:", err)
}
fmt.Printf("Yenilənən sətir sayı: %d\n", affectedRows)
```

### d) Məlumat Silmək (Delete)
Artıq lazım olmayan məlumatı silmək üçün `DELETE` sorğusunu istifadə edirik.

```go
// Məlumat silmə nümunəsi
deleteQuery := "DELETE FROM users WHERE id = $1"
res, err = db.Exec(deleteQuery, 1)
if err != nil {
	log.Fatal("Silmə əməliyyatı zamanı xəta baş verdi:", err)
}

deletedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Silinən sətir sayını əldə edərkən xəta:", err)
}
fmt.Printf("Silinən satır sayı: %d\n", deletedRows)
```

---

## 4. Hazırlanan İfadələr (Prepared Statements)

Təkrarlanan sorğular üçün hazırlanan ifadələr performansı artırır və SQL enjeksiyasının qarşısını alır. Gəlin, bir neçə məlumatı hazırlanan ifadə ilə necə əlavə edəcəyimizə baxaq.

```go
// Hazırlanan ifadə ilə məlumat əlavə etmək
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id")
if err != nil {
	log.Fatal("İfadə hazırlanarkən xəta baş verdi:", err)
}
defer stmt.Close()

// Birdən çox istifadəçi əlavə etmək nümunəsi
users := []struct {
	name  string
	email string
}{
	{"Khanbala", "khanbala@example.com"},
}

for _, user := range users {
	var id int
	err := stmt.QueryRow(user.name, user.email).Scan(&id)
	if err != nil {
		log.Println("Məlumat əlavə edilərkən xəta:", err)
		continue
	}
	fmt.Printf("Əlavə olunan istifadəçi ID-si: %d\n", id)
}
```

---

## 5. Transaction (Əməliyyat) İdarəetməsi

Transaction istifadə edərək, birdən çox əməliyyatın ya hamısının uğurla, ya da heç birinin tətbiq olunmasını təmin edirik. Bu, məlumatın tutarlılığını qorumağa kömək edir.

```go
// Transaction başlatma
tx, err := db.Begin()
if err != nil {
	log.Fatal("Transaction başladılarkən xəta:", err)
}

// İlk əməliyyat: məlumat əlavə etmək
insertQuery := "INSERT INTO users (name, email) VALUES ($1, $2)"
_, err = tx.Exec(insertQuery, "Mushfig", "mushfig@example.com")
if err != nil {
	tx.Rollback() // Xəta olduqda bütün əməliyyatları geri alırıq
	log.Fatal("Transaction zamanı xəta (INSERT):", err)
}
ß
// İkinci əməliyyat: məlumat yeniləmək
updateQuery := "UPDATE users SET email = $1 WHERE name = $2"
_, err = tx.Exec(updateQuery, "mushfig_yeni@example.com", "Mushfig")
if err != nil {
	tx.Rollback()
	log.Fatal("Transaction zamanı xəta (UPDATE):", err)
}

// Transaction-u təsdiqləyirik
if err := tx.Commit(); err != nil {
	log.Fatal("Transaction commit edilərkən xəta:", err)
}

fmt.Println("Transaction uğurla tamamlandı!")
```

---

## 6. Xəta İdarəetməsi və Ən Yaxşı Təcrübələr

- **Xəta Yoxlanışı:** Hər əməliyyatdan sonra `err` yoxlayaraq proqramınızın sabit işləməsini təmin edin.
- **Bağlantı İdarəetməsi:** İşiniz bitdikdən sonra `defer db.Close()` istifadə edərək açıq əlaqələrin düzgün bağlanmasını unutmayın.
- **Hazırlanan İfadələr:** Təkrarlanan sorğularda hazırlanan ifadələrdən istifadə edərək həm təhlükəsizliyi, həm də performansı artırın.
- **Transaction İstifadəsi:** Əməliyyatların bir-birinə bağlı olduğu hallarda transaction istifadə edin ki, hər hansı bir problem yarandıqda bütün əməliyyatlar geri alınsın.
