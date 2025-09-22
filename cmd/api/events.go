package main

import (
	"fmt"
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

	user:=app.GetUserFromContext(c)
	event.OwnerID=user.Id

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
	user:=app.GetUserFromContext(c)
	existingEvent,err:=app.model.Events.Get(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive event"})
		return
	}
	if existingEvent==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
		return
	}
	if existingEvent.OwnerID!=user.Id{
		c.JSON(http.StatusForbidden,gin.H{"error":"You are not authorized to update this event"})
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

	user:=app.GetUserFromContext(c)
	existingEvent,err:=app.model.Events.Get(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the Event"})
		return
	}
	if existingEvent==nil{
		c.JSON(http.StatusNotFound,gin.H{"Error":"Event not Found"})
		return
	}

	if existingEvent.OwnerID!=user.Id{
		c.JSON(http.StatusForbidden,gin.H{"error":"You are not authorized to delete this event"})
		return
	}

	if err:=app.model.Events.Delete(id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to delete the Event"})
		return
	}
	c.JSON(http.StatusNoContent,nil)
}


// func (app *application)addAttendeetoEvent(c *gin.Context){
// 	eventId,err:=strconv.Atoi(c.Param("id"))
// 	if err!=nil{
// 		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for event"})
// 		return
// 	}
// 	userId,err:=strconv.Atoi(c.Param("userId"))
// 	if err!=nil{
// 		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for the user"})
// 		return
// 	}
// 	event,err:=app.model.Events.Get(eventId)
// 	if err!=nil{
// 		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the event"})
// 		return
// 	}
// 	if event==nil{
// 		c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
// 		return
// 	}
// 	usertoAdd,err:=app.model.Users.Get(userId)
// 	if err!=nil{
// 		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the user"})
// 		return
// 	}
// 	if usertoAdd==nil{
// 		c.JSON(http.StatusNotFound,gin.H{"error":"User not Found"})
// 		return
// 	}
// 	existingAttenddes,err:=app.model.Attenddes.GetByEventAndAttendees(event.Id,usertoAdd.Id)
// 	if err!=nil{
// 		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the user"})
// 		return
// 	}
// 	if existingAttenddes==nil{
// 		c.JSON(http.StatusConflict,gin.H{"error":"Attendees already exists"})
// 		return
// 	}
// 	attendees:=database.Attenddes{
// 		EventId: event.Id,
// 		UserId: usertoAdd.Id,
// 	}
// 	_,err=app.model.Attenddes.Insert(&attendees)
// 	if err!=nil{
// 		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to add Attendees"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated,attendees)

// }

// func (app *application)addAttendeetoEvent(c *gin.Context){
//     eventId,err:=strconv.Atoi(c.Param("id"))
//     if err!=nil{
//         c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for event"})
//         return
//     }
//     userId,err:=strconv.Atoi(c.Param("userId"))
//     if err!=nil{
//         c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for the user"})
//         return
//     }
    
//     event,err:=app.model.Events.Get(eventId)
//     if err!=nil{
//         c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the event"})
//         return
//     }
//     if event==nil{
//         c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
//         return
//     }
    
//     usertoAdd,err:=app.model.Users.Get(userId)
//     if err!=nil{
//         c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the user"})
//         return
//     }
//     if usertoAdd==nil{
//         c.JSON(http.StatusNotFound,gin.H{"error":"User not Found"})
//         return
//     }
    
//     // FIX: Check if attendee already exists - CORRECTED LOGIC
//     existingAttenddes,err:=app.model.Attenddes.GetByEventAndAttendees(event.Id,usertoAdd.Id)
//     if err!=nil{
//         c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to check existing attendee"})
//         return
//     }
//     if existingAttenddes != nil{  // Changed from == nil to != nil
//         c.JSON(http.StatusConflict,gin.H{"error":"Attendee already exists"})
//         return
//     }
    
//     attendees:=database.Attenddes{
//         EventId: event.Id,
//         UserId: usertoAdd.Id,
//     }
    
//     _,err=app.model.Attenddes.Insert(&attendees)
//     if err!=nil{
//         c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to add Attendee"})
//         return
//     }
//     c.JSON(http.StatusCreated,attendees)
// }

func (app *application)addAttendeetoEvent(c *gin.Context){
    eventId,err:=strconv.Atoi(c.Param("id"))
    if err!=nil{
        c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for event"})
        return
    }
    userId,err:=strconv.Atoi(c.Param("userId"))
    if err!=nil{
        c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid ID for the user"})
        return
    }

    
    event,err:=app.model.Events.Get(eventId)
    if err!=nil{
        fmt.Printf("Event retrieval error: %v\n", err)
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the event"})
        return
    }
    if event==nil{
        fmt.Printf("Event not found: %d\n", eventId)
        c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
        return
    }
    fmt.Printf("Event found: %+v\n", event)

    usertoAdd,err:=app.model.Users.Get(userId)
    if err!=nil{
        fmt.Printf("User retrieval error: %v\n", err)
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the user"})
        return
    }
    if usertoAdd==nil{
        fmt.Printf("User not found: %d\n", userId)
        c.JSON(http.StatusNotFound,gin.H{"error":"User not Found"})
        return
    }


	user:=app.GetUserFromContext(c)
	if event.OwnerID!=user.Id{
		c.JSON(http.StatusForbidden,gin.H{"error":"Your not authorized to add the attendee"})
		return
	}

    existingAttenddes,err:=app.model.Attenddes.GetByEventAndAttendees(event.Id,usertoAdd.Id)
    if err!=nil{
        fmt.Printf("Attendee check ERROR: %v\n", err)
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to check existing attendee"})
        return
    }
    
    if existingAttenddes != nil{
        fmt.Printf("Attendee already exists: %+v\n", existingAttenddes)
        c.JSON(http.StatusConflict,gin.H{"error":"Attendee already exists"})
        return
    }
    fmt.Printf("No existing attendee found, creating new one...\n")
    
    attendees:=database.Attenddes{
        EventId: event.Id,
        UserId: usertoAdd.Id,
    }
    
    _,err=app.model.Attenddes.Insert(&attendees)
    if err!=nil{
        fmt.Printf("Insertion error: %v\n", err)
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to add Attendee"})
        return
    }
    
    fmt.Printf("=== END addAttendeetoEvent - SUCCESS ===\n")
    c.JSON(http.StatusCreated,attendees)
}


func (app *application)getAttendeeForEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Id inserted"})
		return 
	}
	users,err:=app.model.Attenddes.GetAttenddesByEvent(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive attendee for the event"})
		return	
	}
	c.JSON(http.StatusOK,users)

}

func (app *application)deleteAttendeeFromEvent(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Event ID"})
		return
	}
	userId,err:=strconv.Atoi(c.Param("userId"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid User ID"})
		return
	}

	event,err:=app.model.Events.Get(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Something went wrong"})
		return
	}	
	if event==nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Event not found"})
		return
	}

	user:=app.GetUserFromContext(c)
	if event.OwnerID!=user.Id{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Your not authorized to delete the attendee for the event"})
	}

	err=app.model.Attenddes.Delete(userId,id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to delete Attendee"})
		return
	}

	c.JSON(http.StatusNoContent,nil)
}

func (app *application)getEventsByAttendee(c *gin.Context){
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid Attendee ID"})
		return
	}
	events,err:=app.model.Attenddes.GetEventsByAttendees(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to retrive the event"})
		return
	}
	c.JSON(http.StatusOK,events)
}