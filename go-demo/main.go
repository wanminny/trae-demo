/**
 * @file main.go
 * @brief 这是一个简单的Go程序，用于打印"Hello, World!"并提供获取网页内容的功能。
 * @author ai
 * @date 2023-10-10
 */
package main

// 这是一个简单的Go程序，用于打印"Hello, World!"。
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * @brief 主函数，用于打印"Hello, World!"。
 *
 * @param None
 * @return None
 */
func main() {
	fmt.Println("Hello, World!")
}

/**
 * @brief 获取指定URL的网页内容
 * @param url 要获取的网页URL
 * @return 网页内容字符串和可能的错误
 */
func GetWebPageContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
