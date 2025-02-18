/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer( "%5B", "[", "%5D", "]" )
)

type Requests struct {
    c *APIClient
}

// APIClient manages communication with the Horizon API API v2.7.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {

	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.
    Requests *Requests

	// API Services

	AdocAPI *AdocAPIService

	AutomationExecutionAPI *AutomationExecutionAPIService

	AutomationPolicyAPI *AutomationPolicyAPIService

	CaAPI *CaAPIService

	CacheAPI *CacheAPIService

	CertificateAPI *CertificateAPIService

	CertificateAnalyticsAPI *CertificateAnalyticsAPIService

	CertificateGradingPolicyAPI *CertificateGradingPolicyAPIService

	CertificateGradingRulesetAPI *CertificateGradingRulesetAPIService

	CertificateLabelAPI *CertificateLabelAPIService

	CertificateProfileAPI *CertificateProfileAPIService

	DatasourceAPI *DatasourceAPIService

	DatasourceFlowAPI *DatasourceFlowAPIService

	DiscoveryCampaignAPI *DiscoveryCampaignAPIService

	DiscoveryEventAPI *DiscoveryEventAPIService

	DiscoveryEventAnalyticsAPI *DiscoveryEventAnalyticsAPIService

	DiscoveryFeedAPI *DiscoveryFeedAPIService

	EventAPI *EventAPIService

	EventAnalyticsAPI *EventAnalyticsAPIService

	HttpHttpproxyAPI *HttpHttpproxyAPIService

	LicenseAPI *LicenseAPIService

	PkiConnectorAPI *PkiConnectorAPIService

	PkiQueueAPI *PkiQueueAPIService

	RequestAPI *RequestAPIService

	Rfc5280API *Rfc5280APIService

	SchedulerTaskAPI *SchedulerTaskAPIService

	SecurityCredentialsAPI *SecurityCredentialsAPIService

	SecurityIdentityLocalAPI *SecurityIdentityLocalAPIService

	SecurityIdentityProviderAPI *SecurityIdentityProviderAPIService

	SecurityPasswordpolicyAPI *SecurityPasswordpolicyAPIService

	SecurityPrincipalAPI *SecurityPrincipalAPIService

	SecurityPrincipalinfoAPI *SecurityPrincipalinfoAPIService

	SecurityRoleAPI *SecurityRoleAPIService

	SecurityScimprofileAPI *SecurityScimprofileAPIService

	SecurityTeamAPI *SecurityTeamAPIService

	SystemConfigurationAPI *SystemConfigurationAPIService

	TemplatestringAPI *TemplatestringAPIService

	ThirdpartyConnectorAPI *ThirdpartyConnectorAPIService

	TriggerAPI *TriggerAPIService

	TrustchainAPI *TrustchainAPIService

	WcceAPI *WcceAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c
    c.Requests = &Requests{c: c}

	// API Services
	c.AdocAPI = (*AdocAPIService)(&c.common)
	c.AutomationExecutionAPI = (*AutomationExecutionAPIService)(&c.common)
	c.AutomationPolicyAPI = (*AutomationPolicyAPIService)(&c.common)
	c.CaAPI = (*CaAPIService)(&c.common)
	c.CacheAPI = (*CacheAPIService)(&c.common)
	c.CertificateAPI = (*CertificateAPIService)(&c.common)
	c.CertificateAnalyticsAPI = (*CertificateAnalyticsAPIService)(&c.common)
	c.CertificateGradingPolicyAPI = (*CertificateGradingPolicyAPIService)(&c.common)
	c.CertificateGradingRulesetAPI = (*CertificateGradingRulesetAPIService)(&c.common)
	c.CertificateLabelAPI = (*CertificateLabelAPIService)(&c.common)
	c.CertificateProfileAPI = (*CertificateProfileAPIService)(&c.common)
	c.DatasourceAPI = (*DatasourceAPIService)(&c.common)
	c.DatasourceFlowAPI = (*DatasourceFlowAPIService)(&c.common)
	c.DiscoveryCampaignAPI = (*DiscoveryCampaignAPIService)(&c.common)
	c.DiscoveryEventAPI = (*DiscoveryEventAPIService)(&c.common)
	c.DiscoveryEventAnalyticsAPI = (*DiscoveryEventAnalyticsAPIService)(&c.common)
	c.DiscoveryFeedAPI = (*DiscoveryFeedAPIService)(&c.common)
	c.EventAPI = (*EventAPIService)(&c.common)
	c.EventAnalyticsAPI = (*EventAnalyticsAPIService)(&c.common)
	c.HttpHttpproxyAPI = (*HttpHttpproxyAPIService)(&c.common)
	c.LicenseAPI = (*LicenseAPIService)(&c.common)
	c.PkiConnectorAPI = (*PkiConnectorAPIService)(&c.common)
	c.PkiQueueAPI = (*PkiQueueAPIService)(&c.common)
	c.RequestAPI = (*RequestAPIService)(&c.common)
	c.Rfc5280API = (*Rfc5280APIService)(&c.common)
	c.SchedulerTaskAPI = (*SchedulerTaskAPIService)(&c.common)
	c.SecurityCredentialsAPI = (*SecurityCredentialsAPIService)(&c.common)
	c.SecurityIdentityLocalAPI = (*SecurityIdentityLocalAPIService)(&c.common)
	c.SecurityIdentityProviderAPI = (*SecurityIdentityProviderAPIService)(&c.common)
	c.SecurityPasswordpolicyAPI = (*SecurityPasswordpolicyAPIService)(&c.common)
	c.SecurityPrincipalAPI = (*SecurityPrincipalAPIService)(&c.common)
	c.SecurityPrincipalinfoAPI = (*SecurityPrincipalinfoAPIService)(&c.common)
	c.SecurityRoleAPI = (*SecurityRoleAPIService)(&c.common)
	c.SecurityScimprofileAPI = (*SecurityScimprofileAPIService)(&c.common)
	c.SecurityTeamAPI = (*SecurityTeamAPIService)(&c.common)
	c.SystemConfigurationAPI = (*SystemConfigurationAPIService)(&c.common)
	c.TemplatestringAPI = (*TemplatestringAPIService)(&c.common)
	c.ThirdpartyConnectorAPI = (*ThirdpartyConnectorAPIService)(&c.common)
	c.TriggerAPI = (*TriggerAPIService)(&c.common)
	c.TrustchainAPI = (*TrustchainAPIService)(&c.common)
	c.WcceAPI = (*WcceAPIService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("expected %s to be of type %s but received %s", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

func parameterValueToString( obj interface{}, key string ) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return fmt.Sprintf("%v", obj)
	}
	var param,ok = obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap,err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	var value = ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
			case reflect.Invalid:
				value = "invalid"

			case reflect.Struct:
				if t,ok := obj.(MappedNullable); ok {
					dataMap,err := t.ToMap()
					if err != nil {
						return
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
					return
				}
				if t, ok := obj.(time.Time); ok {
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
					return
				}
				value = v.Type().String() + " value"
			case reflect.Slice:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				var lenIndValue = indValue.Len()
				for i:=0;i<lenIndValue;i++ {
					var arrayValue = indValue.Index(i)
					var keyPrefixForCollectionType = keyPrefix
					if style == "deepObject" {
						keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
					}
					parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
				}
				return

			case reflect.Map:
				var indValue = reflect.ValueOf(obj)
				if indValue == reflect.ValueOf(nil) {
					return
				}
				iter := indValue.MapRange()
				for iter.Next() {
					k,v := iter.Key(), iter.Value()
					parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
				}
				return

			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
				return

			case reflect.Int, reflect.Int8, reflect.Int16,
				reflect.Int32, reflect.Int64:
				value = strconv.FormatInt(v.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				value = strconv.FormatUint(v.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
			case reflect.Bool:
				value = strconv.FormatBool(v.Bool())
			case reflect.String:
				value = v.String()
			default:
				value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
		case url.Values:
			if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
				valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix) + "," + value)
			} else {
				valuesMap.Add(keyPrefix, value)
			}
			break
		case map[string]string:
			valuesMap[keyPrefix] = value
			break
	}
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
		fileBytes []byte
		fileName string
		formFileName string
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
						return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
						return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(*os.File); ok {
		f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = f.Write(b)
		if err != nil {
			return
		}
		_, err = f.Seek(0, io.SeekStart)
		return
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v interface{}) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}


