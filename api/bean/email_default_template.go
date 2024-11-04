package bean

import entity "unibee/internal/model/entity/default"

type Sender struct {
	Name    string `json:"name" description:"name"`
	Address string `json:"address" description:"address"`
}

type EmailDefaultTemplate struct {
	Id                  int64  `json:"id"                 description:""`
	TemplateName        string `json:"templateName"       description:""`
	TemplateDescription string `json:"templateDescription" description:""`
	TemplateTitle       string `json:"templateTitle"      description:""`
	TemplateContent     string `json:"templateContent"    description:""`
	TemplateAttachName  string `json:"templateAttachName" description:""`
}

func SimplifyEmailDefaultTemplate(emailTemplate *entity.EmailDefaultTemplate) *EmailDefaultTemplate {
	if emailTemplate == nil {
		return nil
	}
	return &EmailDefaultTemplate{
		Id:                  emailTemplate.Id,
		TemplateName:        emailTemplate.TemplateName,
		TemplateDescription: emailTemplate.TemplateDescription,
		TemplateTitle:       emailTemplate.TemplateTitle,
		TemplateContent:     emailTemplate.TemplateContent,
		TemplateAttachName:  emailTemplate.TemplateAttachName,
	}
}
