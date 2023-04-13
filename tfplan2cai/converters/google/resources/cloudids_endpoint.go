// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const CloudIdsEndpointAssetType string = "ids.googleapis.com/Endpoint"

func resourceConverterCloudIdsEndpoint() ResourceConverter {
	return ResourceConverter{
		AssetType: CloudIdsEndpointAssetType,
		Convert:   GetCloudIdsEndpointCaiObject,
	}
}

func GetCloudIdsEndpointCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//ids.googleapis.com/projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetCloudIdsEndpointApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: CloudIdsEndpointAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/ids/v1/rest",
				DiscoveryName:        "Endpoint",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetCloudIdsEndpointApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudIdsEndpointName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandCloudIdsEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	descriptionProp, err := expandCloudIdsEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	severityProp, err := expandCloudIdsEndpointSeverity(d.Get("severity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("severity"); !isEmptyValue(reflect.ValueOf(severityProp)) && (ok || !reflect.DeepEqual(v, severityProp)) {
		obj["severity"] = severityProp
	}
	threatExceptionsProp, err := expandCloudIdsEndpointThreatExceptions(d.Get("threat_exceptions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threat_exceptions"); !isEmptyValue(reflect.ValueOf(threatExceptionsProp)) && (ok || !reflect.DeepEqual(v, threatExceptionsProp)) {
		obj["threatExceptions"] = threatExceptionsProp
	}

	return obj, nil
}

func expandCloudIdsEndpointName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
}

func expandCloudIdsEndpointNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointSeverity(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointThreatExceptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}