package application

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/dao"
	dm "github.com/mt1976/ebEstimates/datamodel"
	"github.com/mt1976/ebEstimates/logs"
)

const GET = "GET"
const PUT = "PUT"
const VAL = "VALIDATE"
const NEW = "NEW"

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
func PageTitle(pageTitle string, pageSubTitle string) string {

	pt := dao.Translate("Page", pageTitle)
	pst := dao.Translate("Action", pageSubTitle)
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

func CardTitle(cardTitle string, pageSubTitle string) string {
	return dao.Translate("CardTitle", cardTitle+" - "+pageSubTitle)
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
		logs.Warning("Error executing template: " + err.Error())
	}
}

func ExecuteRedirect(inRoute string, w http.ResponseWriter, r *http.Request, thisID string, thisQuery string, data interface{}) {
	// //fmt.Printf("tname: %v\n", tname)
	//fmt.Printf("inRoute: %v\n", inRoute)
	//fmt.Printf("inRoute: %v\n", thisID)
	//fmt.Printf("thisID: %v\n", thisID)
	//fmt.Printf("thisQuery: %v\n", thisQuery)

	path := firstDir(inRoute)

	//fmt.Printf("path: %v\n", path)

	tname := core.URI_SEP + path + core.URI_SEP + core.URI_QUERY + thisID + "=" + thisQuery
	errText := core.ContextState + "=" + core.ContextState_ERROR
	if !strings.Contains(path, errText) {
		//	path = path + core.ExceptionHandlingPageSuffix
		tname = tname + "&" + errText
	}
	//fmt.Printf("tname: %v\n", tname)
	logs.Redirecting(tname)
	//fmt.Printf("data: %v\n", data)
	gob.Register(data)
	core.SessionManager.Put(r.Context(), thisQuery, data)

	http.Redirect(w, r, tname, http.StatusSeeOther)
}

func addActivity(in string, what string, r *http.Request) string {
	out := addActivity_ForUser(in, what, Session_GetUserName(r))
	return out
}

func addActivity_System(in string, what string) string {
	out := addActivity_ForUser(in, what, dao.Audit_Host())
	return out
}

func addActivity_ForUser(in string, what string, un string) string {
	return core.AddActivity_ForUser(in, what, un)
}

func getCheckBoxValue(field string, r *http.Request) string {
	testField := r.FormValue(field)
	if testField == "on" {
		return core.TRUE
	} else {
		return core.FALSE
	}
}

func firstDir(input_url string) string {

	u, err := url.Parse(input_url)
	if err != nil {
		logs.Warning(err.Error())
	}

	// fmt.Printf("u: %v\n", u)
	// fmt.Printf("u.Host: %v\n", u.Host)
	// fmt.Printf("u.Path: %v\n", u.Path)
	// fmt.Printf("u.RawQuery: %v\n", u.RawQuery)
	// fmt.Printf("u.Fragment: %v\n", u.Fragment)
	// fmt.Printf("u.Scheme: %v\n", u.Scheme)

	workingPath := u.Path
	if workingPath == "" {
		return ""
	}
	noSlashs := strings.Count(workingPath, "/")
	//fmt.Printf("noSlashs: %v\n", noSlashs)
	if noSlashs == 0 {
		return "/" + workingPath
	}
	if noSlashs == 1 {
		return workingPath
	}

	rtnValue := workingPath[1 : strings.Index(workingPath[1:], "/")+1]

	return rtnValue
}

func NextTemplate(source string, action string, defaultTemplate string) string {
	logs.Warning("NextTemplate: " + source + " " + action + " " + defaultTemplate)
	nextTemplate, err := dao.Data_Get(source, action, dm.Data_Category_NextAction)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("nextTemplate: %v\n", nextTemplate)
	if nextTemplate == "" {
		fmt.Printf("defaultTemplate: %v\n", defaultTemplate)
		return defaultTemplate
	}
	fmt.Printf("nextTemplate: %v\n", nextTemplate)
	return nextTemplate
}
