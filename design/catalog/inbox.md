# **Inbox** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**Inbox** (inbox) |
|Endpoint 	    |**/Inbox...** [^1]|
|Endpoint Query |**Message**|
|REST API|**/API/Inbox/**|
Glyph|**fas fa-inbox** ()
Friendly Name|**Inbox**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/Inbox/InboxList) [Exportable]
* **View** (/Inbox/InboxView)

* **Save** (/Inbox/InboxSave)

* **Delete** (/Inbox/InboxDelete)







##  Provides


* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **inboxMessages**
SQL Table Key | **MailId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**MailId**|String|false|true|false|true|||||H|MailId||false|false|false|text||
|**MailTo**|String|false|true|false|false|||||Y|MailTo||false|false|false|text||
|**MailFrom**|String|false|true|false|false|||||Y|MailFrom||false|false|false|text||
|**MailSource**|String|false|true|false|true|||||H|MailSource||false|false|false|text||
|**MailSent**|String|false|true|false|true|||||Y|MailSent||false|false|false|datetime|dd/mm/yyy hh:mm|
|**MailUnread**|String|false|true|false|false|LL|tf|||Y|MailUnread||false|false|false|text||
|**MailSubject**|String|false|true|false|false|||||Y|MailSubject||false|false|false|text||
|**MailContent**|String|false|true|false|true|||||Y|MailContent||false|false|false|textarea||
|**MailAcknowledged**|String|false|true|false|false|LL|tf|||Y|MailAcknowledged||false|false|false|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/inbox_core.go |
| code | **validation** | /dao/inbox_validation.go_template |
| code | **api** | /application/inbox_api.go |
| code | **dao** | /dao/inbox_core.go |
| code | **datamodel** | /datamodel/inbox_core.go |
| code | **menu** | /design/menu/inbox.json |
| html | **list** | /InboxList.html |
| html | **view** | /InboxView.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **15/02/2023** at **10:44:44**
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