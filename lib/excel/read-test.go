package excel

// read excel and print cell contents
import (
	"fmt"
	"io/ioutil"
)

// Test read test
func Test() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
