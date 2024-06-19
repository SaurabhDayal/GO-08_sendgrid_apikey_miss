package providers

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/volatiletech/null"
)

type EmailType string

const (
	EmailTypeContactUser EmailType = "contactUser"
)

type EmailProvider interface {
	Send(dt *DynamicTemplate) error
	GetEmailTemplate(emailType EmailType) (*DynamicTemplate, error)
}

type DynamicTemplate struct {
	TemplateID  string
	DynamicData map[string]interface{}
	Categories  []string
	Recipients  []mail.Email
	FromEmail   null.String
	Attachments []*mail.Attachment
	URL         string
	Subject     string
}

func (d *DynamicTemplate) AddRecipient(name, email string) {
	d.Recipients = append(d.Recipients, mail.Email{
		Name:    name,
		Address: email,
	})
}
