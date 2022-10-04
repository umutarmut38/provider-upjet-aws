/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/cognitoidp/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CognitoIdentityPoolProviderPrincipalTag.
func (mg *CognitoIdentityPoolProviderPrincipalTag) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IdentityPoolIDRef,
		Selector:     mg.Spec.ForProvider.IdentityPoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityPoolID")
	}
	mg.Spec.ForProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityPoolIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityProviderName),
		Extract:      resource.ExtractParamPath("endpoint", true),
		Reference:    mg.Spec.ForProvider.IdentityProviderNameRef,
		Selector:     mg.Spec.ForProvider.IdentityProviderNameSelector,
		To: reference.To{
			List:    &v1beta1.UserPoolList{},
			Managed: &v1beta1.UserPool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityProviderName")
	}
	mg.Spec.ForProvider.IdentityProviderName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityProviderNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Pool.
func (mg *Pool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.CognitoIdentityProviders); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDRef,
			Selector:     mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDSelector,
			To: reference.To{
				List:    &v1beta1.UserPoolClientList{},
				Managed: &v1beta1.UserPoolClient{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID")
		}
		mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.CognitoIdentityProviders[i3].ClientIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SAMLProviderArns),
		Extract:       common.ARNExtractor(),
		References:    mg.Spec.ForProvider.SAMLProviderArnsRefs,
		Selector:      mg.Spec.ForProvider.SAMLProviderArnsSelector,
		To: reference.To{
			List:    &v1beta11.SAMLProviderList{},
			Managed: &v1beta11.SAMLProvider{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SAMLProviderArns")
	}
	mg.Spec.ForProvider.SAMLProviderArns = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SAMLProviderArnsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this PoolRolesAttachment.
func (mg *PoolRolesAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IdentityPoolID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.IdentityPoolIDRef,
		Selector:     mg.Spec.ForProvider.IdentityPoolIDSelector,
		To: reference.To{
			List:    &PoolList{},
			Managed: &Pool{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IdentityPoolID")
	}
	mg.Spec.ForProvider.IdentityPoolID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IdentityPoolIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.RoleMapping); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.RoleMapping[i3].MappingRule); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef,
				Selector:     mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnSelector,
				To: reference.To{
					List:    &v1beta11.RoleList{},
					Managed: &v1beta11.Role{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn")
			}
			mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.RoleMapping[i3].MappingRule[i4].RoleArnRef = rsp.ResolvedReference

		}
	}

	return nil
}
