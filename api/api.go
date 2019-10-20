package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vanigabriel/Projeto-Pismo/entity"
	"github.com/vanigabriel/Projeto-Pismo/usecases"

	"github.com/evalphobia/go-timber/timber"
)

// SetupRouter create router
func SetupRouter(service *usecases.Service, log *timber.Client) *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	v1 := r.Group("/v1")
	{
		v1.POST("/accounts", CreateAccount(service, log))
		v1.PATCH("/accounts/:id", AlterAccount(service, log))
		v1.GET("/accounts/limits", RetriveAccounts(service, log))

		v1.POST("/transactions", InsertTransaction(service, log))
		v1.POST("/payments", InsertPayments(service, log))
		// v1.GET("/transactions")
		// v1.GET("/transactions/account")
	}

	return r
}

// CreateAccount route to creating a new account
func CreateAccount(service *usecases.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var Income Account

		err := c.BindJSON(&Income)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var Acc entity.Account

		Acc.CreditLimit = Income.CreditLimit.Value
		Acc.WithdrawlLimit = Income.WithdrawlLimit.Value

		fmt.Println(Acc)
		created, err := service.AddAccount(&Acc)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, created)

	}
}

// AlterAccount route to patch account
func AlterAccount(service *usecases.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var id string
		id = c.Param("id")
		if len(id) == 0 {
			log.Err("Account ID not informed")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Account ID not informed"})
			return
		}

		idAccont, _ := strconv.Atoi(id)

		var Income Account

		err := c.BindJSON(&Income)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var Acc entity.Account

		Acc.CreditLimit = Income.CreditLimit.Value
		Acc.WithdrawlLimit = Income.WithdrawlLimit.Value

		updated, err := service.UpdateAccount(idAccont, Acc.CreditLimit, Acc.WithdrawlLimit)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, updated)
	}
}

// RetriveAccounts route to get accounts limits
func RetriveAccounts(service *usecases.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		Accounts, err := service.GetAccounts()
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, Accounts)
	}
}

// InsertTransaction route to insert a new transaction
func InsertTransaction(service *usecases.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var Transaction entity.Transaction

		err := c.BindJSON(&Transaction)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.InsertTransaction(&Transaction)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Created"})
	}
}

// InsertPayments route to insert a collection of transactions
func InsertPayments(service *usecases.Service, log *timber.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var Payments []entity.Transaction

		err := c.BindJSON(&Payments)
		if err != nil {
			log.Err(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, p := range Payments {
			err = service.InsertTransaction(&p)
			if err != nil {
				log.Err(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Created"})
	}
}
