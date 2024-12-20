// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure

import "encoding/json"

func unmarshalAgentProfileClassification(rawMsg json.RawMessage) (AgentProfileClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b AgentProfileClassification
	switch m["kind"] {
	case "Stateful":
		b = &Stateful{}
	case "Stateless":
		b = &StatelessAgentProfile{}
	default:
		b = &AgentProfile{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalFabricProfileClassification(rawMsg json.RawMessage) (FabricProfileClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FabricProfileClassification
	switch m["kind"] {
	case "Vmss":
		b = &VmssFabricProfile{}
	default:
		b = &FabricProfile{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalOrganizationProfileClassification(rawMsg json.RawMessage) (OrganizationProfileClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b OrganizationProfileClassification
	switch m["kind"] {
	case "AzureDevOps":
		b = &AzureDevOpsOrganizationProfile{}
	case "GitHub":
		b = &GitHubOrganizationProfile{}
	default:
		b = &OrganizationProfile{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalResourcePredictionsProfileClassification(rawMsg json.RawMessage) (ResourcePredictionsProfileClassification, error) {
	if rawMsg == nil || string(rawMsg) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ResourcePredictionsProfileClassification
	switch m["kind"] {
	case string(ResourcePredictionsProfileTypeManual):
		b = &ManualResourcePredictionsProfile{}
	case string(ResourcePredictionsProfileTypeAutomatic):
		b = &AutomaticResourcePredictionsProfile{}
	default:
		b = &ResourcePredictionsProfile{}
	}
	if err := json.Unmarshal(rawMsg, b); err != nil {
		return nil, err
	}
	return b, nil
}
