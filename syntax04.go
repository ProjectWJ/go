package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 추출된 난수의 소수 판정 프로그램 v0.6
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	// number = 129
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break // 첫 번째 약수가 발견되면 반복문 즉시 종료
		}
		// fmt.Print(i, " ")
	}

	if isPrime { // 비교 연산자 제거
		fmt.Println(number, "는(은) 소수입니다.")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다.")
	}
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 추출된 난수의 소수 판정 프로그램 v0.4
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 		}
// 		fmt.Print(i, " ")
// 	}

// 	if isPrime { // 비교 연산자 제거
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 추출된 난수의 소수 판정 프로그램 v0.3
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	isPrime := true
// 	number := rand.Intn(150) + 2
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			// count = count + 1
// 		}
// 	}

// 	if isPrime == true {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 추출된 난수의 소수 판정 프로그램 v0.2
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1은 제외한다. 범위는 2~151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 2; i < number; i++ { // 1과 nuber일 때 반복하지 않음
// 		if number%i == 0 {
// 			// count++
// 			count = count + 1
// 		}
// 	}

// 	if count == 0 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // 추출된 난수의 소수 판정 프로그램 v0.1
// // 소수는? 1과 자기 자신만을 약수로 갖는 수(0, 1은 제외)
// func main() {
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 // 0과 1은 제외한다. 범위는 2~151 사이의 수
// 	fmt.Println("임의로 추출된 수 : ", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}

// 	if count == 2 {
// 		fmt.Println(number, "는(은) 소수입니다.")
// 	} else {
// 		fmt.Println(number, "는(은) 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" // 시드 생성용 패키지
// )

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }
