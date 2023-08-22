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

package certificatemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CertificateManagerTrustConfigAssetType string = "certificatemanager.googleapis.com/TrustConfig"

func ResourceConverterCertificateManagerTrustConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CertificateManagerTrustConfigAssetType,
		Convert:   GetCertificateManagerTrustConfigCaiObject,
	}
}

func GetCertificateManagerTrustConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/trustConfigs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCertificateManagerTrustConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CertificateManagerTrustConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/certificatemanager/v1/rest",
				DiscoveryName:        "TrustConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCertificateManagerTrustConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	labelsProp, err := expandCertificateManagerTrustConfigLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	descriptionProp, err := expandCertificateManagerTrustConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	trustStoresProp, err := expandCertificateManagerTrustConfigTrustStores(d.Get("trust_stores"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("trust_stores"); !tpgresource.IsEmptyValue(reflect.ValueOf(trustStoresProp)) && (ok || !reflect.DeepEqual(v, trustStoresProp)) {
		obj["trustStores"] = trustStoresProp
	}

	return obj, nil
}

func expandCertificateManagerTrustConfigLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerTrustConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerTrustConfigTrustStores(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedTrustAnchors, err := expandCertificateManagerTrustConfigTrustStoresTrustAnchors(original["trust_anchors"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTrustAnchors); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["trustAnchors"] = transformedTrustAnchors
		}

		transformedIntermediateCas, err := expandCertificateManagerTrustConfigTrustStoresIntermediateCas(original["intermediate_cas"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIntermediateCas); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["intermediateCas"] = transformedIntermediateCas
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresTrustAnchors(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPemCertificate, err := expandCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(original["pem_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pemCertificate"] = transformedPemCertificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresTrustAnchorsPemCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerTrustConfigTrustStoresIntermediateCas(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedPemCertificate, err := expandCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(original["pem_certificate"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["pemCertificate"] = transformedPemCertificate
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerTrustConfigTrustStoresIntermediateCasPemCertificate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}