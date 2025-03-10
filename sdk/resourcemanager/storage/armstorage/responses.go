// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

// AccountsClientAbortHierarchicalNamespaceMigrationResponse contains the response from method AccountsClient.BeginAbortHierarchicalNamespaceMigration.
type AccountsClientAbortHierarchicalNamespaceMigrationResponse struct {
	// placeholder for future response values
}

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	// The CheckNameAvailability operation response.
	CheckNameAvailabilityResult
}

// AccountsClientCreateResponse contains the response from method AccountsClient.BeginCreate.
type AccountsClientCreateResponse struct {
	// The storage account.
	Account
}

// AccountsClientCustomerInitiatedMigrationResponse contains the response from method AccountsClient.BeginCustomerInitiatedMigration.
type AccountsClientCustomerInitiatedMigrationResponse struct {
	// placeholder for future response values
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.Delete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientFailoverResponse contains the response from method AccountsClient.BeginFailover.
type AccountsClientFailoverResponse struct {
	// placeholder for future response values
}

// AccountsClientGetCustomerInitiatedMigrationResponse contains the response from method AccountsClient.GetCustomerInitiatedMigration.
type AccountsClientGetCustomerInitiatedMigrationResponse struct {
	// The parameters or status associated with an ongoing or enqueued storage account migration in order to update its current
	// SKU or region.
	AccountMigration
}

// AccountsClientGetPropertiesResponse contains the response from method AccountsClient.GetProperties.
type AccountsClientGetPropertiesResponse struct {
	// The storage account.
	Account
}

// AccountsClientHierarchicalNamespaceMigrationResponse contains the response from method AccountsClient.BeginHierarchicalNamespaceMigration.
type AccountsClientHierarchicalNamespaceMigrationResponse struct {
	// placeholder for future response values
}

// AccountsClientListAccountSASResponse contains the response from method AccountsClient.ListAccountSAS.
type AccountsClientListAccountSASResponse struct {
	// The List SAS credentials operation response.
	ListAccountSasResponse
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	// The response from the List Storage Accounts operation.
	AccountListResult
}

// AccountsClientListKeysResponse contains the response from method AccountsClient.ListKeys.
type AccountsClientListKeysResponse struct {
	// The response from the ListKeys operation.
	AccountListKeysResult
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	// The response from the List Storage Accounts operation.
	AccountListResult
}

// AccountsClientListServiceSASResponse contains the response from method AccountsClient.ListServiceSAS.
type AccountsClientListServiceSASResponse struct {
	// The List service SAS credentials operation response.
	ListServiceSasResponse
}

// AccountsClientRegenerateKeyResponse contains the response from method AccountsClient.RegenerateKey.
type AccountsClientRegenerateKeyResponse struct {
	// The response from the ListKeys operation.
	AccountListKeysResult
}

// AccountsClientRestoreBlobRangesResponse contains the response from method AccountsClient.BeginRestoreBlobRanges.
type AccountsClientRestoreBlobRangesResponse struct {
	// Blob restore status.
	BlobRestoreStatus
}

// AccountsClientRevokeUserDelegationKeysResponse contains the response from method AccountsClient.RevokeUserDelegationKeys.
type AccountsClientRevokeUserDelegationKeysResponse struct {
	// placeholder for future response values
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.Update.
type AccountsClientUpdateResponse struct {
	// The storage account.
	Account
}

// BlobContainersClientClearLegalHoldResponse contains the response from method BlobContainersClient.ClearLegalHold.
type BlobContainersClientClearLegalHoldResponse struct {
	// The LegalHold property of a blob container.
	LegalHold
}

// BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse contains the response from method BlobContainersClient.CreateOrUpdateImmutabilityPolicy.
type BlobContainersClientCreateOrUpdateImmutabilityPolicyResponse struct {
	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientCreateResponse contains the response from method BlobContainersClient.Create.
type BlobContainersClientCreateResponse struct {
	// Properties of the blob container, including Id, resource name, resource type, Etag.
	BlobContainer
}

// BlobContainersClientDeleteImmutabilityPolicyResponse contains the response from method BlobContainersClient.DeleteImmutabilityPolicy.
type BlobContainersClientDeleteImmutabilityPolicyResponse struct {
	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientDeleteResponse contains the response from method BlobContainersClient.Delete.
type BlobContainersClientDeleteResponse struct {
	// placeholder for future response values
}

// BlobContainersClientExtendImmutabilityPolicyResponse contains the response from method BlobContainersClient.ExtendImmutabilityPolicy.
type BlobContainersClientExtendImmutabilityPolicyResponse struct {
	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientGetImmutabilityPolicyResponse contains the response from method BlobContainersClient.GetImmutabilityPolicy.
type BlobContainersClientGetImmutabilityPolicyResponse struct {
	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientGetResponse contains the response from method BlobContainersClient.Get.
type BlobContainersClientGetResponse struct {
	// Properties of the blob container, including Id, resource name, resource type, Etag.
	BlobContainer
}

// BlobContainersClientLeaseResponse contains the response from method BlobContainersClient.Lease.
type BlobContainersClientLeaseResponse struct {
	// Lease Container response schema.
	LeaseContainerResponse
}

// BlobContainersClientListResponse contains the response from method BlobContainersClient.NewListPager.
type BlobContainersClientListResponse struct {
	// Response schema. Contains list of blobs returned, and if paging is requested or required, a URL to next page of containers.
	ListContainerItems
}

// BlobContainersClientLockImmutabilityPolicyResponse contains the response from method BlobContainersClient.LockImmutabilityPolicy.
type BlobContainersClientLockImmutabilityPolicyResponse struct {
	// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
	ImmutabilityPolicy

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// BlobContainersClientObjectLevelWormResponse contains the response from method BlobContainersClient.BeginObjectLevelWorm.
type BlobContainersClientObjectLevelWormResponse struct {
	// placeholder for future response values
}

// BlobContainersClientSetLegalHoldResponse contains the response from method BlobContainersClient.SetLegalHold.
type BlobContainersClientSetLegalHoldResponse struct {
	// The LegalHold property of a blob container.
	LegalHold
}

// BlobContainersClientUpdateResponse contains the response from method BlobContainersClient.Update.
type BlobContainersClientUpdateResponse struct {
	// Properties of the blob container, including Id, resource name, resource type, Etag.
	BlobContainer
}

// BlobInventoryPoliciesClientCreateOrUpdateResponse contains the response from method BlobInventoryPoliciesClient.CreateOrUpdate.
type BlobInventoryPoliciesClientCreateOrUpdateResponse struct {
	// The storage account blob inventory policy.
	BlobInventoryPolicy
}

// BlobInventoryPoliciesClientDeleteResponse contains the response from method BlobInventoryPoliciesClient.Delete.
type BlobInventoryPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// BlobInventoryPoliciesClientGetResponse contains the response from method BlobInventoryPoliciesClient.Get.
type BlobInventoryPoliciesClientGetResponse struct {
	// The storage account blob inventory policy.
	BlobInventoryPolicy
}

// BlobInventoryPoliciesClientListResponse contains the response from method BlobInventoryPoliciesClient.NewListPager.
type BlobInventoryPoliciesClientListResponse struct {
	// List of blob inventory policies returned.
	ListBlobInventoryPolicy
}

// BlobServicesClientGetServicePropertiesResponse contains the response from method BlobServicesClient.GetServiceProperties.
type BlobServicesClientGetServicePropertiesResponse struct {
	// The properties of a storage account’s Blob service.
	BlobServiceProperties
}

// BlobServicesClientListResponse contains the response from method BlobServicesClient.NewListPager.
type BlobServicesClientListResponse struct {
	BlobServiceItems
}

// BlobServicesClientSetServicePropertiesResponse contains the response from method BlobServicesClient.SetServiceProperties.
type BlobServicesClientSetServicePropertiesResponse struct {
	// The properties of a storage account’s Blob service.
	BlobServiceProperties
}

// DeletedAccountsClientGetResponse contains the response from method DeletedAccountsClient.Get.
type DeletedAccountsClientGetResponse struct {
	// Deleted storage account
	DeletedAccount
}

// DeletedAccountsClientListResponse contains the response from method DeletedAccountsClient.NewListPager.
type DeletedAccountsClientListResponse struct {
	// The response from the List Deleted Accounts operation.
	DeletedAccountListResult
}

// EncryptionScopesClientGetResponse contains the response from method EncryptionScopesClient.Get.
type EncryptionScopesClientGetResponse struct {
	// The Encryption Scope resource.
	EncryptionScope
}

// EncryptionScopesClientListResponse contains the response from method EncryptionScopesClient.NewListPager.
type EncryptionScopesClientListResponse struct {
	// List of encryption scopes requested, and if paging is required, a URL to the next page of encryption scopes.
	EncryptionScopeListResult
}

// EncryptionScopesClientPatchResponse contains the response from method EncryptionScopesClient.Patch.
type EncryptionScopesClientPatchResponse struct {
	// The Encryption Scope resource.
	EncryptionScope
}

// EncryptionScopesClientPutResponse contains the response from method EncryptionScopesClient.Put.
type EncryptionScopesClientPutResponse struct {
	// The Encryption Scope resource.
	EncryptionScope
}

// FileServicesClientGetServicePropertiesResponse contains the response from method FileServicesClient.GetServiceProperties.
type FileServicesClientGetServicePropertiesResponse struct {
	// The properties of File services in storage account.
	FileServiceProperties
}

// FileServicesClientGetServiceUsageResponse contains the response from method FileServicesClient.GetServiceUsage.
type FileServicesClientGetServiceUsageResponse struct {
	// The usage of file service in storage account.
	FileServiceUsage
}

// FileServicesClientListResponse contains the response from method FileServicesClient.List.
type FileServicesClientListResponse struct {
	FileServiceItems
}

// FileServicesClientListServiceUsagesResponse contains the response from method FileServicesClient.NewListServiceUsagesPager.
type FileServicesClientListServiceUsagesResponse struct {
	// List file service usages schema.
	FileServiceUsages
}

// FileServicesClientSetServicePropertiesResponse contains the response from method FileServicesClient.SetServiceProperties.
type FileServicesClientSetServicePropertiesResponse struct {
	// The properties of File services in storage account.
	FileServiceProperties
}

// FileSharesClientCreateResponse contains the response from method FileSharesClient.Create.
type FileSharesClientCreateResponse struct {
	// Properties of the file share, including Id, resource name, resource type, Etag.
	FileShare
}

// FileSharesClientDeleteResponse contains the response from method FileSharesClient.Delete.
type FileSharesClientDeleteResponse struct {
	// placeholder for future response values
}

// FileSharesClientGetResponse contains the response from method FileSharesClient.Get.
type FileSharesClientGetResponse struct {
	// Properties of the file share, including Id, resource name, resource type, Etag.
	FileShare
}

// FileSharesClientLeaseResponse contains the response from method FileSharesClient.Lease.
type FileSharesClientLeaseResponse struct {
	// Lease Share response schema.
	LeaseShareResponse

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// FileSharesClientListResponse contains the response from method FileSharesClient.NewListPager.
type FileSharesClientListResponse struct {
	// Response schema. Contains list of shares returned, and if paging is requested or required, a URL to next page of shares.
	FileShareItems
}

// FileSharesClientRestoreResponse contains the response from method FileSharesClient.Restore.
type FileSharesClientRestoreResponse struct {
	// placeholder for future response values
}

// FileSharesClientUpdateResponse contains the response from method FileSharesClient.Update.
type FileSharesClientUpdateResponse struct {
	// Properties of the file share, including Id, resource name, resource type, Etag.
	FileShare
}

// LocalUsersClientCreateOrUpdateResponse contains the response from method LocalUsersClient.CreateOrUpdate.
type LocalUsersClientCreateOrUpdateResponse struct {
	// The local user associated with the storage accounts.
	LocalUser
}

// LocalUsersClientDeleteResponse contains the response from method LocalUsersClient.Delete.
type LocalUsersClientDeleteResponse struct {
	// placeholder for future response values
}

// LocalUsersClientGetResponse contains the response from method LocalUsersClient.Get.
type LocalUsersClientGetResponse struct {
	// The local user associated with the storage accounts.
	LocalUser
}

// LocalUsersClientListKeysResponse contains the response from method LocalUsersClient.ListKeys.
type LocalUsersClientListKeysResponse struct {
	// The Storage Account Local User keys.
	LocalUserKeys
}

// LocalUsersClientListResponse contains the response from method LocalUsersClient.NewListPager.
type LocalUsersClientListResponse struct {
	// List of local users requested, and if paging is required, a URL to the next page of local users.
	LocalUsers
}

// LocalUsersClientRegeneratePasswordResponse contains the response from method LocalUsersClient.RegeneratePassword.
type LocalUsersClientRegeneratePasswordResponse struct {
	// The secrets of Storage Account Local User.
	LocalUserRegeneratePasswordResult
}

// ManagementPoliciesClientCreateOrUpdateResponse contains the response from method ManagementPoliciesClient.CreateOrUpdate.
type ManagementPoliciesClientCreateOrUpdateResponse struct {
	// The Get Storage Account ManagementPolicies operation response.
	ManagementPolicy
}

// ManagementPoliciesClientDeleteResponse contains the response from method ManagementPoliciesClient.Delete.
type ManagementPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagementPoliciesClientGetResponse contains the response from method ManagementPoliciesClient.Get.
type ManagementPoliciesClientGetResponse struct {
	// The Get Storage Account ManagementPolicies operation response.
	ManagementPolicy
}

// NetworkSecurityPerimeterConfigurationsClientGetResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.Get.
type NetworkSecurityPerimeterConfigurationsClientGetResponse struct {
	// The Network Security Perimeter configuration resource.
	NetworkSecurityPerimeterConfiguration
}

// NetworkSecurityPerimeterConfigurationsClientListResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.NewListPager.
type NetworkSecurityPerimeterConfigurationsClientListResponse struct {
	// Result of the List Network Security Perimeter configuration operation.
	NetworkSecurityPerimeterConfigurationList
}

// NetworkSecurityPerimeterConfigurationsClientReconcileResponse contains the response from method NetworkSecurityPerimeterConfigurationsClient.BeginReconcile.
type NetworkSecurityPerimeterConfigurationsClientReconcileResponse struct {
	// placeholder for future response values
}

// ObjectReplicationPoliciesClientCreateOrUpdateResponse contains the response from method ObjectReplicationPoliciesClient.CreateOrUpdate.
type ObjectReplicationPoliciesClientCreateOrUpdateResponse struct {
	// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
	ObjectReplicationPolicy
}

// ObjectReplicationPoliciesClientDeleteResponse contains the response from method ObjectReplicationPoliciesClient.Delete.
type ObjectReplicationPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ObjectReplicationPoliciesClientGetResponse contains the response from method ObjectReplicationPoliciesClient.Get.
type ObjectReplicationPoliciesClientGetResponse struct {
	// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
	ObjectReplicationPolicy
}

// ObjectReplicationPoliciesClientListResponse contains the response from method ObjectReplicationPoliciesClient.NewListPager.
type ObjectReplicationPoliciesClientListResponse struct {
	// List storage account object replication policies.
	ObjectReplicationPolicies
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Storage operations. It contains a list of operations and a URL link to get the next set of
	// results.
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListResponse contains the response from method PrivateEndpointConnectionsClient.NewListPager.
type PrivateEndpointConnectionsClientListResponse struct {
	// List of private endpoint connection associated with the specified storage account
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientPutResponse contains the response from method PrivateEndpointConnectionsClient.Put.
type PrivateEndpointConnectionsClientPutResponse struct {
	// The Private Endpoint Connection resource.
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListByStorageAccountResponse contains the response from method PrivateLinkResourcesClient.ListByStorageAccount.
type PrivateLinkResourcesClientListByStorageAccountResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// QueueClientCreateResponse contains the response from method QueueClient.Create.
type QueueClientCreateResponse struct {
	Queue
}

// QueueClientDeleteResponse contains the response from method QueueClient.Delete.
type QueueClientDeleteResponse struct {
	// placeholder for future response values
}

// QueueClientGetResponse contains the response from method QueueClient.Get.
type QueueClientGetResponse struct {
	Queue
}

// QueueClientListResponse contains the response from method QueueClient.NewListPager.
type QueueClientListResponse struct {
	// Response schema. Contains list of queues returned
	ListQueueResource
}

// QueueClientUpdateResponse contains the response from method QueueClient.Update.
type QueueClientUpdateResponse struct {
	Queue
}

// QueueServicesClientGetServicePropertiesResponse contains the response from method QueueServicesClient.GetServiceProperties.
type QueueServicesClientGetServicePropertiesResponse struct {
	// The properties of a storage account’s Queue service.
	QueueServiceProperties
}

// QueueServicesClientListResponse contains the response from method QueueServicesClient.List.
type QueueServicesClientListResponse struct {
	ListQueueServices
}

// QueueServicesClientSetServicePropertiesResponse contains the response from method QueueServicesClient.SetServiceProperties.
type QueueServicesClientSetServicePropertiesResponse struct {
	// The properties of a storage account’s Queue service.
	QueueServiceProperties
}

// SKUsClientListResponse contains the response from method SKUsClient.NewListPager.
type SKUsClientListResponse struct {
	// The response from the List Storage SKUs operation.
	SKUListResult
}

// TableClientCreateResponse contains the response from method TableClient.Create.
type TableClientCreateResponse struct {
	// Properties of the table, including Id, resource name, resource type.
	Table
}

// TableClientDeleteResponse contains the response from method TableClient.Delete.
type TableClientDeleteResponse struct {
	// placeholder for future response values
}

// TableClientGetResponse contains the response from method TableClient.Get.
type TableClientGetResponse struct {
	// Properties of the table, including Id, resource name, resource type.
	Table
}

// TableClientListResponse contains the response from method TableClient.NewListPager.
type TableClientListResponse struct {
	// Response schema. Contains list of tables returned
	ListTableResource
}

// TableClientUpdateResponse contains the response from method TableClient.Update.
type TableClientUpdateResponse struct {
	// Properties of the table, including Id, resource name, resource type.
	Table
}

// TableServicesClientGetServicePropertiesResponse contains the response from method TableServicesClient.GetServiceProperties.
type TableServicesClientGetServicePropertiesResponse struct {
	// The properties of a storage account’s Table service.
	TableServiceProperties
}

// TableServicesClientListResponse contains the response from method TableServicesClient.List.
type TableServicesClientListResponse struct {
	ListTableServices
}

// TableServicesClientSetServicePropertiesResponse contains the response from method TableServicesClient.SetServiceProperties.
type TableServicesClientSetServicePropertiesResponse struct {
	// The properties of a storage account’s Table service.
	TableServiceProperties
}

// TaskAssignmentInstancesReportClientListResponse contains the response from method TaskAssignmentInstancesReportClient.NewListPager.
type TaskAssignmentInstancesReportClientListResponse struct {
	// Fetch Storage Tasks Run Summary.
	TaskReportSummary
}

// TaskAssignmentsClientCreateResponse contains the response from method TaskAssignmentsClient.BeginCreate.
type TaskAssignmentsClientCreateResponse struct {
	// The storage task assignment.
	TaskAssignment
}

// TaskAssignmentsClientDeleteResponse contains the response from method TaskAssignmentsClient.BeginDelete.
type TaskAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// TaskAssignmentsClientGetResponse contains the response from method TaskAssignmentsClient.Get.
type TaskAssignmentsClientGetResponse struct {
	// The storage task assignment.
	TaskAssignment
}

// TaskAssignmentsClientListResponse contains the response from method TaskAssignmentsClient.NewListPager.
type TaskAssignmentsClientListResponse struct {
	// List of storage task assignments for the storage account
	TaskAssignmentsList
}

// TaskAssignmentsClientUpdateResponse contains the response from method TaskAssignmentsClient.BeginUpdate.
type TaskAssignmentsClientUpdateResponse struct {
	// The storage task assignment.
	TaskAssignment
}

// TaskAssignmentsInstancesReportClientListResponse contains the response from method TaskAssignmentsInstancesReportClient.NewListPager.
type TaskAssignmentsInstancesReportClientListResponse struct {
	// Fetch Storage Tasks Run Summary.
	TaskReportSummary
}

// UsagesClientListByLocationResponse contains the response from method UsagesClient.NewListByLocationPager.
type UsagesClientListByLocationResponse struct {
	// The response from the List Usages operation.
	UsageListResult
}
