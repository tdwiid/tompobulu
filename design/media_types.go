package design

import(
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Menentukan response data
// Memberi nama pada MediaType
var IntegerMedia = MediaType("application/vnd.integer+json", func() {
	// Gambaran umum
	Description("example")
	// Jenis nilai apa yang ada di sana (tipe multiple dapat ditentukan juga) 
	Attributes(func() {
	// id merupakan tipe integer
		Attribute("id", Integer, "id", func() {
				// Contoh nilai response yang dikembalikan
				Example(1)
		})
		// Elemen penting untuk response (lebih mudah membuat semua menjadi basics mandatory)
		Required("id")
	})
	// Format nilai balik response
	View("default", func() {
	Attribute("id")
	})
})

var MessageMedia = MediaType("application/vnd.message+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("message", String, "message", func() {
			Example("ok")
		})
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description ("example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
					Example(1)
		})
		Attribute("name", String, "nama", func() {
		Example("Dwicahya Sulistyawan")
		})
		Attribute("email", String, "alamat email", func() {
		Example("example@gmail.com")
		})
		Required("id", "name", "email")
	})
	// Kalau atribut tidak ditentukan maka menjadi MediaType default 
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("email")
	})
	// Kalau atributnya "tiny" atau sedikit, format responsenya bisa ringkas 
	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})
})

var ValidationMedia = MediaType ("application/vnd.validation+json", func() {
	Description ("Example")
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			Example(1)
		})
		Attribute("integerType", Integer, "angka (1~10)", func() {
			Example (5)
		})
		Attribute("stringType", String, "characters (berisi 1~10 huruf)", func() {
			Example ("Sulawesi")
		})
		Attribute("mail", String, "alamat email", func() {
			Example ("example@gmail.com")
		})
		Attribute("enumType", String, "tipe enum", func() {
			Example ("A")
		})
		Attribute("defaultType", String, "nilai default", func() {
			Example ("Default")
		})
		Attribute("reg", String, "bahasa formal/formal language", func() {
			Example ("12abc")
		})
	})
	Required("id", "integerType", "stringType", "email", "enumType", "defaultType", "reg")
	View("default", func() {
		Attribute("id")
		Attribute("integerType")
		Attribute("stringType")
		Attribute("email")
		Attribute("enumType")
		Attribute("defaultType")
		Attribute("reg")
	})
})

var OKMedia = Type ("ok", func() {
	Attribute("status", Integer, func() {
		Example (200)
	})
	Required ("status")
})

var ErrorValue = Type("errorValue", func() {
	Attribute("status", Integer, func() {
		Example(400)
	})
	Attribute("code", String, func() {
		Example("bad_request")
	})
	Attribute("details", strings, func() {
		Example("xxx adalah nil")
	})
		Required("status", "code", "detail")
})

var CustomeErrorMedia = MediaType("application/vnd.error+json", func() {
	Attribute("response", ErrorValue)
	Required("response")
	View("default", func() {
		Attribute("response")
	})
})

var ArticleData = Type("data", func() {
	Attribute("title", String)
	Attribute("body", String)
	Required("title", "body")
})

var ArticleMedia = MediaType("application/vnd.article+json", func() {
	Description ("example")
	Attribute("data", ArrayOf (ArticleData))
	Attribute("response", OKMedia)
	Required("data", "response")
	View("default", func() {
		Attribute("data")
		Attribute("response")
	})
})

// --------------------------------
// gorma
// --------------------------------

var AccountData = MediaType("application/vnd.account+json", func() {
	Description("celler account")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "name", func() {
		Example("Dwicahya Sulistyawan")
	})
	Attribute("mail", String, "email address", func() {
		Example("example@gmail.com")
	})
	Required("id", "name", "email")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("mail")
	})
})

var AccountMedia = MediaType("application/vnd.accountMedia+json", func() {
	Description("example")
	Attribute("data", ArrayOf(AccountData))
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("data", "status")
	View("default", func() {
		Attribute("data")
		Attribute("status")
	})
})

// Bottle

var BottleData = MediaType("application/vnd.bottle+json", func() {
	Description("cellar bottles")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "nama bottle", func() {
		Example("Bango soy sauce")
	})
	Attribute("quantity", Integer, "kuantitas", func() {
		Example(4)
	})
	// Account dengan tipe Nested MediaType
	Attribute("account", AccountData)
	Attribute("categories", ArrayOf(CategoryData))
	Required ("id", "name", "quantity", "account", "categories")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("quantity")
		Attribute("account")
	})
	View("relation", func() {
		Attribute ("id")
		Attribute ("name")
		Attribute ("quantity")
		Attribute ("account")
		Attribute ("categories")
	})
})

var BottleMedia = MediaType ("application/vnd.bottleMedia+json", func() {
	Description ("Example")
	Attribute("data", ArrayOf(BottleData))
	Attribute("status", Integer, func() {
		Example (200)
	})
	Required("data", "status")
	View ("default", func() {
	Attribute("data")
	Attribute("Status")
	})
})

// Category

var CategoryData = MediaType("application/vnd.category+json", func() {
	Description("cellar account")
	Attribute("id", Integer, "id", func() {
		Example(1)
	})
	Attribute("name", String, "nama", func() {
		Example("soy sauce")
	})
	Required("id", "name")
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
})

var CategoryMedia = MediaType("application/vnd.categoryMedia+json", func() {
	Description("example")
	Attribute("data", ArrayOf (CategoryData))
	Attribute("status", Integer, func() {
		Example(200)
	})
	Required("data", "status")
	View("default", func(){
		Attribute("data")
		Attribute("status")
	})
}) 
