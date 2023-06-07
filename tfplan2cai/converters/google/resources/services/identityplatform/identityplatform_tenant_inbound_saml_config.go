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

package identityplatform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const IdentityPlatformTenantInboundSamlConfigAssetType string = "identitytoolkit.googleapis.com/TenantInboundSamlConfig"

func ResourceConverterIdentityPlatformTenantInboundSamlConfig() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: IdentityPlatformTenantInboundSamlConfigAssetType,
		Convert:   GetIdentityPlatformTenantInboundSamlConfigCaiObject,
	}
}

func GetIdentityPlatformTenantInboundSamlConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetIdentityPlatformTenantInboundSamlConfigApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: IdentityPlatformTenantInboundSamlConfigAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "TenantInboundSamlConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetIdentityPlatformTenantInboundSamlConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandIdentityPlatformTenantInboundSamlConfigName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandIdentityPlatformTenantInboundSamlConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enabledProp, err := expandIdentityPlatformTenantInboundSamlConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	idpConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfig(d.Get("idp_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("idp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(idpConfigProp)) && (ok || !reflect.DeepEqual(v, idpConfigProp)) {
		obj["idpConfig"] = idpConfigProp
	}
	spConfigProp, err := expandIdentityPlatformTenantInboundSamlConfigSpConfig(d.Get("sp_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("sp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(spConfigProp)) && (ok || !reflect.DeepEqual(v, spConfigProp)) {
		obj["spConfig"] = spConfigProp
	}

	return obj, nil
}

func expandIdentityPlatformTenantInboundSamlConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIdpEntityId, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(original["idp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpEntityId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpEntityId"] = transformedIdpEntityId
	}

	transformedSsoUrl, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(original["sso_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSsoUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ssoUrl"] = transformedSsoUrl
	}

	transformedSignRequest, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(original["sign_request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSignRequest); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["signRequest"] = transformedSignRequest
	}

	transformedIdpCertificates, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(original["idp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIdpCertificates); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["idpCertificates"] = transformedIdpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpEntityId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigSsoUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigSignRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformTenantInboundSamlConfigIdpConfigIdpCertificatesX509Certificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSpEntityId, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(original["sp_entity_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpEntityId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["spEntityId"] = transformedSpEntityId
	}

	transformedCallbackUri, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(original["callback_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCallbackUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["callbackUri"] = transformedCallbackUri
	}

	transformedSpCertificates, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(original["sp_certificates"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSpCertificates); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["spCertificates"] = transformedSpCertificates
	}

	return transformed, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpEntityId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigCallbackUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificates(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedX509Certificate, err := expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(original["x509_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedX509Certificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["x509Certificate"] = transformedX509Certificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIdentityPlatformTenantInboundSamlConfigSpConfigSpCertificatesX509Certificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}