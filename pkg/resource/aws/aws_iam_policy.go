// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/helpers"
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/zclconf/go-cty/cty"
)

const AwsIamPolicyResourceType = "aws_iam_policy"

type AwsIamPolicy struct {
	Arn         *string    `cty:"arn" computed:"true"`
	Description *string    `cty:"description"`
	Id          string     `cty:"id" computed:"true"`
	Name        *string    `cty:"name" computed:"true"`
	NamePrefix  *string    `cty:"name_prefix" diff:"-"`
	Path        *string    `cty:"path"`
	Policy      *string    `cty:"policy" jsonstring:"true"`
	CtyVal      *cty.Value `diff:"-"`
}

func (r *AwsIamPolicy) TerraformId() string {
	return r.Id
}

func (r *AwsIamPolicy) TerraformType() string {
	return AwsIamPolicyResourceType
}

func (r *AwsIamPolicy) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsIAMPolicyMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.UpdateSchema(AwsIamPolicyResourceType, map[string]func(attributeSchema *resource.AttributeSchema){
		"policy": func(attributeSchema *resource.AttributeSchema) {
			attributeSchema.JsonString = true
		},
	})
	resourceSchemaRepository.SetNormalizeFunc(AwsIamPolicyResourceType, func(res *resource.AbstractResource) {
		val := res.Attrs
		jsonString, err := helpers.NormalizeJsonString((*val)["policy"])
		if err == nil {
			val.SafeSet([]string{"policy"}, jsonString)
		}

		val.SafeDelete([]string{"name_prefix"})
	})
}
