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

type UserStatus int

const (
	UserStatusActive   UserStatus = 1
	UserStatusInactive UserStatus = 2
)

/**
 * @brief 用户结构体
 */
type User struct {
	UserID int        `json:"userID"` // 用户ID //将ID修改为userID
	Emails string     `json:"emails"` // 用户邮箱
	Phone  string     `json:"phone"`  // 用户电话
	Status UserStatus `json:"status"` // 用户状态
}

/**
 * 根据ID获取用户信息
 * @param id 用户ID
 * @param users 用户列表
 * @returns {User} 用户对象
 */
func GetUserById(id int, users []User) (User, error) {
	for _, user := range users {
		if user.UserID == id {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with ID %d not found", id)
}

// 根据邮箱获取用户信息
func GetUserByEmail(email string, users []User) (User, error) {
	for _, user := range users {
		if user.Emails == email {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with email %s not found", email)
}

// 根据电话获取用户信息
func GetUserByPhone(phone string, users []User) (User, error) {
	for _, user := range users {
		if user.Phone == phone {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with phone %s not found", phone)
}

// 查找非活跃用户
func FindInactiveUsers(users []User) []User {
	var inactiveUsers []User
	for _, user := range users {
		if user.Status == UserStatusInactive {
			inactiveUsers = append(inactiveUsers, user)
		}
	}
	return inactiveUsers
}
