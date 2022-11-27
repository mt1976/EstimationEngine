package application

// Translation_Lookup  will translate a string from its source version to user-defined version.
func Translation_Lookup(class string, message string) string {
	// translation_ID := dao.Translation_BuildID(class, message)
	// //log.Println("translationStoreID", translationStoreID)
	// _, t, err := dao.Translation_GetByID(translation_ID)
	// //spew.Dump(translationItem)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// // Check if tis null
	// if (dm.Translation{}) == t {
	// 	// Create One
	// 	t.Id = translation_ID
	// 	t.Class = class
	// 	t.Message = message
	// 	t.Translation = message
	// 	dao.Translation_StoreSystem(t)
	// }
	// //fmt.Printf("message: %v\n", message)
	// return t.Translation
	return message
}
