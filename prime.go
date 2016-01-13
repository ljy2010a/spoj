package main

import (
	"fmt"
	"math"
)

func main() {
	var k, j, i, max_m, max_n, test_cases, kase int64
	fmt.Scanln(&test_cases)
	case_m, case_n := make([]int64, test_cases), make([]int64, test_cases)
	EratosthenesArray := make(map[int64][]bool)
	max_m = 0
	max_n = 0
	for i = 0; i < test_cases; i++ {
		fmt.Scanf("%d %d", &case_m[i], &case_n[i])
		if case_m[i] > case_n[i] {
			case_m[i] = 0
			case_n[i] = 0
		}
		if max_m < case_m[i] {
			max_m = case_m[i]
		}
		if max_n < case_n[i] {
			max_n = case_n[i]
		}
		length := case_n[i] - case_m[i] + 1
		EratosthenesArray[i] = make([]bool, length)
	}

	if max_m <= max_n {
		upperbound := int64(math.Sqrt(float64(max_n)))
		UpperboundArray := make([]bool, upperbound + 1)
		for i = 2; i <= upperbound; i++ {
			if !UpperboundArray[i] {
				for k = i * i; k <= upperbound; k += i {
					UpperboundArray[k] = true
				}
				for kase = 0; kase < test_cases; kase++ {
					start := (case_m[kase] - i * i) / i

					if case_m[kase] - i * i < 0 {
						start = i
					}
					for k = start * i; k <= case_n[kase]; k += i {
						if k >= case_m[kase] && k <= case_n[kase] {
							EratosthenesArray[kase][k - case_m[kase]] = true
						}
					}
				}
			}
		}
	}

	for i = 0; i < test_cases; i++ {
		k = 0
		for j = 0; j <= case_n[i] - case_m[i]; j++ {
			if !EratosthenesArray[i][j] {
				ret := case_m[i] + j
				if ret > 1 {
					fmt.Println(ret)
				}
			}
		}
		fmt.Println()
	}
}
