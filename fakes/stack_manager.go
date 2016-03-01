package fakes

import (
	"time"

	"github.com/pivotal-cf-experimental/bosh-bootloader/aws/cloudformation"
)

type StackManager struct {
	CreateOrUpdateCall struct {
		Receives struct {
			StackName string
			Template  cloudformation.Template
			Client    cloudformation.Client
		}
		Returns struct {
			Error error
		}
	}

	WaitForCompletionCall struct {
		Receives struct {
			Client        cloudformation.Client
			StackName     string
			SleepInterval time.Duration
		}
		Returns struct {
			Error error
		}
	}
}

func (m *StackManager) CreateOrUpdate(client cloudformation.Client, stackName string, template cloudformation.Template) error {
	m.CreateOrUpdateCall.Receives.Client = client
	m.CreateOrUpdateCall.Receives.StackName = stackName
	m.CreateOrUpdateCall.Receives.Template = template

	return m.CreateOrUpdateCall.Returns.Error
}

func (m *StackManager) WaitForCompletion(client cloudformation.Client, stackName string, sleepInterval time.Duration) error {
	m.WaitForCompletionCall.Receives.Client = client
	m.WaitForCompletionCall.Receives.StackName = stackName
	m.WaitForCompletionCall.Receives.SleepInterval = sleepInterval

	return m.WaitForCompletionCall.Returns.Error
}
