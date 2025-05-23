/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package webhook

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815"
	"github.com/Azure/azure-service-operator/v2/internal/util/randextensions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

var _ genruntime.Defaulter = &SqlRoleAssignment{}

func (webhook *SqlRoleAssignment) CustomDefault(ctx context.Context, obj runtime.Object) error {
	resource, ok := obj.(*v1api20240815.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815/SqlRoleAssignment, but got %T", obj)
	}

	webhook.defaultAzureName(ctx, resource)
	return nil
}

// defaultAzureName performs special AzureName defaulting for RoleAssignment by generating a stable GUID
// based on the assignment name.
// We generate the UUID using UUIDv5 with a seed string based on the group, kind, namespace and name.
// We include the namespace and name to ensure no two RoleAssignments in the same cluster can end up
// with the same UUID.
// We include the group and kind to ensure that different kinds of resources get different UUIDs. This isn't
// entirely required by Azure, but it makes sense to avoid collisions between two resources of different types
// even if they have the same namespace and name.
// We include the owner group, kind, and name to avoid collisions between resources with the same name in different
// clusters that actually point to different Azure resources.
// In the rare case users have multiple ASO instances with resources in the same namespace in each cluster
// having the same name but not actually pointing to the same Azure resource (maybe in a different subscription?)
// they can avoid name conflicts by explicitly specifying AzureName for their RoleAssignment.
func (webhook *SqlRoleAssignment) defaultAzureName(_ context.Context, assignment *v1api20240815.SqlRoleAssignment) {
	// If owner is not set we can't default AzureName, but the request will be rejected anyway for lack of owner.
	if assignment.Spec.Owner == nil {
		return
	}

	if assignment.AzureName() == "" {
		gk := assignment.GroupVersionKind().GroupKind()
		assignment.Spec.AzureName = randextensions.MakeUUIDName(
			assignment.Name,
			randextensions.MakeUniqueOwnerScopedStringLegacy(
				assignment.Owner(),
				gk,
				assignment.Namespace,
				assignment.Name))
	}
}
