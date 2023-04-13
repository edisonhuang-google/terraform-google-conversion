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

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const workforcePoolProviderIdRegexp = `^[a-z0-9-]{4,32}$`

func ValidateWorkforcePoolProviderId(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "gcp-") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) can not start with \"gcp-\". "+
				"The prefix `gcp-` is reserved for use by Google, and may not be specified.", k, value))
	}

	if !regexp.MustCompile(workforcePoolProviderIdRegexp).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be 4-32 characters, and may contain the characters [a-z0-9-].", k, value))
	}

	return
}

const IAMWorkforcePoolWorkforcePoolProviderAssetType string = "iam.googleapis.com/WorkforcePoolProvider"

func resourceConverterIAMWorkforcePoolWorkforcePoolProvider() ResourceConverter {
	return ResourceConverter{
		AssetType: IAMWorkforcePoolWorkforcePoolProviderAssetType,
		Convert:   GetIAMWorkforcePoolWorkforcePoolProviderCaiObject,
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//iam.googleapis.com/locations/{{location}}/workforcePools/{{workforce_pool_id}}/providers/{{provider_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetIAMWorkforcePoolWorkforcePoolProviderApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: IAMWorkforcePoolWorkforcePoolProviderAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v1/rest",
				DiscoveryName:        "WorkforcePoolProvider",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetIAMWorkforcePoolWorkforcePoolProviderApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	disabledProp, err := expandIAMWorkforcePoolWorkforcePoolProviderDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	attributeMappingProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(d.Get("attribute_mapping"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attribute_mapping"); !isEmptyValue(reflect.ValueOf(attributeMappingProp)) && (ok || !reflect.DeepEqual(v, attributeMappingProp)) {
		obj["attributeMapping"] = attributeMappingProp
	}
	attributeConditionProp, err := expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(d.Get("attribute_condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attribute_condition"); !isEmptyValue(reflect.ValueOf(attributeConditionProp)) && (ok || !reflect.DeepEqual(v, attributeConditionProp)) {
		obj["attributeCondition"] = attributeConditionProp
	}
	samlProp, err := expandIAMWorkforcePoolWorkforcePoolProviderSaml(d.Get("saml"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("saml"); !isEmptyValue(reflect.ValueOf(samlProp)) && (ok || !reflect.DeepEqual(v, samlProp)) {
		obj["saml"] = samlProp
	}
	oidcProp, err := expandIAMWorkforcePoolWorkforcePoolProviderOidc(d.Get("oidc"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("oidc"); !isEmptyValue(reflect.ValueOf(oidcProp)) && (ok || !reflect.DeepEqual(v, oidcProp)) {
		obj["oidc"] = oidcProp
	}

	return obj, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeMapping(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderAttributeCondition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSaml(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpMetadataXml, err := expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(original["idp_metadata_xml"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpMetadataXml); val.IsValid() && !isEmptyValue(val) {
		transformed["idpMetadataXml"] = transformedIdpMetadataXml
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderSamlIdpMetadataXml(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidc(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIssuerUri, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(original["issuer_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIssuerUri); val.IsValid() && !isEmptyValue(val) {
		transformed["issuerUri"] = transformedIssuerUri
	}

	transformedClientId, err := expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(original["client_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedClientId); val.IsValid() && !isEmptyValue(val) {
		transformed["clientId"] = transformedClientId
	}

	return transformed, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcIssuerUri(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolWorkforcePoolProviderOidcClientId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}