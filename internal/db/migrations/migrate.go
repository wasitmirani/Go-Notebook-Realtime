package migrations

import(

	"gorm.io/gorm"
	"GoNotebookRealtime/internal/app/models"
)


func RunMigrations(db *gorm.DB) {
	logs.Println("Running database migrations...")
	err := db.AutoMigrate(
		&models.User{},
		&models.Note{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully ✅")

}