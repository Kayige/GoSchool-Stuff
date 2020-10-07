package main

// ItemInfo Struct
type ItemInfo struct {
	Category string
	Quantity int
	UnitCost float64
}

// ShoppingList and ItemInfo
type ShoppingList map[string]ItemInfo

func (shoppingList ShoppingList) addItem(name string, item ItemInfo) {
	shoppingList[name] = item
}

func (shoppingList ShoppingList) getItem(name string) (ItemInfo, bool) {
	item, flag := shoppingList[name]
	return item, flag
}

func (shoppingList ShoppingList) deleteItem(name string) bool {
	if _, flag := shoppingList.getItem(name); flag {
		delete(shoppingList, name)
		return true
	}
	return false
}

func (shoppingList ShoppingList) modifyItem(name string, n string, item ItemInfo) {
	shoppingList.addItem(n, item)
	if name != n {
		shoppingList.deleteItem(name)
	}
}
