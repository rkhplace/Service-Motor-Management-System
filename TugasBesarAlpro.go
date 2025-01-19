package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 1000

// untuk barang yang sering di beli simpan di transaksi

type dataService struct {
	Gudang    sparePart
	Pelanggan pelanggan
}

type sparePart [NMAX]tabSpareP

type tabSpareP struct {
	ID     string
	barang string
	jumlah int
	harga  int
	sold   int
}

type pelanggan [NMAX]identitas

type identitas struct {
	noID       string
	nama       string
	alamat     lokasi
	no_telp    string
	merk       string
	tipe_motor string
	tipe_mesin int
	no_polisi  string
	tanggal    int
	bulan      int
	tahun      int
	ganti      transaksi
	idxG       int
}

type lokasi struct {
	kabKot string
	kec    string
	komDes string
	NoRum  string
}

type transaksi struct {
	barang [100]int
	tarif  int
}

func main() {
	var n, m int
	var choice string
	var tabservice dataService
	n = 0
	m = 0
	for choice != "4" {
		MainMenu()
		fmt.Scan(&choice)
		MenuSelect(choice, &tabservice, &n, &m)
	}
}
func header() {
	fmt.Println("================================================")
	fmt.Println("|                                              |")
	fmt.Println("|                  TUGAS BESAR                 |")
	fmt.Println("|               Irham Kurnia Putra             |")
	fmt.Println("|             Muhammad Rakha Pratama           |")
	fmt.Println("|                 Service Motor                |")
	fmt.Println("|                                              |")
	fmt.Println("================================================")
}
func MainMenu() {
	clearScreen()
	header()
	fmt.Println("1. Service")
	fmt.Println("2. Gudang")
	fmt.Println("3. Transaksi")
	fmt.Println("4. Keluar")
	fmt.Println("------------------------------------------------")
	fmt.Print("masukkan pilihan (1/2/3/4): ")
}
func subMenu() {
	header()
	fmt.Println("1. Input data")
	fmt.Println("2. Ubah data")
	fmt.Println("3. Hapus data")
	fmt.Println("4. Tampilkan data")
	fmt.Println("5. Kembali")
	fmt.Println("------------------------------------------------")
	fmt.Print("masukkan pilihan (1/2/3/4/5): ")
}
func serviceMenu(A *dataService, n *int, m *int) {
	var Iu string
	clearScreen()
	subMenu()
	fmt.Scan(&Iu)
	switch Iu {
	case "1":
		InputDataPelanggan(A, n, *m)
	case "2":
		UpdateDataService(A, n)
	case "3":
		DeleteDataService(A, n)
	case "4":
		ReadDataPelanggan(A, *n, *m)
	case "5":
	default:
		fmt.Println("pilihan tidak tersedia")
	}
}
func MenuSelect(choice string, A *dataService, n *int, m *int) {
	switch choice {
	case "1":
		serviceMenu(A, n, m)
	case "2":
		gudangMenu(A, m)
	case "3":
		transaksiMenu(A, n, m)
	case "4":
	default:
		fmt.Println("pilihan tidak tersedia")
	}
}
func gudangMenu(A *dataService, m *int) {
	var input string
	clearScreen()
	subMenu()
	fmt.Scan(&input)
	switch input {
	case "1":
		CreateDataSparePart(A, m)
	case "2":
		ChangeDataSparePart(A, m)
	case "3":
		DeleteDataSparePart(A, m)
	case "4":
		ReadDataSparePart(*A, *m)
	default:
		fmt.Println("pilihan tidak tersedia")
	}
}
func transaksiMenu(A *dataService, n *int, m *int) {
	var choice string
	clearScreen()
	header()
	fmt.Println("1. ubah data transaksi")
	fmt.Println("2. tampilkan data transaksi")
	fmt.Println("3. tampilkan penjualan barang")
	fmt.Println("------------------------------------------------")
	fmt.Print("masukkan pilihan (1/2/3): ")
	fmt.Scan(&choice)
	switch choice {
	case "1":
		changeDataTransaksi(A, *n, *m)
	case "2":
		readDataTransaksi(A, *n)
	case "3":
		readPenjualanBarang(A, *m)
	}

}
func readDataTransaksi(A *dataService, n int) {
	var tabT dataService = *A
	var i, j int
	var pass int = n - 1
	var temp identitas
	var end string
	clearScreen()
	for i = 0; i < pass; i++ {
		j = i
		temp = tabT.Pelanggan[i+1]
		for j >= 0 && ((tabT.Pelanggan[j].tahun > temp.tahun) || ((tabT.Pelanggan[j].tahun == temp.tahun) && (tabT.Pelanggan[j].bulan > temp.bulan)) || ((tabT.Pelanggan[j].tahun == temp.tahun) && (tabT.Pelanggan[j].bulan == temp.bulan) && (tabT.Pelanggan[j].tanggal > temp.tanggal))) {
			tabT.Pelanggan[j+1] = tabT.Pelanggan[j]
			j--
		}
		tabT.Pelanggan[j+1] = temp
	}
	fmt.Println("berikut hasil transaksi: ")
	fmt.Println("____________________________________")
	for k := 0; k < n; k++ {
		fmt.Println("transaksi oleh: ", tabT.Pelanggan[k].nama)
		fmt.Println("periode: ", tabT.Pelanggan[k].tanggal, "-", tabT.Pelanggan[k].bulan, "-", tabT.Pelanggan[k].tahun)
		fmt.Println("dengan spare-part : ")
		for l := 0; l < tabT.Pelanggan[k].idxG; l++ {
			fmt.Println("-", tabT.Gudang[tabT.Pelanggan[k].ganti.barang[l]].barang)
		}
		fmt.Println("jumlah transaksi: ", tabT.Pelanggan[k].ganti.tarif)
		fmt.Println("____________________________________")
	}
	fmt.Println()
	fmt.Print("(masukan apapun untuk keluar)")
	fmt.Scan(&end)
}
func changeDataTransaksi(A *dataService, n int, m int) {
	var input, find, choice string
	var idx1, idx2 int
	var loop bool = false
	clearScreen()
	fmt.Print("masukan ID pelanggan: ")
	fmt.Scan(&input)
	FindDataService(*A, input, n, &idx1)
	clearScreen()
	fmt.Println("pilih:")
	fmt.Println("1. tambah pergantian spare part")
	fmt.Println("2. hapus pergantian spare part")
	fmt.Scan(&input)
	switch input {
	case "1":
		for !loop {
			clearScreen()
			fmt.Println("List Barang :")
			for i := 0; i < m; i++ {
				fmt.Println(i+1, ".", A.Gudang[i].barang, "      ", A.Gudang[i].ID)
			}
			fmt.Print("masukan spare part yang ingin ditambahkan: ")
			fmt.Scan(&find)
			idx2 = FindDataSparePart(*A, n, find)
			if idx2 == -1 {
				fmt.Println("data tidak ditemukan")
			} else {
				A.Pelanggan[idx1].idxG++
				A.Gudang[idx2].jumlah--
				A.Gudang[idx2].sold++
				A.Pelanggan[idx1].ganti.tarif += A.Gudang[idx2].harga
				fmt.Print("ada tambahan lain? (Y/N): ")
				fmt.Scan(&choice)
				if choice == "N" {
					loop = true
				}
			}
		}
	case "2":
		for !loop {
			clearScreen()
			fmt.Println("list pergantian spare part")
			for i := 0; i < A.Pelanggan[idx1].idxG; i++ {
				fmt.Println(i+1, ".", A.Gudang[A.Pelanggan[idx1].ganti.barang[i]].barang, A.Gudang[A.Pelanggan[idx1].ganti.barang[i]].ID)
			}
			fmt.Println("masukan ID barang yang ingin di hapus: ")
			fmt.Scan(&find)
			idx2 = FindDataSparePart(*A, n, find)
			A.Gudang[idx2].jumlah++
			A.Gudang[idx2].sold--
			A.Pelanggan[idx1].ganti.tarif -= A.Gudang[idx2].harga
			for j := idx2; j < A.Pelanggan[idx1].idxG; j++ {
				A.Pelanggan[idx1].ganti.barang[j] = A.Pelanggan[idx1].ganti.barang[j+1]
			}
			A.Pelanggan[idx1].idxG--
			fmt.Print("hapus data kembali? (Y/N): ")
			fmt.Scan(&choice)
			switch choice {
			case "N", "n":
				loop = true
			case "Y", "y":
			default:
				fmt.Println("tidak ada dalam pilihan")
			}
		}

	}
}
func readPenjualanBarang(A *dataService, m int) {
	var pass int = m - 1
	var tabT dataService = *A
	var max int
	var end string
	var temp tabSpareP
	clearScreen()
	for i := 0; i < pass; i++ {
		max = i
		for j := i; j < m; j++ {
			if tabT.Gudang[j].sold > tabT.Gudang[max].sold {
				max = j
			}
		}
		temp = tabT.Gudang[max]
		tabT.Gudang[max] = tabT.Gudang[i]
		tabT.Gudang[i] = temp
	}
	fmt.Printf("%20s %10s \n", "barang", "terjual")
	for k := 0; k < m; k++ {
		fmt.Printf("%20s %10d \n", tabT.Gudang[k].barang, tabT.Gudang[k].sold)
	}
	fmt.Println()
	fmt.Print("(masukan apapun untuk keluar)")
	fmt.Scan(&end)
}
func CreateDataSparePart(A *dataService, n *int) {
	var x, jum, harga int
	var nama, ID string
	fmt.Print("masukkan jumlah data yang ingin di input: ")
	fmt.Scan(&x)
	clearScreen()
	for i := 0; i < x; i++ {
		fmt.Print("masukkan ID sparepart:")
		fmt.Scan(&ID)
		A.Gudang[*n].ID = ID
		fmt.Print("masukkan nama sparepart: ")
		fmt.Scan(&nama)
		A.Gudang[*n].barang = nama
		fmt.Print("masukkan jumlah sparepart: ")
		fmt.Scan(&jum)
		A.Gudang[*n].jumlah = jum
		fmt.Print("masukkan harga sparepart: ")
		fmt.Scan(&harga)
		A.Gudang[*n].harga = harga
		fmt.Println("Data sudah terinput")
		*n = *n + 1
	}
}
func ChangeDataSparePart(A *dataService, n *int) {
	var x string
	var IDb, Name string
	var idx, Amount int
	clearScreen()
	fmt.Println("masukan ID sparepart: ")
	fmt.Scan(&IDb)
	idx = FindDataSparePart(*A, *n, IDb)
	if idx != -1 {
		clearScreen()
		fmt.Println("pilih data yang ingin di ubah: ")
		fmt.Println("1. sparepart: ")
		fmt.Println("2. jumlah")
		fmt.Println("3. harga")
		fmt.Scan(&x)
		switch x {
		case "1":
			clearScreen()
			fmt.Print("masukkan nama sparepart yang ingin di ganti: ")
			fmt.Scan(&Name)
			A.Gudang[idx].barang = Name
		case "2":
			clearScreen()
			fmt.Print("masukkan jumlah sparepart yang ingin di ganti: ")
			fmt.Scan(&Amount)
			A.Gudang[idx].jumlah = Amount
		case "3":
			clearScreen()
			fmt.Print("masukkan harga sparepart yang ingin di ganti: ")
			fmt.Scan(&Amount)
			A.Gudang[idx].harga = Amount
		default:
			fmt.Println("pilihan tidak tersedia")
		}
	} else {
		fmt.Println("data tidak ditemukan")
	}
}
func DeleteDataSparePart(A *dataService, n *int) {
	var IDb string
	var idx int
	clearScreen()
	fmt.Println("masukkan ID barang yang ingin di hapus: ")
	fmt.Scan(&IDb)
	idx = FindDataSparePart(*A, *n, IDb)
	for i := idx; i < *n; i++ {
		A.Gudang[i] = A.Gudang[i+1]
	}
	*n--
}
func ReadDataSparePart(A dataService, n int) {
	var tempS string
	clearScreen()
	for i := 0; i < n; i++ {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Print("ID barang: ")
		fmt.Println(A.Gudang[i].ID)
		fmt.Print("Nama Barang: ")
		fmt.Println(A.Gudang[i].barang)
		fmt.Print("Jumlah Barang: ")
		fmt.Println(A.Gudang[i].jumlah)
		fmt.Print("Harga Barang: ")
		fmt.Println(A.Gudang[i].harga)
		fmt.Print("Terjual: ")
		fmt.Println(A.Gudang[i].sold)

	}
	fmt.Println()
	fmt.Println("(masukan apapun untuk keluar)")
	fmt.Scan(&tempS)
}
func InputDataPelanggan(A *dataService, n *int, m int) {
	var x, tempI int
	var tempS string
	clearScreen()
	fmt.Print("masukkan jumlah data yang ingin di input: ")
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		clearScreen()
		fmt.Print("masukkan ID pelanggan: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].noID = tempS
		fmt.Print("masukkan nama pelanggan: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].nama = tempS
		clearScreen()
		fmt.Println("masukkan alamat pelanggan ")
		fmt.Print("masukkan kabupaten/kota: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].alamat.kabKot = tempS
		fmt.Print("masukkan kecamatan: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].alamat.kec = tempS
		fmt.Print("masukkan komplek/desa: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].alamat.komDes = tempS
		fmt.Print("masukkan nomor rumah: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].alamat.NoRum = tempS
		clearScreen()
		fmt.Print("masukkan nomor telepon pelanggan: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].no_telp = tempS
		clearScreen()
		fmt.Print("masukkan merk/tipe motor: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].merk = tempS
		clearScreen()
		fmt.Println("jenis motor: ")
		fmt.Println("1. bebek")
		fmt.Println("2. matic")
		fmt.Println("3. sport")
		fmt.Print("pilih jenis motor: ")
		fmt.Scan(&tempS)
		switch tempS {
		case "1":
			A.Pelanggan[i].tipe_motor = "bebek"
		case "2":
			A.Pelanggan[i].tipe_motor = "matic"
		case "3":
			A.Pelanggan[i].tipe_motor = "sport"
		default:
			fmt.Println("pilihan tidak tersedia")
		}
		clearScreen()
		fmt.Println("tipe mesin: ")
		fmt.Println("1. 125")
		fmt.Println("2. 150")
		fmt.Println("3. 200")
		fmt.Println("4. 250")
		fmt.Println("5.lainnya")
		fmt.Println("pilih tipe mesin")
		fmt.Scan(&tempS)
		switch tempS {
		case "1":
			A.Pelanggan[i].tipe_mesin = 125
		case "2":
			A.Pelanggan[i].tipe_mesin = 150
		case "3":
			A.Pelanggan[i].tipe_mesin = 200
		case "4":
			A.Pelanggan[i].tipe_mesin = 250
		case "5":
			fmt.Print("masukkan tipe mesin: ")
			fmt.Scan(&tempI)
			A.Pelanggan[i].tipe_mesin = tempI
		default:
			fmt.Println("pilihan tidak tersedia")
		}
		A.Pelanggan[i].tipe_mesin = tempI
		clearScreen()
		fmt.Print("masukkan nomor polisi: ")
		fmt.Scan(&tempS)
		A.Pelanggan[i].no_polisi = tempS
		clearScreen()
		fmt.Println("masukkan tanggal sevice")
		fmt.Print("tanggal: ")
		fmt.Scan(&tempI)
		A.Pelanggan[i].tanggal = tempI
		fmt.Print("bulan(dengan angka): ")
		fmt.Scan(&tempI)
		A.Pelanggan[i].bulan = tempI
		fmt.Print("tahun: ")
		fmt.Scan(&tempI)
		A.Pelanggan[i].tahun = tempI
		sparePartsChange(A, i, m)
		*n = *n + 1
	}
}
func sparePartsChange(A *dataService, n int, m int) {
	var SP, choice string
	var idx int
	var loop bool = false
	A.Pelanggan[n].ganti.tarif += ServicePrice(A, n)
	for !loop {
		clearScreen()
		fmt.Println("====================================================================")
		fmt.Println()
		fmt.Println("                            List Barang                             ")
		fmt.Println()
		fmt.Println("====================================================================")
		for i := 0; i < m; i++ {
			fmt.Printf("%6d %30s %20s \n", i+1, A.Gudang[i].barang, A.Gudang[i].ID)
		}
		fmt.Print("masukkan spare-parts yang akan di ganti(gunakan ID barang):")
		fmt.Scan(&SP)
		idx = FindDataSparePart(*A, m, SP)
		if idx == -1 {
			fmt.Println("data tidak ditemukan")
		} else {
			A.Pelanggan[n].ganti.barang[A.Pelanggan[n].idxG] = idx
			A.Pelanggan[n].idxG++
			A.Gudang[idx].jumlah--
			A.Gudang[idx].sold++
			A.Pelanggan[n].ganti.tarif += A.Gudang[idx].harga
			fmt.Print("ada tambahan lain? (Y/N): ")
			fmt.Scan(&choice)
			switch choice {
			case "Y", "y":
				loop = false
			case "N", "n":
				loop = true
			default:
				fmt.Println("input salah")
			}
		}
	}
	fmt.Println("------------------------------------")
	fmt.Println("tarif: ", A.Pelanggan[n].ganti.tarif)
	fmt.Println()
	fmt.Print("(masukan apapun untuk keluar): ")
	fmt.Scan(&choice)
}
func FindDataSparePart(A dataService, n int, x string) int {
	var hasil bool = false
	idx := -1
	i := 0
	for i < n && !hasil {
		if A.Gudang[i].ID == x {
			hasil = true
			idx = i
		} else {
			i++
		}
	}
	return idx
}
func ServicePrice(A *dataService, n int) int {
	var X int
	switch A.Pelanggan[n].tipe_motor {
	case "bebek":
		X = 60000
	case "matic":
		switch {
		case A.Pelanggan[n].tipe_mesin == 125:
			X = 75000
		case A.Pelanggan[n].tipe_mesin == 150:
			X = 90000
		case A.Pelanggan[n].tipe_mesin > 150:
			X = 100000
		default:
			X = 0
		}
	case "sport":
		switch {
		case A.Pelanggan[n].tipe_mesin == 150:
			X = 100000
		case A.Pelanggan[n].tipe_mesin == 250:
			X = 150000
		case A.Pelanggan[n].tipe_mesin > 250:
			X = 200000
		}
	default:
		X = 0
	}
	return X
}
func FindDataService(A dataService, x string, n int, idx *int) {
	var kir, kan, ten int
	kir = 0
	kan = n - 1
	*idx = -1
	for i := 0; i < n; i++ {
		ten = (kir + kan) / 2
		if A.Pelanggan[ten].noID == x {
			*idx = ten
		} else if A.Pelanggan[ten].noID > x {
			kan = ten - 1
		} else if A.Pelanggan[ten].noID < x {
			kir = ten + 1
		}
	}
}
func UpdateDataService(A *dataService, n *int) {
	var search, tempS string
	var idx, tempI int
	clearScreen()
	fmt.Print("masukan ID pelanggan: ")
	fmt.Scan(&search)
	FindDataService(*A, search, *n, &idx)
	if idx == -1 {
		fmt.Println("data tidak ditemukan")
	} else {
		clearScreen()
		fmt.Println("pilih data yang ingin di ubah: ")
		fmt.Println("1. nama")
		fmt.Println("2. alamat")
		fmt.Println("3. nomor telepon")
		fmt.Println("4. merk motor")
		fmt.Println("5. tipe motor")
		fmt.Println("6. tipe mesin")
		fmt.Println("7. nomor polisi")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("masukan nama: ")
			fmt.Scan(&tempS)
			A.Pelanggan[idx].nama = tempS
		case 2:
			var choice2 int
			fmt.Print("pilih alamat yang ingin di ganti: ")
			fmt.Print("1. kabupaten/kota")
			fmt.Print("2. kecamatan")
			fmt.Print("3. komplek/desa")
			fmt.Print("4. nomor rumah")
			fmt.Scan(&choice2)
			switch choice2 {
			case 1:
				clearScreen()
				fmt.Print("masukan kabupaten/kota: ")
				fmt.Scan(&tempS)
				A.Pelanggan[idx].alamat.kabKot = tempS
			case 2:
				clearScreen()
				fmt.Print("masukan kecamatan: ")
				fmt.Scan(&tempS)
				A.Pelanggan[idx].alamat.kec = tempS
			case 3:
				clearScreen()
				fmt.Print("masukan komplek/desa: ")
				fmt.Scan(&tempS)
				A.Pelanggan[idx].alamat.komDes = tempS
			case 4:
				clearScreen()
				fmt.Print("masukan nomor rumah: ")
				fmt.Scan(&tempS)
				A.Pelanggan[idx].alamat.NoRum = tempS
			}
		case 3:
			fmt.Print("masukan nomor telepon: ")
			fmt.Scan(&tempS)
			A.Pelanggan[idx].no_telp = tempS
		case 4:
			fmt.Print("masukan merk motor: ")
			fmt.Scan(&tempS)
			A.Pelanggan[idx].merk = tempS
		case 5:
			fmt.Print("masukan tipe motor: ")
			fmt.Scan(&tempS)
			A.Pelanggan[idx].tipe_motor = tempS
		case 6:
			fmt.Print("masukan tipe mesin: ")
			fmt.Scan(&tempI)
			A.Pelanggan[idx].tipe_mesin = tempI
		case 7:
			fmt.Print("masukan nomor polisi: ")
			fmt.Scan(&tempS)
			A.Pelanggan[idx].no_polisi = tempS
		}
	}
}
func DeleteDataService(A *dataService, n *int) {
	var idx int
	var del string
	clearScreen()
	fmt.Print("masukan ID pelanggan yang ingin di hapus: ")
	fmt.Scan(&del)
	FindDataService(*A, del, *n, &idx)
	for i := idx; i < *n; i++ {
		A.Pelanggan[i] = A.Pelanggan[i+1]
	}
	*n--
}
func ReadDataPelanggan(A *dataService, n, m int) {
	var choice, tempI, x, y, z, i int
	var tempS string
	var tabDate1, tabDate2, tabDate3, tabSlice [NMAX]int
	clearScreen()
	fmt.Println("cari berdasarkan: ")
	fmt.Println("1. tanggal")
	fmt.Println("2. bulan")
	fmt.Println("3. tahun")
	fmt.Println("4. spare-part")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		clearScreen()
		fmt.Println("masukan tahun: ")
		fmt.Scan(&tempI)
		date3(*A, &tabDate3, n, tempI, &x)
		fmt.Println("masukan bulan: ")
		fmt.Scan(&tempI)
		date2(*A, &tabDate2, n, tempI, &y)
		fmt.Println("masukan tanggal: ")
		fmt.Scan(&tempI)
		clearScreen()
		date1(*A, &tabDate1, n, tempI, &z)
		sliceArr1(tabDate1, tabDate2, tabDate3, &tabSlice, x, y, z, &i)
		ReadData(*A, tabSlice, i)
		fmt.Println("(masukan apapun untuk keluar)")
		fmt.Scan(&tempS)
	case 2:
		clearScreen()
		fmt.Println("masukan tahun: ")
		fmt.Scan(&tempI)
		date3(*A, &tabDate3, n, tempI, &x)
		fmt.Println("masukan bulan: ")
		fmt.Scan(&tempS)
		clearScreen()
		date2(*A, &tabDate2, n, tempI, &y)
		sliceArr2(tabDate2, tabDate3, &tabSlice, x, y, &i)
		ReadData(*A, tabSlice, i)
		fmt.Println("(masukan apapun untuk keluar)")
		fmt.Scan(&tempS)
	case 3:
		fmt.Println("masukan tahun: ")
		fmt.Scan(&tempI)
		clearScreen()
		date3(*A, &tabDate3, n, tempI, &x)
		ReadData(*A, tabDate3, x)
		fmt.Println("(masukan apapun untuk keluar)")
		fmt.Scan(&tempS)

	case 4:
		ReadDataBySparePart(*A, n, m)
	}
}
func ReadDataBySparePart(A dataService, n, m int) {
	var tempS string
	var tempI int
	clearScreen()
	fmt.Println("berikut data barang yang tercantum pada gudang: ")
	for i := 0; i < m; i++ {
		fmt.Println(i+1, ".", A.Gudang[i].barang, "      ", A.Gudang[i].ID)
	}
	fmt.Println("masukan nama spare-part(gunakan ID barang): ")
	fmt.Scan(&tempS)
	clearScreen()
	tempI = FindDataSparePart(A, m, tempS)
	if tempI == -1 {
		fmt.Println("data tidak ditemukan")
	} else {
		fmt.Println("berikut nama pembeli spare-part: ")
		for j := 0; j < n; j++ {
			for k := 0; k < A.Pelanggan[j].idxG; k++ {
				if A.Gudang[tempI].barang == A.Gudang[A.Pelanggan[j].ganti.barang[k]].barang {
					fmt.Println(j+1, ".", A.Pelanggan[j].nama, "pada periode", A.Pelanggan[j].tanggal, A.Pelanggan[j].bulan, A.Pelanggan[j].tahun)
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("masukan huruf atau angka apapun untuk keluar")
	fmt.Scan(&tempS)
}
func date1(A dataService, date *[NMAX]int, n, temp int, x *int) {
	for i := 0; i < n; i++ {
		if A.Pelanggan[i].tanggal == temp {
			date[*x] = i
			*x++
		}
	}
}
func date2(A dataService, date *[NMAX]int, n int, temp int, x *int) {
	for i := 0; i < n; i++ {
		if A.Pelanggan[i].bulan == temp {
			date[*x] = i
			*x++
		}
	}
}
func date3(A dataService, date *[NMAX]int, n, temp int, x *int) {
	for i := 0; i < n; i++ {
		if A.Pelanggan[i].tahun == temp {
			date[*x] = i
			*x++
		}
	}
}
func sliceArr1(date1, date2, date3 [NMAX]int, slice *[NMAX]int, x, y, z int, n *int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			for k := 0; k < z; k++ {
				if date1[i] == date2[j] && date2[j] == date3[k] {
					slice[i] = date1[i]
					*n++
				}
			}
		}
	}
}
func sliceArr2(date1, date2 [NMAX]int, slice *[NMAX]int, x, y int, n *int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if date1[i] == date2[j] {
				slice[i] = date1[i]
				*n++
			}
		}
	}
}
func ReadData(A dataService, date [NMAX]int, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("nama: ", A.Pelanggan[date[i]].nama)
		fmt.Println("alamat: ", A.Pelanggan[date[i]].alamat.kabKot, A.Pelanggan[date[i]].alamat.kec, A.Pelanggan[date[i]].alamat.komDes, A.Pelanggan[date[i]].alamat.NoRum)
		fmt.Println("nomor telepon: ", A.Pelanggan[date[i]].no_telp)
		fmt.Println("merk motor: ", A.Pelanggan[date[i]].merk)
		fmt.Println("tipe motor: ", A.Pelanggan[date[i]].tipe_motor)
		fmt.Println("tipe mesin: ", A.Pelanggan[date[i]].tipe_mesin)
		fmt.Println("nomor polisi: ", A.Pelanggan[date[i]].no_polisi)
		fmt.Println("periode waktu:", A.Pelanggan[date[i]].tanggal, A.Pelanggan[date[i]].bulan)
		fmt.Println("spare-part yang diganti: ")
		for j := 0; j < A.Pelanggan[date[i]].idxG; j++ {
			fmt.Println(j+1, ".", A.Gudang[A.Pelanggan[date[i]].ganti.barang[j]].barang)
		}

	}
}
func clearScreen() {
	/* IS: -
	   FS: Mengosongkan layar.
	*/
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Platform tidak didukung.")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
