package application

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	core "github.com/mt1976/ebEstimates/core"
)

var sqlSYSDToday sql.NullString

func setChecked(inValue string) string {
	var outValue string
	outValue = ""
	if inValue == "true" {
		outValue = "checked"
	}
	if inValue == "false" {
		outValue = ""
	}
	return outValue
}

func isChecked(inValue string) string {
	var outValue string
	//fmt.Println(inValue)
	outValue = ""
	if inValue == "true" {
		outValue = "checked"
	}
	if inValue == "false" {
		outValue = ""
	}
	return outValue
}

// PageTitle: core.GetAppName() + core.Character_Break + Translation_Lookup("Page", dm.Translation_Title) + core.Character_Break + Translation_Lookup("Action", "New"),
func PageTitle(
	pageTitle string,
	pageSubTitle string) string {

	pt := Translation_Lookup("Page", pageTitle)
	pst := Translation_Lookup("Action", pageSubTitle)
	appName := core.ApplicationName()
	if len(appName) == 0 {
		appName = "Application Name"
	}
	pageTitle = appName + core.Character_Break + pt
	if len(pst) > 0 {
		pageTitle = appName + core.Character_Break + pt + core.Character_Break + pst
	}

	return pageTitle
}

func CardTitle(
	cardTitle string,
	pageSubTitle string) string {

	pt := Translation_Lookup("CardTitle", cardTitle+" - "+pageSubTitle)

	cardTitle = pt

	return cardTitle
}

func ExecuteTemplate(tname string, w http.ResponseWriter, r *http.Request, data interface{}) {
	//fmt.Printf("tname: %v\n", tname)

	t := make(map[string]*template.Template)
	baseTemplateID := core.GetTemplateID(tname, Session_GetUserRole(r))
	headerTemplateID := core.GetTemplateID("core/imports", Session_GetUserRole(r))

	//fmt.Printf("baseTemplateID: %v\n", baseTemplateID)
	//fmt.Printf("headerTemplateID: %v\n", headerTemplateID)

	t[tname] = template.Must(template.ParseFiles(baseTemplateID, headerTemplateID))
	err := t[tname].Execute(w, data)
	if err != nil {
		log.Println("Error executing template: " + err.Error())
	}
}

func addActivity(in string, what string, r *http.Request) string {
	//fmt.Println("addActivity")
	//fmt.Println(in)
	//fmt.Println("addActivity")
	if what == "" {
		return in
	}

	tm := time.Now().Format("02/01/2006 15:04:05")
	out := in + "\n" + tm + " " + Session_GetUserName(r) + ":" + what
	return out
}
