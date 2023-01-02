# **Credentials** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Credentials** (credentials) |
|Endpoint 	    |**/Credentials...** [^1]|
|Endpoint Query |**Id**|
|REST API|**/API/Credentials/**|
Glyph|**fas fa-key** (text-danger)
Friendly Name|**Credentials**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Credentials/CredentialsList) [Exportable]
* **View** (/Credentials/CredentialsView)
* **Edit** (/Credentials/CredentialsEdit)
* **Save** (/Credentials/CredentialsSave)
* **New** (/Credentials/CredentialsNew)
* **Delete** (/Credentials/CredentialsDelete)







##  Provides
 * Lookup (Id Username)

* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **credentialsStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|false|true|false|true|||||N|Id||true|false|false|text||
|**Username**|String|true|true|false|false|||||Y|Username||false|false|false|text||
|**Password**|String|false|true|false|true|||||NH|Password||false|false|false|password||
|**Firstname**|String|false|true|false|false|||||Y|Firstname||false|false|false|text||
|**Lastname**|String|false|true|false|false|||||Y|Lastname||false|false|false|text||
|**Knownas**|String|false|true|false|false|||||Y|Knownas||false|false|false|text||
|**Email**|String|false|true|false|true|||||Y|Email||false|false|false|email||
|**Issued**|String|false|true|false|true|||||N|Issued||false|false|false|datetime||
|**Expiry**|String|false|true|false|true|||||N|Expiry||false|false|false|datetime||
|**RoleType**|String|false|true|false|false|OL|UserRole|||Y|RoleType||false|false|false|text||
|**Brand**|String|false|true|false|true|||||H|Brand||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**State**|String|false|true|false|true|LL|credentialStates|||Y|State||false|true|false|text||
|**Notes**|String|false|true|false|true|||||Y|Notes||false|false|false|textarea||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSActivity**|String|false|true|false|false|||||NH|_activity||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||
|**EmailNotifications**|String|false|true|false|false|||||Y|emailNotifications||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentials_core.go |
| code | **adaptor** | /adaptor/credentials_impl.go_template |
| code | **api** | /application/credentials_api.go |
| code | **dao** | /dao/credentials_core.go |
| code | **datamodel** | /datamodel/credentials_core.go |
| code | **job** | /jobs/credentials_core.go |
| code | **menu** | /design/menu/credentials.json |
| html | **list** | /Credentials_List.html |
| html | **view** | /Credentials_View.html |
| html | **edit** | /Credentials_Edit.html |
| html | **new** | /Credentials_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **Dysprosium [r4-21.12.31]**
Date & Time		     | **02/01/2023** at **15:17:13**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field