package main

import (
	"net/http"
	"rest-api/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application)createEvent(c *gin.Context){
	var event database.Event
	if err:=c.ShouldBindJSON(&event); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"Error":err.Error()})
		return
	}
	err:=app.model.Events.Insert(&event)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create the event"})
		return
	}
	c.JSON(http.StatusCreated,event)

}

func (app *application)getAllEvent(c *gin.Context){
	events,err:=app.model.Events.GetAll()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the event"})
		return
	}

	c.JSON(http.StatusOK,events)
}

func (app *application) getEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid event Id"})
		return
	}
	event,err:=app.model.Events.Get(id)
	if event==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
		return
	}
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive event"})
		return
	}
	c.JSON(http.StatusOK,event)
}

func (app *application) updateEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid event id"})
		return
	}
	existingEvent,err:=app.model.Events.Get(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive event"})
		return
	}
	if existingEvent==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
		return
	}

	updatedEvent:= &database.Event{}

	if err:=c.ShouldBindJSON(updatedEvent);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	updatedEvent.Id=id

	if err:=app.model.Events.Update(updatedEvent);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to update event"})
		return
	}

	c.JSON(http.StatusOK,updatedEvent)
}

func (app *application)deleteEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid event Id"})
		return
	}
	if err:=app.model.Events.Delete(id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to delete the Event"})
		return
	}
	c.JSON(http.StatusNoContent,nil)
}