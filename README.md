# GORUNTINE

goruntine mirip dengan thread, tapi sebenarnya bukan.
Goruntine sangat ringan, hanya dibutuhkan memori sekitar 2kb memori saja untuk satu buah goruntine.
Eksekusi goruntine ini bersifat asynchronous, menjadikan prosesnya tidak saling tunggu satu sama lain.

Untuk menerapkan goruntine, proses yang akan dieksekusi sebagai goruntine harus dibungkus ke dalam sebuah function.
Pada saat pemanggilan function tersebut, ditambahkan keyword `go` didepannya.

contoh :

```go
package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 1; i < till; i++ {
		fmt.Println(i, message)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	go print(5, "Hello")
	print(5, "Apa Kabar")

	var input string
	fmt.Scanln(&input)
}
```

pada kode di atas, function `runtime.GOMAXPROCS(4)` digunakan untuk menentukan jumplah core yang akan digunakan untuk eksekusi program, jika tidak dituliskan maka secara default akan menggunakan jumplah seluruh core yang ada di komputer.

function `fmt.Scanln()` mengakibatkan proses jalannya aplikasi berhenti di baris itu (**blocking**) hingga user menekan enter.
Hal ini dilakukan karena kemungkinan waktu selesainya eksekusi goruntine `print()` lebih lama dibanding dengan eksekusinya goruntine utama (`main()`).
Karena keduanya berjalan secara pararel / asynchronous, jika itu terjadi, kemungkinan goruntine yang belum selesai dipaksa berhenti karena goruntine utama sudah selesai dijalankan.