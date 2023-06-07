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

package cloudfunctions2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const Cloudfunctions2functionIAMAssetType string = "cloudfunctions.googleapis.com/function"

func ResourceConverterCloudfunctions2functionIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudfunctions2functionIamPolicy,
	}
}

func ResourceConverterCloudfunctions2functionIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamBindingCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamBinding,
		MergeDelete:       MergeCloudfunctions2functionIamBindingDelete,
	}
}

func ResourceConverterCloudfunctions2functionIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamMemberCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamMember,
		MergeDelete:       MergeCloudfunctions2functionIamMemberDelete,
	}
}

func GetCloudfunctions2functionIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetCloudfunctions2functionIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetCloudfunctions2functionIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeCloudfunctions2functionIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudfunctions2functionIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeCloudfunctions2functionIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newCloudfunctions2functionIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: Cloudfunctions2functionIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudfunctions2functionIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("cloud_function"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		Cloudfunctions2functionIamUpdaterProducer,
		d,
		config,
		"//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}",
		Cloudfunctions2functionIAMAssetType,
	)
}