//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorageimportexport

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type DeliveryPackageInformation.
func (d DeliveryPackageInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "carrierName", d.CarrierName)
	populate(objectMap, "driveCount", d.DriveCount)
	populate(objectMap, "shipDate", d.ShipDate)
	populate(objectMap, "trackingNumber", d.TrackingNumber)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DeliveryPackageInformation.
func (d *DeliveryPackageInformation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "carrierName":
			err = unpopulate(val, "CarrierName", &d.CarrierName)
			delete(rawMsg, key)
		case "driveCount":
			err = unpopulate(val, "DriveCount", &d.DriveCount)
			delete(rawMsg, key)
		case "shipDate":
			err = unpopulate(val, "ShipDate", &d.ShipDate)
			delete(rawMsg, key)
		case "trackingNumber":
			err = unpopulate(val, "TrackingNumber", &d.TrackingNumber)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DriveBitLockerKey.
func (d DriveBitLockerKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "bitLockerKey", d.BitLockerKey)
	populate(objectMap, "driveId", d.DriveID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DriveBitLockerKey.
func (d *DriveBitLockerKey) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bitLockerKey":
			err = unpopulate(val, "BitLockerKey", &d.BitLockerKey)
			delete(rawMsg, key)
		case "driveId":
			err = unpopulate(val, "DriveID", &d.DriveID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DriveStatus.
func (d DriveStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "bitLockerKey", d.BitLockerKey)
	populate(objectMap, "bytesSucceeded", d.BytesSucceeded)
	populate(objectMap, "copyStatus", d.CopyStatus)
	populate(objectMap, "driveHeaderHash", d.DriveHeaderHash)
	populate(objectMap, "driveId", d.DriveID)
	populate(objectMap, "errorLogUri", d.ErrorLogURI)
	populate(objectMap, "manifestFile", d.ManifestFile)
	populate(objectMap, "manifestHash", d.ManifestHash)
	populate(objectMap, "manifestUri", d.ManifestURI)
	populate(objectMap, "percentComplete", d.PercentComplete)
	populate(objectMap, "state", d.State)
	populate(objectMap, "verboseLogUri", d.VerboseLogURI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DriveStatus.
func (d *DriveStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "bitLockerKey":
			err = unpopulate(val, "BitLockerKey", &d.BitLockerKey)
			delete(rawMsg, key)
		case "bytesSucceeded":
			err = unpopulate(val, "BytesSucceeded", &d.BytesSucceeded)
			delete(rawMsg, key)
		case "copyStatus":
			err = unpopulate(val, "CopyStatus", &d.CopyStatus)
			delete(rawMsg, key)
		case "driveHeaderHash":
			err = unpopulate(val, "DriveHeaderHash", &d.DriveHeaderHash)
			delete(rawMsg, key)
		case "driveId":
			err = unpopulate(val, "DriveID", &d.DriveID)
			delete(rawMsg, key)
		case "errorLogUri":
			err = unpopulate(val, "ErrorLogURI", &d.ErrorLogURI)
			delete(rawMsg, key)
		case "manifestFile":
			err = unpopulate(val, "ManifestFile", &d.ManifestFile)
			delete(rawMsg, key)
		case "manifestHash":
			err = unpopulate(val, "ManifestHash", &d.ManifestHash)
			delete(rawMsg, key)
		case "manifestUri":
			err = unpopulate(val, "ManifestURI", &d.ManifestURI)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, "PercentComplete", &d.PercentComplete)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &d.State)
			delete(rawMsg, key)
		case "verboseLogUri":
			err = unpopulate(val, "VerboseLogURI", &d.VerboseLogURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type EncryptionKeyDetails.
func (e EncryptionKeyDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "kekType", e.KekType)
	populate(objectMap, "kekUrl", e.KekURL)
	populate(objectMap, "kekVaultResourceID", e.KekVaultResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EncryptionKeyDetails.
func (e *EncryptionKeyDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "kekType":
			err = unpopulate(val, "KekType", &e.KekType)
			delete(rawMsg, key)
		case "kekUrl":
			err = unpopulate(val, "KekURL", &e.KekURL)
			delete(rawMsg, key)
		case "kekVaultResourceID":
			err = unpopulate(val, "KekVaultResourceID", &e.KekVaultResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseError.
func (e ErrorResponseError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populateAny(objectMap, "innererror", e.Innererror)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponseError.
func (e *ErrorResponseError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "details":
			err = unpopulate(val, "Details", &e.Details)
			delete(rawMsg, key)
		case "innererror":
			err = unpopulate(val, "Innererror", &e.Innererror)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseErrorDetailsItem.
func (e ErrorResponseErrorDetailsItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponseErrorDetailsItem.
func (e *ErrorResponseErrorDetailsItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "target":
			err = unpopulate(val, "Target", &e.Target)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Export.
func (e Export) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "blobList", e.BlobList)
	populate(objectMap, "blobListBlobPath", e.BlobListBlobPath)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Export.
func (e *Export) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "blobList":
			err = unpopulate(val, "BlobList", &e.BlobList)
			delete(rawMsg, key)
		case "blobListBlobPath":
			err = unpopulate(val, "BlobListBlobPath", &e.BlobListBlobPath)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExportBlobList.
func (e ExportBlobList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "blobPath", e.BlobPath)
	populate(objectMap, "blobPathPrefix", e.BlobPathPrefix)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExportBlobList.
func (e *ExportBlobList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "blobPath":
			err = unpopulate(val, "BlobPath", &e.BlobPath)
			delete(rawMsg, key)
		case "blobPathPrefix":
			err = unpopulate(val, "BlobPathPrefix", &e.BlobPathPrefix)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type GetBitLockerKeysResponse.
func (g GetBitLockerKeysResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", g.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type GetBitLockerKeysResponse.
func (g *GetBitLockerKeysResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", g, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &g.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", g, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IdentityDetails.
func (i IdentityDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "principalId", i.PrincipalID)
	populate(objectMap, "tenantId", i.TenantID)
	populate(objectMap, "type", i.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IdentityDetails.
func (i *IdentityDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "principalId":
			err = unpopulate(val, "PrincipalID", &i.PrincipalID)
			delete(rawMsg, key)
		case "tenantId":
			err = unpopulate(val, "TenantID", &i.TenantID)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &i.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobDetails.
func (j JobDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "backupDriveManifest", j.BackupDriveManifest)
	populate(objectMap, "cancelRequested", j.CancelRequested)
	populate(objectMap, "deliveryPackage", j.DeliveryPackage)
	populate(objectMap, "diagnosticsPath", j.DiagnosticsPath)
	populate(objectMap, "driveList", j.DriveList)
	populate(objectMap, "encryptionKey", j.EncryptionKey)
	populate(objectMap, "export", j.Export)
	populate(objectMap, "incompleteBlobListUri", j.IncompleteBlobListURI)
	populate(objectMap, "jobType", j.JobType)
	populate(objectMap, "logLevel", j.LogLevel)
	populate(objectMap, "percentComplete", j.PercentComplete)
	populate(objectMap, "provisioningState", j.ProvisioningState)
	populate(objectMap, "returnAddress", j.ReturnAddress)
	populate(objectMap, "returnPackage", j.ReturnPackage)
	populate(objectMap, "returnShipping", j.ReturnShipping)
	populate(objectMap, "shippingInformation", j.ShippingInformation)
	populate(objectMap, "state", j.State)
	populate(objectMap, "storageAccountId", j.StorageAccountID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobDetails.
func (j *JobDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backupDriveManifest":
			err = unpopulate(val, "BackupDriveManifest", &j.BackupDriveManifest)
			delete(rawMsg, key)
		case "cancelRequested":
			err = unpopulate(val, "CancelRequested", &j.CancelRequested)
			delete(rawMsg, key)
		case "deliveryPackage":
			err = unpopulate(val, "DeliveryPackage", &j.DeliveryPackage)
			delete(rawMsg, key)
		case "diagnosticsPath":
			err = unpopulate(val, "DiagnosticsPath", &j.DiagnosticsPath)
			delete(rawMsg, key)
		case "driveList":
			err = unpopulate(val, "DriveList", &j.DriveList)
			delete(rawMsg, key)
		case "encryptionKey":
			err = unpopulate(val, "EncryptionKey", &j.EncryptionKey)
			delete(rawMsg, key)
		case "export":
			err = unpopulate(val, "Export", &j.Export)
			delete(rawMsg, key)
		case "incompleteBlobListUri":
			err = unpopulate(val, "IncompleteBlobListURI", &j.IncompleteBlobListURI)
			delete(rawMsg, key)
		case "jobType":
			err = unpopulate(val, "JobType", &j.JobType)
			delete(rawMsg, key)
		case "logLevel":
			err = unpopulate(val, "LogLevel", &j.LogLevel)
			delete(rawMsg, key)
		case "percentComplete":
			err = unpopulate(val, "PercentComplete", &j.PercentComplete)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &j.ProvisioningState)
			delete(rawMsg, key)
		case "returnAddress":
			err = unpopulate(val, "ReturnAddress", &j.ReturnAddress)
			delete(rawMsg, key)
		case "returnPackage":
			err = unpopulate(val, "ReturnPackage", &j.ReturnPackage)
			delete(rawMsg, key)
		case "returnShipping":
			err = unpopulate(val, "ReturnShipping", &j.ReturnShipping)
			delete(rawMsg, key)
		case "shippingInformation":
			err = unpopulate(val, "ShippingInformation", &j.ShippingInformation)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &j.State)
			delete(rawMsg, key)
		case "storageAccountId":
			err = unpopulate(val, "StorageAccountID", &j.StorageAccountID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JobResponse.
func (j JobResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", j.ID)
	populate(objectMap, "identity", j.Identity)
	populate(objectMap, "location", j.Location)
	populate(objectMap, "name", j.Name)
	populate(objectMap, "properties", j.Properties)
	populate(objectMap, "systemData", j.SystemData)
	populateAny(objectMap, "tags", j.Tags)
	populate(objectMap, "type", j.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JobResponse.
func (j *JobResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &j.ID)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &j.Identity)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &j.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &j.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &j.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &j.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &j.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &j.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListJobsResponse.
func (l ListJobsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListJobsResponse.
func (l *ListJobsResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ListOperationsResponse.
func (l ListOperationsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ListOperationsResponse.
func (l *ListOperationsResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Location.
func (l Location) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "name", l.Name)
	populate(objectMap, "properties", l.Properties)
	populate(objectMap, "type", l.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Location.
func (l *Location) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &l.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &l.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &l.Properties)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &l.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocationProperties.
func (l LocationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalShippingInformation", l.AdditionalShippingInformation)
	populate(objectMap, "alternateLocations", l.AlternateLocations)
	populate(objectMap, "city", l.City)
	populate(objectMap, "countryOrRegion", l.CountryOrRegion)
	populate(objectMap, "phone", l.Phone)
	populate(objectMap, "postalCode", l.PostalCode)
	populate(objectMap, "recipientName", l.RecipientName)
	populate(objectMap, "stateOrProvince", l.StateOrProvince)
	populate(objectMap, "streetAddress1", l.StreetAddress1)
	populate(objectMap, "streetAddress2", l.StreetAddress2)
	populate(objectMap, "supportedCarriers", l.SupportedCarriers)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocationProperties.
func (l *LocationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalShippingInformation":
			err = unpopulate(val, "AdditionalShippingInformation", &l.AdditionalShippingInformation)
			delete(rawMsg, key)
		case "alternateLocations":
			err = unpopulate(val, "AlternateLocations", &l.AlternateLocations)
			delete(rawMsg, key)
		case "city":
			err = unpopulate(val, "City", &l.City)
			delete(rawMsg, key)
		case "countryOrRegion":
			err = unpopulate(val, "CountryOrRegion", &l.CountryOrRegion)
			delete(rawMsg, key)
		case "phone":
			err = unpopulate(val, "Phone", &l.Phone)
			delete(rawMsg, key)
		case "postalCode":
			err = unpopulate(val, "PostalCode", &l.PostalCode)
			delete(rawMsg, key)
		case "recipientName":
			err = unpopulate(val, "RecipientName", &l.RecipientName)
			delete(rawMsg, key)
		case "stateOrProvince":
			err = unpopulate(val, "StateOrProvince", &l.StateOrProvince)
			delete(rawMsg, key)
		case "streetAddress1":
			err = unpopulate(val, "StreetAddress1", &l.StreetAddress1)
			delete(rawMsg, key)
		case "streetAddress2":
			err = unpopulate(val, "StreetAddress2", &l.StreetAddress2)
			delete(rawMsg, key)
		case "supportedCarriers":
			err = unpopulate(val, "SupportedCarriers", &l.SupportedCarriers)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocationsResponse.
func (l LocationsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocationsResponse.
func (l *LocationsResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Operation.
func (o Operation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "display", o.Display)
	populate(objectMap, "name", o.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Operation.
func (o *Operation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &o.Display)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &o.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PackageInformation.
func (p PackageInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "carrierName", p.CarrierName)
	populate(objectMap, "driveCount", p.DriveCount)
	populate(objectMap, "shipDate", p.ShipDate)
	populate(objectMap, "trackingNumber", p.TrackingNumber)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PackageInformation.
func (p *PackageInformation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "carrierName":
			err = unpopulate(val, "CarrierName", &p.CarrierName)
			delete(rawMsg, key)
		case "driveCount":
			err = unpopulate(val, "DriveCount", &p.DriveCount)
			delete(rawMsg, key)
		case "shipDate":
			err = unpopulate(val, "ShipDate", &p.ShipDate)
			delete(rawMsg, key)
		case "trackingNumber":
			err = unpopulate(val, "TrackingNumber", &p.TrackingNumber)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PutJobParameters.
func (p PutJobParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "properties", p.Properties)
	populateAny(objectMap, "tags", p.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PutJobParameters.
func (p *PutJobParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "location":
			err = unpopulate(val, "Location", &p.Location)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &p.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &p.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReturnAddress.
func (r ReturnAddress) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "city", r.City)
	populate(objectMap, "countryOrRegion", r.CountryOrRegion)
	populate(objectMap, "email", r.Email)
	populate(objectMap, "phone", r.Phone)
	populate(objectMap, "postalCode", r.PostalCode)
	populate(objectMap, "recipientName", r.RecipientName)
	populate(objectMap, "stateOrProvince", r.StateOrProvince)
	populate(objectMap, "streetAddress1", r.StreetAddress1)
	populate(objectMap, "streetAddress2", r.StreetAddress2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReturnAddress.
func (r *ReturnAddress) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "city":
			err = unpopulate(val, "City", &r.City)
			delete(rawMsg, key)
		case "countryOrRegion":
			err = unpopulate(val, "CountryOrRegion", &r.CountryOrRegion)
			delete(rawMsg, key)
		case "email":
			err = unpopulate(val, "Email", &r.Email)
			delete(rawMsg, key)
		case "phone":
			err = unpopulate(val, "Phone", &r.Phone)
			delete(rawMsg, key)
		case "postalCode":
			err = unpopulate(val, "PostalCode", &r.PostalCode)
			delete(rawMsg, key)
		case "recipientName":
			err = unpopulate(val, "RecipientName", &r.RecipientName)
			delete(rawMsg, key)
		case "stateOrProvince":
			err = unpopulate(val, "StateOrProvince", &r.StateOrProvince)
			delete(rawMsg, key)
		case "streetAddress1":
			err = unpopulate(val, "StreetAddress1", &r.StreetAddress1)
			delete(rawMsg, key)
		case "streetAddress2":
			err = unpopulate(val, "StreetAddress2", &r.StreetAddress2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ReturnShipping.
func (r ReturnShipping) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "carrierAccountNumber", r.CarrierAccountNumber)
	populate(objectMap, "carrierName", r.CarrierName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ReturnShipping.
func (r *ReturnShipping) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "carrierAccountNumber":
			err = unpopulate(val, "CarrierAccountNumber", &r.CarrierAccountNumber)
			delete(rawMsg, key)
		case "carrierName":
			err = unpopulate(val, "CarrierName", &r.CarrierName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ShippingInformation.
func (s ShippingInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "additionalInformation", s.AdditionalInformation)
	populate(objectMap, "city", s.City)
	populate(objectMap, "countryOrRegion", s.CountryOrRegion)
	populate(objectMap, "phone", s.Phone)
	populate(objectMap, "postalCode", s.PostalCode)
	populate(objectMap, "recipientName", s.RecipientName)
	populate(objectMap, "stateOrProvince", s.StateOrProvince)
	populate(objectMap, "streetAddress1", s.StreetAddress1)
	populate(objectMap, "streetAddress2", s.StreetAddress2)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ShippingInformation.
func (s *ShippingInformation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "additionalInformation":
			err = unpopulate(val, "AdditionalInformation", &s.AdditionalInformation)
			delete(rawMsg, key)
		case "city":
			err = unpopulate(val, "City", &s.City)
			delete(rawMsg, key)
		case "countryOrRegion":
			err = unpopulate(val, "CountryOrRegion", &s.CountryOrRegion)
			delete(rawMsg, key)
		case "phone":
			err = unpopulate(val, "Phone", &s.Phone)
			delete(rawMsg, key)
		case "postalCode":
			err = unpopulate(val, "PostalCode", &s.PostalCode)
			delete(rawMsg, key)
		case "recipientName":
			err = unpopulate(val, "RecipientName", &s.RecipientName)
			delete(rawMsg, key)
		case "stateOrProvince":
			err = unpopulate(val, "StateOrProvince", &s.StateOrProvince)
			delete(rawMsg, key)
		case "streetAddress1":
			err = unpopulate(val, "StreetAddress1", &s.StreetAddress1)
			delete(rawMsg, key)
		case "streetAddress2":
			err = unpopulate(val, "StreetAddress2", &s.StreetAddress2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateJobParameters.
func (u UpdateJobParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "properties", u.Properties)
	populateAny(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateJobParameters.
func (u *UpdateJobParameters) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &u.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &u.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateJobParametersProperties.
func (u UpdateJobParametersProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]any)
	populate(objectMap, "backupDriveManifest", u.BackupDriveManifest)
	populate(objectMap, "cancelRequested", u.CancelRequested)
	populate(objectMap, "deliveryPackage", u.DeliveryPackage)
	populate(objectMap, "driveList", u.DriveList)
	populate(objectMap, "logLevel", u.LogLevel)
	populate(objectMap, "returnAddress", u.ReturnAddress)
	populate(objectMap, "returnShipping", u.ReturnShipping)
	populate(objectMap, "state", u.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateJobParametersProperties.
func (u *UpdateJobParametersProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "backupDriveManifest":
			err = unpopulate(val, "BackupDriveManifest", &u.BackupDriveManifest)
			delete(rawMsg, key)
		case "cancelRequested":
			err = unpopulate(val, "CancelRequested", &u.CancelRequested)
			delete(rawMsg, key)
		case "deliveryPackage":
			err = unpopulate(val, "DeliveryPackage", &u.DeliveryPackage)
			delete(rawMsg, key)
		case "driveList":
			err = unpopulate(val, "DriveList", &u.DriveList)
			delete(rawMsg, key)
		case "logLevel":
			err = unpopulate(val, "LogLevel", &u.LogLevel)
			delete(rawMsg, key)
		case "returnAddress":
			err = unpopulate(val, "ReturnAddress", &u.ReturnAddress)
			delete(rawMsg, key)
		case "returnShipping":
			err = unpopulate(val, "ReturnShipping", &u.ReturnShipping)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &u.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

func populate(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func populateAny(m map[string]any, k string, v any) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v any) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
