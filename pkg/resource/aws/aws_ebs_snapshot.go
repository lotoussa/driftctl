// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/zclconf/go-cty/cty"
)

const AwsEbsSnapshotResourceType = "aws_ebs_snapshot"

type AwsEbsSnapshot struct {
	Arn                 *string           `cty:"arn" computed:"true"`
	DataEncryptionKeyId *string           `cty:"data_encryption_key_id" computed:"true"`
	Description         *string           `cty:"description"`
	Encrypted           *bool             `cty:"encrypted" computed:"true"`
	Id                  string            `cty:"id" diff:"Id,identifier" computed:"true"`
	KmsKeyId            *string           `cty:"kms_key_id" computed:"true"`
	OwnerAlias          *string           `cty:"owner_alias" computed:"true"`
	OwnerId             *string           `cty:"owner_id" computed:"true"`
	Tags                map[string]string `cty:"tags"`
	VolumeId            *string           `cty:"volume_id"`
	VolumeSize          *int              `cty:"volume_size" computed:"true"`
	Timeouts            *struct {
		Create *string `cty:"create"`
		Delete *string `cty:"delete"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsEbsSnapshot) TerraformId() string {
	return r.Id
}

func (r *AwsEbsSnapshot) TerraformType() string {
	return AwsEbsSnapshotResourceType
}

func (r *AwsEbsSnapshot) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsEbsSnapshotMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsEbsSnapshotResourceType, func(res *resource.AbstractResource) {
		val := res.Attrs
		val.SafeDelete([]string{"timeouts"})
	})
}
