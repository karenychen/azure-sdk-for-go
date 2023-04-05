//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple8000series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetAlertSettings.json
func ExampleDeviceSettingsClient_GetAlertSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetAlertSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AlertSettings = armstorsimple8000series.AlertSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/alertSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/alertSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.AlertNotificationProperties{
	// 		AdditionalRecipientEmailList: []*string{
	// 		},
	// 		AlertNotificationCulture: to.Ptr("en-US"),
	// 		EmailNotification: to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
	// 		NotificationToServiceOwners: to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsCreateOrUpdateAlertSettings.json
func ExampleDeviceSettingsClient_BeginCreateOrUpdateAlertSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginCreateOrUpdateAlertSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.AlertSettings{
		Properties: &armstorsimple8000series.AlertNotificationProperties{
			AdditionalRecipientEmailList: []*string{},
			AlertNotificationCulture:     to.Ptr("en-US"),
			EmailNotification:            to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
			NotificationToServiceOwners:  to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
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
	// res.AlertSettings = armstorsimple8000series.AlertSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/alertSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/alertSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.AlertNotificationProperties{
	// 		AdditionalRecipientEmailList: []*string{
	// 		},
	// 		AlertNotificationCulture: to.Ptr("en-US"),
	// 		EmailNotification: to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
	// 		NotificationToServiceOwners: to.Ptr(armstorsimple8000series.AlertEmailNotificationStatusEnabled),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetNetworkSettings.json
func ExampleDeviceSettingsClient_GetNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetNetworkSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSettings = armstorsimple8000series.NetworkSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/networkSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/networkSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.NetworkSettingsProperties{
	// 		DNSSettings: &armstorsimple8000series.DNSSettings{
	// 			PrimaryDNSServer: to.Ptr("10.171.65.60"),
	// 			SecondaryDNSServers: []*string{
	// 				to.Ptr("8.8.8.8")},
	// 				SecondaryIPv6DNSServers: []*string{
	// 				},
	// 			},
	// 			NetworkAdapters: &armstorsimple8000series.NetworkAdapterList{
	// 				Value: []*armstorsimple8000series.NetworkAdapters{
	// 					{
	// 						InterfaceID: to.Ptr(armstorsimple8000series.NetInterfaceIDData0),
	// 						IsDefault: to.Ptr(true),
	// 						IscsiAndCloudStatus: to.Ptr(armstorsimple8000series.ISCSIAndCloudStatusIscsiAndCloudEnabled),
	// 						Mode: to.Ptr(armstorsimple8000series.NetworkModeIPV4),
	// 						NetInterfaceStatus: to.Ptr(armstorsimple8000series.NetInterfaceStatusEnabled),
	// 						NicIPv4Settings: &armstorsimple8000series.NicIPv4{
	// 							Controller0IPv4Address: to.Ptr("10.168.241.143"),
	// 							Controller1IPv4Address: to.Ptr("10.168.241.121"),
	// 							IPv4Address: to.Ptr("10.168.241.187"),
	// 							IPv4Gateway: to.Ptr("10.168.241.1"),
	// 							IPv4Netmask: to.Ptr("255.255.252.0"),
	// 						},
	// 						NicIPv6Settings: &armstorsimple8000series.NicIPv6{
	// 							Controller0IPv6Address: to.Ptr(""),
	// 							Controller1IPv6Address: to.Ptr(""),
	// 							IPv6Address: to.Ptr(""),
	// 							IPv6Gateway: to.Ptr(""),
	// 							IPv6Prefix: to.Ptr(""),
	// 						},
	// 						Speed: to.Ptr[int64](1000000000),
	// 				}},
	// 			},
	// 			WebproxySettings: &armstorsimple8000series.WebproxySettings{
	// 				Authentication: to.Ptr(armstorsimple8000series.AuthenticationTypeNone),
	// 				ConnectionURI: to.Ptr(""),
	// 				Username: to.Ptr(""),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsUpdateNetworkSettings.json
func ExampleDeviceSettingsClient_BeginUpdateNetworkSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginUpdateNetworkSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.NetworkSettingsPatch{
		Properties: &armstorsimple8000series.NetworkSettingsPatchProperties{
			DNSSettings: &armstorsimple8000series.DNSSettings{
				PrimaryDNSServer: to.Ptr("10.171.65.60"),
				SecondaryDNSServers: []*string{
					to.Ptr("8.8.8.8")},
			},
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
	// res.NetworkSettings = armstorsimple8000series.NetworkSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/encryptionSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/encryptionSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.NetworkSettingsProperties{
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetSecuritySettings.json
func ExampleDeviceSettingsClient_GetSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetSecuritySettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SecuritySettings = armstorsimple8000series.SecuritySettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/securitySettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.SecuritySettingsProperties{
	// 		ChapSettings: &armstorsimple8000series.ChapSettings{
	// 			InitiatorUser: to.Ptr("test-initiator-user"),
	// 			TargetUser: to.Ptr("test-target-user"),
	// 		},
	// 		RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettings{
	// 			RemoteManagementCertificate: to.Ptr("-----BEGIN CERTIFICATE-----\r\nMIIDfjCCAmagAwIBAgIQR9UA510Lj4FPZGVXI0oyQDANBgkqhkiG9w0BAQUFADBWMSAwHgYDVQQL\r\nDBdIY3NSZW1vdGVNYW5hZ2VtZW50Q2VydDEeMBwGA1UECgwVTWljcm9zb2Z0IENvcnBvcmF0aW9u\r\nMRIwEAYDVQQDDAkxMjM0NTY3ODkwIBcNMTcwNjA3MTExMDUzWhgPMjExNjA2MDcxMTEwNTNaMFYx\r\nIDAeBgNVBAsMF0hjc1JlbW90ZU1hbmFnZW1lbnRDZXJ0MR4wHAYDVQQKDBVNaWNyb3NvZnQgQ29y\r\ncG9yYXRpb24xEjAQBgNVBAMMCTEyMzQ1Njc4OTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoC\r\nggEBAMp25cGMmEJdwur4y2Qgs9AKimTYGHdL5R9o2rcf+rGpgLC2wMY4qWsz3Ub4EWDwgT26rCiX\r\nv/24FK3xGB0OLxzTC3EbBJ7g7uCkapAKp/cFcmr+uBZiBhXe3hNnkeP8H2IfXNRSKa54XUMSqoFi\r\naavSBstS8FMuDMKopIVe6o4j+oDZBJGbb4a2anxuSyBM3cNhFTaOCLnfKynREC4sbbBL6enaFs0i\r\nRdIobT1FDSb2sLSUz4ONa3ZN7Os25UFvj1OcWeDjT1vJKcjx+u5owILvifRGoGbPETtCjwJ306+7\r\nihl+0GQ7teD2+O1h/Q7qdBm3gQAX9ZTsQAoiRYyJ14UCAwEAAaNGMEQwEwYDVR0lBAwwCgYIKwYB\r\nBQUHAwEwHQYDVR0OBBYEFJNK6cTqsGCcZkjP1q0LF53B2+G1MA4GA1UdDwEB/wQEAwIFIDANBgkq\r\nhkiG9w0BAQUFAAOCAQEAl+2ApGz9hpt6+sND+Q3TGczBqh2BSprgK//KwPdWthSLnxPf7f06MITN\r\ng1sTxFlbrFb46Y0UPN5YKTMb7wxRd+Nqrz6l7X4kp6IR+2sMX1ydkWyJ6HZrosc96fkFvFUt9B+J\r\n25gb9YVRotNUeborkznDFW6VeeS0kxJVps5r3fcuhz2vOPW1q/U0UWNV+LnrzRT+7MauUTYyLe0q\r\nuobmnD8NC2HhZcJ04xYhCKOJQ/NNJyaURwGJ2TZSYwS0HOnQViZp/SipjhyOOwwB6h0W62ElOovN\r\nAxVgTBEQjbG0jsMOm079hLUKuIBYmH4HJDmF3QKml0er1NcXeieATPLOWQ==\r\n-----END CERTIFICATE-----\r\n"),
	// 			RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsUpdateSecuritySettings.json
func ExampleDeviceSettingsClient_BeginUpdateSecuritySettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginUpdateSecuritySettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.SecuritySettingsPatch{
		Properties: &armstorsimple8000series.SecuritySettingsPatchProperties{
			ChapSettings: &armstorsimple8000series.ChapSettings{
				InitiatorSecret: &armstorsimple8000series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
					Value:                    to.Ptr("V/uVfWk5OcXfMC0HvUV89o9+cmF636jBnqhFM1pD/zHhmh8Z1KB5/LhVV3T53uGjIlKL3wjhwg+9NIQrIbYuKhl/r8jSftSSH+WqUnQHTRDWazjPAeMu6ozrL5RYzP1h5mgw7XtidZPaaV9ae/uF1KQPkK6TIARaOTdr8I/BLWUg7WdDrfARNYHnW6ezXek1M9Qhv1sL9fZY+JrGB58LF6D2aC2Xjed4K4Jk6v2T1ieneNV27uIdnt21TajuM7w90UlRiVZJZtq/KdEUfqI28C7VoUdcXluAwzR95Ho8hmyIJDqeW3/Wxymdjv+Rctwqtmcka9i2G85Hj8SVV3g4kA=="),
				},
				InitiatorUser: to.Ptr("test-initiator-user"),
				TargetSecret: &armstorsimple8000series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
					Value:                    to.Ptr("OTR4uwVpy+pf0zthnCIAUXurC8NdSh8RpRG5GWL9TSv4WtkVmpeU/U2A4vjkrchfQOzI1x+uooWikWW9txwwQOM+/N3NG44+/dlHoaEe7AxjmItCKhNj8K2RM6D1mb45wicbF/M4uanuXnGXuT+JmZ+1Lcy2k1GXsk67ejplz2K08h37B+oIW85qMUHLdKuuQlAA/fFS+q6qMti3j2Q8Fr+Sh4U76/2AZVkKRtFeqPB1QhC12dFx6TFoZJkMFzdQz4WNvWVelIK2McKNnOiH0/Z5lAXC7164uzReAoTEfqoNU7qqqRrHhsdwWPu6jbeUn8BQnr7A/X6NWvgeax+HGA=="),
				},
				TargetUser: to.Ptr("test-target-user"),
			},
			DeviceAdminPassword: &armstorsimple8000series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    to.Ptr("<value>"),
			},
			RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettingsPatch{
				RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
			},
			SnapshotPassword: &armstorsimple8000series.AsymmetricEncryptedSecret{
				EncryptionAlgorithm:      to.Ptr(armstorsimple8000series.EncryptionAlgorithmRSAESPKCS1V15),
				EncryptionCertThumbprint: to.Ptr("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    to.Ptr("<value>"),
			},
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
	// res.SecuritySettings = armstorsimple8000series.SecuritySettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/securitySettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/securitySettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.SecuritySettingsProperties{
	// 		ChapSettings: &armstorsimple8000series.ChapSettings{
	// 			InitiatorUser: to.Ptr("test-initiator-user"),
	// 			TargetUser: to.Ptr("test-target-user"),
	// 		},
	// 		RemoteManagementSettings: &armstorsimple8000series.RemoteManagementSettings{
	// 			RemoteManagementCertificate: to.Ptr(""),
	// 			RemoteManagementMode: to.Ptr(armstorsimple8000series.RemoteManagementModeConfigurationHTTPSAndHTTPEnabled),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsSyncRemotemanagementCertificate.json
func ExampleDeviceSettingsClient_BeginSyncRemotemanagementCertificate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginSyncRemotemanagementCertificate(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsGetTimeSettings.json
func ExampleDeviceSettingsClient_GetTimeSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeviceSettingsClient().GetTimeSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TimeSettings = armstorsimple8000series.TimeSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/timeSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/timeSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.TimeSettingsProperties{
	// 		PrimaryTimeServer: to.Ptr("time.windows.com"),
	// 		SecondaryTimeServer: []*string{
	// 			to.Ptr("8.8.8.8")},
	// 			TimeZone: to.Ptr("Pacific Standard Time"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/storsimple8000series/resource-manager/Microsoft.StorSimple/stable/2017-06-01/examples/DeviceSettingsCreateOrUpdateTimeSettings.json
func ExampleDeviceSettingsClient_BeginCreateOrUpdateTimeSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorsimple8000series.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDeviceSettingsClient().BeginCreateOrUpdateTimeSettings(ctx, "Device05ForSDKTest", "ResourceGroupForSDKTest", "ManagerForSDKTest1", armstorsimple8000series.TimeSettings{
		Properties: &armstorsimple8000series.TimeSettingsProperties{
			PrimaryTimeServer: to.Ptr("time.windows.com"),
			SecondaryTimeServer: []*string{
				to.Ptr("8.8.8.8")},
			TimeZone: to.Ptr("Pacific Standard Time"),
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
	// res.TimeSettings = armstorsimple8000series.TimeSettings{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.StorSimple/managers/devices/timeSettings"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/timeSettings/default"),
	// 	Kind: to.Ptr("Series8000"),
	// 	Properties: &armstorsimple8000series.TimeSettingsProperties{
	// 		PrimaryTimeServer: to.Ptr("time.windows.com"),
	// 		SecondaryTimeServer: []*string{
	// 			to.Ptr("8.8.8.8")},
	// 			TimeZone: to.Ptr("Pacific Standard Time"),
	// 		},
	// 	}
}