package seed

import (
	"fmt"
	"mh_api_go/database"
	"mh_api_go/models"
)

func ResetDatabase() {
	// Drop all tables
	database.DB.Migrator().DropTable(
		&models.User{},
		&models.Client{},
		&models.Location{},
		&models.ServiceOrder{},
		&models.Tool{},
	)

	fmt.Println("Tables dropped successfully!")
}

func SeedDatabase() {

	ResetDatabase()

	// Auto Migrate the User model
	database.DB.AutoMigrate(&models.User{}, &models.Client{}, &models.Location{}, &models.ServiceOrder{}, &models.Tool{})

	// Insert a new User
	database.DB.Create(&models.User{Name: "Admin", Email: "admin@example.com"})

	// Insert a new Client
	client := &models.Client{Name: "Mh Client", PhoneNumber: "1234567890", Metadata: map[string]any{"whatsapp": true, "test": "test"}}
	database.DB.Create(client)

	// Insert a new Service Order
	service_order := &models.ServiceOrder{ClientID: client.ID, Details: "Repair the tools"}
	database.DB.Create(service_order)

	// Insert a new Tool
	tool := &models.Tool{Model: "HP1350", ToolType: "Taladro", Brand: "Makita", Location: "Location 1", ServiceOrderID: service_order.ID}
	database.DB.Create(tool)

	fmt.Println("Database connected & migrated!")
}
