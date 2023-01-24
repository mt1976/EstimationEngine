# **CredentialsAction** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**CredentialsAction** (credentialsaction) |
|Endpoint 	    |**/CredentialsAction...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-user-cog** ()
Friendly Name|**Credentials Actions**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}

* **View** (/CredentialsAction/CredentialsActionView)
* **Edit** (/CredentialsAction/CredentialsActionEdit)
* **Save** (/CredentialsAction/CredentialsActionSave)
* **New** (/CredentialsAction/CredentialsActionNew)








##  Provides







##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **inboxMessages**
SQL Table Key | **MailId**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func CredentialsAction_GetList_impl() (int, []dm.CredentialsAction, error) {return 0, nil, nil}</li><li>func CredentialsAction_GetByID_impl(id string) (int, dm.CredentialsAction, error) {return 0, dm.CredentialsAction{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func CredentialsAction_NewID_impl(rec dm.CredentialsAction) (string) { return rec.ID } </li><li>func CredentialsAction_Delete_impl(id string) (error) {return nil}</li><li>func CredentialsAction_Update_impl(id string,rec dm.CredentialsAction, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|true|||||NH|ID||false|false|false|text||
|**User**|String|true|true|false|false|OL|Credentials|Credentials_Username|Credentials_Id|Y|User||false|false|false|text||
|**Action**|String|true|true|false|false|LL|credentialStates|||Y|Action||false|false|false|text|true|
|**Notes**|String|false|true|false|true|||||Y|Notes||false|false|false|textarea||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentialsAction_core.go |
| code | **adaptor** | /dao/credentialsAction_adaptor.go_template |
| code | **dao** | /dao/credentialsAction_core.go |
| code | **datamodel** | /datamodel/credentialsAction_core.go |
| code | **menu** | /design/menu/credentialsAction.json |
| html | **view** | /CredentialsAction_View.html |
| html | **edit** | /CredentialsAction_Edit.html |
| html | **new** | /CredentialsAction_New.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **24/01/2023** at **13:18:08**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

---
### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
    * ∀ = This lookup has a filter that can be defined in the Data Object
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field