////////////////// Request API Helpers //////////////////

// region Webra Enroll

type WebraEnrollTemplate struct {
	// Describes how certificates will be enrolled on this profile
	Capabilities CertificateProfileCryptoPolicy `json:"capabilities,omitempty"`
	// The password policy that will be used to generate the certificate's PKCS#12 password
	PasswordPolicy PasswordPolicy `json:"passwordPolicy,omitempty"`
	// The type of key that will be used to generate the certificate, if in centralized mode
	KeyType string `json:"keyType,omitempty"`
	// If decentralized enrollment is enabled, this field will contain the CSR that will be used to generate the certificate
	Csr string `json:"csr,omitempty"`
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElement `json:"subject,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElement `json:"sans,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner CertificateOwnerElement `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team CertificateTeamElement `json:"team,omitempty"`
	// The technical metadata for this certificate
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
}

type GetEnrollTemplateRequest struct {
	Profile string `json:"profile"`
	Csr     string `json:"csr"`
	c       *APIClient
}

func (r GetEnrollTemplateRequest) Execute() (*WebraEnrollTemplate, *http.Response, error) {
	var req WebRAEnrollRequestOnTemplate
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_WEBRA)
	req.Profile = *NewNullableString(&r.Profile)
	if r.Csr != "" {
		templateMap := make(map[string]interface{})
		req.Template = &templateMap
		templateMap["csr"] = r.Csr
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRAEnrollRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraEnrollTemplate
	templateData, err := resp.WebRAEnrollRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetEnrollTemplate(profile string) GetEnrollTemplateRequest {
	return GetEnrollTemplateRequest{
		Profile: profile,
		c:       r.c,
	}
}

func (r *Requests) GetEnrollTemplateWithCsr(profile, csr string) GetEnrollTemplateRequest {
	return GetEnrollTemplateRequest{
		Profile: profile,
		Csr:     csr,
		c:       r.c,
	}
}

type WebRAEnrollResponse struct {
	// The module that will be used to process this request. For a WebRA request, this is always `webra`
	Module string `json:"module"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template WebRAEnrollRequestTemplate `json:"template"`
	// The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode
	Pkcs12 SecretString `json:"pkcs12,omitempty"`
	// The password to decrypt the PKCS12 file.
	Password SecretString `json:"password,omitempty"`
	// The certificate that was generated for this request. This is only available after the request has been approved
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitEnrollRequest struct {
	Profile          string
	Template         *WebraEnrollTemplate
	RequesterComment string
	c                *APIClient
}

func (r SubmitEnrollRequest) Execute() (*WebRAEnrollResponse, *http.Response, error) {
	var req WebRAEnrollRequestOnSubmit
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_WEBRA)
	req.Profile = r.Profile
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	if r.Template.KeyType != "" {
		req.Template.KeyType = *NewNullableString(&r.Template.KeyType)
	}
	if r.Template.Csr != "" {
		req.Template.Csr = *NewNullableString(&r.Template.Csr)
	}
	req.Template.Subject = r.Template.Subject
	req.Template.Sans = r.Template.Sans
	req.Template.Extensions = r.Template.Extensions
	req.Template.Labels = r.Template.Labels
	if r.Template.ContactEmail.HasValue() {
		req.Template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
	}
	if r.Template.Owner.HasValue() {
		req.Template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
	}
	if r.Template.Team.HasValue() {
		req.Template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
	}
	req.Template.Metadata = r.Template.Metadata

	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRAEnrollRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAEnrollResponse
	responseData, err := resp.WebRAEnrollRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitEnroll(profile, requesterComment string, template *WebraEnrollTemplate) SubmitEnrollRequest {
	return SubmitEnrollRequest{
		Profile:          profile,
		Template:         template,
		RequesterComment: requesterComment,
		c:                r.c,
	}
}

type WebRAEnrollRequest struct {
	Module   string `json:"module"`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template WebRAEnrollRequestTemplate `json:"template"`
	// The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode
	Pkcs12 SecretString `json:"pkcs12,omitempty"`
	// The password to decrypt the PKCS12 file.
	Password SecretString `json:"password,omitempty"`
	// The certificate that was generated for this request. This is only available after the request has been approved
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type GetEnrollRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetEnrollRequestRequest) Execute() (*WebRAEnrollResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAEnrollResponse
	responseData, err := resp.WebRAEnrollRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetEnrollRequest(id string) GetEnrollRequestRequest {
	return GetEnrollRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelEnrollRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelEnrollRequestRequest) Execute() (*WebRAEnrollResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_ENROLL
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAEnrollResponse
	responseData, err := resp.WebRAEnrollRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelEnrollRequest(id string, module string) CancelEnrollRequestRequest {
	return CancelEnrollRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type WebRAEnrollApprove struct {
	// Object internal ID
	Id string `json:"_id"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The module that will be used to process this request. For a WebRA request, this is always `webra`
	Module string `json:"module"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template *WebRAEnrollRequestTemplate `json:"template,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type ApproveEnrollRequestRequest struct {
	Id              string
	ApproverComment string
	Template        *WebraEnrollTemplate
	c               *APIClient
}

func (r ApproveEnrollRequestRequest) Execute() (*WebRAEnrollResponse, *http.Response, error) {
	var req WebRAEnrollRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_WEBRA)
	if r.ApproverComment != "" {
		req.ApproverComment = *NewNullableString(&r.ApproverComment)
	}
	if r.Template != nil {
		var template WebRAEnrollRequestTemplate
		if r.Template.KeyType != "" {
			template.KeyType = *NewNullableString(&r.Template.KeyType)
		}
		if r.Template.Csr != "" {
			template.Csr = *NewNullableString(&r.Template.Csr)
		}
		template.Subject = r.Template.Subject
		template.Sans = r.Template.Sans
		template.Extensions = r.Template.Extensions
		template.Labels = r.Template.Labels
		if r.Template.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
		}
		if r.Template.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
		}
		if r.Template.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
		}
		template.Metadata = r.Template.Metadata
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRAEnrollRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAEnrollResponse
	responseData, err := resp.WebRAEnrollRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveEnrollRequest(id, approverComment string) ApproveEnrollRequestRequest {
	return ApproveEnrollRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Template:        &WebraEnrollTemplate{},
		c:               r.c,
	}
}

