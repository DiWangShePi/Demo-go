package initialize

import (
	_ "Demo/docs"
	"Demo/global"
	"Demo/router"
	"Demo/table"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var configFile = flag.String("config", "config.json", "Configuration file for server. Defaults to config.json.")

type DbConifg struct {
	DbType   string `json:"DataBaseType"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	User     string `json:"User"`
	PassWord string `json:"PassWord"`
	DbName   string `json:"DbName"`
	// Admin information
	AdminName     string `json:"AdminName"`
	AdminPassword string `json:"AdminPassWord"`
	// Encrypt information
	Salt string `json:"Salt"`
}

func (c DbConifg) Dsn() string {
	return "host=" + c.Host + " " + "user=" + c.User + " " + "password=" + c.PassWord + " " + "dbname=" + c.DbName + " " + "port=" + c.Port
}

//func Initial() {
//	// Init DataBase
//
//	// Init Routers
//	router.AdminRouter{}.InitAdminRouter()
//}

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	router.AdminRouter{}.InitAdminRouter(Router)
	router.UserRouter{}.InitUserRouter(Router)
	return Router
}

func InitDb() {
	// Read config file
	file, err := os.Open(*configFile)
	if err != nil {
		fmt.Println("Error while trying to read config file: ", err)
		os.Exit(0)
	}
	defer file.Close()

	// Decode config file
	var dbConfig DbConifg
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dbConfig)
	if err != nil {
		fmt.Println("Error while trying to load config file: ", err)
		os.Exit(0)
	}
	global.GLO_SALT = dbConfig.Salt

	// Check database
	if dbConfig.DbType != "postgresql" {
		fmt.Println("We need to use PostgreSQL")
		os.Exit(0)
	}

	// Connect to database
	pgsqlConfig := postgres.Config{
		DSN:                  dbConfig.Dsn(),
		PreferSimpleProtocol: false,
	}
	global.GLO_DB, err = gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("Error while trying to connect to databese: ", err)
		os.Exit(0)
	}

	// Create table
	if err := global.GLO_DB.AutoMigrate(
		&table.BlankUser{},
		&table.BlankAdmin{},
	); err != nil {
		fmt.Println("Error while trying to auto migrate all tables: ", err)
		os.Exit(0)
	}

	result := global.GLO_DB.Take(&table.BlankAdmin{})
	if result.RowsAffected > 0 {
		fmt.Println("Admin already exists")
	} else {
		// Insert Admin data, if any
		if (dbConfig.AdminName != "") && (dbConfig.AdminPassword != "") {
			err = table.CreateAdmin(table.BlankAdmin{
				Name:     dbConfig.AdminName,
				PassWord: dbConfig.AdminPassword,
			})
			if err != nil {
				fmt.Println("Error while trying to insert admin info: ", err)
			}
		}
	}
}
