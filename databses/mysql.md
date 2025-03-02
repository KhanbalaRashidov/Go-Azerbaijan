# Mysql

Golang ilə MySQL verilənlər bazasında necə işləmək barədə təcrübəmi paylaşacağam. Aşağıdakı addımlarla həm MySQL serverinə necə qoşulacağınızı, həm də CRUD (yaratma, oxuma, yeniləmə, silmə) əməliyyatlarını, həmçinin hazırlanan ifadələr və transaction idarəetməsini nümunələrlə izah edəcəyəm.

---

## 1. Başlamazdan Əvvəl: Tələblər və Quraşdırma

### Nə Lazımdır:
- **MySQL Serveri:** İşlək bir MySQL verilənlər bazası (lokal və ya uzaq server).
- **MySQL Go Driver:** [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) paketini istifadə edəcəyik.

### Quraşdırma:
Terminalınızda aşağıdakı əmri işə salın:
```bash
go get -u github.com/go-sql-driver/mysql
```
Bu driver Golang-ın `database/sql` paketinin MySQL ilə işləməsini təmin edir.

---

## 2. Verilənlər Bazası ilə Əlaqə

İlk addımımız MySQL serverinə qoşulmaqdır. Bunun üçün DSN (Data Source Name) adlanan məlumatı müəyyən etmək lazımdır. DSN-də istifadəçi adı, parol, serverin ünvanı, port və verilənlər bazasının adı yer alır.

### Məsələn:
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driverini əlavə edirik
)

func main() {
	// DSN: "user:pass@tcp(127.0.0.1:3306)/database_name"
	dsn := "istifadeci:password@tcp(127.0.0.1:3306)/database_name"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Bağlantı xətası:", err)
	}
	defer db.Close()

	// Əlaqəni sınayırıq
	if err := db.Ping(); err != nil {
		log.Fatal("Verilənlər bazasına qoşularkən problem:", err)
	}

	fmt.Println("MySQL verilənlər bazasına uğurla qoşuldunuz!")
}
```

> **Qeyd:** `sql.Open` funksiya ilə əlaqə yaradılır, amma əlaqənin həqiqətən işlədiyini `db.Ping()` ilə yoxlamaq vacibdir.

---

## 3. CRUD Əməliyyatları

### a) Məlumat Əlavə Etmək (Create)
Yeni məlumat əlavə etmək üçün `INSERT` sorğusundan istifadə edirik.

```go
// Məlumat əlavə etmək nümunəsi
query := "INSERT INTO users (name, email) VALUES (?, ?)"
result, err := db.Exec(query, "Ahmet", "ahmet@example.com")
if err != nil {
	log.Fatal("Məlumat əlavə edilərkən xəta baş verdi:", err)
}

// Əlavə olunan son ID-ni əldə edirik
lastInsertID, err := result.LastInsertId()
if err != nil {
	log.Fatal("Son əlavə olunan ID əldə edilərkən xəta:", err)
}
fmt.Println("Əlavə olunan kaydın ID-si:", lastInsertID)
```

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
updateQuery := "UPDATE users SET email = ? WHERE id = ?"
res, err := db.Exec(updateQuery, "yeni_email@example.com", 1)
if err != nil {
	log.Fatal("Yeniləmə zamanı xəta baş verdi:", err)
}

// Neçə sətirin yeniləndiyini öyrənirik
affectedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Etkilenen satır sayını əldə edərkən xəta:", err)
}
fmt.Printf("Yenilənən satır sayı: %d\n", affectedRows)
```

### d) Məlumat Silmək (Delete)
Artıq lazım olmayan məlumatı silmək üçün `DELETE` sorğusu istifadə olunur.

```go
// Məlumat silmə nümunəsi
deleteQuery := "DELETE FROM users WHERE id = ?"
res, err = db.Exec(deleteQuery, 1)
if err != nil {
	log.Fatal("Silmə əməliyyatı zamanı xəta baş verdi:", err)
}

deletedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Silinən satır sayını əldə edərkən xəta:", err)
}
fmt.Printf("Silinən satır sayı: %d\n", deletedRows)
```

---

## 4. Hazırlanan İfadələr (Prepared Statements)

Təkrarlanan sorğular üçün hazırlanan ifadələr performansı artırır və SQL enjeksiyasının qarşısını alır. Gəlin, bir neçə məlumatı hazırlanan ifadə ilə necə əlavə edəcəyimizə baxaq.

```go
// Hazırlanan ifadə ilə məlumat əlavə etmək
stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
if err != nil {
	log.Fatal("İfadə hazırlanarkən xəta baş verdi:", err)
}
defer stmt.Close()

// Birdən çox istifadəçi əlavə etmək nümunəsi
users := []struct {
	name  string
	email string
}{
	{"Mehmet", "mehmet@example.com"},
	{"Ayşə", "ayse@example.com"},
}

for _, user := range users {
	res, err := stmt.Exec(user.name, user.email)
	if err != nil {
		log.Println("Məlumat əlavə edilərkən xəta:", err)
		continue
	}
	id, _ := res.LastInsertId()
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
_, err = tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Emre", "emre@example.com")
if err != nil {
	tx.Rollback() // Xəta olduqda bütün əməliyyatları geri alırıq
	log.Fatal("Transaction zamanı xəta (INSERT):", err)
}

// İkinci əməliyyat: məlumat yeniləmək
_, err = tx.Exec("UPDATE users SET email = ? WHERE name = ?", "emre_yeni@example.com", "Emre")
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

- **Xəta Yoxlanışı:** Hər əməliyyatdan sonra `err` yoxlayaraq proqramınızın möhkəm işləməsini təmin edin.
- **Bağlantı İdarəetməsi:** İşiniz bitdikdən sonra `defer db.Close()` istifadə edərək açıq əlaqələrin düzgün bağlanmasını unutmayın.
- **Hazırlanan İfadələr:** Təkrarlanan sorğularda hazırlanan ifadələrdən istifadə edərək həm təhlükəsizliyi, həm də performansı artırın.
- **Transaction İstifadəsi:** Əməliyyatların bir-birinə bağlı olduğu hallarda transaction istifadə edin ki, hər hansı bir problem yarandıqda bütün əməliyyatlar geri alınsın.
