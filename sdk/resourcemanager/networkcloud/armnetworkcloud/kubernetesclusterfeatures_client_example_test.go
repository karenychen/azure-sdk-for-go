//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/KubernetesClusterFeatures_ListByKubernetesCluster.json
func ExampleKubernetesClusterFeaturesClient_NewListByKubernetesClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKubernetesClusterFeaturesClient().NewListByKubernetesClusterPager("resourceGroupName", "kubernetesClusterName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KubernetesClusterFeatureList = armnetworkcloud.KubernetesClusterFeatureList{
		// 	Value: []*armnetworkcloud.KubernetesClusterFeature{
		// 		{
		// 			Name: to.Ptr("featureName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/features"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/features/featureName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
		// 				AvailabilityLifecycle: to.Ptr(armnetworkcloud.KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
		// 				DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterFeatureDetailedStatusInstalled),
		// 				DetailedStatusMessage: to.Ptr("The feature is successfully installed."),
		// 				Options: []*armnetworkcloud.StringKeyValuePair{
		// 					{
		// 						Key: to.Ptr("featureOptionName"),
		// 						Value: to.Ptr("featureOptionValue"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterFeatureProvisioningStateSucceeded),
		// 				Required: to.Ptr(armnetworkcloud.KubernetesClusterFeatureRequiredTrue),
		// 				Version: to.Ptr("1.2.3"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/KubernetesClusterFeatures_Get.json
func ExampleKubernetesClusterFeaturesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewKubernetesClusterFeaturesClient().Get(ctx, "resourceGroupName", "kubernetesClusterName", "featureName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesClusterFeature = armnetworkcloud.KubernetesClusterFeature{
	// 	Name: to.Ptr("featureName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/features"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/features/featureName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
	// 		AvailabilityLifecycle: to.Ptr(armnetworkcloud.KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterFeatureDetailedStatusInstalled),
	// 		DetailedStatusMessage: to.Ptr("The feature is successfully installed."),
	// 		Options: []*armnetworkcloud.StringKeyValuePair{
	// 			{
	// 				Key: to.Ptr("featureOptionName"),
	// 				Value: to.Ptr("featureOptionValue"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterFeatureProvisioningStateSucceeded),
	// 		Required: to.Ptr(armnetworkcloud.KubernetesClusterFeatureRequiredTrue),
	// 		Version: to.Ptr("1.2.3"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/KubernetesClusterFeatures_Create.json
func ExampleKubernetesClusterFeaturesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubernetesClusterFeaturesClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "kubernetesClusterName", "featureName", armnetworkcloud.KubernetesClusterFeature{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
			Options: []*armnetworkcloud.StringKeyValuePair{
				{
					Key:   to.Ptr("featureOptionName"),
					Value: to.Ptr("featureOptionValue"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesClusterFeature = armnetworkcloud.KubernetesClusterFeature{
	// 	Name: to.Ptr("featureName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/features"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/features/featureName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
	// 		AvailabilityLifecycle: to.Ptr(armnetworkcloud.KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterFeatureDetailedStatusInstalled),
	// 		DetailedStatusMessage: to.Ptr("The feature is successfully installed."),
	// 		Options: []*armnetworkcloud.StringKeyValuePair{
	// 			{
	// 				Key: to.Ptr("featureOptionName"),
	// 				Value: to.Ptr("featureOptionValue"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterFeatureProvisioningStateSucceeded),
	// 		Required: to.Ptr(armnetworkcloud.KubernetesClusterFeatureRequiredTrue),
	// 		Version: to.Ptr("1.2.3"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/KubernetesClusterFeatures_Delete.json
func ExampleKubernetesClusterFeaturesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubernetesClusterFeaturesClient().BeginDelete(ctx, "resourceGroupName", "kubernetesClusterName", "featureName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c0a12a75b702054cf1e7fcd8c014d0fc116dea6e/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2024-07-01/examples/KubernetesClusterFeatures_Patch.json
func ExampleKubernetesClusterFeaturesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKubernetesClusterFeaturesClient().BeginUpdate(ctx, "resourceGroupName", "kubernetesClusterName", "featureName", armnetworkcloud.KubernetesClusterFeaturePatchParameters{
		Properties: &armnetworkcloud.KubernetesClusterFeaturePatchProperties{
			Options: []*armnetworkcloud.StringKeyValuePair{
				{
					Key:   to.Ptr("featureOptionName"),
					Value: to.Ptr("featureOptionValue"),
				}},
		},
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.KubernetesClusterFeature = armnetworkcloud.KubernetesClusterFeature{
	// 	Name: to.Ptr("featureName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/features"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/features/featureName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	Properties: &armnetworkcloud.KubernetesClusterFeatureProperties{
	// 		AvailabilityLifecycle: to.Ptr(armnetworkcloud.KubernetesClusterFeatureAvailabilityLifecycleGenerallyAvailable),
	// 		DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterFeatureDetailedStatusInstalled),
	// 		DetailedStatusMessage: to.Ptr("The feature is successfully installed."),
	// 		Options: []*armnetworkcloud.StringKeyValuePair{
	// 			{
	// 				Key: to.Ptr("featureOptionName"),
	// 				Value: to.Ptr("featureOptionValue"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterFeatureProvisioningStateSucceeded),
	// 		Required: to.Ptr(armnetworkcloud.KubernetesClusterFeatureRequiredTrue),
	// 		Version: to.Ptr("1.2.3"),
	// 	},
	// }
}
