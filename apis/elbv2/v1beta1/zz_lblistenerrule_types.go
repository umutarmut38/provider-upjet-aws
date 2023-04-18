/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionAuthenticateCognitoObservation struct {

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// The behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// The set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// The ARN of the Cognito user pool.
	UserPoolArn *string `json:"userPoolArn,omitempty" tf:"user_pool_arn,omitempty"`

	// The ID of the Cognito user pool client.
	UserPoolClientID *string `json:"userPoolClientId,omitempty" tf:"user_pool_client_id,omitempty"`

	// The domain prefix or fully-qualified domain name of the Cognito user pool.
	UserPoolDomain *string `json:"userPoolDomain,omitempty" tf:"user_pool_domain,omitempty"`
}

type ActionAuthenticateCognitoParameters struct {

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// The behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// The set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// The ARN of the Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	UserPoolArn *string `json:"userPoolArn,omitempty" tf:"user_pool_arn,omitempty"`

	// Reference to a UserPool in cognitoidp to populate userPoolArn.
	// +kubebuilder:validation:Optional
	UserPoolArnRef *v1.Reference `json:"userPoolArnRef,omitempty" tf:"-"`

	// Selector for a UserPool in cognitoidp to populate userPoolArn.
	// +kubebuilder:validation:Optional
	UserPoolArnSelector *v1.Selector `json:"userPoolArnSelector,omitempty" tf:"-"`

	// The ID of the Cognito user pool client.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolClient
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	UserPoolClientID *string `json:"userPoolClientId,omitempty" tf:"user_pool_client_id,omitempty"`

	// Reference to a UserPoolClient in cognitoidp to populate userPoolClientId.
	// +kubebuilder:validation:Optional
	UserPoolClientIDRef *v1.Reference `json:"userPoolClientIdRef,omitempty" tf:"-"`

	// Selector for a UserPoolClient in cognitoidp to populate userPoolClientId.
	// +kubebuilder:validation:Optional
	UserPoolClientIDSelector *v1.Selector `json:"userPoolClientIdSelector,omitempty" tf:"-"`

	// The domain prefix or fully-qualified domain name of the Cognito user pool.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidp/v1beta1.UserPoolDomain
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("domain",false)
	// +kubebuilder:validation:Optional
	UserPoolDomain *string `json:"userPoolDomain,omitempty" tf:"user_pool_domain,omitempty"`

	// Reference to a UserPoolDomain in cognitoidp to populate userPoolDomain.
	// +kubebuilder:validation:Optional
	UserPoolDomainRef *v1.Reference `json:"userPoolDomainRef,omitempty" tf:"-"`

	// Selector for a UserPoolDomain in cognitoidp to populate userPoolDomain.
	// +kubebuilder:validation:Optional
	UserPoolDomainSelector *v1.Selector `json:"userPoolDomainSelector,omitempty" tf:"-"`
}

type ActionAuthenticateOidcObservation struct {

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// The authorization endpoint of the IdP.
	AuthorizationEndpoint *string `json:"authorizationEndpoint,omitempty" tf:"authorization_endpoint,omitempty"`

	// The OAuth 2.0 client identifier.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The OIDC issuer identifier of the IdP.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// The set of user claims to be requested from the IdP.
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The name of the cookie used to maintain session information.
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// The token endpoint of the IdP.
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`

	// The user info endpoint of the IdP.
	UserInfoEndpoint *string `json:"userInfoEndpoint,omitempty" tf:"user_info_endpoint,omitempty"`
}

type ActionAuthenticateOidcParameters struct {

	// The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// The authorization endpoint of the IdP.
	// +kubebuilder:validation:Required
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// The OAuth 2.0 client identifier.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The OAuth 2.0 client secret.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The OIDC issuer identifier of the IdP.
	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// The behavior if the user is not authenticated. Valid values: deny, allow and authenticate
	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// The set of user claims to be requested from the IdP.
	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// The name of the cookie used to maintain session information.
	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// The maximum duration of the authentication session, in seconds.
	// +kubebuilder:validation:Optional
	SessionTimeout *float64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// The token endpoint of the IdP.
	// +kubebuilder:validation:Required
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// The user info endpoint of the IdP.
	// +kubebuilder:validation:Required
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint,omitempty"`
}

