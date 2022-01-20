package order

// func TestProductRepo(t *testing.T) {
// 	config := configs.GetConfig()
// 	db := utils.InitDB(config)

// 	db.AutoMigrate(&entities.Order{})

// 	orderRepo := NewOrderRepo(db)

// 	db.Migrator().DropTable(&entities.Order{})
// 	t.Run("Error Delete User", func(t *testing.T) {
// 		_, err := orderRepo.Delete(1, 1)
// 		assert.Error(t, err)

// 	})
// 	t.Run("Error Update User", func(t *testing.T) {
// 		var mockProduct entities.Order
// 		_, err := orderRepo.Update(mockProduct, 1, 1)
// 		assert.Error(t, err)

// 	})
// 	t.Run("Error Select Product from Database", func(t *testing.T) {
// 		_, err := orderRepo.Get(1, 1)
// 		assert.Error(t, err)

// 	})
// 	t.Run("Error Select Product from Database", func(t *testing.T) {
// 		_, err := orderRepo.GetAll(1)
// 		assert.Error(t, err)

// 	})
// 	db.AutoMigrate(&entities.Order{})

// }
