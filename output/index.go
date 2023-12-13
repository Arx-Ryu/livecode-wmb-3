package main

import (
	"fmt"
	//"livecode-3-wmb/data"
)

type Meja struct {
	NomorMeja  int
	StatusMeja bool
}

type Menu struct {
	KodeItem int
	NamaItem string
}

type Pesanan struct {
	JumlahItem int
	KodeItem int
}

type Transaksi struct {
	DaftarPesanan []Pesanan
	NamaCustomer string
	Table int
	Paid bool
}

func main() {
	pilihan := 9
	//var me data.InterMenu
	var TableNum int
	var DaftarMeja []Meja
	var DaftarMenu []Menu
	var DaftarTransaksi []Transaksi
	DaftarMeja = TableSetup(DaftarMeja)
	DaftarMenu = MenuSetup(DaftarMenu)
	for pilihan > 0 {
		fmt.Println(`Kasir: 
	1 - Status Meja
	2 - Menu Warung Makan
	3 - Pemesanan
	4 - Pembayaran
	0 - Exit`)
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		switch {
		case pilihan == 1:
			TableStatus(DaftarMeja)			
		case pilihan == 2:
			OpenMenu(DaftarMenu)
		case pilihan == 3:			
			TableNum = NextEmptyTable(DaftarMeja)
			if TableNum == 0 {
				fmt.Println("Maaf, Tidak Ada Meja Kosong")
			} else {
				OpenMenu(DaftarMenu)
				DaftarTransaksi = Pemesanan(DaftarTransaksi, DaftarMenu, TableNum)
				DaftarMeja = MejaDitempati(DaftarMeja, TableNum)
			}
		case pilihan == 4:
			fmt.Print("Table Number: ")
			fmt.Scanln(&TableNum)			
			Pembayaran(DaftarTransaksi, TableNum)
			MejaAvailable(DaftarMeja, TableNum)
		case pilihan == 0:
			fmt.Print("Goodbye")
		default:
			fmt.Printf("Tidak ada Pilihan dengan Nomor %d", pilihan)
		}
	}	
}


func TableSetup(DaftarMeja []Meja) []Meja {
	tempMeja := Meja{}
	for x := 0; x < 30; x++ {
		tempMeja = Meja{NomorMeja: x + 1, StatusMeja: false}
		DaftarMeja = append(DaftarMeja, tempMeja)	
	}
	return DaftarMeja
}

func MenuSetup(DaftarMenu []Menu) []Menu {
	tempMenu := Menu{}
	listItem := []string{"Nasi Goreng", "Nasi dan Telor", "Nasi Ayam Bakar", "Indomie Goreng", "Teh Tawar", "Kopi"}
	for x, item := range listItem {
		tempMenu = Menu{KodeItem: x+1, NamaItem: item}
		DaftarMenu = append(DaftarMenu, tempMenu)
	}
	return DaftarMenu 
}

func TableStatus(DaftarMeja []Meja) {
	EmptyTable := 0
	for _, tables := range DaftarMeja {
		if tables.StatusMeja {
			EmptyTable++
		}
	}
	fmt.Printf("Jumlah Meja Kosong: %d/30\n", EmptyTable)	
	for _, tables := range DaftarMeja {
		if tables.StatusMeja {
			fmt.Printf("Meja %d: Available\n", tables.NomorMeja)
		} else {
			fmt.Printf("Meja %d: Occupied\n", tables.NomorMeja)
		}	
	}
}

func OpenMenu(DaftarMenu []Menu) {
	fmt.Println("MENU:")
	for _, items := range DaftarMenu {
		fmt.Printf("\t%d - %s\n", items.KodeItem, items.NamaItem)
	}
}

func NextEmptyTable(DaftarMeja []Meja) int {
	EmptyTable := 0
	for _, tables := range DaftarMeja {
		if tables.StatusMeja {
			EmptyTable = tables.NomorMeja
			break
		}
	}
	return EmptyTable
}

func Pemesanan(DaftarTransaksi []Transaksi, DaftarMenu []Menu, TableNum int) []Transaksi {
	tempKode := 9
	kodeItemInvalid := true
	var nama string
	var jumlahItem int
	var tempPesanan Pesanan
	var daftarPesanan []Pesanan
	var tempTransaksi Transaksi
	for tempKode > 0 {
		fmt.Print("Kode Item Menu yang ingin dipesan (0 untuk exit): ")
		fmt.Scanln(&tempKode)		
		for _, items := range DaftarMenu {
			if tempKode == items.KodeItem {
				kodeItemInvalid = false
				fmt.Print("Jumlah? (0 untuk cancel) : ")
				fmt.Scanln(&jumlahItem)
				if jumlahItem == 0 {
					break
				}
				tempPesanan = Pesanan{JumlahItem: jumlahItem, KodeItem: items.KodeItem}
				daftarPesanan = append(daftarPesanan, tempPesanan)
				fmt.Printf("%s dengan jumlah %d masuk pada pesanan\n", items.NamaItem, jumlahItem)
			} 
		}
		if kodeItemInvalid && tempKode != 0 {
			fmt.Println("Tidak Ada Kode Item Menu ", tempKode)
		}
	}
	if daftarPesanan != nil {
		fmt.Print("Pesanan Atas Nama? : ")
		fmt.Scanln(&nama)
		tempTransaksi = Transaksi{
			DaftarPesanan: daftarPesanan,
			NamaCustomer: nama,
			Table: TableNum,
			Paid: false,
		}
		//fmt.Println(tempTransaksi)
		DaftarTransaksi = append(DaftarTransaksi, tempTransaksi)
	}	
	return DaftarTransaksi
}

func MejaDitempati(DaftarMeja []Meja, TableNum int) []Meja {
	mejaNotFound := true
	for _, tables := range DaftarMeja {
		if tables.NomorMeja == TableNum {
			mejaNotFound = false
			DaftarMeja[tables.NomorMeja-1] = Meja{NomorMeja: tables.NomorMeja, StatusMeja: false}
			break
		}
	}
	if mejaNotFound {
		fmt.Println("Meja Not Found")
	}
	return DaftarMeja
}

func Pembayaran(DaftarTransaksi []Transaksi, TableNum int) []Transaksi {
	transaksiInvalid := true
	for _, transaksi := range DaftarTransaksi {
		if transaksi.Table == TableNum && !transaksi.Paid {
			transaksiInvalid = false
			transaksi.Paid = true
			fmt.Printf("Transaksi Meja %d berhasil dibayar\n", TableNum)
			break
		}
		
	}
	if transaksiInvalid {
		fmt.Println("Tidak ada transaksi untuk meja tersebut")
	}
	return DaftarTransaksi
}

func MejaAvailable(DaftarMeja []Meja, TableNum int) []Meja {
	for _, tables := range DaftarMeja {
		if tables.NomorMeja == TableNum {
			DaftarMeja[tables.NomorMeja-1] = Meja{NomorMeja: tables.NomorMeja, StatusMeja: true}
			break
		}
	}
	return DaftarMeja
}