type ActionFixedResponseObservation struct {

	// The content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// The message body.
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// The HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ActionFixedResponseParameters struct {

	// The content type. Valid values are text/plain, text/css, text/html, application/javascript and application/json.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// The message body.
	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// The HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ActionForwardObservation struct {

	// The target group stickiness for the rule.
	Stickiness []ForwardStickinessObservation `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// One or more target groups block.
	TargetGroup []ForwardTargetGroupObservation `json:"targetGroup,omitempty" tf:"target_group,omitempty"`
}

type ActionForwardParameters struct {

	// The target group stickiness for the rule.
	// +kubebuilder:validation:Optional
	Stickiness []ForwardStickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// One or more target groups block.
	// +kubebuilder:validation:Required
	TargetGroup []ForwardTargetGroupParameters `json:"targetGroup" tf:"target_group,omitempty"`
}

type ActionObservation struct {

	// Information for creating an authenticate action using Cognito. Required if type is authenticate-cognito.
	AuthenticateCognito []ActionAuthenticateCognitoObservation `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Information for creating an authenticate action using OIDC. Required if type is authenticate-oidc.
	AuthenticateOidc []ActionAuthenticateOidcObservation `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	FixedResponse []ActionFixedResponseObservation `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Information for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn.
	Forward []ActionForwardObservation `json:"forward,omitempty" tf:"forward,omitempty"`

	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Information for creating a redirect action. Required if type is redirect.
	Redirect []ActionRedirectObservation `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// The ARN of the Target Group to which to route traffic. Specify only if type is forward and you want to route to a single target group. To route to one or more target groups, use a forward block instead.
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// The type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionParameters struct {

	// Information for creating an authenticate action using Cognito. Required if type is authenticate-cognito.
	// +kubebuilder:validation:Optional
	AuthenticateCognito []ActionAuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// Information for creating an authenticate action using OIDC. Required if type is authenticate-oidc.
	// +kubebuilder:validation:Optional
	AuthenticateOidc []ActionAuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if type is fixed-response.
	// +kubebuilder:validation:Optional
	FixedResponse []ActionFixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// Information for creating an action that distributes requests among one or more target groups. Specify only if type is forward. If you specify both forward block and target_group_arn attribute, you can specify only one target group using forward and it must be the same target group specified in target_group_arn.
	// +kubebuilder:validation:Optional
	Forward []ActionForwardParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// +kubebuilder:validation:Optional
	Order *float64 `json:"order,omitempty" tf:"order,omitempty"`

	// Information for creating a redirect action. Required if type is redirect.
	// +kubebuilder:validation:Optional
	Redirect []ActionRedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// The ARN of the Target Group to which to route traffic. Specify only if type is forward and you want to route to a single target group. To route to one or more target groups, use a forward block instead.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnRef *v1.Reference `json:"targetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate targetGroupArn.
	// +kubebuilder:validation:Optional
	TargetGroupArnSelector *v1.Selector `json:"targetGroupArnSelector,omitempty" tf:"-"`

	// The type of routing action. Valid values are forward, redirect, fixed-response, authenticate-cognito and authenticate-oidc.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ActionRedirectObservation struct {

	// The hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ActionRedirectParameters struct {

	// The hostname. This component is not percent-encoded. The hostname can contain #{host}. Defaults to #{host}.
	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to /#{path}.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The port. Specify a value from 1 to 65535 or #{port}. Defaults to #{port}.
	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// The protocol. Valid values are HTTP, HTTPS, or #{protocol}. Defaults to #{protocol}.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to #{query}.
	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The HTTP redirect code. The redirect is either permanent (HTTP_301) or temporary (HTTP_302).
	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type ConditionObservation struct {

	// HTTP headers to match. HTTP Header block fields documented below.
	HTTPHeader []HTTPHeaderObservation `json:"httpHeader,omitempty" tf:"http_header,omitempty"`

	// Contains a single values item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
	HTTPRequestMethod []HTTPRequestMethodObservation `json:"httpRequestMethod,omitempty" tf:"http_request_method,omitempty"`

	// Contains a single values item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
	HostHeader []HostHeaderObservation `json:"hostHeader,omitempty" tf:"host_header,omitempty"`

	// Contains a single values item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a query_string condition.
	PathPattern []PathPatternObservation `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// Query strings to match. Query String block fields documented below.
	QueryString []QueryStringObservation `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// Contains a single values item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the X-Forwarded-For header, use http_header condition instead.
	SourceIP []SourceIPObservation `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`
}

type ConditionParameters struct {

	// HTTP headers to match. HTTP Header block fields documented below.
	// +kubebuilder:validation:Optional
	HTTPHeader []HTTPHeaderParameters `json:"httpHeader,omitempty" tf:"http_header,omitempty"`

	// Contains a single values item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
	// +kubebuilder:validation:Optional
	HTTPRequestMethod []HTTPRequestMethodParameters `json:"httpRequestMethod,omitempty" tf:"http_request_method,omitempty"`

	// Contains a single values item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
	// +kubebuilder:validation:Optional
	HostHeader []HostHeaderParameters `json:"hostHeader,omitempty" tf:"host_header,omitempty"`

	// Contains a single values item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a query_string condition.
	// +kubebuilder:validation:Optional
	PathPattern []PathPatternParameters `json:"pathPattern,omitempty" tf:"path_pattern,omitempty"`

	// Query strings to match. Query String block fields documented below.
	// +kubebuilder:validation:Optional
	QueryString []QueryStringParameters `json:"queryString,omitempty" tf:"query_string,omitempty"`

	// Contains a single values item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the X-Forwarded-For header, use http_header condition instead.
	// +kubebuilder:validation:Optional
	SourceIP []SourceIPParameters `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`
}

type ForwardStickinessObservation struct {

	// The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Indicates whether target group stickiness is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ForwardStickinessParameters struct {

	// The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	// +kubebuilder:validation:Required
	Duration *float64 `json:"duration" tf:"duration,omitempty"`

	// Indicates whether target group stickiness is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type ForwardTargetGroupObservation struct {

	// The Amazon Resource Name (ARN) of the target group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The weight. The range is 0 to 999.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type ForwardTargetGroupParameters struct {

	// The Amazon Resource Name (ARN) of the target group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate arn.
	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate arn.
	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// The weight. The range is 0 to 999.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type HTTPHeaderObservation struct {

	// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a host-header condition instead.
	HTTPHeaderName *string `json:"httpHeaderName,omitempty" tf:"http_header_name,omitempty"`

	// List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type HTTPHeaderParameters struct {

	// Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a host-header condition instead.
	// +kubebuilder:validation:Required
	HTTPHeaderName *string `json:"httpHeaderName" tf:"http_header_name,omitempty"`

	// List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type HTTPRequestMethodObservation struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type HTTPRequestMethodParameters struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type HostHeaderObservation struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type HostHeaderParameters struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type LBListenerRuleObservation struct {

	// An Action block. Action blocks are documented below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// The ARN of the rule (matches id)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// The ARN of the rule (matches arn)
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the listener to which to attach the rule.
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`

	// The priority for the rule between 1 and 50000. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBListenerRuleParameters struct {

	// An Action block. Action blocks are documented below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// The ARN of the listener to which to attach the rule.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBListener
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`

	// Reference to a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnRef *v1.Reference `json:"listenerArnRef,omitempty" tf:"-"`

	// Selector for a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnSelector *v1.Selector `json:"listenerArnSelector,omitempty" tf:"-"`

	// The priority for the rule between 1 and 50000. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PathPatternObservation struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type PathPatternParameters struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type QueryStringObservation struct {

	// Query string key pattern to match.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Query string value pattern to match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type QueryStringParameters struct {

	// Query string key pattern to match.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Query string value pattern to match.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SourceIPObservation struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type SourceIPParameters struct {

	// Query string pairs or values to match. Query String Value blocks documented below. Multiple values blocks can be specified, see example above. Maximum size of each string is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). To search for a literal '*' or '?' character in a query string, escape the character with a backslash (\). Only one pair needs to match for the condition to be satisfied.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

// LBListenerRuleSpec defines the desired state of LBListenerRule
type LBListenerRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBListenerRuleParameters `json:"forProvider"`
}

// LBListenerRuleStatus defines the observed state of LBListenerRule.
type LBListenerRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBListenerRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerRule is the Schema for the LBListenerRules API. Provides a Load Balancer Listener Rule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.action)",message="action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.condition)",message="condition is a required parameter"
	Spec   LBListenerRuleSpec   `json:"spec"`
	Status LBListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerRuleList contains a list of LBListenerRules
type LBListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBListenerRule `json:"items"`
}

// Repository type metadata.
var (
	LBListenerRule_Kind             = "LBListenerRule"
	LBListenerRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBListenerRule_Kind}.String()
	LBListenerRule_KindAPIVersion   = LBListenerRule_Kind + "." + CRDGroupVersion.String()
	LBListenerRule_GroupVersionKind = CRDGroupVersion.WithKind(LBListenerRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LBListenerRule{}, &LBListenerRuleList{})
}
