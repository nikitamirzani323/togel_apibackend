package helpers

import (
	"fmt"
	"strconv"
)

func RumusTogel(angka, tipe, nomorkeluaran string, idcomppasaran, idtrxkeluarandetail int) string {
	result := "LOSE"
	temp := angka
	temp4d := string([]byte(temp)[0]) + string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp3d := string([]byte(temp)[1]) + string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp2d := string([]byte(temp)[2]) + string([]byte(temp)[3])
	temp2dd := string([]byte(temp)[0]) + string([]byte(temp)[1])
	temp2dt := string([]byte(temp)[1]) + string([]byte(temp)[2])

	switch tipe {
	case "4D":
		if temp4d == nomorkeluaran {
			result = "WINNER"
		}
	case "3D":
		if temp3d == nomorkeluaran {
			result = "WINNER"
		}
	case "2D":
		if temp2d == nomorkeluaran {
			result = "WINNER"
		}
	case "2DD":
		if temp2dd == nomorkeluaran {
			result = "WINNER"
		}
	case "2DT":
		if temp2dt == nomorkeluaran {
			result = "WINNER"
		}
	case "COLOK_BEBAS":
		flag := false
		count := 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == nomorkeluaran {
				flag = true
				count = count + 1
			}
		}
		if flag == true {
			win := 0
			if count == 1 {
				win = count * 1
			}
			if count == 2 {
				win = count * 2
			}
			if count == 3 {
				win = count * 3
			}
			if count == 4 {
				win = count * 4
			}
			fmt.Println(win)
			result = "WINNER"
		}
	case "COLOK_MACAU":
		flag_1 := false
		flag_2 := false
		count_1 := 0
		count_2 := 0
		totalcount := 0
		win := 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
		}
		if flag_1 == true && flag_2 == true {
			totalcount = count_1 + count_2
			if totalcount == 2 {
				win = totalcount * 2
			}
			if totalcount == 3 {
				win = totalcount * 3
			}
			if totalcount == 4 {
				win = totalcount * 4
			}
			fmt.Println(win)
			result = "WINNER"
		}
	case "COLOK_NAGA":
		flag_1 := false
		flag_2 := false
		flag_3 := false
		count_1 := 0
		count_2 := 0
		count_3 := 0
		totalcount := 0
		win := 0
		for i := 0; i < len(temp); i++ {
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[0]) {
				flag_1 = true
				count_1 = count_1 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[1]) {
				flag_2 = true
				count_2 = count_2 + 1
			}
			if string([]byte(temp)[i]) == string([]byte(nomorkeluaran)[2]) {
				flag_3 = true
				count_3 = count_3 + 1
			}
		}
		if flag_1 == true && flag_2 == true {
			if flag_3 == true {
				totalcount = count_1 + count_2 + count_3

				if totalcount == 3 {
					win = totalcount * 3
				}
				if totalcount == 4 {
					win = totalcount * 4
				}
				fmt.Println(win)
				result = "WINNER"
			}
		}
	case "COLOK_JITU":
		flag := false
		as := string([]byte(temp)[0]) + "_AS"
		kop := string([]byte(temp)[1]) + "_KOP"
		kepala := string([]byte(temp)[2]) + "_KEPALA"
		ekor := string([]byte(temp)[3]) + "_KEKOR"

		if as == nomorkeluaran {
			flag = true
		}
		if kop == nomorkeluaran {
			flag = true
		}
		if kepala == nomorkeluaran {
			flag = true
		}
		if ekor == nomorkeluaran {
			flag = true
		}
		if flag == true {
			result = "WINNER"
		}
	case "50_50_UMUM":
		flag := false
		data := []string{}
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		dasar := kepala_2 + ekor_2
		//BESARKECIL
		if kepala_2 <= 4 {
			data = append(data, "KECIL")
		} else {
			data = append(data, "BESAR")
		}
		//GENAPGANJIL
		if ekor_2%2 == 0 {
			data = append(data, "GENAP")
		} else {
			data = append(data, "GANJIL")
		}
		//TEPITENGAH
		if dasar >= 0 && dasar <= 24 {
			data = append(data, "TEPI")
		}
		if dasar >= 25 && dasar <= 74 {
			data = append(data, "TENGAH")
		}
		if dasar >= 75 && dasar <= 99 {
			data = append(data, "TEPI")
		}
		for i := 0; i < len(data); i++ {
			if data[i] == nomorkeluaran {
				flag = true
			}
		}
		if flag == true {
			result = "WINNER"
		}
		fmt.Println(data)
	case "50_50_SPECIAL":
		flag := false
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)
		//AS - BESARKECIL == GENAPGANJIL
		if as_2 <= 4 {
			if "AS_KECIL" == nomorkeluaran {
				flag = true
			}
		} else {
			if "AS_BESAR" == nomorkeluaran {
				flag = true
			}
		}
		if as_2%2 == 0 {
			if "AS_GENAP" == nomorkeluaran {
				flag = true
			}
		} else {
			if "AS_GANJIL" == nomorkeluaran {
				flag = true
			}
		}

		//KOP - BESARKECIL == GENAPGANJIL
		if kop_2 <= 4 {
			if "KOP_KECIL" == nomorkeluaran {
				flag = true
			}
		} else {
			if "KOP_BESAR" == nomorkeluaran {
				flag = true
			}
		}
		if kop_2%2 == 0 {
			if "KOP_GENAP" == nomorkeluaran {
				flag = true
			}
		} else {
			if "KOP_GANJIL" == nomorkeluaran {
				flag = true
			}
		}

		//KEPALA - BESARKECIL == GENAPGANJIL
		if kepala_2 <= 4 {
			if "KEPALA_KECIL" == nomorkeluaran {
				flag = true
			}
		} else {
			if "KEPALA_BESAR" == nomorkeluaran {
				flag = true
			}
		}
		if kepala_2%2 == 0 {
			if "KEPALA_GENAP" == nomorkeluaran {
				flag = true
			}
		} else {
			if "KEPALA_GANJIL" == nomorkeluaran {
				flag = true
			}
		}

		//EKOR - BESARKECIL == GENAPGANJIL
		if ekor_2 <= 4 {
			if "EKOR_KECIL" == nomorkeluaran {
				flag = true
			}
		} else {
			if "EKOR_BESAR" == nomorkeluaran {
				flag = true
			}
		}
		if ekor_2%2 == 0 {
			if "EKOR_GENAP" == nomorkeluaran {
				flag = true
			}
		} else {
			if "EKOR_GANJIL" == nomorkeluaran {
				flag = true
			}
		}

		if flag == true {
			result = "WINNER"
		}
	case "50_50_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		belakang := ""
		depan_1 := ""
		tengah_1 := ""
		belakang_1 := ""
		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2%2 == 0 {
			data_1 = "GENAP"
		} else {
			data_1 = "GANJIL"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2%2 == 0 {
			data_3 = "GENAP"
		} else {
			data_3 = "GANJIL"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}
		depan = data_1 + "-" + data_2
		tengah = data_2 + "-" + data_3
		belakang = data_3 + "-" + data_4

		if depan == "GENAP-GANJIL" || depan == "GANJIL-GENAP" {
			depan = "DEPAN_STEREO"
		} else {
			depan = "DEPAN_MONO"
		}
		if tengah == "GENAP-GANJIL" || tengah == "GANJIL-GENAP" {
			tengah = "TENGAH_STEREO"
		} else {
			tengah = "TENGAH_MONO"
		}
		if belakang == "GENAP-GANJIL" || belakang == "GANJIL-GENAP" {
			belakang = "BELAKANG_STEREO"
		} else {
			belakang = "BELAKANG_MONO"
		}
		if as_2 < kop_2 {
			depan_1 = "DEPAN_KEMBANG"
		}
		if as_2 > kop_2 {
			depan_1 = "DEPAN_KEMPIS"
		}
		if as_2 == kop_2 {
			depan_1 = "DEPAN_KEMBAR"
		}
		if kop_2 < kepala_2 {
			tengah_1 = "TENGAH_KEMBANG"
		}
		if kop_2 > kepala_2 {
			tengah_1 = "TENGAH_KEMPIS"
		}
		if kop_2 == kepala_2 {
			tengah_1 = "TENGAH_KEMBAR"
		}
		if kepala_2 < ekor_2 {
			belakang_1 = "BELAKANG_KEMBANG"
		}
		if kepala_2 > ekor_2 {
			belakang_1 = "BELAKANG_KEMPIS"
		}
		if kepala_2 == ekor_2 {
			belakang_1 = "BELAKANG_KEMBAR"
		}

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}
		if depan_1 == nomorkeluaran {
			flag = true
		}
		if tengah_1 == nomorkeluaran {
			flag = true
		}
		if belakang_1 == nomorkeluaran {
			flag = true
		}

		if flag == true {
			result = "WINNER"
		}
	case "MACAU_KOMBINASI":
		flag := false
		data_1 := ""
		data_2 := ""
		data_3 := ""
		data_4 := ""
		depan := ""
		tengah := ""
		belakang := ""

		as := string([]byte(temp)[0])
		kop := string([]byte(temp)[1])
		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		as_2, _ := strconv.Atoi(as)
		kop_2, _ := strconv.Atoi(kop)
		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		if as_2 <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if kop_2%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}
		if kepala_2 <= 4 {
			data_3 = "KECIL"
		} else {
			data_3 = "BESAR"
		}
		if ekor_2%2 == 0 {
			data_4 = "GENAP"
		} else {
			data_4 = "GANJIL"
		}

		depan = "DEPAN_" + data_1 + "_" + data_2
		tengah = "TENGAH_" + data_2 + "_" + data_3
		belakang = "BELAKANG_" + data_3 + "_" + data_4

		if depan == nomorkeluaran {
			flag = true
		}
		if tengah == nomorkeluaran {
			flag = true
		}
		if belakang == nomorkeluaran {
			flag = true
		}

		if flag == true {
			result = "WINNER"
		}
	case "DASAR":
		flag := false
		data_1 := ""
		data_2 := ""

		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])

		kepala_2, _ := strconv.Atoi(kepala)
		ekor_2, _ := strconv.Atoi(ekor)

		dasar := kepala_2 + ekor_2

		if dasar > 9 {
			temp2 := strconv.Itoa(dasar) //int to string
			temp21 := string([]byte(temp2)[0])
			temp22 := string([]byte(temp2)[1])

			temp21_2, _ := strconv.Atoi(temp21)
			temp22_2, _ := strconv.Atoi(temp22)
			dasar = temp21_2 + temp22_2
		}
		if dasar <= 4 {
			data_1 = "KECIL"
		} else {
			data_1 = "BESAR"
		}
		if dasar%2 == 0 {
			data_2 = "GENAP"
		} else {
			data_2 = "GANJIL"
		}

		if data_1 == nomorkeluaran {
			flag = true
		}
		if data_2 == nomorkeluaran {
			flag = true
		}

		if flag == true {
			result = "WINNER"
		}
	case "SHIO":
		flag := false

		kepala := string([]byte(temp)[2])
		ekor := string([]byte(temp)[3])
		data := Tableshio(kepala + ekor)

		if data == nomorkeluaran {
			flag = true
		}

		if flag == true {
			result = "WINNER"
		}
	}
	return result
}
func Tableshio(shiodata string) string {
	result := ""
	tikus := []string{"01", "13", "25", "37", "49", "61", "73", "85", "97"}
	babi := []string{"02", "14", "26", "38", "50", "62", "74", "86", "98"}
	anjing := []string{"03", "15", "27", "39", "51", "63", "75", "87", "99"}
	ayam := []string{"04", "16", "28", "40", "52", "64", "76", "88", "00"}
	monyet := []string{"05", "17", "29", "41", "53", "65", "77", "89", ""}
	kambing := []string{"06", "18", "30", "42", "54", "66", "78", "90", ""}
	kuda := []string{"07", "19", "31", "43", "55", "67", "79", "91", ""}
	ular := []string{"08", "20", "32", "44", "56", "68", "80", "92", ""}
	naga := []string{"09", "21", "33", "45", "57", "69", "81", "93", ""}
	kelinci := []string{"10", "22", "34", "46", "58", "70", "82", "94", ""}
	harimau := []string{"11", "23", "35", "47", "59", "71", "83", "95", ""}
	kerbau := []string{"12", "24", "36", "48", "60", "72", "84", "96", ""}
	for i := 0; i < len(babi); i++ {
		if shiodata == babi[i] {
			result = "BABI"
		}
	}
	for i := 0; i < len(ular); i++ {
		if shiodata == ular[i] {
			result = "ULAR"
		}
	}
	for i := 0; i < len(anjing); i++ {
		if shiodata == anjing[i] {
			result = "ANJING"
		}
	}
	for i := 0; i < len(ayam); i++ {
		if shiodata == ayam[i] {
			result = "AYAM"
		}
	}
	for i := 0; i < len(monyet); i++ {
		if shiodata == monyet[i] {
			result = "MONYET"
		}
	}
	for i := 0; i < len(kambing); i++ {
		if shiodata == kambing[i] {
			result = "KAMBING"
		}
	}
	for i := 0; i < len(kuda); i++ {
		if shiodata == kuda[i] {
			result = "KUDA"
		}
	}
	for i := 0; i < len(naga); i++ {
		if shiodata == naga[i] {
			result = "NAGA"
		}
	}
	for i := 0; i < len(kelinci); i++ {
		if shiodata == kelinci[i] {
			result = "KELINCI"
		}
	}
	for i := 0; i < len(harimau); i++ {
		if shiodata == harimau[i] {
			result = "HARIMAU"
		}
	}
	for i := 0; i < len(kerbau); i++ {
		if shiodata == kerbau[i] {
			result = "KERBAU"
		}
	}
	for i := 0; i < len(tikus); i++ {
		if shiodata == tikus[i] {
			result = "TIKUS"
		}
	}
	return result
}
