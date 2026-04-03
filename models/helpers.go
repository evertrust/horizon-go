package models

func TemplateIndexElementsFromResponse(input []IndexedDNElementResponse) []IndexedDNElement {
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
	output := make([]CertificateMetadataElement, len(input))

	for i := range input {
		output[i] = CertificateMetadataElement{
			Metadata: input[i].Metadata,
			Value:    input[i].Value,
		}
	}

	return output
}
