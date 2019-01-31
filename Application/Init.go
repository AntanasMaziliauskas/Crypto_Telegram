package Application

////BOT Authentication paleidimas

func Init() {
	//FromFile paleidimas
	app.CryptoRule, err = app.FromFile()

	////URL ID funkcijos paleidimas
	r := app.URLID()
}
