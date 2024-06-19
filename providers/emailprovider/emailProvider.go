package emailprovider

import (
	"GO-08/providers"
	"GO-08/utils"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type emailProvider struct {
	sendgridClient *sendgrid.Client
}

func NewSendGridEmailProvider(apiKey string) providers.EmailProvider {
	return &emailProvider{
		sendgridClient: sendgrid.NewSendClient(apiKey),
	}
}

func (e emailProvider) Send(dt *providers.DynamicTemplate) error {

	for _, recipient := range dt.Recipients {
		newMail := mail.NewV3Mail()

		fromEmail := "noreply@circonomy.co"

		if dt.FromEmail.Valid {
			fromEmail = dt.FromEmail.String
		}
		newMail.From = &mail.Email{
			Name:    "Team Circonomy",
			Address: fromEmail,
		}

		newMail.Personalizations = []*mail.Personalization{}

		personalization := mail.NewPersonalization()
		personalization.To = []*mail.Email{
			{
				Name:    recipient.Name,
				Address: recipient.Address,
			},
		}

		personalization.Categories = []string{"Circonomy"}
		for i := range dt.Categories {
			personalization.Categories = append(personalization.Categories, dt.Categories[i])
		}

		if dt.TemplateID != "" {
			newMail.TemplateID = dt.TemplateID
			for key, value := range dt.DynamicData {
				personalization.DynamicTemplateData[key] = value
			}
		} else { // todo check this part from URL
			htmlContent, err := utils.GetHTMLContent(dt.URL, dt.DynamicData)
			if err != nil {
				logrus.Errorf("error generating content email: %v", err)
				return err
			}
			newMail.AddContent(mail.NewContent("text/html", htmlContent))
			newMail.Subject = getSubject(dt.Subject)
		}

		newMail.AddPersonalizations(personalization)

		// if there is attachment, add it
		if len(dt.Attachments) > 0 {
			newMail.Attachments = dt.Attachments
		}

		resp, err := e.sendgridClient.Send(newMail)
		if resp.StatusCode > 300 {
			logrus.Errorf("error sending email: %v", resp.Body)
			return errors.New("error sending email")
		}
		if err != nil {
			logrus.Errorf("error sending email: %v", err)
			return err
		}
	}

	return nil
}

func getSubject(subject string) string {
	return subject
}

func (e emailProvider) GetEmailTemplate(emailType providers.EmailType) (*providers.DynamicTemplate, error) {
	switch emailType {
	case providers.EmailTypeContactUser:
		return contactUserTemplate()
	default:
		return nil, errors.New("email type invalid")
	}
}
