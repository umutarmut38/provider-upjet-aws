/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/glue/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Permissions.
func (mg *Permissions) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DataLocation); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataLocation[i3].Arn),
			Extract:      resource.ExtractParamPath("arn", false),
			Reference:    mg.Spec.ForProvider.DataLocation[i3].ArnRef,
			Selector:     mg.Spec.ForProvider.DataLocation[i3].ArnSelector,
			To: reference.To{
				List:    &ResourceList{},
				Managed: &Resource{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DataLocation[i3].Arn")
		}
		mg.Spec.ForProvider.DataLocation[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DataLocation[i3].ArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Database); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Database[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Database[i3].NameRef,
			Selector:     mg.Spec.ForProvider.Database[i3].NameSelector,
			To: reference.To{
				List:    &v1beta1.CatalogDatabaseList{},
				Managed: &v1beta1.CatalogDatabase{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Database[i3].Name")
		}
		mg.Spec.ForProvider.Database[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Database[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.TableWithColumns); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TableWithColumns[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.TableWithColumns[i3].NameRef,
			Selector:     mg.Spec.ForProvider.TableWithColumns[i3].NameSelector,
			To: reference.To{
				List:    &v1beta1.CatalogTableList{},
				Managed: &v1beta1.CatalogTable{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.TableWithColumns[i3].Name")
		}
		mg.Spec.ForProvider.TableWithColumns[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.TableWithColumns[i3].NameRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Resource.
func (mg *Resource) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.RoleArnRef,
		Selector:     mg.Spec.ForProvider.RoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleArn")
	}
	mg.Spec.ForProvider.RoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleArnRef = rsp.ResolvedReference

	return nil
}
