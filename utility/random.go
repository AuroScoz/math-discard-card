package utility

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
)

// RandomFloatBetweenInts 從兩個整數之間生成一個隨機float64
func RandomFloatBetweenInts(min, max int) (float64, error) {
	if min > max {
		return 0, fmt.Errorf("RandomFloatBetweenInts傳入值不符合規則 最小值<=最大值")
	}
	return float64(min) + rnd.Float64()*(float64(max)-float64(min)), nil
}

// 從兩個整數之間生成一個隨機int 傳入0,100會回傳0到100(包含0和100)的隨機整數
func GetRandomIntFromMinMax(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("RandomIntBetweenInts傳入值不符合規則 最小值<=最大值")
	}
	return rnd.Intn(max-min+1) + min, nil
}

// GetRandomTFromSlice 傳入泛型切片，返回隨機1個元素。
func GetRandomTFromSlice[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var value T
		return value, fmt.Errorf("GetRandomTFromSlice傳入參數錯誤")
	}
	return slice[rnd.Intn(len(slice))], nil
}

// 從map中取隨機key值出來
func GetRndKeyFromMap[K comparable, V any](m map[K]V) K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		var defaultK K
		return defaultK // 如果map為空, 返回K類型的零值
	}

	return keys[rnd.Intn(len(keys))] // 隨機選擇一個鍵並返回
}

// 從map中取隨機value值出來
func GetRndValueFromMap[K comparable, V any](m map[K]V) V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}

	if len(values) == 0 {
		var defaultV V
		return defaultV // 如果map為空, 返回V類型的零值
	}
	return values[rnd.Intn(len(values))] // 隨機選擇一個值並返回
}

// 傳入機率回傳結果 EX. 傳入0.3就是有30%機率返回true
func GetProbResult(prob float64) bool {
	return rnd.Float64() < prob
}

// 範例: 傳入"100~200" 回傳100~199之間的int
func GetRndIntFromRangeStr(input string, delimiter string) (int, error) {
	parts := strings.Split(input, delimiter)
	if len(parts) != 2 {
		return 0, fmt.Errorf("傳入字串要剛好只有一個分隔符號")
	}
	min, errMin := strconv.Atoi(parts[0])
	max, errMax := strconv.Atoi(parts[1])
	if errMin != nil || errMax != nil {
		return 0, fmt.Errorf("傳入字串的最小獲最大值無法轉為數字")
	}
	if min > max {
		return 0, fmt.Errorf("傳入字串的最小不可大於最大值")
	}
	rndInt, err := GetRandomIntFromMinMax(min, max)
	if err != nil {
		return 0, err
	}
	return rndInt, nil
}

// 範例: 傳入"100,200,300" 回傳隨機一個值, 例如200
func GetRndIntFromString(input string, delimiter string) (int, error) {
	// 檢查 input 不為空字串
	if input == "" {
		return 0, fmt.Errorf("input string is empty or incorrect format")
	}

	// 收集合法 parts
	parts := strings.Split(input, delimiter)
	validNumbers := make([]int, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			number, err := strconv.Atoi(part)
			if err == nil {
				validNumbers = append(validNumbers, number)
			}
		}
	}

	// 檢查是否不存在任何合法 parts
	if len(validNumbers) == 0 {
		return 0, fmt.Errorf("no valid numbers found in the input string")
	}

	return validNumbers[rnd.Intn(len(validNumbers))], nil
}

// 範例: 傳入"100,200,300" 回傳隨機一個字串, 例如"200"
func GetRndStrFromString(input string, delimiter string) (string, error) {
	// 檢查 input 不為空字串
	if input == "" {
		return "", fmt.Errorf("input string is empty or incorrect format")
	}

	// 收集合法 parts
	parts := strings.Split(input, delimiter)
	validParts := make([]string, 0, len(parts))
	for _, part := range parts {
		if part != "" {
			validParts = append(validParts, part)
		}
	}

	// 檢查是否不存在任何合法 parts
	if len(validParts) == 0 {
		return "", fmt.Errorf("no valid parts after string splits")
	}

	return validParts[rnd.Intn(len(validParts))], nil
}
