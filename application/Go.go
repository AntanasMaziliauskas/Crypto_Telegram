package application

// Go function starts two Go-Routines: PriceCheker and Sender
func (a *App) Go() {

	go a.CS.PriceChecker()

	go a.CS.Sender()
}
