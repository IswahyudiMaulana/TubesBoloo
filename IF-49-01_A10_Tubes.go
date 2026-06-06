package main
import "fmt"

const NMAX = 100
type Peserta struct {
	namaPeserta, namaKursus string
	id, WaktuDaftar       int
}
type arrayPeserta [NMAX]Peserta

func main() {
	var peserta arrayPeserta
	var n int

	DummyData(&peserta, &n)
	Menu(&peserta, &n)
}

func Menu(Peserta *arrayPeserta, n *int) {
	/*	I.S : Data peserta dan jumlah data peserta telah terdefinisi
   		F.S : Pengguna dapat memilih dan menjalankan menu aplikasi(Entry Data, Mengubah Data, Menghapus Data, Mencari Data, Menampilkan Data, Sorting Data), atau keluar dari program
	*/
	var pilihan, id, idx int
	var berjalan bool
	berjalan = true

	for berjalan {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n%15s────────────────────────────────────────────────────────────────────────────────────────\n", "")
		fmt.Printf("%30s      	PROGRAM PENDAFTARAN PESERTA KURSUS\n", "")
		fmt.Printf("%15s────────────────────────────────────────────────────────────────────────────────────────\n", "")
		fmt.Printf("%15s 1. Entry Data\n", "")
		fmt.Printf("%15s 2. Mengubah Data\n", "")
		fmt.Printf("%15s 3. Menghapus Data\n", "")
		fmt.Printf("%15s 4. Mencari Data\n", "")
		fmt.Printf("%15s 5. Tampilkan Data\n", "")
		fmt.Printf("%15s 6. Sorting Data\n", "")
		fmt.Printf("%15s 7. Keluar\n", "")
		fmt.Printf("%15s Pilih menu : ", "")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			EntryData(Peserta, n)
		case 2:
			fmt.Printf("%15sMasukkan ID peserta yang ingin diubah : ", "")
			fmt.Scan(&id)
			MengubahData(Peserta, *n, id)
			MenampilkanData(*Peserta, *n)

		case 3:
			fmt.Printf("%15sMasukkan ID peserta yang ingin dihapus : ", "")
			fmt.Scan(&id)
			MenghapusData(Peserta, n, id)
			MenampilkanData(*Peserta, *n)

		case 4:
			idx = MencariData(*Peserta, *n)

			if idx == -1 {
				fmt.Printf("%15s──────────────────────────────────────── DATA TIDAK DITEMUKAN ──────────────────────────────────────── \n", "")
			} 
		case 5:
			if *n == 0 {
				fmt.Printf("%15s──────────────────────────────────────── BELUM ADA DATA PESERTA ──────────────────────────────────────── \n", "")
			} else {
				MenampilkanData(*Peserta, *n)
			}
			MenampilkanStatistik(*Peserta, *n)
		case 6:
			var metodeSort, pilihSort,pilih int

			fmt.Printf("%15s──────────────────────────────────────── SORTING DATA ──────────────────────────────────────── \n", "")
			fmt.Printf("%15s1. Selection Sort\n", "")
			fmt.Printf("%15s2. Insertion Sort\n", "")
			fmt.Printf("%15sMetode : ", "")
			fmt.Scan(&metodeSort)

			fmt.Printf("%15s\n", "")
			fmt.Printf("%15s1. ID\n", "")
			fmt.Printf("%15s2. Nama\n", "")
			fmt.Printf("%15s3. Waktu (DD-MM-YYYY)\n", "")
			fmt.Printf("%15sPilihan : ", "")
			fmt.Scan(&pilihSort)


			if metodeSort == 1 {

				switch pilihSort {
				case 1:
					fmt.Printf("\n\n%15s1. Membesar [↑]\n","")
					fmt.Printf("%15s2. Mengecil [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)
					
					if pilih == 1 {
						SelectionSortID(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih == 2{
						SelectionSortID(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}

				case 2:
					fmt.Printf("\n\n%15s1. Membesar (A-Z) [↑]\n","")
					fmt.Printf("%15s2. Mengecil (Z-A) [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)

					if pilih == 1{
						SelectionSortNama(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih ==2{
						SelectionSortNama(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}
					
				case 3:
					fmt.Printf("\n\n%15s1. Membesar [↑]\n","")
					fmt.Printf("%15s2. Mengecil [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)

					if pilih == 1 {
						SelectionSortWaktu(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih ==2 {
						SelectionSortWaktu(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}
				}
			} else {
				switch pilihSort {
				case 1:
					fmt.Printf("\n\n%15s1. Membesar [↑]\n","")
					fmt.Printf("%15s2. Mengecil [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)
					
					if pilih == 1 {
						InsertionSortID(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih ==2 {
						InsertionSortID(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}
					
				case 2:
					fmt.Printf("\n\n%15s1. Membesar (A-Z) [↑]\n","")
					fmt.Printf("%15s2. Mengecil (Z-A) [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)

					if pilih == 1 {
						InsertionSortNama(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih ==2{
						InsertionSortNama(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}
					
				case 3:
					fmt.Printf("\n\n%15s1. Membesar [↑]\n","")
					fmt.Printf("%15s2. Mengecil [↓]\n","")
					fmt.Printf("%15sPilihan : ","")
					fmt.Scan(&pilih)

					if pilih == 1 {
						InsertionSortWaktu(Peserta, *n,true)
						MenampilkanData(*Peserta, *n)
					}else if pilih == 2 {
						InsertionSortWaktu(Peserta, *n,false)
						MenampilkanData(*Peserta, *n)
					}else {
						fmt.Printf("%15sPilihan tidak valid\n","")
					}
				}
			}
		case 7:
			fmt.Printf("%15s──────────────────────────────────── TERIMA KASIH ────────────────────────────────────\n\n\n\n\n\n\n", "")
			berjalan = false
		default:
			fmt.Printf("%15sPilihan menu tidak valid\n", "")
		}
		fmt.Println()
	}
}

func EntryData(peserta *arrayPeserta, n *int) {
	/* I.S : peserta belum terisi, n belum terdefinisi
	   F.S : Array peserta terisi dengan data peserta yang dimasukkan */
	var inputID, inputWaktu int
	var lanjut bool
	lanjut = true

	fmt.Printf("%15s──────────────────────── Masukkan Data Peserta ────────────────────────\n", "")
	for *n < NMAX-1 && lanjut {
		fmt.Printf("\n%15sPeserta ke %d\n\n", "", *n+1)
		fmt.Printf("%15sID Peserta : ", "")
		fmt.Scan(&inputID)

		if inputID == 0  {
			lanjut = false
		} else {
			fmt.Printf("%15sWaktu Daftar (YYYYMMDD) : ", "")
			fmt.Scan(&inputWaktu)

			(*peserta)[*n].id = inputID
			(*peserta)[*n].WaktuDaftar = inputWaktu
			fmt.Printf("%15sNama Peserta : ", "")
			fmt.Scan(&peserta[*n].namaPeserta)
			fmt.Printf("%15sNama Kursus : ", "")
			fmt.Scan(&peserta[*n].namaKursus)
			*n = *n + 1
		}
	}
}

func MengubahData(peserta *arrayPeserta, n, id int) {
	/* I.S : peserta terdefinisi, n terdefinisi, id terdefinisi
	   F.S : Data peserta dengan id yang sesuai telah diubah */
	var i int
	i = 0
	for i < n {
		if peserta[i].id == id {
			fmt.Printf("%15sNama Peserta : ", "")
			fmt.Scan(&peserta[i].namaPeserta)
			fmt.Printf("%15sNama Kursus : ", "")
			fmt.Scan(&peserta[i].namaKursus)
			fmt.Printf("%15sWaktu Daftar (YYYYMMDD) : ", "")
			fmt.Scan(&peserta[i].WaktuDaftar)
		}
		i = i + 1
	}
}

func MenghapusData(peserta *arrayPeserta, n *int, id int) {
	/* I.S : peserta terdefinisi, n terdefinisi, id terdefinisi
	   F.S : Data peserta terhapus dan n berkurang 1 jika ditemukan */
	var i, j, idx int
	var ketemu bool

	i = 0
	idx = -1
	ketemu = false

	for i < *n && !ketemu {
		if peserta[i].id == id {
			idx = i
			ketemu = true
		}
		i = i + 1
	}

	if ketemu {
		j = idx
		for j < *n-1 {
			peserta[j] = peserta[j+1]
			j = j + 1
		}
		*n = *n - 1
		fmt.Printf("%15sData berhasil dihapus!\n", "")
	} else {
		fmt.Printf("%15sData dengan ID tersebut tidak ditemukan.\n", "")
	}
}

func MencariData(Peserta arrayPeserta, n int) int {
	/* I.S : array Peserta berisi n data peserta
	   F.S : mengembalikan indeks data yang dicari dan menampilkan data jika ditemukan, atau mengembalikan -1 jika tidak ditemukan
	*/
	var kategori int
	var idx, idCari int
	var cari string

	fmt.Printf("%15s──────────────────────── PENCARIAN DATA ────────────────────────\n", "")
	fmt.Printf("%15sCari berdasarkan:\n", "")
	fmt.Printf("%15s1. ID\n", "")
	fmt.Printf("%15s2. Nama Peserta\n", "")
	fmt.Printf("%15s3. Nama Kursus\n", "")
	fmt.Printf("%15s4. Waktu Daftar (YYMMDD)\n", "")
	fmt.Printf("%15sPilihan : ", "")
	fmt.Scan(&kategori)

	switch kategori {
	case 1:
		fmt.Printf("%15sMasukkan ID yang ingin dicari : ", "")
		fmt.Scan(&idCari)
		idx = SequentialSearchInt(Peserta, n, idCari, "id")

		if idx !=-1 {
			Peserta[0] = Peserta[idx]
			MenampilkanData(Peserta,1)
		}
	case 2:
		fmt.Printf("%15sMasukkan nama peserta yang ingin dicari : ", "")
		fmt.Scan(&cari)
		InsertionSortNama(&Peserta, n,true)
		idx = BinarySearchString(Peserta, n, cari, "peserta")
		
		if idx !=-1{
			Peserta[0] = Peserta[idx]
			MenampilkanData(Peserta,1)
		}
	case 3:
		fmt.Printf("%15sMasukkan nama kursus yang ingin dicari : ", "")
		fmt.Scan(&cari)
		idx = SequentialSearchString(Peserta, n, "kursus", cari)

		if idx !=-1{
			MenampilkanKursus( Peserta, n, cari)
		}
	case 4:
		fmt.Printf("%15sMasukkan waktu daftar yang ingin dicari (YYYYMMDD) : ", "")
		fmt.Scan(&idCari)
		SelectionSortWaktu(&Peserta, n,true)
		idx = BinarySearchInt(Peserta, n, idCari, "Waktu")

		if idx !=-1{
			MenampilkanWaktu(Peserta, n, idCari)
		}
	default : 
		fmt.Printf("%15s Pilihan Tidak Valid","")
	}
	return idx
}

func MenampilkanData(Peserta arrayPeserta, n int) {
	/* I.S : Peserta terdefinisi, n terdefinisi
	   F.S : Data peserta ditampilkan ke layar */
	var i, hari, bulan, tahun, tanggal int
	fmt.Printf("\n%15s┌────┬────────────┬─────────────────────────┬──────────────────────────────┬──────────────┐\n", "")
	fmt.Printf("%15s│ %-2s │ %-10s │ %-23s │ %-28s │ %-12s │\n", "", "No", "ID", "Nama Peserta", "Nama Kursus", "Tanggal")
	fmt.Printf("%15s├────┼────────────┼─────────────────────────┼──────────────────────────────┼──────────────┤\n", "")

	for i = 0; i < n; i++ {
		tanggal = Peserta[i].WaktuDaftar

		hari = tanggal % 100
		bulan = (tanggal % 10000) / 100
		tahun = tanggal / 10000

		fmt.Printf("%15s│ %-2d │ %-10d │ %-23s │ %-28s │ %02d-%02d-%04d   │\n","",i+1,Peserta[i].id,Peserta[i].namaPeserta,Peserta[i].namaKursus,hari, bulan, tahun)
	}
	fmt.Printf("%15s└────┴────────────┴─────────────────────────┴──────────────────────────────┴──────────────┘\n", "")
}

func MenampilkanKursus(Peserta arrayPeserta, n int, kursus string) {
	/* 	 I.S : Array Peserta berisi n data peserta dan nama kursus yang dicari terdefinisi
  		 F.S : Seluruh data peserta yang mengikuti kursus tersebut ditampilkan ke layar
	 */
	var hasil arrayPeserta
	var jumlah, i int
	jumlah = 0

	for i = 0; i < n; i++ {
		if Peserta[i].namaKursus == kursus {
			hasil[jumlah] = Peserta[i]
			jumlah = jumlah + 1
		}
	}

	if jumlah > 0 {
		MenampilkanData(hasil, jumlah)
	}
}

func MenampilkanWaktu(Peserta arrayPeserta, n,waktu int){
	/* 	 I.S : Array Peserta berisi n data peserta dan tanggal yang dicari terdefinisi
  		 F.S : Seluruh data peserta dengan tanggal pendaftaran yang sesuai ditampilkan ke layar
	*/
	var hasil arrayPeserta
	var jumlah, i int
	jumlah = 0 
	for i = 0; i < n; i++ {
		if Peserta[i].WaktuDaftar == waktu {
			hasil[jumlah] = Peserta[i]
			jumlah = jumlah + 1
		}
	}
	if jumlah > 0{
		MenampilkanData(hasil,jumlah)
	}
}

func SequentialSearchString(Peserta arrayPeserta, n int, pilihan, cari string) int {
	/* I.S : Data peserta tersimpan dalam array arrayPeserta dan jumlah n terdefinisi
	   F.S : mengembalikan index data peserta yang dicari menggunakan Sequential Search
	*/
	var i, idx int
	idx = -1

	for i = 0; i < n; i++ {
		if pilihan == "peserta" {
			if Peserta[i].namaPeserta == cari {
				idx = i
			}
		} else if pilihan == "kursus" {
			if Peserta[i].namaKursus == cari {
				idx = i
			}
		}
	}
	return idx
}

func SequentialSearchInt(Peserta arrayPeserta, n int, cari int, pilihan string) int { //ini ada tanggalnya
	/* I.S : Data peserta tersimpan dalam array arrayPeserta dan jumlah n terdefinisi
	   F.S : mengembalikan index data peserta yang dicari menggunakan Sequential Search berdasarkan ID
	*/
	var i, idx int
	idx = -1

	for i = 0; i < n; i++ {
		if pilihan == "id" {
			if Peserta[i].id == cari {
				idx = i
			}
		} else if pilihan == "Waktu" {
			if Peserta[i].WaktuDaftar == cari {
				idx = i
			}
		}
	}
	return idx
}

func BinarySearchString(Peserta arrayPeserta, n int, cari, pilihan string) int {
	/* I.S : Data peserta tersimpan dalam array arrayPeserta dan jumlah n terdefinisi
	   F.S : mengembalikan index data peserta yang dicari menggunakan Binary Search
	*/
	var idx, left, right, mid int
	idx = -1
	left = 0
	right = n - 1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if pilihan == "peserta" {
			if cari < Peserta[mid].namaPeserta {
				right = mid - 1
			} else if cari > Peserta[mid].namaPeserta {
				left = mid + 1
			} else {
				idx = mid
			}
		} else if pilihan == "kursus" {
			if cari < Peserta[mid].namaKursus {
				right = mid - 1
			} else if cari > Peserta[mid].namaKursus {
				left = mid + 1
			} else {
				idx = mid
			}
		}
	}
	return idx
}

func BinarySearchInt(Peserta arrayPeserta, n int, cari int, pilihan string) int { // ini ada tanggalnya
	/* I.S : Data peserta tersimpan dalam array arrayPeserta dan jumlah n terdefinisi, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : mengembalikan index data peserta yang dicari menggunakan Binary Search berdasarkan ID
	*/
	var idx, left, right, mid int
	idx = -1
	left = 0
	right = n - 1

	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if pilihan == "id" {
			if cari < Peserta[mid].id {
				right = mid - 1
			} else if cari > Peserta[mid].id {
				left = mid + 1
			} else {
				idx = mid
			}
		} else if pilihan == "Waktu" {
			if cari < Peserta[mid].WaktuDaftar {
				right = mid - 1
			} else if cari > Peserta[mid].WaktuDaftar {
				left = mid + 1
			} else {
				idx = mid
			}
		}
	}
	return idx
}

func InsertionSortID(data *arrayPeserta, n int, isAscend bool) {
	/* I.S : terdefinisi array data yang berisi n data peserta, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : array data terurut berdasarkan ID menggunakan Insertion Sort. Jika isAscend = true maka data terurut menaik (ascending), jika isAscend = false maka data terurut menurun (descending)
	*/
	var pass, i int
	var temp Peserta
	pass = 1

	if isAscend {
		for pass <= n-1 {
			i = pass
			temp = (*data)[pass]
			for i > 0 && temp.id < (*data)[i-1].id {
				(*data)[i] = (*data)[i-1]
				i = i - 1
			}
			(*data)[i] = temp
			pass = pass + 1
		}
	} else {
		for pass <= n-1 {
			i = pass
			temp = (*data)[pass]
			for i > 0 && temp.id > (*data)[i-1].id {
				(*data)[i] = (*data)[i-1]
				i = i - 1
			}
			(*data)[i] = temp
			pass = pass + 1
		}
	}
}

func InsertionSortNama(data *arrayPeserta, n int,isAscend bool) {
	/* I.S : array data berisi n data peserta yang belum tentu terurut berdasarkan nama peserta, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : array data terurut berdasarkan nama peserta menggunakan Insertion Sort, ascending jika isAscend = true dan descending jika isAscend = false 
	*/
	var pass, i int
	var temp Peserta
	pass = 1

	if isAscend {
		for pass <= n-1 {
			i = pass
			temp = (*data)[pass]
			for i > 0 && temp.namaPeserta < (*data)[i-1].namaPeserta {
				(*data)[i] = (*data)[i-1]
				i = i - 1
			}
			(*data)[i] = temp
			pass = pass + 1
		}
	} else {
		for pass <= n-1 {
			i = pass
			temp = (*data)[pass]
			for i > 0 && temp.namaPeserta > (*data)[i-1].namaPeserta {
				(*data)[i] = (*data)[i-1]
				i = i - 1
			}
			(*data)[i] = temp
			pass = pass + 1
		}
	}
}


func InsertionSortWaktu(waktu *arrayPeserta, n int,isAscend bool) {
	/* I.S : array data berisi n data peserta yang belum tentu terurut berdasarkan tanggal daftar, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : array data terurut berdasarkan tanggal daftar menggunakan Insertion Sort, ascending jika isAscend = true dan descending jika isAscend = false
	*/
	var pass, i int
	var temp Peserta
	pass = 1

	if isAscend {
		for pass <= n-1 {
			i = pass
			temp = (*waktu)[pass]
			for i > 0 && temp.WaktuDaftar < (*waktu)[i-1].WaktuDaftar {
				(*waktu)[i] = (*waktu)[i-1]
				i = i - 1
			}
			(*waktu)[i] = temp
			pass = pass + 1
		}
	}else {
		for pass <= n-1 {
			i = pass
			temp = (*waktu)[pass]
			for i > 0 && temp.WaktuDaftar > (*waktu)[i-1].WaktuDaftar {
				(*waktu)[i] = (*waktu)[i-1]
				i = i - 1
			}
			(*waktu)[i] = temp
			pass = pass + 1
		}
	}
}

func SelectionSortID(data *arrayPeserta, n int,isAscend bool) {
	/* I.S : array data berisi n data peserta yang belum tentu terurut berdasarkan ID, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : array data terurut berdasarkan ID menggunakan Selection Sort, ascending jika isAscend = true dan descending jika isAscend = false
	*/
	var pass, i, idx int
	var temp Peserta
	if isAscend {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*data)[i].id < (*data)[idx].id {
				idx = i
				}
			}
			temp = (*data)[pass]
			(*data)[pass] = (*data)[idx]
			(*data)[idx] = temp
		}
	}else {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*data)[i].id > (*data)[idx].id {
				idx = i
				}
			}
			temp = (*data)[pass]
			(*data)[pass] = (*data)[idx]
			(*data)[idx] = temp
		}
	}
}

func SelectionSortNama(data *arrayPeserta, n int, isAscend bool) {
	/* I.S : array data berisi n data peserta yang belum tentu terurut berdasarkan nama peserta, isAscend bernilai true untuk urut menaik dan false untuk urut menurun
	   F.S : array data terurut berdasarkan nama peserta menggunakan Selection Sort, ascending jika isAscend = true dan descending jika isAscend = false
	*/
	var pass, i, idx int
	var temp Peserta

	if isAscend {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*data)[i].namaPeserta < (*data)[idx].namaPeserta {
					idx = i
				}
			}
			temp = (*data)[pass]
			(*data)[pass] = (*data)[idx]
			(*data)[idx] = temp
		}
	}else {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*data)[i].namaPeserta > (*data)[idx].namaPeserta {
					idx = i
				}
			}
			temp = (*data)[pass]
			(*data)[pass] = (*data)[idx]
			(*data)[idx] = temp
		}
	}
}

func SelectionSortWaktu(waktu *arrayPeserta, n int,isAscend bool) {
	/* I.S : array data berisi n data peserta yang belum tentu terurut berdasarkan tanggal daftar, isAscend bernilai true untuk urut membesar dan false untuk urut mengecil
	   F.S : array data terurut berdasarkan tanggal daftar menggunakan Selection Sort, ascending jika isAscend = true dan descending jika isAscend = false
	*/
	var pass, i, idx int
	var temp Peserta

	if isAscend {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*waktu)[i].WaktuDaftar < (*waktu)[idx].WaktuDaftar {
					idx = i
				}
			}
			temp = (*waktu)[pass]
			(*waktu)[pass] = (*waktu)[idx]
			(*waktu)[idx] = temp
		}
	}else {
		for pass = 0; pass < n-1; pass++ {
			idx = pass
			for i = pass; i < n; i++ {
				if (*waktu)[i].WaktuDaftar > (*waktu)[idx].WaktuDaftar {
					idx = i
				}
			}
			temp = (*waktu)[pass]
			(*waktu)[pass] = (*waktu)[idx]
			(*waktu)[idx] = temp
		}
	}
}

func MenampilkanStatistik(data arrayPeserta, n int) {
	/* I.S : Array data berisi n data peserta yang tersimpan dalam sistem
	   F.S : Menampilkan jumlah peserta pada setiap kursus serta total peserta aktif yang terdaftar dalam sistem
	*/
	var i, j, k, count int
	var dihitung bool

	fmt.Printf("\n%15s┌───────────────────────────────────────────────────────┐\n", "")
	fmt.Printf("%15s│ %-53s │\n", "", "STATISTIK KURSUS")
	fmt.Printf("%15s├───────────────────────────────────────┬───────────────┤\n", "")
	fmt.Printf("%15s│ %-37s │ %-13s │\n", "", "Nama Kursus", "Jumlah")
	fmt.Printf("%15s├───────────────────────────────────────┼───────────────┤\n", "")

	if n == 0 {
		fmt.Printf("%15s│ %-53s │\n", "", "Belum ada data peserta")
	} else {
		for i = 0; i < n; i++ {
			dihitung = false

			for j = 0; j < i; j++ {
				if data[j].namaKursus == data[i].namaKursus {
					dihitung = true
				}
			}
			if !dihitung {
				count = 0
				for k = 0; k < n; k++ {
					if data[k].namaKursus == data[i].namaKursus {
						count++
					}
				}
				fmt.Printf("%15s│ %-37s │ %-13d │\n","",data[i].namaKursus,count)
			}
		}
	}
	fmt.Printf("%15s├───────────────────────────────────────┴───────────────┤\n", "")
	fmt.Printf("%15s│ Total Peserta Aktif                     %-d            │\n", "", n)
	fmt.Printf("%15s└───────────────────────────────────────────────────────┘\n", "")
}

func DummyData(peserta *arrayPeserta, n *int) {
	var i int
	i=0
	(*peserta)[i] = Peserta{"Benedict_Cumberbatch", "Akuntansi_Dasar", 73829104, 20211105}
	i++
	(*peserta)[i] = Peserta{"Joe_Cole", "Desain_Grafis", 19283745, 20190822}
	i++
	(*peserta)[i] = Peserta{"Scarlett_Johansson", "Public_Speaking", 56473829, 20230214}
	i++
	(*peserta)[i] = Peserta{"Arthur_Shelby", "Manajemen_Bisnis", 90128374, 20200709}
	i++
	(*peserta)[i] = Peserta{"Brie_Larson", "Fotografi_Dasar", 48572910, 20251201}
	i++
	(*peserta)[i] = Peserta{"Tom_Holland", "Kewirausahaan", 83920156, 20180430}
	i++
	(*peserta)[i] = Peserta{"Cillian_Murphy", "Ilmu_Komunikasi", 27485930, 20221015}
	i++
	(*peserta)[i] = Peserta{"Chris_Evans", "Manajemen_Waktu", 61029384, 20240118}
	i++
	(*peserta)[i] = Peserta{"Helen_McCrory", "Bahasa_Inggris", 38475619, 20190822}
	i++
	(*peserta)[i] = Peserta{"Paul_Rudd", "Digital_Marketing", 94857102, 20210611}
	i++
	(*peserta)[i] = Peserta{"Tom_Hardy", "Akuntansi_Dasar", 50392817, 20230925}
	i++
	(*peserta)[i] = Peserta{"Elizabeth_Olsen","Desain_Grafis",12938475,20200304}
	i++
	(*peserta)[i] = Peserta{"Thomas_Shelby","Public_Speaking",85746392,20251201}
	i++
	(*peserta)[i] = Peserta{"Anya_Taylor_Joy","Manajemen_Bisnis",41029385 ,20220520}
	i++
	(*peserta)[i] = Peserta{"Robert_Downey","Fotografi_Dasar", 76859403, 20181111}
	i++
	*n = i
}