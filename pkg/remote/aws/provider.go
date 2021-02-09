package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/cloudskiff/driftctl/pkg/remote/terraform"
	tf "github.com/cloudskiff/driftctl/pkg/terraform"
)

type awsConfig struct {
	AccessKey     string
	SecretKey     string
	CredsFilename string
	Profile       string
	Token         string
	Region        string `cty:"region"`
	MaxRetries    int

	AssumeRoleARN         string
	AssumeRoleExternalID  string
	AssumeRoleSessionName string
	AssumeRolePolicy      string

	AllowedAccountIds   []string
	ForbiddenAccountIds []string

	Endpoints        map[string]string
	IgnoreTagsConfig map[string]string
	Insecure         bool

	SkipCredsValidation     bool
	SkipGetEC2Platforms     bool
	SkipRegionValidation    bool
	SkipRequestingAccountId bool
	SkipMetadataApiCheck    bool
	S3ForcePathStyle        bool
}

type AWSTerraformProvider struct {
	*terraform.TerraformProvider
	session *session.Session
}

func NewAWSTerraformProvider() (*AWSTerraformProvider, error) {
	p := &AWSTerraformProvider{}
	installer, err := tf.NewProviderInstaller(tf.ProviderConfig{
		Key:     "aws",
		Version: "3.19.0",
		Postfix: "x5",
	})
	if err != nil {
		return nil, err
	}
	p.session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	tfProvider, err := terraform.NewTerraformProvider(installer, terraform.TerraformProviderConfig{
		DefaultAlias: *p.session.Config.Region,
		GetProviderConfig: func(alias string) interface{} {
			return awsConfig{
				Region:     alias,
				MaxRetries: 10, // TODO make this configurable
			}
		},
	})
	if err != nil {
		return nil, err
	}
	p.TerraformProvider = tfProvider
	return p, err
}
