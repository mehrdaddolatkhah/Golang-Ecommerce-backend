package model

// type Base struct {
// 	ID uuid.UUID `db:"type:uuid;primary_key;not null;" json:"-"`
// }

// func (base *Base) BeforeCreate(scope *sql.Scope) error {
// 	uuid, err := uuid.NewV4()
// 	if err != nil {
// 		return err
// 	}
// 	return scope.SetColumn("ID", uuid)
// }

// func DBMigrate(db *gorm.DB) *gorm.DB {
// 	db.AutoMigrate(&Product{})
// 	db.AutoMigrate(&User{})
// 	return db
// }
