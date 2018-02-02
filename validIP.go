package validIP

import (
	"strconv"
	"strings"
)

/*func main() {
	var input string
	fmt.Scanln(&input)
	if isValidIP(input) {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong!")
	}
	//fmt.Println(removeRandomChar(input))
}*/

func isValidIP(ip string) bool {
	ip_slice := strings.Split(ip, ".")
	if len(ip_slice) == 4 {
		for i := 0; i < 4; i++ {
			if len(ip_slice[i]) > 1 && (ip_slice[i][0] == '0' && ip_slice[i][1] == '0') {
				return false
			} else if bit, err := strconv.Atoi(ip_slice[i]); err != nil {
				return false
			} else if bit < 0 || bit > 255 {
				return false
			}
		}
		return true
	}
	return false
}

func removeRandomChar(str string) string {
	var return_str = ""
	for i := 0; i < len(str); i++ {
		if (str[i]-'a' > 25 || str[i]-'a' < 0) && (str[i]-'A' > 25 || str[i]-'A' < 0) && (str[i]-'0' > 9 || str[i]-'0' < 0) {
			continue
		}
		return_str += string(str[i])
	}
	return return_str
}
