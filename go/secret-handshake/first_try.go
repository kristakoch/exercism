package secret

// import (
// 	"math"
// )

// var hscodes = []string{
// 	"wink",
// 	"double blink",
// 	"close your eyes",
// 	"jump",
// 	"reverse",
// }

// // Handshake ...
// func Handshake(num uint) []string {
// 	var handshake []string

// 	if num == 0 {
// 		return handshake
// 	}

// 	binary := convert(num)

// 	for i, bin := range binary {
// 		if string(bin) == "1" {
// 			if i != 4 {
// 				handshake = append(handshake, hscodes[i])
// 				continue
// 			}
// 			handshake = reverse(handshake)
// 		}
// 		if i > 5 {
// 			break
// 		}
// 	}

// 	return handshake
// }

// func convert(num uint) []string {
// 	n := float64(num)
// 	binary := make([]string, int(math.Log2(n))+1)

// 	for i := range binary {
// 		binary[i] = "0"
// 	}

// 	for n > 0 {
// 		log := math.Floor(math.Log2(n))
// 		binary[int(log)] = "1"

// 		n = n - math.Pow(2, log)

// 		if n == 1 {
// 			binary[0] = "1"
// 			break
// 		}
// 	}

// 	return binary
// }

// func reverse(ss []string) []string {
// 	last := len(ss) - 1
// 	for i := 0; i < len(ss)/2; i++ {
// 		ss[i], ss[last-i] = ss[last-i], ss[i]
// 	}

// 	return ss
// }
