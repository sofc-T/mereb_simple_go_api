package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sofc-t/mereb_simple_go_api/internals/models"
	"github.com/sofc-t/mereb_simple_go_api/internals/store"
)

var inMemoryStore = store.NewInMemoryStore()

func GetPersons(c *gin.Context) {
	persons := inMemoryStore.GetAllPersons()
	c.JSON(http.StatusOK, persons)
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")
	person, err := inMemoryStore.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
}

func CreatePerson(c *gin.Context) {
	var newPerson models.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	person := models.NewPerson(newPerson.Name, newPerson.Age, newPerson.Hobbies)
	inMemoryStore.CreatePerson(person)
	c.JSON(http.StatusCreated, person)
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var updatedPerson models.Person
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	updatedPerson.ID = id
	if err := inMemoryStore.UpdatePerson(id, &updatedPerson); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPerson)
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if err := inMemoryStore.DeletePerson(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
}
