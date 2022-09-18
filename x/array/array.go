package array

func InsertAt[t interface{}](arr []t, index int, value t) []t {
	if index > len(arr) || index < 0 {
		return arr
	}

	if index == len(arr) {
		return append(arr, value)
	}

	newArr := append(arr[:index+1], arr[index:]...)
	newArr[index] = value
	return newArr
}

func RemoveAt[t interface{}](arr []t, index int) []t {
	if index > len(arr)-1 || index < 0 {
		return arr
	}

	if index == len(arr)-1 {
		return arr[:len(arr)-1]
	}

	if index == 0 {
		return arr[1:]
	}

	return append(arr[:index], arr[index+1:]...)
}
