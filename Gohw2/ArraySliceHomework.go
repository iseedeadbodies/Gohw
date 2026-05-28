package main

import "fmt"

func task1() {
	fmt.Println("1. Фиксированный массив")

	arr := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Массив:", arr)
	fmt.Println("Длина:", len(arr))
}

func task2() {
	fmt.Println("\n2. Индексы")

	arr := [4]int{1, 2, 3, 4}

	fmt.Println("До:", arr)

	arr[1] = 22
	arr[3] = 44

	fmt.Println("После:", arr)
}

func task3() {
	fmt.Println("\n3. Копирование массива")

	a := [3]int{1, 2, 3}
	b := a

	a[0] = 100

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("b не изменился, потому что массив копируется полностью")
}

func task4() {
	fmt.Println("\n4. Пустой slice")

	var s []int

	fmt.Println("Значение:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println("s == nil:", s == nil)
}

func task5() {
	fmt.Println("\n5. Slice через литерал")

	s := []int{5, 10, 15}

	fmt.Println("До append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, 20)

	fmt.Println("После append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
}

func task6() {
	fmt.Println("\n6. Slice через make")

	s := make([]int, 3, 5)

	s[0] = 10
	s[1] = 20
	s[2] = 30

	fmt.Println("До append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, 40)
	s = append(s, 50)

	fmt.Println("После append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println("cap не изменился, потому что вместимость была 5")
}

func task7() {
	fmt.Println("\n7. Slice из массива")

	arr := [6]int{1, 2, 3, 4, 5, 6}

	s := arr[2:5]

	fmt.Println("До изменения:")
	fmt.Println("Массив:", arr)
	fmt.Println("Slice:", s)

	s[0] = 100

	fmt.Println("После изменения:")
	fmt.Println("Массив:", arr)
	fmt.Println("Slice:", s)
	fmt.Println("Массив изменился, потому что slice ссылается на тот же массив")
}

func task8() {
	fmt.Println("\n8. Независимый slice")

	arr := [6]int{1, 2, 3, 4, 5, 6}

	s := arr[2:5]

	newSlice := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		newSlice[i] = s[i]
	}

	newSlice[0] = 999

	fmt.Println("Массив:", arr)
	fmt.Println("Старый slice:", s)
	fmt.Println("Новый slice:", newSlice)
	fmt.Println("Массив не изменился, потому что новый slice скопирован отдельно")
}

func changeFirst(s []int) {
	s[0] = 777
}

func task9() {
	fmt.Println("\n9. Функция и slice")

	s := []int{1, 2, 3}

	fmt.Println("До:", s)

	changeFirst(s)

	fmt.Println("После:", s)
	fmt.Println("Slice изменился, потому что функция получает ссылку на общий массив")
}

func appendInside(s []int) {
	s = append(s, 4)

	fmt.Println("Внутри функции:", s)
	fmt.Println("Внутри len:", len(s))
	fmt.Println("Внутри cap:", cap(s))
}

func task10() {
	fmt.Println("\n10. Append и подводный камень")

	s := make([]int, 3, 3)

	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Println("До функции:", s)
	fmt.Println("До len:", len(s))
	fmt.Println("До cap:", cap(s))

	appendInside(s)

	fmt.Println("Снаружи функции:", s)
	fmt.Println("Снаружи len:", len(s))
	fmt.Println("Снаружи cap:", cap(s))
	fmt.Println("Снаружи slice не изменился, потому что append создал новый массив")
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()
	task9()
	task10()
}