package service

import (
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/tukaelu/opsworks-cli/credentials"
)

// OpsWorksResource id/name
type OpsWorksResource struct {
	ID   string
	Name string
}

// Service service client
type Service struct {
	credentials *credentials.Credentials
	profile     string
	client      *opsworks.OpsWorks
}

func New(creds *credentials.Credentials, profile string) *Service {
	service := &Service{}

	return service
}
