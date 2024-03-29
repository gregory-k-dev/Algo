package main

type SimpleArray struct {
	length int
	array  []interface{}
}

const capResizeSA = 1

func (a *SimpleArray) Add(element interface{}, index int) {
	newArray := make([]interface{}, a.length+capResizeSA)
	for i := 0; i < index; i++ {
		newArray[i] = a.array[i]
	}
	newArray[index] = element
	for i := index + 1; i < a.length+capResizeSA; i++ {
		newArray[i] = a.array[i-1]
	}
	a.length = a.length + capResizeSA
	a.array = newArray
}

func (a *SimpleArray) GetLength() int {
	return a.length
}

func (a *SimpleArray) GetArray() []interface{} {
	return a.array
}

func (a *SimpleArray) Remove(index int) interface{} {
	var removeElement = a.array[index]
	a.length = a.length - capResizeSA
	newArray := make([]interface{}, a.length)
	for i := 0; i < index; i++ {
		newArray[i] = a.array[i]
	}
	for i := index; i < a.length; i++ {
		newArray[i] = a.array[i+1]
	}
	a.array = newArray
	return removeElement
}
