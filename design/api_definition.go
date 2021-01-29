package design

import (
	. "github.com/goadesign/goa/design"
    . "github.com/goadesign/goa/design/apidsl"
)

var _ = API ("goa - the infinitive approach of developing API", func() {
	    // Jejuluk API
	    Title("tdwiid/tompobulu")
	    // Penjelasan API
	    Description("Contoh develop API menggunakan Goa framework.")
        // Informasi lengkap, hubungi ke penulis API
        Contact(func() {
			Name("Dwicahya Sulistyawan")
			Email("dwi.chy_sul@gmail.com")
			URL("https://github.com/tdwiid/tompobulu/issues")
		})
		// Hak pakai API
		License(func() {
			Name("MIT")
			URL("https://www.academia.edu/45004935/MIT_license")
		})
		// Dokumentasi API
		Docs(func() {
			Description("wiki")
			URL("https://github.com/tdwiid/tompobulu/wiki")
		})
		// Pengaturan host
		Host("localhost:8080")
		// Protokol yang didukung, http, https dan/atau keduanya
		Scheme("http", "https")
		// Akar path untuk semua endpoint
		// Jika seumpama ada endpoint yang disebut /items, maka ini maksudnya /api/v1/items
		BasePath("/api/v1")
		
		// Definisi kebijakan CORS
		Origin("http://localhost:8080/swagger", func() {
		})
		// Satu atau lebih headers yang diekspose ke client
		Expose("X-Time")
		// Satu atau lebih method HTTP yang diijinkan
		Methods("GET","POST","PUT","DELETE")
		// Waktu yang dibutuhkan untuk menyimpan respons dari Preflight Request
		MaxAge(600)
		// Penentuan setting Access-Control-Allow-Credentials
		Credentials()
		
		
		
