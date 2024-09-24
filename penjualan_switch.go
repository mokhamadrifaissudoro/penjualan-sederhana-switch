package main

import "fmt"

func main() {

	// deklarasi variabel
	var nama_customer, nama_barang string
	var harga float32 = 2500
	var quantity int32
	var hasil_discount, total_harga float32

	// input nama customer
	fmt.Print("Input nama customer: ")
	fmt.Scanf("%s\n", &nama_customer)

	// input nama barang
	fmt.Print("Input nama barang: ")
	fmt.Scanf("%s\n", &nama_barang)

	// input quantity barang
	fmt.Print("Input quantity barang: ")
	fmt.Scanf("%d\n", &quantity)

	// kondisi diskon, kalau lebih dari 5 dapat 10%, selain itu 2%
	switch {
	case quantity > 5:
		hasil_discount = 0.1 // 10%
	default:
		hasil_discount = 0.02 // 2%
	}

	// proses
	sub_total := float32(quantity) * harga
	total_harga = sub_total - (hasil_discount * sub_total)

	// tampilkan hasil harga
	fmt.Println("hasil_discount: ", hasil_discount)
	fmt.Println("quantity: ", quantity)
	fmt.Println("Total Harga adalah: ", total_harga)

}
