package router

import (
	"booking-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	roomHandler *handlers.RoomHandler,
	bookingHandler *handlers.BookingHandler,
) *gin.Engine {
	r := gin.Default()

	// админ
	r.POST("/rooms", roomHandler.CreateRoom)
	r.GET("/rooms", roomHandler.GetAllRooms)
	r.GET("/rooms/:id", roomHandler.GetRoom)
	r.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	// пользователь
	r.POST("/bookings", bookingHandler.CreateBooking)
	r.GET("/bookings/:id", bookingHandler.GetBooking)
	r.GET("/bookings", bookingHandler.GetUserBookings)

	return r
}