func (r *Requests) ApproveEnrollRequestWithTemplate(id, approverComment string, template *WebraEnrollTemplate) ApproveEnrollRequestRequest {
	return ApproveEnrollRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Template:        template,
		c:               r.c,
	}
}

type DenyEnrollRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyEnrollRequestRequest) Execute() (*WebRAEnrollResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_ENROLL
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAEnrollResponse
	responseData, err := resp.WebRAEnrollRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyEnrollRequest(id, approverComment, module string) DenyEnrollRequestRequest {
	return DenyEnrollRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Revoke

type WebraRevokeTemplate struct {
	// One of: `unspecified`, `keycompromise`, `cacompromise`, `affiliationchange`, `superseded`, `cessationofoperation`
	RevocationReason string `json:"revocationReason,omitempty"`
}

type GetRevokeTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Profile        string
	Module         string
	c              *APIClient
}

func (r GetRevokeTemplateRequest) Execute() (*WebraRevokeTemplate, *http.Response, error) {
	var req WebRARevokeRequestOnTemplate
	req.Workflow = string(WORKFLOW_REVOKE)
	req.Module = &r.Module
	req.Profile = &r.Profile
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRARevokeRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraRevokeTemplate
	templateData, err := resp.WebRARevokeRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetRevokeTemplate(profile, module string) GetRevokeTemplateRequest {
	return GetRevokeTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Profile:        profile,
		Module:         module,
		c:              r.c,
	}
}

func (r *Requests) GetRevokeTemplateWithCertificateId(certificateId string) GetRevokeTemplateRequest {
	return GetRevokeTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		c:              r.c,
	}
}

func (r *Requests) GetRevokeTemplateWithCertificatePem(certificatePem string) GetRevokeTemplateRequest {
	return GetRevokeTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		c:              r.c,
	}
}

type WebRARevokeResponse struct {
	// The module of the certificate revoked.
	Module Module `json:"module"`
	// What this request will do. For a revocation request, this is always `revoke`
	Workflow string `json:"workflow"`
	// The user-data that was used to revoke the certificate
	Template WebRARevokeRequestTemplate `json:"template"`
	// The certificate that was revoked for this request. This is only available after the request has been approved
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitRevokeRequest struct {
	CertificateId  string
	CertificatePem string
	Template       *WebraRevokeTemplate
	c              *APIClient
}

func (r SubmitRevokeRequest) Execute() (*WebRARevokeResponse, *http.Response, error) {
	var req WebRARevokeRequestOnSubmit
	req.Workflow = string(WORKFLOW_REVOKE)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRARevokeRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARevokeResponse
	responseData, err := resp.WebRARevokeRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitRevokeWithCertificateId(certificateId string, template *WebraRevokeTemplate) SubmitRevokeRequest {
	return SubmitRevokeRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Template:       template,
		c:              r.c,
	}
}

func (r *Requests) SubmitRevokeWithCertificatePem(certificatePem string, template *WebraRevokeTemplate) SubmitRevokeRequest {
	return SubmitRevokeRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Template:       template,
		c:        r.c,
	}
}

type GetRevokeRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetRevokeRequestRequest) Execute() (*WebRARevokeResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARevokeResponse
	responseData, err := resp.WebRARevokeRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetRevokeRequest(id string) GetRevokeRequestRequest {
	return GetRevokeRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelRevokeRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelRevokeRequestRequest) Execute() (*WebRARevokeResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_REVOKE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARevokeResponse
	responseData, err := resp.WebRARevokeRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelRevokeRequest(id string, module string) CancelRevokeRequestRequest {
	return CancelRevokeRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveRevokeRequestRequest struct {
	Id             string
	AppoverComment string
	Template       *WebraRevokeTemplate
	c              *APIClient
}

func (r ApproveRevokeRequestRequest) Execute() (*WebRARevokeResponse, *http.Response, error) {
	var req WebRARevokeRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_REVOKE)
	if r.AppoverComment != "" {
		req.ApproverComment = *NewNullableString(&r.AppoverComment)
	}
	if r.Template != nil {
		var template WebRARevokeRequestTemplate
		if r.Template.RevocationReason != "" {
			template.RevocationReason = *NewNullableString(&r.Template.RevocationReason)
		}
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRARevokeRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARevokeResponse
	responseData, err := resp.WebRARevokeRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveRevokeRequest(id string, template *WebraRevokeTemplate) ApproveRevokeRequestRequest {
	return ApproveRevokeRequestRequest{
		Id:       id,
		Template: template,
		c:        r.c,
	}
}

func (r *Requests) ApproveRevokeRequestWithComment(id, comment string, template *WebraRevokeTemplate) ApproveRevokeRequestRequest {
	return ApproveRevokeRequestRequest{
		Id:             id,
		AppoverComment: comment,
		Template:       template,
		c:              r.c,
	}
}

type DenyRevokeRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyRevokeRequestRequest) Execute() (*WebRARevokeResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_REVOKE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARevokeResponse
	responseData, err := resp.WebRARevokeRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyRevokeRequest(id, approverComment, module string) DenyRevokeRequestRequest {
	return DenyRevokeRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Update

type WebraUpdateTemplate struct {
	// Information about the certificate's labels and how to edit them
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// Information about the certificate's metadata and how to edit them
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner CertificateOwnerElement `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team CertificateTeamElement `json:"team,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
}

type GetUpdateTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Profile        string
	Module         string
	c              *APIClient
}

func (r GetUpdateTemplateRequest) Execute() (*WebraUpdateTemplate, *http.Response, error) {
	var req WebRAUpdateRequestOnTemplate
	req.Workflow = string(WORKFLOW_UPDATE)
	req.Module = &r.Module
	if r.Profile != "" {
		req.Profile = *NewNullableString(&r.Profile)
	}
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRAUpdateRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraUpdateTemplate
	templateData, err := resp.WebRAUpdateRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetUpdateTemplate(profile, module string) GetUpdateTemplateRequest {
	return GetUpdateTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Profile:        profile,
		Module:         module,
		c:              r.c,
	}
}

func (r *Requests) GetUpdateTemplateWithCertificateId(certificateId string) GetUpdateTemplateRequest {
	return GetUpdateTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Profile:        "",
		c:              r.c,
	}
}

func (r *Requests) GetUpdateTemplateWithCertificatePem(certificatePem string) GetUpdateTemplateRequest {
	return GetUpdateTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Profile:        "",
		c:              r.c,
	}
}

type WebRAUpdateResponse struct {
	// The module of the certificate updated.
	Module Module `json:"module"`
	// What this request will do. For an update request, this is always `update`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template WebRAUpdateRequestTemplate `json:"template"`
	// The certificate that was updated for this request. This is only available after the request has been approved
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitUpdateRequest struct {
	CertificateId  string
	CertificatePem string
	Template       *WebraUpdateTemplate
	c              *APIClient
}

func (r SubmitUpdateRequest) Execute() (*WebRAUpdateResponse, *http.Response, error) {
	var req WebRAUpdateRequestOnSubmit
	req.Workflow = string(WORKFLOW_UPDATE)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	if r.Template != nil {
		var template WebRAUpdateRequestTemplate
		template.Labels = r.Template.Labels
		template.Metadata = r.Template.Metadata
		if r.Template.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
		}
		if r.Template.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
		}
		if r.Template.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
		}
		req.Template = template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRAUpdateRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAUpdateResponse
	responseData, err := resp.WebRAUpdateRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitUpdateWithCertificateId(certificateId string, template *WebraUpdateTemplate) SubmitUpdateRequest {
	return SubmitUpdateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Template:       template,
		c:              r.c,
	}
}

func (r *Requests) SubmitUpdateWithCertificatePem(certificatePem string, template *WebraUpdateTemplate) SubmitUpdateRequest {
	return SubmitUpdateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Template:       template,
		c:              r.c,
	}
}

type GetUpdateRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetUpdateRequestRequest) Execute() (*WebRAUpdateResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAUpdateResponse
	responseData, err := resp.WebRAUpdateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetUpdateRequest(id string) GetUpdateRequestRequest {
	return GetUpdateRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelUpdateRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelUpdateRequestRequest) Execute() (*WebRAUpdateResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_UPDATE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAUpdateResponse
	responseData, err := resp.WebRAUpdateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelUpdateRequest(id string, module string) CancelUpdateRequestRequest {
	return CancelUpdateRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveUpdateRequestRequest struct {
	Id              string
	ApproverComment string
	Template        *WebraUpdateTemplate
	c               *APIClient
}

func (r ApproveUpdateRequestRequest) Execute() (*WebRAUpdateResponse, *http.Response, error) {
	var req WebRAUpdateRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_UPDATE)
	if r.ApproverComment != "" {
		req.ApproverComment = *NewNullableString(&r.ApproverComment)
	}
	if r.Template != nil {
		var template WebRAUpdateRequestTemplate
		template.Labels = r.Template.Labels
		template.Metadata = r.Template.Metadata
		if r.Template.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
		}
		if r.Template.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
		}
		if r.Template.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
		}
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRAUpdateRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAUpdateResponse
	responseData, err := resp.WebRAUpdateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveUpdateRequest(id, approverComment string) ApproveUpdateRequestRequest {
	return ApproveUpdateRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		c:               r.c,
	}
}

func (r *Requests) ApproveUpdateRequestWithTemplate(id, approverComment string, template *WebraUpdateTemplate) ApproveUpdateRequestRequest {
	return ApproveUpdateRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Template:        template,
		c:               r.c,
	}
}

type DenyUpdateRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyUpdateRequestRequest) Execute() (*WebRAUpdateResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_UPDATE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAUpdateResponse
	responseData, err := resp.WebRAUpdateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyUpdateRequest(id, approverComment, module string) DenyUpdateRequestRequest {
	return DenyUpdateRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Recover

type WebraRecoverTemplate struct {
	// The password mode of the certificate
	PasswordMode string `json:"passwordMode,omitempty"`
	// The selected password policy for this profile. If none is defined and the password mode is `manual`, there is no constraint on the password. In `random` mode, the `Horizon-Default` policy is used
	PasswordPolicy PasswordPolicy `json:"passwordPolicy,omitempty"`
}

type GetRecoverTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Profile        string
	Module         string
	c              *APIClient
}

func (r GetRecoverTemplateRequest) Execute() (*WebraRecoverTemplate, *http.Response, error) {
	var req WebRARecoverRequestOnTemplate
	req.Workflow = string(WORKFLOW_RECOVER)
	req.Module = &r.Module
	if r.Profile != "" {
		req.Profile = r.Profile
	}
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRARecoverRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraRecoverTemplate
	templateData, err := resp.WebRARecoverRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetRecoverTemplate(profile, module string) GetRecoverTemplateRequest {
	return GetRecoverTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Profile:        profile,
		Module:         module,
		c:              r.c,
	}
}

func (r *Requests) GetRecoverTemplateWithCertificateId(certificateId string) GetRecoverTemplateRequest {
	return GetRecoverTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		c:              r.c,
	}
}

func (r *Requests) GetRecoverTemplateWithCertificatePem(certificatePem string) GetRecoverTemplateRequest {
	return GetRecoverTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		c:              r.c,
	}
}

