package main

import (
	"fmt"
	"os"

	"booking-api/internal/config"
	"booking-api/internal/handlers"
	"booking-api/internal/models"
	"booking-api/internal/repository"
	"booking-api/internal/router"
	"booking-api/internal/service"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Запуск сервера...")

	cfg := config.LoadConfig()

	dsn := "host=" + cfg.DBHost +
		" port=" + cfg.DBPort +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" sslmode=disable"

	fmt.Println("Подключение к базе")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка подключения к базе:", err)
		os.Exit(1)
	}
	fmt.Println("База подключена")

	err = db.AutoMigrate(&models.Room{}, &models.Booking{})
	if err != nil {
		fmt.Println("Ошибка при создании таблиц:", err)
		os.Exit(1)
	}
	fmt.Println("Таблицы готовы")

	roomRepo := repository.NewRoomRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	roomService := service.NewRoomService(roomRepo)
	bookingService := service.NewBookingService(bookingRepo, roomRepo)

	roomHandler := handlers.NewRoomHandler(roomService)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	r := router.SetupRouter(roomHandler, bookingHandler)

	fmt.Println("Сервер запущен на порту " + cfg.AppPort)
	err = r.Run(":" + cfg.AppPort)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
