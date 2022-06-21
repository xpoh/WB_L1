/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры
	Human (аналог наследования).
*/
package main

/*
	Родительская структура
*/
type Human struct {
	Name string
}

/*
	Метод родительской структуры
*/
func (h *Human) Rename(newName string) {
	h.Name = newName
}

/*
	В дочернюю структуру встроен указатель на предка
*/
type Action struct {
	h   *Human
	Run func(speed float32) error
}

/*
	Конструктор с указателем на предка
*/
func NewAction(h *Human) *Action {
	return &Action{h: h}
}

// Пример:
func main() {
	//Создаем предка c полями и
	h := Human{Name: "Bob"}

	a := NewAction(&h)
	a.h.Rename("Alice")
}
