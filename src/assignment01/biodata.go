package main

import (
	"fmt"
	"os"
)

func main() {

	var people = []map[string]string{
		{"name": "Husin", "alamat": "Jakarta", "pekerjaan": "Karyawan"},
		{"name": "Ahmad", "alamat": "Jakarta", "pekerjaan": "Karyawan"},
		{"name": "Abdurrahman", "alamat": "Jakarta", "pekerjaan": "Karyawan"},
		{"name": "Fattumah", "alamat": "Jakarta", "pekerjaan": "Karyawan"},
	}

	var args string

	if len(os.Args) > 1 {
		args = os.Args[1]
	}

	for i, person := range people {
		if args == person["name"] {
			fmt.Printf("ID : %d\nNama : %s\nAlamat : %s\nPekerjaan : %s\n", i, person["name"], person["alamat"], person["pekerjaan"])
		}
	}
}
