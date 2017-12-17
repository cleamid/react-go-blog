package router

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/cleamid/react-go-blog/db"
    "strconv"
    "net/http"
)

func getUser(c *gin.Context) {
    param := c.Query("id")
    id, err := strconv.ParseUint(param, 10, 64)
    if err != nil {
        fmt.Println(0)
        c.JSON(http.StatusNotAcceptable, gin.H{"err": err.Error()})
        fmt.Println(1)
    }
    fmt.Println(2)
    user, err := db.DB.UserTable.SelectById(id)
    if user != nil {
        //c.JSON(200, gin.H{"id": user.Id, "email": user.Email, "name": user.Username,})
        c.JSON(200, user)
    } else {
        c.JSON(http.StatusOK, gin.H{})
    }
}