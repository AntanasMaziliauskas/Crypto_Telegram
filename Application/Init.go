package Application

func Init() {

	//BOT Authentication iskvietimas

	//FromFile paleidimas
	app.CryptoRule, err = app.FromFile()

	////URL ID funkcijos paleidimas
	r := app.URLID()

}
