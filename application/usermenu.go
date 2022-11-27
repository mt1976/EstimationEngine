package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	core "github.com/mt1976/ebEstimates/core"
	dm "github.com/mt1976/ebEstimates/datamodel"
)

// UserMenu_Get
func UserMenu_Get(r *http.Request) dm.AppMenuItem {
	session_role := core.SessionManager.GetString(r.Context(), core.SessionRole)
	_, thisMenuStructure := userMenu_Fetch(session_role)
	//fmt.Printf("thisMenuList: %v\n", thisMenuStructure)
	return thisMenuStructure
}

// fetchappMenuData read all employees
func userMenu_Fetch(inRole string) (int, dm.AppMenuItem) {
	file, _ := ioutil.ReadFile(core.GetMenuID("menu", inRole))
	//	fmt.Printf("file: %v\n", file)

	//fmt.Printf("s: %v\n", string(file))
	data := dm.AppMenuItem{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	// for i := 0; i < len(data.SectionTrees); i++ {
	// 	fmt.Printf("i: %v\n", i)
	// 	fmt.Printf("data.MenuStructure[i]: %v\n", data.MenuHeaderText)
	// 	// 	data.MenuItem[i].MenuText = Translation_Lookup("Menu", data.MenuItem[i].MenuText)
	// 	// 	data.MenuItem[i].MenuHeaderText = core.GetAppName()
	// }
	//fmt.Printf("data: %v\n", data)
	//fmt.Printf("file: %v\n", file)

	return len(data.SectionTrees), data
}
