package main

func main() {
	var (
		mineCraft        = game{title: "Minecraft", price: 5}
		worldOfWarcraft  = game{title: "World Of Warcraft", price: 19}
		eliteDangerous   = game{title: "Elite Dangerous", price: 54}
		candleInTheTomb  = book{title: "Candle In The Tomb", price: 20}
		barneyAndFriends = book{title: "Barney And Friends", price: 10}
		rBTEar           = compAcc{title: "Razer BT Earpiece", price: 159}
		rKeyboard        = compAcc{title: "Razer Keyboard", price: 110}
		logitechMouse    = compAcc{title: "Logitech Mouse", price: 80}
	)

	var store list
	store = append(store, mineCraft, worldOfWarcraft, eliteDangerous, candleInTheTomb, barneyAndFriends, rBTEar, rKeyboard, logitechMouse)
	store.print()
}
