package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Menentukan path dari berkas Swagger di lokal
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})

var _ = Resource("js", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/js/*filepath", "/js/")
})

// Penentuan /actions
var _ = Resource("actions", func() {
	// Akar path untuk resource /actions
	BasePath("/actions")
	/*
		Tentukan apa yang harus dilakukan dengan resource, seperti :
		+ add ; MENAMBAHKAN resource
		+ list ; MENAMPILKAN resources pada sebuah list
		+ delete ; MENGHAPUS resource
		Berupa penentuan resource yang biasa dituliskan, kalau mau tiru saja.
	*/
	Action("ping", func() {
		// Penjelasan action
		Description("Memeriksa ketersambungan dengan server.")
		Routing(
			// Endpoint-> GET http://localhost/api/v1/actions/ping
			GET("/ping"),
		)
		// Response yang dikembalikan
		// 200 OK + MediaType ditentukan oleh MessageType
		Response(OK, MessageMedia)
		// 400 BadRequest + ErrorMedia, MediaType yang muncul secara default.
		// Secara otomatis mengembalikan params yang hilang, dll.
		Response(BadRequest, ErrorMedia)
	})
	Action("hello", func() {
		Description("Greet, biasanya 'hello'.")
		Routing(
			GET("/hello"),
		)
		//  Params yang bisa ditambahkan ke request
		Params(func() {
			// Kita dapat memasukkan params dengan nama 'Nama Anda' (string).
			Param("name", String, "Nama Anda", func() {
			    // ("") akan berguna untuk menyimpan string kosong.	
				Default("")
			})
			// Params yang diatur (ada nilai defaultnya yakni 'name', kosongkan kalau tidak ada)
			Required("name")
		})
		Response(OK, MessageMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("id", func() {
		Description("Penentuan actions untuk（:id）")
		Routing(
			// Kita tentukan resource untuk endpoint-nya
			// GET http://localhost:8080/api/v1/actions/1
			GET("/:id"),
		)
		Params(func() {
			// :id harus bertipe Integer.
			Param("id", Integer, "id")
			// Required merupakan endpoint yang menyertakan resource, tidak usah kita buat.
			// Required("id")
		})
		Response(OK, IntegerMedia)
		// Mengembalikan NotFound kalau resource yang ditentukan tidak ada.
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("security", func() {
	BasePath("/securiy")
	Security(UserAuth)
	Action("security", func() {
		Description("Penentuan akar path untuk fitur security.")
		Routing(
			GET("/"),
		)
		Response(OK, MessageMedia)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("method", func() {
	BasePath("/method")
	Action("method", func() {
		Description("Penentuan akar path untuk fitur method.")
		Routing(
			GET("/get"),
			POST("/post"),
			DELETE("/delete"),
			PUT("/put"),
		)
		Response(OK, MessageMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Description("Penentuan akar path untuk fitur list.")
		Routing(
			GET("/list"),
			GET("/list/new"),
			GET("/list/topic"),
		)
		Response(OK, CollectionOf(UserMedia))
		Response(BadRequest, ErrorMedia)
	})
	Action("follow", func() {
		Description("Penentuan akar path untuk fitur follow.")
		Routing(
			PUT("/users/follow"),
			DELETE("/users/follow"),
		)
		Response(OK, MessageMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("etc", func() {
		Routing(GET("/users/:id/follow/:type"))
		Description("Penentuan akar path untuk fitur etc.")
		Params(func() {
			Param("id", Integer, "id")
			Param("type", Integer, "tipe", func() {
				Enum(1, 2, 3)
			})
		})
		Response(OK, "plain/text")
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("response", func() {
	BasePath("/response")
	Action("list", func() {
		Description("User (multiple record), semua ditampilkan.")
		Routing(
			GET("/users"),
		)
		// Mengembalikan nilai record yang bersifat multiple
		Response(OK, CollectionOf(UserMedia))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("User (tunggal), cuma yang ber-ID :id yang tampil.")
		Routing(
			GET("/users/:id"),
		)
		// Tunggal
		Response(OK, UserMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("hash", func() {
		Description("User (hash).")
		Routing(
			GET("/users/hash"),
		)
		// Associative array
		Response(OK, HashOf(String, Integer))
		Response(BadRequest, ErrorMedia)
	})
	Action("array", func() {
		Description("User (array).")
		Routing(
			GET("/users/array"),
		)
		// Array
		Response(OK, ArrayOf(Integer))
		Response(BadRequest, ErrorMedia)
	})
	Action("nested", func() {
		Description("Nested Media Type.")
		Routing(
			GET("/users/nested"),
		)
		Params(func() {
			Param("test", String, func() {
				MinLength(1)
			})
			Required("test")
		})
		// Mengembalikan nested elements
		Response(OK, ArticleMedia)
		Response(BadRequest, CustomeErrorMedia)
	})
})

var _ = Resource("validation", func() {
	BasePath("/validation")
	Action("validation", func() {
		Description("Validation")
		Routing(
			GET("/"),
		)
		Params(func() {
			// bertipe Integer
			Param("id", Integer, "id", func() {
				Example(1)
			})
			// bertipe Integer, antara 0 sd. 10
			Param("integerType", Integer, "angka (1~10)", func() {
				Minimum(0)
				Maximum(10)
				Example(2)
			})
			// bertipe String, berisi character antara 0 sd. 10
			Param("stringType", String, "characters (berisi 1~10 huruf)", func() {
				MinLength(1)
				MaxLength(10)
				Example("AIUEO")
			})
			// bertipe String, berformat email
			Param("email", String, "alamat email", func() {
				Format("email")
				Example("example@gmail.com")
			})
			// bertipe String, dan menjadi elemen dalam suatu Enum
			Param("enumType", String, "tipe enum", func() {
				Enum("A", "B", "C")
				Example("A")
			})
			// jika tidak ditentukan tipe String-nya, maka character string default otomatis diatur
			Param("defaultType", String, "nilai default", func() {
				Default("Default")
				Example("Default")
			})
			// bertipe String, berpola Regex (Regular Expression)
			Param("reg", String, "bahasa formal/formail language", func() {
				Pattern("^[a-z0-9]{5}$")
				Example("12abc")
			})

			// semua params yang diperlukan
			Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
		})
		Response(OK, ValidationMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("accounts", func() {
	BasePath("/accounts")
	Action("list", func() {
		Description("Multiple")
		Routing(
			GET("/"),
		)
		// Mengembalikan records dari account yang bersifat multiple
		Response(OK, CollectionOf(AccountData))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Singular")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		// Mengembalikan 1 record, jika ditemukan
		Response(OK, AccountData)
		// Kalau tidak, mengembalikan NotFound
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("Menambahkan record.")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("name", String, "name", func() {
				Example("Dwicahya Sulistyawan")
			})
			Attribute("email", String, "email", func() {
				Format("email")
				Example("example@gmail.com")
			})
			Required("name", "email")
		})
		Response(OK, AccountData)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("Menghapus record.")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "name", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Memperbaharui record.")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			// Mencari record menggunakan :id
			Param("id", Integer, "id")
		})

		Payload(func() {
			// Pengaturan default, abaikan kalau nihil pengaturan
			Param("name", String, "name", func() {
				Default("")
			})
			// Idem sda., abaikan kalau nihil pengaturan
			Param("email", String, "email", func() {
				Format("email")
				Default("")
			})
		})
		// Response dari server, OK
		Response(OK)
		// Response tidak ditemukan, kesalahan 404.
		Response(NotFound)
		// Response berupa BadRequest
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("bottles", func() {
	BasePath("/bottles")
	Action("list", func() {
		Description("Multiple record")
		Routing(
			GET("/"),
		)
		Response(OK, CollectionOf(BottleData))
		Response(BadRequest, ErrorMedia)
	})
	Action("listRelation", func() {
		Description("Multiple record dan memiliki relasi")
		Routing(
			GET("/relation"),
		)
		Response(OK, CollectionOf(BottleData))
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Record tunggal")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK, BottleData)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("Menambahkan record.")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("account_id", Integer, "Account ID", func() {
				Example(1)
			})
			Attribute("name", String, "Bottle name", func() {
				Default("")
				Example("Something Bango soy sauce")
			})
			Attribute("quantity", Integer, "Kuantitas", func() {
				Example(0)
			})
			Required("account_id", "name", "quantity")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("Menghapus record.")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Memperbaharui record.")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Payload(func() {
			Attribute("name", String, "Bottle name", func() {
				Default("")
				Example("Something Bango soy sauce.")
			})
			Attribute("quantity", Integer, "Kuantitas", func() {
				Default(0)
				Minimum(0)
				Example(0)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("accounts_data", func() {
	BasePath("/accounts_data")
	Action("list", func() {
		Description("Multiple")
		Routing(
			GET("/"),
		)
		Response(OK, AccountMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Tunggal")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK, AccountMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("Menambahkan record.")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("name", String, "name", func() {
				Example("Dwicahya Sulistyawan")
			})
			Attribute("email", String, "email", func() {
				Format("email")
				Example("example@gmail.com")
			})
			Required("name", "email")
		})
		Response(OK, AccountMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("Menghapus record.")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "name", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Memperbaharui record.")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id")
		})
		Payload(func() {
			Attribute("name", String, "name", func() {
				Default("")
			})
			Attribute("email", String, "email", func() {
				Format("email")
				Default("")
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("bottles_data", func() {
	BasePath("/bottles_data")
	Action("list", func() {
		Description("Multiple record")
		Routing(
			GET("/"),
		)
		Response(OK, BottleMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("show", func() {
		Description("Record tunggal")
		Routing(
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK, BottleMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("add", func() {
		Description("Menambahkan record.")
		Routing(
			POST("/"),
		)
		Payload(func() {
			Attribute("account_id", Integer, "Account ID", func() {
				Example(1)
			})
			Attribute("name", String, "Bottle name", func() {
				Default("")
				Example("Something Bango soy sauce")
			})
			Attribute("quantity", Integer, "Kuantitas", func() {
				Example(0)
			})
			Required("account_id", "name", "quantity")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})
	Action("delete", func() {
		Description("Menghapus record.")
		Routing(
			DELETE("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
	Action("update", func() {
		Description("Memperbaharui record.")
		Routing(
			PUT("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id", func() {
				Example(1)
			})
		})
		Payload(func() {
			Param("name", String, "Bottle name", func() {
				Default("")
				Example("Something Bango soy sauce")
			})
			Param("quantity", Integer, "Kuantitas", func() {
				Default(0)
				Minimum(0)
				Example(10)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
