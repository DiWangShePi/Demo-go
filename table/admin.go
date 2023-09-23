package table

import (
	"Demo/global"
	"Demo/utils"
	"gorm.io/gorm"
)

type BlankAdmin struct {
	gorm.Model
	Name     string `json:"name"`
	PassWord string `json:"passWord"`
}

func (BlankAdmin) TableName() string {
	return "blank_admin"
}

func checkExist() (currentAdmin BlankAdmin, err error) {
	result := global.GLO_DB.Take(&currentAdmin)
	if result.RowsAffected > 0 {
		err = &global.Error{
			Code:    500,
			Message: "Admin already exist",
		}
	}
	return currentAdmin, err
}

// CreateAdmin method to create one Admin account, one and only one admin
func CreateAdmin(admin BlankAdmin) (err error) {
	if _, err = checkExist(); err != nil {
		return err
	}

	//hashedPassword := argon2.IDKey([]byte(admin.PassWord), []byte(global.GLO_SALT), 1, 64*1024, 4, 32)
	//hexHash := fmt.Sprintf("%x", hashedPassword)
	newAdmin := BlankAdmin{Name: admin.Name, PassWord: utils.EncryptPassWord(admin.PassWord)}
	return global.GLO_DB.Create(&newAdmin).Error
}

// CheckAdmin For there will be only one admin
func CheckAdmin() (checkResult BlankAdmin, err error) {
	if checkResult, err = checkExist(); err != nil {
		return checkResult, nil
	}

	return checkResult, &global.Error{
		Code:    600,
		Message: "There are no admin in the current table",
	}
}

// UpdateAdmin method to update Admin information
func UpdateAdmin(newAdmin BlankAdmin) (err error) {
	currentAdmin, err := checkExist()
	if err == nil {
		return &global.Error{
			Code:    500,
			Message: "There are no admin in the current table",
		}
	}

	updateData := make(map[string]interface{})
	if newAdmin.Name != "" {
		updateData["Name"] = newAdmin.Name
	}
	if newAdmin.PassWord != "" {
		//hashedPassword := argon2.IDKey([]byte(newAdmin.PassWord), []byte(global.GLO_SALT), 1, 64*1024, 4, 32)
		//updateData["PassWord"] = fmt.Sprintf("%x", hashedPassword)
		updateData["PassWord"] = utils.EncryptPassWord(newAdmin.PassWord)
	}

	if len(updateData) == 0 {
		return nil
	}
	result := global.GLO_DB.Model(&BlankAdmin{}).Where("id = ?", currentAdmin.ID).Updates(updateData)
	return result.Error
}

// DeleteAdmin method to delete admin account
func DeleteAdmin() (err error) {
	currentAdmin, err := checkExist()
	if err == nil {
		return &global.Error{
			Code:    500,
			Message: "There are no admin in the current table",
		}
	}

	// real delete
	return global.GLO_DB.Delete(&currentAdmin).Scopes().Error
}
