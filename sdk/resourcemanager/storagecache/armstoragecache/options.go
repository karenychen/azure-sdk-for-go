//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

// AmlFilesystemsClientArchiveOptions contains the optional parameters for the AmlFilesystemsClient.Archive method.
type AmlFilesystemsClientArchiveOptions struct {
	// Information about the archive operation
	ArchiveInfo *AmlFilesystemArchiveInfo
}

// AmlFilesystemsClientBeginCreateOrUpdateOptions contains the optional parameters for the AmlFilesystemsClient.BeginCreateOrUpdate
// method.
type AmlFilesystemsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AmlFilesystemsClientBeginDeleteOptions contains the optional parameters for the AmlFilesystemsClient.BeginDelete method.
type AmlFilesystemsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AmlFilesystemsClientBeginUpdateOptions contains the optional parameters for the AmlFilesystemsClient.BeginUpdate method.
type AmlFilesystemsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AmlFilesystemsClientCancelArchiveOptions contains the optional parameters for the AmlFilesystemsClient.CancelArchive method.
type AmlFilesystemsClientCancelArchiveOptions struct {
	// placeholder for future optional parameters
}

// AmlFilesystemsClientGetOptions contains the optional parameters for the AmlFilesystemsClient.Get method.
type AmlFilesystemsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AmlFilesystemsClientListByResourceGroupOptions contains the optional parameters for the AmlFilesystemsClient.NewListByResourceGroupPager
// method.
type AmlFilesystemsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AmlFilesystemsClientListOptions contains the optional parameters for the AmlFilesystemsClient.NewListPager method.
type AmlFilesystemsClientListOptions struct {
	// placeholder for future optional parameters
}

// AscOperationsClientGetOptions contains the optional parameters for the AscOperationsClient.Get method.
type AscOperationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AscUsagesClientListOptions contains the optional parameters for the AscUsagesClient.NewListPager method.
type AscUsagesClientListOptions struct {
	// placeholder for future optional parameters
}

// CachesClientBeginCreateOrUpdateOptions contains the optional parameters for the CachesClient.BeginCreateOrUpdate method.
type CachesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginDebugInfoOptions contains the optional parameters for the CachesClient.BeginDebugInfo method.
type CachesClientBeginDebugInfoOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginDeleteOptions contains the optional parameters for the CachesClient.BeginDelete method.
type CachesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginFlushOptions contains the optional parameters for the CachesClient.BeginFlush method.
type CachesClientBeginFlushOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginPausePrimingJobOptions contains the optional parameters for the CachesClient.BeginPausePrimingJob method.
type CachesClientBeginPausePrimingJobOptions struct {
	// Object containing the priming job ID.
	PrimingJobID *PrimingJobIDParameter

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginResumePrimingJobOptions contains the optional parameters for the CachesClient.BeginResumePrimingJob method.
type CachesClientBeginResumePrimingJobOptions struct {
	// Object containing the priming job ID.
	PrimingJobID *PrimingJobIDParameter

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginSpaceAllocationOptions contains the optional parameters for the CachesClient.BeginSpaceAllocation method.
type CachesClientBeginSpaceAllocationOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string

	// List containing storage target cache space percentage allocations.
	SpaceAllocation []*StorageTargetSpaceAllocation
}

// CachesClientBeginStartOptions contains the optional parameters for the CachesClient.BeginStart method.
type CachesClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginStartPrimingJobOptions contains the optional parameters for the CachesClient.BeginStartPrimingJob method.
type CachesClientBeginStartPrimingJobOptions struct {
	// Object containing the definition of a priming job.
	Primingjob *PrimingJob

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginStopOptions contains the optional parameters for the CachesClient.BeginStop method.
type CachesClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginStopPrimingJobOptions contains the optional parameters for the CachesClient.BeginStopPrimingJob method.
type CachesClientBeginStopPrimingJobOptions struct {
	// Object containing the priming job ID.
	PrimingJobID *PrimingJobIDParameter

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginUpdateOptions contains the optional parameters for the CachesClient.BeginUpdate method.
type CachesClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientBeginUpgradeFirmwareOptions contains the optional parameters for the CachesClient.BeginUpgradeFirmware method.
type CachesClientBeginUpgradeFirmwareOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// CachesClientGetOptions contains the optional parameters for the CachesClient.Get method.
type CachesClientGetOptions struct {
	// placeholder for future optional parameters
}

// CachesClientListByResourceGroupOptions contains the optional parameters for the CachesClient.NewListByResourceGroupPager
// method.
type CachesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// CachesClientListOptions contains the optional parameters for the CachesClient.NewListPager method.
type CachesClientListOptions struct {
	// placeholder for future optional parameters
}

// ManagementClientCheckAmlFSSubnetsOptions contains the optional parameters for the ManagementClient.CheckAmlFSSubnets method.
type ManagementClientCheckAmlFSSubnetsOptions struct {
	// Information about the subnets to validate.
	AmlFilesystemSubnetInfo *AmlFilesystemSubnetInfo
}

// ManagementClientGetRequiredAmlFSSubnetsSizeOptions contains the optional parameters for the ManagementClient.GetRequiredAmlFSSubnetsSize
// method.
type ManagementClientGetRequiredAmlFSSubnetsSizeOptions struct {
	// Information to determine the number of available IPs a subnet will need to host the AML file system.
	RequiredAMLFilesystemSubnetsSizeInfo *RequiredAmlFilesystemSubnetsSizeInfo
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// SKUsClientListOptions contains the optional parameters for the SKUsClient.NewListPager method.
type SKUsClientListOptions struct {
	// placeholder for future optional parameters
}

// StorageTargetClientBeginFlushOptions contains the optional parameters for the StorageTargetClient.BeginFlush method.
type StorageTargetClientBeginFlushOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetClientBeginInvalidateOptions contains the optional parameters for the StorageTargetClient.BeginInvalidate
// method.
type StorageTargetClientBeginInvalidateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetClientBeginResumeOptions contains the optional parameters for the StorageTargetClient.BeginResume method.
type StorageTargetClientBeginResumeOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetClientBeginSuspendOptions contains the optional parameters for the StorageTargetClient.BeginSuspend method.
type StorageTargetClientBeginSuspendOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageTargetsClient.BeginCreateOrUpdate
// method.
type StorageTargetsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientBeginDNSRefreshOptions contains the optional parameters for the StorageTargetsClient.BeginDNSRefresh
// method.
type StorageTargetsClientBeginDNSRefreshOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientBeginDeleteOptions contains the optional parameters for the StorageTargetsClient.BeginDelete method.
type StorageTargetsClientBeginDeleteOptions struct {
	// Boolean value requesting the force delete operation for a storage target. Force delete discards unwritten-data in the cache
	// instead of flushing it to back-end storage.
	Force *string

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientBeginRestoreDefaultsOptions contains the optional parameters for the StorageTargetsClient.BeginRestoreDefaults
// method.
type StorageTargetsClientBeginRestoreDefaultsOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// StorageTargetsClientGetOptions contains the optional parameters for the StorageTargetsClient.Get method.
type StorageTargetsClientGetOptions struct {
	// placeholder for future optional parameters
}

// StorageTargetsClientListByCacheOptions contains the optional parameters for the StorageTargetsClient.NewListByCachePager
// method.
type StorageTargetsClientListByCacheOptions struct {
	// placeholder for future optional parameters
}

// UsageModelsClientListOptions contains the optional parameters for the UsageModelsClient.NewListPager method.
type UsageModelsClientListOptions struct {
	// placeholder for future optional parameters
}