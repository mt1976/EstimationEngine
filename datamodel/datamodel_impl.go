package datamodel

import (
	"github.com/mt1976/ebEstimates/core"
	"github.com/mt1976/ebEstimates/logs"
)

type Lookup_Item struct {
	ID   string
	Name string
}

type FieldProperties struct {
	MsgType         string
	MsgMessage      string
	MsgLog          bool
	MsgGlyph        string
	MsgFeedBackType string
	InputType       string
	InputMessage    string
	InputLog        bool
	InputGlyph      string
	DefaultValue    string
}

func NewFieldProperties() *FieldProperties {
	return &FieldProperties{}
}

func AddFieldMessage(fP FieldProperties, typeStr string, message string, log bool, glyph string) FieldProperties {
	fP.MsgType = typeStr
	fP.MsgMessage = message
	fP.MsgLog = log
	fP.MsgGlyph = glyph
	fP.MsgFeedBackType = "form-helper"
	if fP.MsgType == core.FieldMessage_POSITIVE {
		fP.MsgFeedBackType = "valid-feedback"

	}
	if fP.MsgType == core.FieldMessage_ERROR {
		fP.MsgFeedBackType = "invalid-feedback"
	}
	if log || fP.MsgType == core.FieldMessage_ERROR {
		logs.Warning(message)
	}
	return fP
}

type PageProperties struct {
	Title    string
	Glyph    string
	MsgType  string
	MsgClass string
	MsgTitle string
	MsgBody  string
	MsgLog   bool
	MsgGlyph string
}

func NewPageProperties() *PageProperties {
	return &PageProperties{}
}

func AddPageMessage(typeStr string, class string, title string, body string, log bool, glyph string) *PageProperties {
	return &PageProperties{MsgType: typeStr, MsgClass: class, MsgTitle: title, MsgBody: body, MsgLog: log, MsgGlyph: glyph}
}

type appContext struct {
	OnError    string
	OnSuccess  string
	OnRedirect string
	Content    string
	State      string
}

func AppContext_SetState(ac appContext, in string) appContext {
	ac.State = core.EncodeString(in)
	logs.Information("AppContext_SetState: "+in, ac.State)
	return ac
}

func AppContext_GetState(ac appContext) string {
	logs.Information("AppContext_GetState: "+ac.State, core.DecodeString(ac.State))
	return core.DecodeString(ac.State)
}

func AppContext_OnErrorSet(ac appContext, in string) appContext {
	ac.OnError = core.EncodeString(in)
	logs.Information("AppContext_OnErrorSet: "+in, ac.OnError)
	return ac
}

func AppContext_OnSuccessSet(ac appContext, in string) appContext {
	ac.OnSuccess = core.EncodeString(in)
	logs.Information("AppContext_OnSucessSet: "+in, ac.OnSuccess)
	logs.Information("AppContext_OnSuccessSet: "+ac.OnSuccess, AppContext_OnSuccessGet(ac))
	return ac
}

func AppContext_OnRedirectSet(ac appContext, in string) appContext {
	ac.OnRedirect = core.EncodeString(in)
	logs.Information("AppContext_OnRedirectSet: "+in, ac.OnRedirect)
	return ac
}

func AppContext_ContentSet(ac appContext, in string) appContext {
	ac.Content = core.EncodeString(in)
	logs.Information("AppContext_ContentSet: "+in, ac.Content)
	logs.Information("AppContext_ContentSet: "+in, AppContext_ContentGet(ac))
	return ac
}

func AppContext_ContentGet(ac appContext) string {
	logs.Information("AppContext_ContentGet: "+ac.Content, core.DecodeString(ac.Content))
	return core.DecodeString(ac.Content)
}

func AppContext_OnSuccessGet(ac appContext) string {
	logs.Information("AppContext_OnSuccessGet: "+ac.OnSuccess, core.DecodeString(ac.OnSuccess))
	return core.DecodeString(ac.OnSuccess)
}