type WebRARecoverResponse struct {
	// The module of the certificate recovered.
	Module Module `json:"module"`
	// What this request will do. For a recovery request, this is always `recover`
	Workflow string `json:"workflow"`
	// The generated PKCS#12 for this request. This is only available after the request has been approved.
	Pkcs12 SecretString `json:"pkcs12,omitempty"`
	// The password to decrypt the PKCS12 file.
	Password SecretString `json:"password,omitempty"`
	// The certificate that was recovered.
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitRecoverRequest struct {
	CertificateId    string
	CertificatePem   string
	RequesterComment string
	Password         string
	c                *APIClient
}

func (r SubmitRecoverRequest) Execute() (*WebRARecoverResponse, *http.Response, error) {
	var req WebRARecoverRequestOnSubmit
	req.Workflow = string(WORKFLOW_RECOVER)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	if r.Password != "" {
		req.Password = *NewNullableSecretString(&SecretString{Value: *NewNullableString(&r.Password)})
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRARecoverRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARecoverResponse
	responseData, err := resp.WebRARecoverRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitRecoverWithCertificateId(certificateId, password, requesterComment string) SubmitRecoverRequest {
	return SubmitRecoverRequest{
		CertificateId:    certificateId,
		CertificatePem:   "",
		RequesterComment: requesterComment,
		Password:         password,
		c:                r.c,
	}
}

func (r *Requests) SubmitRecoverWithCertificatePem(certificatePem, password, requesterComment string) SubmitRecoverRequest {
	return SubmitRecoverRequest{
		CertificateId:    "",
		CertificatePem:   certificatePem,
		RequesterComment: requesterComment,
		Password:         password,
		c:                r.c,
	}
}

type GetRecoverRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetRecoverRequestRequest) Execute() (*WebRARecoverResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARecoverResponse
	responseData, err := resp.WebRARecoverRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetRecoverRequest(id string) GetRecoverRequestRequest {
	return GetRecoverRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelRecoverRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelRecoverRequestRequest) Execute() (*WebRARecoverResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_RECOVER
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARecoverResponse
	responseData, err := resp.WebRARecoverRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelRecoverRequest(id string, module string) CancelRecoverRequestRequest {
	return CancelRecoverRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveRecoverRequestRequest struct {
	Id             string
	AppoverComment string
	c              *APIClient
}

func (r ApproveRecoverRequestRequest) Execute() (*WebRARecoverResponse, *http.Response, error) {
	var req WebRARecoverRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_RECOVER)
	if r.AppoverComment != "" {
		req.ApproverComment = *NewNullableString(&r.AppoverComment)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRARecoverRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARecoverResponse
	responseData, err := resp.WebRARecoverRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveRecoverRequest(id, comment string) ApproveRecoverRequestRequest {
	return ApproveRecoverRequestRequest{
		Id:             id,
		AppoverComment: comment,
		c:              r.c,
	}
}

type DenyRecoverRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyRecoverRequestRequest) Execute() (*WebRARecoverResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_RECOVER
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARecoverResponse
	responseData, err := resp.WebRARecoverRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyRecoverRequest(id, approverComment, module string) DenyRecoverRequestRequest {
	return DenyRecoverRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Migrate

type WebraMigrateTemplate struct {
	Labels       []RequestLabelElement          `json:"labels,omitempty"`
	Owner        CertificateOwnerElement        `json:"owner,omitempty"`
	Team         CertificateTeamElement         `json:"team,omitempty"`
	Metadata     []CertificateMetadataElement   `json:"metadata,omitempty"`
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
}

type GetMigrateTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Profile        string
	Module         string
}

func (r GetMigrateTemplateRequest) Execute(c *APIClient) (*WebraMigrateTemplate, *http.Response, error) {
	var req WebRAMigrateRequestOnTemplate
	req.Workflow = string(WORKFLOW_MIGRATE)
	req.Module = &r.Module
	if r.Profile != "" {
		req.Profile = r.Profile
	}
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRAMigrateRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraMigrateTemplate
	templateData, err := resp.WebRAMigrateRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetMigrateTemplate(profile, module string) GetMigrateTemplateRequest {
	return GetMigrateTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Profile:        profile,
		Module:         module,
	}
}

func (r *Requests) GetMigrateTemplateWithCertificateId(certificateId string) GetMigrateTemplateRequest {
	return GetMigrateTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Profile:        "",
		Module:         "",
	}
}

func (r *Requests) GetMigrateTemplateWithCertificatePem(certificatePem string) GetMigrateTemplateRequest {
	return GetMigrateTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Profile:        "",
		Module:         "",
	}
}

type WebRAMigrateResponse struct {
	// The module of the certificate migrated.
	Module Module `json:"module"`
	// What this request will do. For a migration request, this is always `migrate`
	Workflow string `json:"workflow"`
	// The target profile name
	Profile interface{} `json:"profile"`
	// The user-data that will be used to generate the certificate
	Template WebRAMigrateRequestTemplate `json:"template"`
	// The certificate that was updated for this request. This is only available after the request has been approved
	Certificate NullableCertificate `json:"certificate,omitempty"`
	// If true, the request is validated, but will not result in a migration
	DryRun NullableBool `json:"dryRun,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester NullableString `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team NullableString `json:"team,omitempty"`
	// The approver's principal identifier
	Approver NullableString `json:"approver,omitempty"`
	// The request's contact email
	Contact NullableString `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment NullableString `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment NullableString `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount NullableInt64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount NullableInt64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
}

type SubmitMigrateRequest struct {
	CertificateId    string
	CertificatePem   string
	Template         *WebraMigrateTemplate
	TargetModule     string
	TargerProfile    string
	RequesterComment string
}

func (r SubmitMigrateRequest) Execute(c *APIClient) (*WebRAMigrateResponse, *http.Response, error) {
	var req WebRAMigrateRequestOnSubmit
	req.Workflow = string(WORKFLOW_MIGRATE)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	if r.TargetModule != "" {
		req.Module = r.TargetModule
	}
	if r.TargerProfile != "" {
		req.Profile = r.TargerProfile
	}
	if r.Template != nil {
		var template WebRAMigrateRequestTemplate
		template.Labels = r.Template.Labels
		if r.Template.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
		}
		if r.Template.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
		}
		template.Metadata = r.Template.Metadata
		if r.Template.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
		}
		req.Template = template
	}
	resp, httpResp, err := c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRAMigrateRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAMigrateResponse
	responseData, err := resp.WebRAMigrateRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitMigrateWithCertificateId(certificateId, targetModule, targetProfile, requesterComment string, template *WebraMigrateTemplate) SubmitMigrateRequest {
	return SubmitMigrateRequest{
		CertificateId:    certificateId,
		CertificatePem:   "",
		Template:         template,
		TargetModule:     targetModule,
		TargerProfile:    targetProfile,
		RequesterComment: requesterComment,
	}
}

func (r *Requests) SubmitMigrateWithCertificatePem(certificatePem, targetModule, targetProfile, requesterComment string, template *WebraMigrateTemplate) SubmitMigrateRequest {
	return SubmitMigrateRequest{
		CertificateId:    "",
		CertificatePem:   certificatePem,
		Template:         template,
		TargetModule:     targetModule,
		TargerProfile:    targetProfile,
		RequesterComment: requesterComment,
	}
}

type GetMigrateRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetMigrateRequestRequest) Execute() (*WebRAMigrateResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAMigrateResponse
	responseData, err := resp.WebRAMigrateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetMigrateRequest(id string) GetMigrateRequestRequest {
	return GetMigrateRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelMigrateRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelMigrateRequestRequest) Execute() (*WebRAMigrateResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_MIGRATE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAMigrateResponse
	responseData, err := resp.WebRAMigrateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelMigrateRequest(id, module string) CancelMigrateRequestRequest {
	return CancelMigrateRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveMigrateRequestRequest struct {
	Id              string
	ApproverComment string
	c               *APIClient
}

func (r ApproveMigrateRequestRequest) Execute() (*WebRAMigrateResponse, *http.Response, error) {
	var req WebRAMigrateRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_MIGRATE)
	if r.ApproverComment != "" {
		req.ApproverComment = *NewNullableString(&r.ApproverComment)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRAMigrateRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAMigrateResponse
	responseData, err := resp.WebRAMigrateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveMigrateRequest(id, comment string) ApproveMigrateRequestRequest {
	return ApproveMigrateRequestRequest{
		Id:              id,
		ApproverComment: comment,
		c:               r.c,
	}
}

type DenyMigrateRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyMigrateRequestRequest) Execute() (*WebRAMigrateResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_MIGRATE
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAMigrateResponse
	responseData, err := resp.WebRAMigrateRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyMigrateRequest(id, approverComment, module string) DenyMigrateRequestRequest {
	return DenyMigrateRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Renew

type WebraRenewTemplate struct {
	// Describes how certificates will be enrolled on this profile
	Capabilities CertificateProfileCryptoPolicy `json:"capabilities,omitempty"`
	// The password policy that will be used to generate the certificate's PKCS#12 password
	PasswordPolicy PasswordPolicy `json:"passwordPolicy,omitempty"`
	// The CSR used to renew the certificate, if in decentralized mode
	Csr string `json:"csr,omitempty"`
	// The key type of the certificate, if in centralized mode
	KeyType string `json:"keyType,omitempty"`
}

type GetRenewTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Module         string
	Profile        string
	c              *APIClient
}

func (r GetRenewTemplateRequest) Execute() (*WebraRenewTemplate, *http.Response, error) {
	var req WebRARenewRequestOnTemplate
	req.Workflow = string(WORKFLOW_RENEW)
	req.Module = r.Module
	if r.Profile != "" {
		req.Profile = &r.Profile
	}
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRARenewRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraRenewTemplate
	templateData, err := resp.WebRARenewRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetRenewTemplate(profile, module string) GetRenewTemplateRequest {
	return GetRenewTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Module:         module,
		Profile:        profile,
		c:              r.c,
	}
}

func (r *Requests) GetRenewTemplateWithCertificateId(certificateId string) GetRenewTemplateRequest {
	return GetRenewTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Module:         "",
		Profile:        "",
		c:              r.c,
	}
}

func (r *Requests) GetRenewTemplateWithCertificatePem(certificatePem string) GetRenewTemplateRequest {
	return GetRenewTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Module:         "",
		Profile:        "",
		c:              r.c,
	}
}

type WebRARenewResponse struct {
	// The module that will be used to process this request. For a WebRA request, this is always `webra`
	Module string `json:"module"`
	// What this request will do. For a renewal request, this is always `renew`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template *WebRARenewRequestTemplate `json:"template,omitempty"`
	// The generated PKCS#12 for this request. This is only available after the request has been approved in centralized mode
	Pkcs12 SecretString `json:"pkcs12,omitempty"`
	// The password to decrypt the PKCS12 file.
	Password SecretString `json:"password,omitempty"`
	// The certificate that was generated for this request. This is only available after the request has been approved
	Certificate Certificate `json:"certificate,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitRenewRequest struct {
	CertificateId    string
	CertificatePem   string
	Template         *WebraRenewTemplate
	RequesterComment string
	c                *APIClient
}

func (r SubmitRenewRequest) Execute() (*WebRARenewResponse, *http.Response, error) {
	var req WebRARenewRequestOnSubmit
	req.Workflow = string(WORKFLOW_RENEW)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	if r.Template != nil {
		var template WebRARenewRequestTemplate
		template.Csr = *NewNullableString(&r.Template.Csr)
		template.KeyType = *NewNullableString(&r.Template.KeyType)
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRARenewRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARenewResponse
	responseData, err := resp.WebRARenewRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitRenewWithCertificateId(certificateId, requesterComment string, template *WebraRenewTemplate) SubmitRenewRequest {
	return SubmitRenewRequest{
		CertificateId:    certificateId,
		CertificatePem:   "",
		Template:         template,
		RequesterComment: requesterComment,
		c:                r.c,
	}
}

func (r *Requests) SubmitRenewWithCertificatePem(certificatePem, requesterComment string, template *WebraRenewTemplate) SubmitRenewRequest {
	return SubmitRenewRequest{
		CertificateId:    "",
		CertificatePem:   certificatePem,
		Template:         template,
		RequesterComment: requesterComment,
		c:                r.c,
	}
}

type GetRenewRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetRenewRequestRequest) Execute() (*WebRARenewResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARenewResponse
	responseData, err := resp.WebRARenewRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetRenewRequest(id string) GetRenewRequestRequest {
	return GetRenewRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelRenewRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelRenewRequestRequest) Execute() (*WebRARenewResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_RENEW
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARenewResponse
	responseData, err := resp.WebRARenewRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelRenewRequest(id, module string) CancelRenewRequestRequest {
	return CancelRenewRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveRenewRequestRequest struct {
	Id             string
	AppoverComment string
	c              *APIClient
}

func (r ApproveRenewRequestRequest) Execute() (*WebRARenewResponse, *http.Response, error) {
	var req WebRARenewRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_RENEW)
	if r.AppoverComment != "" {
		req.ApproverComment = *NewNullableString(&r.AppoverComment)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRARenewRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARenewResponse
	responseData, err := resp.WebRARenewRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveRenewRequest(id, comment string) ApproveRenewRequestRequest {
	return ApproveRenewRequestRequest{
		Id:             id,
		AppoverComment: comment,
		c:              r.c,
	}
}

type DenyRenewRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyRenewRequestRequest) Execute() (*WebRARenewResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_RENEW
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRARenewResponse
	responseData, err := resp.WebRARenewRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyRenewRequest(id, approverComment, module string) DenyRenewRequestRequest {
	return DenyRenewRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Webra Import

type WebraImportTemplate struct {
	// The owner for this certificate
	Owner CertificateOwnerElement `json:"owner,omitempty"`
	// The team for this certificate
	Team CertificateTeamElement `json:"team,omitempty"`
	// The contact email for this certificate
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
	// The labels for this certificate
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// The technical metadata for this certificate
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
	// The third party data associated with the certificate
	ThirdPartyData []ThirdPartyItem `json:"thirdPartyData,omitempty"`
	// Information about the discovery of this certificate
	DiscoveryInfo DiscoveryInfo `json:"discoveryInfo,omitempty"`
	// The host discovery data associated with the certificate (discovery metadata)
	DiscoveryData *HostDiscoveryData `json:"discoveryData,omitempty"`
	// The PEM-encoded private key associated with the certificate. Mandatory if target profile has escrow enabled, forbidden otherwise
	PrivateKey string `json:"privateKey,omitempty"`
}

type GetImportTemplateRequest struct {
	CertificateId  string
	CertificatePem string
	Profile        string
	Module         string
	c              *APIClient
}

func (r GetImportTemplateRequest) Execute() (*WebraImportTemplate, *http.Response, error) {
	var req WebRAImportRequestOnTemplate
	req.Workflow = string(WORKFLOW_IMPORT)
	req.Module = &r.Module
	if r.Profile != "" {
		req.Profile = r.Profile
	}
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(WebRAImportRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template WebraImportTemplate
	templateData, err := resp.WebRAImportRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetImportTemplate(profile, module string) GetImportTemplateRequest {
	return GetImportTemplateRequest{
		CertificateId:  "",
		CertificatePem: "",
		Profile:        profile,
		Module:         module,
		c:              r.c,
	}
}

func (r *Requests) GetImportTemplateWithCertificateId(certificateId string) GetImportTemplateRequest {
	return GetImportTemplateRequest{
		CertificateId:  certificateId,
		CertificatePem: "",
		Profile:        "",
		Module:         "",
		c:              r.c,
	}
}

func (r *Requests) GetImportTemplateWithCertificatePem(certificatePem string) GetImportTemplateRequest {
	return GetImportTemplateRequest{
		CertificateId:  "",
		CertificatePem: certificatePem,
		Profile:        "",
		Module:         "",
		c:              r.c,
	}
}

type WebRAImportResponse struct {
	// The module of the certificate imported.
	Module Module `json:"module"`
	// What this request will do. For an import request, this is always `import`
	Workflow string `json:"workflow"`
	// The user-data that will be added on certificate import
	Template *WebRAImportRequestTemplate `json:"template,omitempty"`
	// The certificate that was generated for this request.
	Certificate Certificate `json:"certificate"`
	// If true, the request is validated, but will not result in an import
	DryRun bool `json:"dryRun,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
}

type SubmitImportRequest struct {
	CertificateId    string
	CertificatePem   string
	Template         *WebraImportTemplate
	RequesterComment string
	Profile          string
	c                *APIClient
}

func (r SubmitImportRequest) Execute() (*WebRAImportResponse, *http.Response, error) {
	var req WebRAImportRequestOnSubmit
	req.Workflow = string(WORKFLOW_IMPORT)
	if r.CertificatePem != "" {
		req.CertificatePem = *NewNullableString(&r.CertificatePem)
	}
	if r.CertificateId != "" {
		req.CertificateId = *NewNullableString(&r.CertificateId)
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	if r.Profile != "" {
		req.Profile = *NewNullableString(&r.Profile)
	}
	if r.Template != nil {
		var template WebRAImportRequestTemplate
		if r.Template.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Template.Owner)
		}
		if r.Template.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Template.Team)
		}
		if r.Template.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Template.ContactEmail)
		}
		template.Labels = r.Template.Labels
		template.Metadata = r.Template.Metadata
		template.ThirdPartyData = r.Template.ThirdPartyData
		if r.Template.DiscoveryInfo.Identifier.IsSet() {
			template.DiscoveryInfo = *NewNullableDiscoveryInfo(&r.Template.DiscoveryInfo)
		}
		if r.Template.DiscoveryData != nil {
			template.DiscoveryData = r.Template.DiscoveryData
		}
		if r.Template.PrivateKey != "" {
			template.PrivateKey = *NewNullableString(&r.Template.PrivateKey)
		}
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(WebRAImportRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAImportResponse
	responseData, err := resp.WebRAImportRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitImportWithCertificateId(certificateId, requesterComment, profile string, template *WebraImportTemplate) SubmitImportRequest {
	return SubmitImportRequest{
		CertificateId:    certificateId,
		CertificatePem:   "",
		Template:         template,
		RequesterComment: requesterComment,
		Profile:          profile,
		c:                r.c,
	}
}

func (r *Requests) SubmitImportWithCertificatePem(certificatePem, requesterComment, profile string, template *WebraImportTemplate) SubmitImportRequest {
	return SubmitImportRequest{
		CertificateId:    "",
		CertificatePem:   certificatePem,
		Template:         template,
		RequesterComment: requesterComment,
		Profile:          profile,
		c:                r.c,
	}
}

type GetImportRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetImportRequestRequest) Execute() (*WebRAImportResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAImportResponse
	responseData, err := resp.WebRAImportRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetImportRequest(id string) GetImportRequestRequest {
	return GetImportRequestRequest{
		Id: id,
		c:  r.c,
	}
}

type CancelImportRequestRequest struct {
	Id     string
	Module string
	c      *APIClient
}

func (r CancelImportRequestRequest) Execute() (*WebRAImportResponse, *http.Response, error) {
	var req RequestCancelRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_IMPORT
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	resp, httpResp, err := r.c.RequestAPI.RequestCancel(context.Background()).RequestCancelRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAImportResponse
	responseData, err := resp.WebRAImportRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) CancelImportRequest(id, module string) CancelImportRequestRequest {
	return CancelImportRequestRequest{
		Id:     id,
		Module: module,
		c:      r.c,
	}
}

type ApproveImportRequestRequest struct {
	Id              string
	ApproverComment string
	c               *APIClient
}

func (r ApproveImportRequestRequest) Execute() (*WebRAImportResponse, *http.Response, error) {
	var req WebRAImportRequestOnApprove
	req.Id = r.Id
	req.Workflow = string(WORKFLOW_IMPORT)
	if r.ApproverComment != "" {
		req.ApproverComment = *NewNullableString(&r.ApproverComment)
	}
	resp, httpResp, err := r.c.RequestAPI.RequestApprove(context.Background()).RequestApproveRequest(WebRAImportRequestOnApproveAsRequestApproveRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAImportResponse
	responseData, err := resp.WebRAImportRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) ApproveImportRequest(id, comment string) ApproveImportRequestRequest {
	return ApproveImportRequestRequest{
		Id:              id,
		ApproverComment: comment,
		c:               r.c,
	}
}

type DenyImportRequestRequest struct {
	Id              string
	ApproverComment string
	Module          string
	c               *APIClient
}

func (r DenyImportRequestRequest) Execute() (*WebRAImportResponse, *http.Response, error) {
	var req RequestDenyRequest
	req.Id = r.Id
	req.Workflow = WORKFLOW_IMPORT
	t, _ := NewModuleFromValue(r.Module)
	req.Module = *t
	req.ApproverComment = &r.ApproverComment
	resp, httpResp, err := r.c.RequestAPI.RequestDeny(context.Background()).RequestDenyRequest(req).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response WebRAImportResponse
	responseData, err := resp.WebRAImportRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) DenyImportRequest(id, approverComment, module string) DenyImportRequestRequest {
	return DenyImportRequestRequest{
		Id:              id,
		ApproverComment: approverComment,
		Module:          module,
		c:               r.c,
	}
}

// endregion

// region Scep Challenge

type ScepChallengeTemplate struct {
	// DN whitelist is enabled on this request
	DnWhitelist bool `json:"dnWhitelist,omitempty"`
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElement `json:"subject,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElement `json:"sans,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner CertificateOwnerElement `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team CertificateTeamElement `json:"team,omitempty"`
}

type GetScepChallengeTemplateRequest struct {
	Profile string
	c       *APIClient
}

func (r GetScepChallengeTemplateRequest) Execute() (*ScepChallengeTemplate, *http.Response, error) {
	var req ScepEnrollRequestOnTemplate
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_SCEP)
	req.Profile = r.Profile
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(ScepEnrollRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template ScepChallengeTemplate
	templateData, err := resp.ScepEnrollRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetScepChallengeTemplate(profile string) GetScepChallengeTemplateRequest {
	return GetScepChallengeTemplateRequest{
		Profile: profile,
		c:       r.c,
	}
}

type ScepChallengeResponse struct {
	// The module that will be used to process this request. For a SCEP request, this is always `scep`
	Module string `json:"module"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow string `json:"workflow"`
	// The DN of the challenge
	Dn *string `json:"dn,omitempty"`
	// The user-data that will be used to generate the certificate
	Template *ScepEnrollRequestTemplate `json:"template,omitempty"`
	// The password of the challenge. Must be set if password mode is `manual`
	Password SecretString `json:"password,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId *string `json:"holderId,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitScepChallengeRequest struct {
	Dn               string
	Password         string
	RequesterComment string
	Profile          string
	Templates        *ScepChallengeTemplate
	c                *APIClient
}

func (r SubmitScepChallengeRequest) Execute() (*ScepChallengeResponse, *http.Response, error) {
	var req ScepEnrollRequestOnSubmit
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_SCEP)
	req.Dn = r.Dn
	if r.Password != "" {
		req.Password = &SecretString{Value: *NewNullableString(&r.Password)}
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	req.Profile = r.Profile
	if r.Templates != nil {
		var template ScepEnrollRequestTemplate
		template.Subject = r.Templates.Subject
		template.Sans = r.Templates.Sans
		template.Extensions = r.Templates.Extensions
		template.Labels = r.Templates.Labels
		if r.Templates.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Templates.ContactEmail)
		}
		if r.Templates.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Templates.Owner)
		}
		if r.Templates.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Templates.Team)
		}
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(ScepEnrollRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response ScepChallengeResponse
	responseData, err := resp.ScepEnrollRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) SubmitScepChallenge(dn, requesterComment, profile string, template *ScepChallengeTemplate) SubmitScepChallengeRequest {
	return SubmitScepChallengeRequest{
		Dn:               dn,
		Password:         "",
		RequesterComment: requesterComment,
		Profile:          profile,
		Templates:        template,
		c:                r.c,
	}
}

func (r *Requests) SubmitScepChallengeWithPassword(dn, password, requesterComment, profile string, template *ScepChallengeTemplate) SubmitScepChallengeRequest {
	return SubmitScepChallengeRequest{
		Dn:               dn,
		Password:         password,
		RequesterComment: requesterComment,
		Profile:          profile,
		Templates:        template,
		c:                r.c,
	}
}

type GetScepChallengeRequestRequest struct {
	Id string
	c  *APIClient
}

func (r GetScepChallengeRequestRequest) Execute() (*ScepChallengeResponse, *http.Response, error) {
	resp, httpResp, err := r.c.RequestAPI.RequestGet(context.Background(), r.Id).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response ScepChallengeResponse
	responseData, err := resp.ScepEnrollRequestOnApproveResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

func (r *Requests) GetScepChallengeRequest(id string) GetScepChallengeRequestRequest {
	return GetScepChallengeRequestRequest{
		Id: id,
		c:  r.c,
	}
}

// endregion

// region Est Challenge

type EstChallengeTemplate struct {
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElement `json:"subject,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElement `json:"sans,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail CertificateContactEmailElement `json:"contactEmail,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner CertificateOwnerElement `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team CertificateTeamElement `json:"team,omitempty"`
}

type GetEstChallengeTemplateRequest struct {
	Profile string
	c       *APIClient
}

func (r GetEstChallengeTemplateRequest) Execute() (*EstChallengeTemplate, *http.Response, error) {
	var req EstEnrollRequestOnTemplate
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_EST)
	req.Profile = r.Profile
	resp, httpResp, err := r.c.RequestAPI.RequestTemplate(context.Background()).RequestTemplateRequest(EstEnrollRequestOnTemplateAsRequestTemplateRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var template EstChallengeTemplate
	templateData, err := resp.EstEnrollRequestOnTemplateResponse.Template.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(templateData, &template)
	if err != nil {
		return nil, httpResp, err
	}
	return &template, httpResp, nil
}

func (r *Requests) GetEstChallengeTemplate(profile string) GetEstChallengeTemplateRequest {
	return GetEstChallengeTemplateRequest{
		Profile: profile,
		c:       r.c,
	}
}

type EstChallengeResponse struct {
	// The module that will be used to process this request. For an EST request, this is always `est`
	Module string `json:"module"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow string `json:"workflow"`
	// The DN of the challenge
	Dn *string `json:"dn,omitempty"`
	// The user-data that will be used to generate the challenge
	Template *EstEnrollRequestTemplate `json:"template,omitempty"`
	// The password of the challenge.
	Password SecretString `json:"password,omitempty"`
	// Object internal ID
	Id     string        `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// The requester's principal identifier
	Requester string `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team string `json:"team,omitempty"`
	// The approver's principal identifier
	Approver string `json:"approver,omitempty"`
	// The request's contact email
	Contact string `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment string `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment string `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId *string `json:"holderId,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount int64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount int64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun bool `json:"dryRun,omitempty"`
}

type SubmitEstChallengeRequest struct {
	Dn               string
	Password         string
	RequesterComment string
	Profile          string
	Templates        *EstChallengeTemplate
	c                *APIClient
}

func (r SubmitEstChallengeRequest) Execute() (*EstChallengeResponse, *http.Response, error) {
	var req EstEnrollRequestOnSubmit
	req.Workflow = string(WORKFLOW_ENROLL)
	req.Module = string(MODULE_EST)
	req.Dn = r.Dn
	if r.Password != "" {
		req.Password = *NewNullableSecretString(&SecretString{Value: *NewNullableString(&r.Password)})
	}
	if r.RequesterComment != "" {
		req.RequesterComment = *NewNullableString(&r.RequesterComment)
	}
	req.Profile = r.Profile
	if r.Templates != nil {
		var template EstEnrollRequestTemplate
		template.Subject = r.Templates.Subject
		template.Sans = r.Templates.Sans
		template.Extensions = r.Templates.Extensions
		template.Labels = r.Templates.Labels
		if r.Templates.ContactEmail.HasValue() {
			template.ContactEmail = *NewNullableCertificateContactEmailElement(&r.Templates.ContactEmail)
		}
		if r.Templates.Owner.HasValue() {
			template.Owner = *NewNullableCertificateOwnerElement(&r.Templates.Owner)
		}
		if r.Templates.Team.HasValue() {
			template.Team = *NewNullableCertificateTeamElement(&r.Templates.Team)
		}
		req.Template = &template
	}
	resp, httpResp, err := r.c.RequestAPI.RequestSubmit(context.Background()).RequestSubmitRequest(EstEnrollRequestOnSubmitAsRequestSubmitRequest(&req)).Execute()
	if err != nil {
		return nil, httpResp, err
	}
	var response EstChallengeResponse
	responseData, err := resp.EstEnrollRequestOnSubmitResponse.MarshalJSON()
	if err != nil {
		return nil, httpResp, err
	}
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, httpResp, err
	}
	return &response, httpResp, nil
}

// endregion
