package control

import (
	"hw/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Data)
}

func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for _, value := range dto.Data {
		if value.ID == uint(id) {
			c.JSON(http.StatusOK, value)
			break
		}
	}
}

func Post(c *gin.Context) {
	var r dto.Role
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	r.ID = uint(len(dto.Data) + 1)
	dto.Data = append(dto.Data, r)
	c.Status(http.StatusOK)
}
func Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var r dto.Role
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i, val := range dto.Data {
		if val.ID == uint(id) {
			r.ID = val.ID
			dto.Data[i] = r
			break
		}
	}
	c.JSON(http.StatusNoContent, nil)
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	for i, val := range dto.Data {
		if val.ID == uint(id) {
			dto.Data = append(dto.Data[:i], dto.Data[i+1:]...)
			break
		}
	}
	c.JSON(http.StatusNoContent, nil)
}
