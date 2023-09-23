package services

import (
	"Demo/table"
	"Demo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// AdminLogin 管理员登录
// @Summary Admin Login
// @Description Verify an admin account
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Admin's name"
// @Param password formData string true "Admin's password"
// @Success 200 {object} LoginResponse  "Login Success"
// @Failure 200 {object} Response "Error"
// @Router /padmin/login [post]
func AdminLogin(c *gin.Context) {
	tableAdmin, err := table.CheckAdmin()
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	adminName := c.PostForm("name")
	if tableAdmin.Name != adminName {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": "Admin name not found in the database",
		})
		return
	}

	adminPassWord := c.PostForm("password")
	if tableAdmin.PassWord != utils.EncryptPassWord(adminPassWord) {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": "Admin password not correct",
		})
		return
	}

	claims := utils.CreateClaims(utils.BaseClaims{
		Identity: "admin",
		ID:       "1",
		Username: adminName,
	})
	token, err := utils.CreateToken(claims)
	if err != nil {
		fmt.Println("Error")
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Login Success",
		"token":   token,
	})
	return
}

// AdminUpdateAdmin 更新管理员信息
// @Summary Update an admin's info
// @Description Update an admin's name and password
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "Admin's name"
// @Param password formData string true "Admin's password"
// @Param token formData string true "Admin's current token"
// @Success 200 {object} Response "Modify Success"
// @Failure 200 {object} Response "Error"
// @Router /admin/update [post]
func AdminUpdateAdmin(c *gin.Context) {
	newAdmin := table.BlankAdmin{
		Name:     c.PostForm("name"),
		PassWord: c.PostForm("password"),
	}
	err := table.UpdateAdmin(newAdmin)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Modify Success",
	})
	return
}

// AdminCreateUser
// @Summary Create a user account
// @Description Create a user account: name, password, studentId and grade
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name formData string true "User's name"
// @Param password formData string true "User's password"
// @Param studentId formData string true "User's ID"
// @Param grade formData string true "User's grade"
// @Param token formData string true "Admin's current token"
// @Success 200 {object} Response "Create Success"
// @Failure 200 {object} Response "Error"
// @Router /admin/createUser [post]
func AdminCreateUser(c *gin.Context) {
	newUser := table.BlankUser{
		Name:      c.PostForm("name"),
		PassWord:  c.PostForm("password"),
		StudentId: c.PostForm("studentId"),
		Grade:     c.PostForm("grade"),
	}

	err := table.CreateUser(newUser)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Create Success",
	})
	return
}

//// AdminUpdateUser
//// @Summary Update a user account
//// @Description Update a user account: name, password, studentId and grade
//// @Tags Admins
//// @Accept application/x-www-form-urlencoded
//// @Produce json
//// @Param name formData string true "User's name"
//// @Param password formData string true "User's password"
//// @Param studentId formData string true "User's ID"
//// @Param grade formData string true "User's grade"
//// @Param token formData string true "Admin's current token"
//// @Success 200 {object} Response "Update Success"
//// @Failure 200 {object} Response "Error"
//// @Router /admin/updateUser [post]
//func AdminUpdateUser(c *gin.Context) {
//	newUser := table.BlankUser{
//		Name:      c.PostForm("name"),
//		PassWord:  c.PostForm("password"),
//		StudentId: c.PostForm("studentid"),
//		Grade:     c.PostForm("grade"),
//	}
//
//	err := table.UpdateUser(newUser)
//	if err != nil {
//		c.JSON(200, gin.H{
//			"status":  "Error",
//			"message": err,
//		})
//		return
//	}
//
//	c.JSON(200, gin.H{
//		"status":  "Success",
//		"message": "Modify Success",
//	})
//	return
//}

// AdminDeleteUser
// @Summary Delete a user account
// @Description Delete a user account
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param studentId formData string true "User's ID"
// @Param token formData string true "Admin's current token"
// @Success 200 {object} Response "delete Success"
// @Failure 200 {object} Response "Error"
// @Router /admin/deleteUser [post]
func AdminDeleteUser(c *gin.Context) {
	StudentId := c.PostForm("studentid")
	err := table.DeleteUser(StudentId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Success",
		"message": "Delete Success",
	})
	return
}

// AdminFindUser
// @Summary Find a user account
// @Description Find a user account
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param studentId query string true "User's ID"
// @Param token query string true "Admin's current token"
// @Success 200 {object} Response "Find Success"
// @Failure 200 {object} Response "Error"
// @Router /admin/findUser [get]
func AdminFindUser(c *gin.Context) {
	StudentId := c.Query("studentid")
	foundUser, err := table.FindUser(StudentId)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":        "Success",
		"message":       "Success in finding User",
		"name":          foundUser.Name,
		"password hash": foundUser.PassWord,
		"student id":    foundUser.StudentId,
		"grade":         foundUser.Grade,
	})
	return
}

// AdminFindUserList
// @Summary Find all user account
// @Description Find all user account
// @Tags Admins
// @Accept application/x-www-form-urlencoded
// @Param token query string true "Admin's current token"
// @Produce json
// @Success 200 {object} Response "Find Success"
// @Failure 200 {object} Response "Error"
// @Router /admin/findUserList [get]
func AdminFindUserList(c *gin.Context) {
	userList, err := table.FindUserList()
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "Error",
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   "Success",
		"message":  "Success in finding User",
		"userlist": userList,
	})
	return
}
