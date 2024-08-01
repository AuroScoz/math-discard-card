package utility

import (
	"fmt"
	"strconv"
	"strings"
)

// 將傳入的字串以傳入的字元分隔並轉為 Int 陣列
//
//	"1,5,6" -> []int{1, 5, 6}
func SplitInt(str string, char string) ([]int, error) {
	if str == "" {
		return []int{}, nil
	}
	parts := strings.Split(str, char)
	nums := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// 將傳入的字串以傳入的字元分隔並轉為 Float64 陣列
//
//	"1.1,5.4,7.6" -> []float64{1.1, 5.4, 7.6}
func SplitFloat(str string, char string) ([]float64, error) {
	if str == "" {
		return []float64{}, nil
	}
	parts := strings.Split(str, char)
	nums := make([]float64, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// 將字串最後一個字轉為數字
func ExtractLastDigit(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("ExtractLastDigit時傳入字串為空")
	}
	lastChar := s[len(s)-1:]
	num, err := strconv.Atoi(lastChar)
	if err != nil {
		return 0, fmt.Errorf("ExtractLastDigit最後一個字串 '%s' 不能轉為數字", lastChar)
	}
	return num, nil
}
