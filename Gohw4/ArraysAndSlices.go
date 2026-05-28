package main

import "fmt"

func main() {

	// 1. Массив из 10 чисел
	fmt.Println("1. Массив")

	arr := [10]int{5, 8, 12, 3, 7, 20, 15, 1, 10, 6}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 2. Сумма массива
	fmt.Println("\n2. Сумма массива")

	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println("Сумма:", sum)

	// 3. Максимум и минимум
	fmt.Println("\n3. Максимум и минимум")

	max := arr[0]
	min := arr[0]

	for i := 0; i < len(arr); i++ {

		if arr[i] > max {
			max = arr[i]
		}

		if arr[i] < min {
			min = arr[i]
		}
	}

	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)

	// 4. Количество четных
	fmt.Println("\n4. Четные числа")

	count := 0

	for i := 0; i < len(arr); i++ {

		if arr[i]%2 == 0 {
			count++
		}
	}

	fmt.Println("Количество четных:", count)

	// 5. Умножение на 2
	fmt.Println("\n5. Умножение на 2")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i] * 2)
	}

	// 6. Slice + append
	fmt.Println("\n6. Slice + append")

	s := []int{1, 2, 3, 4, 5}

	s = append(s, 6)
	s = append(s, 7)
	s = append(s, 8)

	fmt.Println(s)

	// 7. Удаление элемента
	fmt.Println("\n7. Удаление элемента")

	s2 := []int{10, 20, 30, 40, 50}

	index := 2

	s2 = append(s2[:index], s2[index+1:]...)

	fmt.Println(s2)

	// 8. Среднее арифметическое
	fmt.Println("\n8. Среднее арифметическое")

	s3 := []int{10, 20, 30, 40, 50}

	total := 0

	for i := 0; i < len(s3); i++ {
		total += s3[i]
	}

	average := float64(total) / float64(len(s3))

	fmt.Println("Среднее:", average)

	// 9. Самое короткое имя
	fmt.Println("\n9. Самое короткое имя")

	names := []string{"Ali", "Dias", "Alibek", "Aruzhan"}

	shortest := names[0]

	for i := 0; i < len(names); i++ {

		if len(names[i]) < len(shortest) {
			shortest = names[i]
		}
	}

	fmt.Println("Самое короткое имя:", shortest)

	// 10. Числа больше 10
	fmt.Println("\n10. Числа больше 10")

	nums := []int{5, 12, 8, 25, 30, 2, 15}

	var result []int

	for i := 0; i < len(nums); i++ {

		if nums[i] > 10 {
			result = append(result, nums[i])
		}
	}

	fmt.Println(result)
}