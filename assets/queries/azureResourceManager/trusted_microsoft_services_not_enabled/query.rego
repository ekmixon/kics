package Cx

CxPolicy[result] {
	doc := input.document[i]
	[path, value] = walk(doc)

	value.type == "Microsoft.Storage/storageAccounts"

	value.properties.networkAcls.defaultAction == "Deny"
	not contains_azure_service(value.properties.networkAcls.bypass)

	result := {
		"documentId": input.document[i].id,
		"searchKey": "resources.type={{Microsoft.Storage/storageAccounts}}.properties.networkAcls",
		"issueType": "IncorrectValue",
		"keyExpectedValue": "resource with type 'Microsoft.Storage/storageAccounts' has 'Trusted Microsoft Services' enabled",
		"keyActualValue": "resource with type 'Microsoft.Storage/storageAccounts' doesn't have 'Trusted Microsoft Services' enabled",
	}
}

contains_azure_service(bypass) {
	values := split(bypass, ",")
	some j
	values[j] == "AzureServices"
}
