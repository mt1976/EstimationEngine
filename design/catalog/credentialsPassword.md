# **CredentialsPassword** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**CredentialsPassword** (credentialspassword) |
|Endpoint 	    |**/CredentialsPassword...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-user-cog** ()
Friendly Name|**Credentials Password**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}

* **View** (/CredentialsPassword/CredentialsPasswordView)
* **Edit** (/CredentialsPassword/CredentialsPasswordEdit)
* **Save** (/CredentialsPassword/CredentialsPasswordSave)
* **New** (/CredentialsPassword/CredentialsPasswordNew)








##  Provides







##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **inboxMessages**
SQL Table Key | **MailId**
Fetch|<ul><li>**Implement in Adaptor**</li><li> func CredentialsPassword_GetList_impl() (int, []dm.CredentialsPassword, error) {return 0, nil, nil}</li><li>func CredentialsPassword_GetByID_impl(id string) (int, dm.CredentialsPassword, error) {return 0, dm.CredentialsPassword{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func CredentialsPassword_NewID_impl(rec dm.CredentialsPassword) (string) { return rec.ID } </li><li>func CredentialsPassword_Delete_impl(id string) (error) {return nil}</li><li>func CredentialsPassword_Update_impl(id string,rec dm.CredentialsPassword, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|true|||||H|ID||false|false|false|text||
|**UserName**|String|true|true|false|true|||||Y|UserName||false|true|false|text||
|**PasswordOld**|String|true|true|false|true|||||Y|PasswordOld||false|true|false|password||
|**PasswordNew**|String|true|true|false|true|||||Y|PasswordNew||false|true|false|password||
|**PasswordConfirm**|String|true|true|false|true|||||Y|PasswordConfirm||false|true|false|password||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentialsPassword_core.go |
| code | **adaptor** | /dao/credentialsPassword_adaptor.go_template |
| code | **validation** | /dao/credentialsPassword_validation.go_template |
| code | **dao** | /dao/credentialsPassword_core.go |
| code | **datamodel** | /datamodel/credentialsPassword_core.go |
| code | **job** | /jobs/credentialsPassword_core.go |
| code | **menu** | /design/menu/credentialsPassword.json |
| html | **view** | /CredentialsPasswordView.html |
| html | **edit** | /CredentialsPasswordEdit.html |
| html | **new** | /CredentialsPasswordNew.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **08/03/2023** at **18:42:23**
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