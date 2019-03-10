package credentials

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

// Credentials aws credentials
type Credentials struct {
	file     *ini.File
	filePath string
}

// NewCredentials returns new Credentials
func NewCredentials(credsfile string) (*Credentials, error) {
	credentials := &Credentials{}
	if credsfile == "" || credentials == nil {
		credsfile = os.Getenv("HOME") + "/.aws/credentials"
	}

	if file, err := ini.Load(credsfile); err != nil {
		return nil, err
	} else {
		credentials.file = file
		credentials.filePath = credsfile
		return credentials, nil
	}
}

// GetAwsProfileNames returns list of AWS profile names
func (creds *Credentials) GetAwsProfileNames() ([]string, error) {
	sections := creds.file.SectionStrings()
	if len(sections) == 0 {
		return nil, fmt.Errorf("no profiles on %s", creds.filePath)
	}

	return sections[1:], nil
}

// IsAssumeRole returns true if profile is AssumeRole
func (creds *Credentials) IsAssumeRole(profile string) (bool, error) {
	section, err := creds.file.GetSection(profile)
	if err != nil {
		return false, err
	}

	return section.HasKey("role_arn"), nil
}

// GetValue returns credential value by key
func (creds *Credentials) GetValue(profile string, key string) (string, error) {
	section, err := creds.file.GetSection(profile)
	if err != nil {
		return "", err
	}

	return section.Key(key).String(), err
}
