# **ExternalMessage** - Object Definition
##  Information
| Information  | Value  |
|---|---|
|Object         |**ExternalMessage** (externalmessage) |
|Endpoint 	    |**/ExternalMessage...** [^1]|
|Endpoint Query |**Message**|
|REST API|**/API/ExternalMessage/**|
Glyph|**fas fa-mail-bulk** (text-info)
Friendly Name|**ExternalMessage**|
|For Project    |github.com/mt1976/ebEstimates/|

##  Actions {#action-id}
* **List** (/ExternalMessage/ExternalMessageList) [Exportable]
* **View** (/ExternalMessage/ExternalMessageView)
* **Edit** (/ExternalMessage/ExternalMessageEdit)
* **Save** (/ExternalMessage/ExternalMessageSave)

* **Delete** (/ExternalMessage/ExternalMessageDelete)







##  Provides


* Auditing 




##  Data Source 
| Information  | Value  |
|---|---|
SQL Table Name       | **externalMessageStore**
SQL Table Key | **MessageID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**MessageID**|String|false|true|false|true|||||N|messageID||false|false|false|text||
|**MessageFormat**|String|false|true|false|true|||||N|messageFormat||false|false|false|text||
|**MessageDeliveredTo**|String|false|true|false|true|||||N|messageDeliveredTo||false|false|false|text||
|**MessageBody**|String|false|true|false|true|||||N|messageBody||false|false|false|textarea||
|**MessageFilename**|String|false|true|false|true|||||N|messageFilename||false|false|false|text||
|**MessageLife**|String|false|true|false|true|||||N|messageLife||false|false|false|number||
|**MessageDate**|String|false|true|false|true|||||N|messageDate||false|false|false|date|yyyy-mm-dd|
|**MessageTime**|String|false|true|false|true|||||N|messageTime||false|false|false|text||
|**MessageTimeoutAction**|String|false|true|false|true|||||N|messageTimeoutAction||false|false|false|text||
|**MessageACKNAK**|String|false|true|false|false|LL|messageACKNAK|||Y|messageACKNAK||false|false|false|text||
|**ResponseID**|String|false|true|false|false|||||Y|responseID||false|false|false|text||
|**ResponseFilename**|String|false|true|false|false|||||Y|responseFilename||false|false|false|text||
|**ResponseBody**|String|false|true|false|true|||||Y|responseBody||false|false|false|textarea||
|**ResponseDate**|String|false|true|false|true|||||Y|responseDate||false|false|false|date|yyyy-mm-dd|
|**ResponseTime**|String|false|true|false|true|||||Y|responseTime||false|false|false|time||
|**ResponseAdditionalInfo**|String|false|true|false|true|||||Y|responseAdditionalInfo||false|false|false|textarea||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**MessageTimeout**|String|false|true|false|true|||||N|messageTimeout||false|false|false|datetime||
|**MessageEmitted**|String|false|true|false|true|||||N|messageEmitted||false|false|false|text||
|**ResponseRecieved**|String|false|true|false|false|LL|tf|||Y|responseRecieved||false|false|false|text||
|**MessageClass**|String|false|true|false|true|||||N|messageClass||false|false|false|text||
|**AppID**|String|false|true|false|true|||||NH|appID||false|false|false|text||
|**SYSDeleted**|String|false|true|false|false|||||NH|_deleted||false|false|true|text||
|**SYSDeletedBy**|String|false|true|false|false|||||NH|_deletedBy||false|false|true|text||
|**SYSDeletedHost**|String|false|true|false|false|||||NH|_deletedHost||false|false|true|text||
|**SYSDbVersion**|String|false|true|false|false|||||NH|_dbVersion||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/externalMessage_core.go |
| code | **validation** | /dao/externalMessage_validation.go_template |
| code | **api** | /application/externalMessage_api.go |
| code | **dao** | /dao/externalMessage_core.go |
| code | **datamodel** | /datamodel/externalMessage_core.go |
| code | **menu** | /design/menu/externalMessage.json |
| html | **list** | /ExternalMessageList.html |
| html | **view** | /ExternalMessageView.html |
| html | **edit** | /ExternalMessageEdit.html |


## Audit Information
| Information  | Value |
|---|---|
Template Generator Version   | **Einsteinium [r5-23.01.23]**
Date & Time		     | **13/03/2023** at **14:22:27**
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