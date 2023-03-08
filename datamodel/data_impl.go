package datamodel

const (
	//Define Constants for the Class
	Data_Category_Setting    = "Setting"
	Data_Category_Data       = "Data"
	Data_Category_Database   = "Database"
	Data_Category_User       = Data_Category_Setting + ".User"
	Data_Category_FilterRule = Data_Category_Setting + ".Rule.Filter"
	Data_Category_StateRule  = Data_Category_Setting + ".Rule.State"
	Data_Category_Importer   = Data_Category_Setting + ".Importer"
	Data_Category_Indexer    = Data_Category_Setting + ".Index.URI"
	Data_Category_URI        = Data_Category_Setting + ".URI"
	Data_Category_Path       = Data_Category_Setting + ".Path"
	Data_Category_Info       = Data_Category_Data + ".Info"
	Data_Category_Sequence   = Data_Category_Data + ".Sequence"
	Data_Category_State      = Data_Category_Data + ".State"
	Data_Category_NextAction = Data_Category_Setting + ".Next.Action"
)

var Data_Category_List = []string{
	Data_Category_Setting,
	Data_Category_Data,
	Data_Category_Database,
	Data_Category_User,
	Data_Category_FilterRule,
	Data_Category_StateRule,
	Data_Category_Importer,
	Data_Category_Indexer,
	Data_Category_URI,
	Data_Category_Path,
	Data_Category_Info,
	Data_Category_Sequence,
	Data_Category_State,
	Data_Category_NextAction,
}
