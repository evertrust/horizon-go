package models

func TemplateIndexElementsFromResponse(input []IndexedDNElementResponse) []IndexedDNElement {
	if input == nil {
		return nil
	}
	output := make([]IndexedDNElement, len(input))

	for i := range input {
		output[i] = IndexedDNElement{
			Element: input[i].Element,
			Value:   input[i].Value,
		}
	}

	return output
}

func TemplateSansFromResponse(input []ListSANElementResponse) []ListSANElement {
	if input == nil {
		return nil
	}
	output := make([]ListSANElement, len(input))

	for i := range input {
		output[i] = ListSANElement{
			Type:  input[i].Type,
			Value: input[i].Value,
		}
	}

	return output
}

func TemplateExtensionsFromResponse(input []CertificateExtensionElementResponse) []CertificateExtensionElement {
	if input == nil {
		return nil
	}
	output := make([]CertificateExtensionElement, len(input))

	for i := range input {
		output[i] = CertificateExtensionElement{
			Type:  input[i].Type,
			Value: input[i].Value,
		}
	}

	return output
}

func TemplateLabelsFromResponse(input []RequestLabelElementResponse) []RequestLabelElement {
	if input == nil {
		return nil
	}
	output := make([]RequestLabelElement, len(input))

	for i := range input {
		output[i] = RequestLabelElement{
			Label: input[i].Label,
			Value: input[i].Value,
		}
	}

	return output
}

func TemplateContactEmailFromResponse(input CertificateContactEmailElementResponse) CertificateContactEmailElement {
	return CertificateContactEmailElement{
		Value: input.Value,
	}
}

func TemplateOwnerFromResponse(input CertificateOwnerElementResponse) CertificateOwnerElement {
	return CertificateOwnerElement{
		Value: input.Value,
	}
}

func TemplateTeamFromResponse(input CertificateTeamElementResponse) CertificateTeamElement {
	return CertificateTeamElement{
		Value: input.Value,
	}
}

func TemplateMetadataFromResponse(input []CertificateMetadataElementResponse) []CertificateMetadataElement {
	if input == nil {
		return nil
	}
	output := make([]CertificateMetadataElement, len(input))

	for i := range input {
		output[i] = CertificateMetadataElement{
			Metadata: input[i].Metadata,
			Value:    input[i].Value,
		}
	}

	return output
}

func TemplateAutoRenewFromResponse(input *CertificateAutoRenewElementResponse) *CertificateAutoRenewElement {
	if input == nil {
		return nil
	}
	value := input.Value
	return &CertificateAutoRenewElement{Value: &value}
}

// Template returns the submit template carrying the defaults of the response.
// dnWhitelist and termsOfService are not part of a submit template.
func (o EstEnrollRequestTemplateResponse) Template() EstEnrollRequestTemplate {
	return EstEnrollRequestTemplate{
		ContactEmail: o.ContactEmail,
		Extensions:   o.Extensions,
		Labels:       o.Labels,
		Owner:        o.Owner,
		Sans:         o.Sans,
		Subject:      o.Subject,
		Team:         o.Team,
	}
}

// Template returns the submit template carrying the defaults of the response.
// dnWhitelist and termsOfService are not part of a submit template.
func (o ScepEnrollRequestTemplateResponse) Template() ScepEnrollRequestTemplate {
	return ScepEnrollRequestTemplate{
		ContactEmail: o.ContactEmail,
		Extensions:   o.Extensions,
		Labels:       o.Labels,
		Owner:        o.Owner,
		Sans:         o.Sans,
		Subject:      o.Subject,
		Team:         o.Team,
	}
}

// Template returns the submit template carrying the defaults of the response.
// Capabilities, passwordPolicy and termsOfService are not part of a submit
// template; csr, keyType and asynchronous are left to the caller.
func (o WebRAEnrollRequestTemplateResponse) Template() WebRAEnrollRequestTemplate {
	template := WebRAEnrollRequestTemplate{
		AutoRenew:  TemplateAutoRenewFromResponse(o.AutoRenew),
		Extensions: TemplateExtensionsFromResponse(o.Extensions),
		Labels:     TemplateLabelsFromResponse(o.Labels),
		Metadata:   TemplateMetadataFromResponse(o.Metadata),
		Sans:       TemplateSansFromResponse(o.Sans),
		Subject:    TemplateIndexElementsFromResponse(o.Subject),
	}
	if element := o.ContactEmail.Get(); element != nil {
		contactEmail := TemplateContactEmailFromResponse(*element)
		template.ContactEmail.Set(&contactEmail)
	}
	if element := o.Owner.Get(); element != nil {
		owner := TemplateOwnerFromResponse(*element)
		template.Owner.Set(&owner)
	}
	if element := o.Team.Get(); element != nil {
		team := TemplateTeamFromResponse(*element)
		template.Team.Set(&team)
	}
	return template
}

// Template returns the submit template carrying the defaults of the response.
// Capabilities and passwordPolicy are not part of a submit template; csr and
// asynchronous are left to the caller.
func (o WebRARenewRequestTemplateResponse) Template() WebRARenewRequestTemplate {
	return WebRARenewRequestTemplate{AutoRenew: TemplateAutoRenewFromResponse(o.AutoRenew)}
}
