Aşağıda Golang ilə MSSQL (Microsoft SQL Server) verilənlər bazasında necə işləmək barədə təcrübəmi paylaşacağam. Addım-addım izah edəcəyəm ki, MSSQL serverinə necə qoşulursunuz, CRUD (yaratma, oxuma, yeniləmə, silmə) əməliyyatlarını necə yerinə yetirirsiniz, həmçinin hazırlanan ifadələr və transaction idarəetməsini necə tətbiq edirsiniz.

---

## 1. Başlamazdan Əvvəl: Tələblər və Quraşdırma

### Nə Lazımdır:
- **MSSQL Serveri:** İşlək bir Microsoft SQL Server (lokal və ya uzaq server).
- **MSSQL Go Driver:** [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) paketini istifadə edəcəyik.

### Quraşdırma:
Terminalınızda aşağıdakı əmri işə salın:
```bash
go get -u github.com/denisenkom/go-mssqldb
```
Bu driver, Golang-ın `database/sql` paketinin MSSQL ilə işləməsini təmin edir.

---

## 2. Verilənlər Bazası ilə Əlaqə

İlk addım MSSQL serverinə qoşulmaqdır. Bunun üçün DSN (Data Source Name) formatında məlumatları təyin etmək lazımdır. DSN-də istifadəçi adı, parol, serverin ünvanı, port və verilənlər bazasının adı yer alır.

### Məsələn:
```go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // MSSQL driverini əlavə edirik
)

func main() {
	// DSN formatı: "sqlserver://user_name:password@localhost:1433?database=database_name"
	dsn := "sqlserver://user_name:password@localhost:1433?database=database_name"
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal("Bağlantı xətası:", err)
	}
	defer db.Close()

	// Əlaqəni sınayırıq
	if err := db.Ping(); err != nil {
		log.Fatal("Verilənlər bazasına qoşularkən problem:", err)
	}

	fmt.Println("MSSQL verilənlər bazasına uğurla qoşuldunuz!")
}
```

> **Qeyd:** `sql.Open` funksiyası əlaqəni dərhal yoxlamır, ona görə də `db.Ping()` istifadə edərək əlaqənin aktiv olduğunu təsdiq etmək vacibdir.

---

## 3. CRUD Əməliyyatları

### a) Məlumat Əlavə Etmək (Create)
Yeni məlumat əlavə etmək üçün `INSERT` sorğusundan istifadə edirik. MSSQL-də əlavə olunan ID-ni əldə etmək üçün `OUTPUT INSERTED.id` klauzulası istifadə olunur.

```go
// Məlumat əlavə etmək nümunəsi
query := "INSERT INTO users (name, email) OUTPUT INSERTED.id VALUES (?, ?)"
var lastInsertID int
err := db.QueryRow(query, "khanbala", "khanbala@example.com").Scan(&lastInsertID)
if err != nil {
	log.Fatal("Məlumat əlavə edilərkən xəta baş verdi:", err)
}
fmt.Printf("Əlavə olunan məlumatın ID-si: %d\n", lastInsertID)
```

### b) Məlumat Oxumaq (Read)
Mövcud məlumatları oxumaq üçün `SELECT` sorğusundan istifadə edirik.

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

affectedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Yenilənən satır sayını əldə edərkən xəta:", err)
}
fmt.Printf("Yenilənən sətir sayı: %d\n", affectedRows)
```

### d) Məlumat Silmək (Delete)
Artıq lazım olmayan məlumatı silmək üçün `DELETE` sorğusundan istifadə edirik.

```go
// Məlumat silmə nümunəsi
deleteQuery := "DELETE FROM users WHERE id = ?"
res, err = db.Exec(deleteQuery, 1)
if err != nil {
	log.Fatal("Silmə əməliyyatı zamanı xəta baş verdi:", err)
}

deletedRows, err := res.RowsAffected()
if err != nil {
	log.Fatal("Silinən sətir sayını əldə edərkən xəta:", err)
}
fmt.Printf("Silinən sətir sayı: %d\n", deletedRows)
```

---

## 4. Hazırlanan İfadələr (Prepared Statements)

Təkrarlanan sorğular üçün hazırlanan ifadələr performansı artırır və SQL enjeksiyasının qarşısını alır. Gəlin, bir neçə məlumatı hazırlanan ifadə ilə necə əlavə edəcəyimizə baxaq.

```go
// Hazırlanan ifadə ilə məlumat əlavə etmək
stmt, err := db.Prepare("INSERT INTO users (name, email) OUTPUT INSERTED.id VALUES (?, ?)")
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
insertQuery := "INSERT INTO users (name, email) VALUES (?, ?)"
_, err = tx.Exec(insertQuery, "Mushfig", "mushfig@example.com")
if err != nil {
	tx.Rollback() // Xəta olduqda bütün əməliyyatları geri alırıq
	log.Fatal("Transaction zamanı xəta (INSERT):", err)
}

// İkinci əməliyyat: məlumat yeniləmək
updateQuery := "UPDATE users SET email = ? WHERE name = ?"
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
- **Bağlantı İdarəetməsi:** İşiniz bitdikdən sonra `defer db.Close()` istifadə edərək açıq əlaqələrin düzgün bağlanmasına diqqət yetirin.
- **Hazırlanan İfadələr:** Təkrarlanan sorğularda hazırlanan ifadələrdən istifadə edərək həm təhlükəsizliyi, həm də performansı artırın.
- **Transaction İstifadəsi:** Əməliyyatların bir-birinə bağlı olduğu hallarda transaction istifadə edin ki, hər hansı bir problem yarandıqda bütün əməliyyatlar geri alınsın.
