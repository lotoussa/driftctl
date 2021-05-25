// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
	"github.com/zclconf/go-cty/cty"
)

const AwsEipResourceType = "aws_eip"

type AwsEip struct {
	AllocationId           *string           `cty:"allocation_id" computed:"true"`
	AssociateWithPrivateIp *string           `cty:"associate_with_private_ip"`
	AssociationId          *string           `cty:"association_id" computed:"true"`
	CustomerOwnedIp        *string           `cty:"customer_owned_ip" computed:"true"`
	CustomerOwnedIpv4Pool  *string           `cty:"customer_owned_ipv4_pool"`
	Domain                 *string           `cty:"domain" computed:"true"`
	Id                     string            `cty:"id" computed:"true"`
	Instance               *string           `cty:"instance" computed:"true"`
	NetworkBorderGroup     *string           `cty:"network_border_group" computed:"true"`
	NetworkInterface       *string           `cty:"network_interface" computed:"true"`
	PrivateDns             *string           `cty:"private_dns" computed:"true"`
	PrivateIp              *string           `cty:"private_ip" computed:"true"`
	PublicDns              *string           `cty:"public_dns" computed:"true"`
	PublicIp               *string           `cty:"public_ip" computed:"true"`
	PublicIpv4Pool         *string           `cty:"public_ipv4_pool" computed:"true"`
	Tags                   map[string]string `cty:"tags"`
	Vpc                    *bool             `cty:"vpc" computed:"true"`
	Timeouts               *struct {
		Delete *string `cty:"delete"`
		Read   *string `cty:"read"`
		Update *string `cty:"update"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsEip) TerraformId() string {
	return r.Id
}

func (r *AwsEip) TerraformType() string {
	return AwsEipResourceType
}

func (r *AwsEip) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsEipMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsEipResourceType, func(res *resource.AbstractResource) {
		val := res.Attrs
		val.SafeDelete([]string{"timeouts"})
	})
}
