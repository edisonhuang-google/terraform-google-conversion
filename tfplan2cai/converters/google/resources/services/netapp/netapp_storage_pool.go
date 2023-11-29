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

package netapp

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetappstoragePoolAssetType string = "netapp.googleapis.com/storagePool"

func ResourceConverterNetappstoragePool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetappstoragePoolAssetType,
		Convert:   GetNetappstoragePoolCaiObject,
	}
}

func GetNetappstoragePoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//netapp.googleapis.com/projects/{{project}}/locations/{{location}}/storagePools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetappstoragePoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetappstoragePoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/netapp/v1beta1/rest",
				DiscoveryName:        "storagePool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetappstoragePoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	serviceLevelProp, err := expandNetappstoragePoolServiceLevel(d.Get("service_level"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_level"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceLevelProp)) && (ok || !reflect.DeepEqual(v, serviceLevelProp)) {
		obj["serviceLevel"] = serviceLevelProp
	}
	capacityGibProp, err := expandNetappstoragePoolCapacityGib(d.Get("capacity_gib"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("capacity_gib"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityGibProp)) && (ok || !reflect.DeepEqual(v, capacityGibProp)) {
		obj["capacityGib"] = capacityGibProp
	}
	descriptionProp, err := expandNetappstoragePoolDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandNetappstoragePoolNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	activeDirectoryProp, err := expandNetappstoragePoolActiveDirectory(d.Get("active_directory"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("active_directory"); !tpgresource.IsEmptyValue(reflect.ValueOf(activeDirectoryProp)) && (ok || !reflect.DeepEqual(v, activeDirectoryProp)) {
		obj["activeDirectory"] = activeDirectoryProp
	}
	kmsConfigProp, err := expandNetappstoragePoolKmsConfig(d.Get("kms_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsConfigProp)) && (ok || !reflect.DeepEqual(v, kmsConfigProp)) {
		obj["kmsConfig"] = kmsConfigProp
	}
	ldapEnabledProp, err := expandNetappstoragePoolLdapEnabled(d.Get("ldap_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ldap_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(ldapEnabledProp)) && (ok || !reflect.DeepEqual(v, ldapEnabledProp)) {
		obj["ldapEnabled"] = ldapEnabledProp
	}
	labelsProp, err := expandNetappstoragePoolEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetappstoragePoolServiceLevel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolCapacityGib(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolActiveDirectory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolKmsConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolLdapEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappstoragePoolEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}