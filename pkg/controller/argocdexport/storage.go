// Copyright 2019 ArgoCD Operator Developers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package argocdexport

import (
	"context"

	argoprojv1a1 "github.com/argoproj-labs/argocd-operator/pkg/apis/argoproj/v1alpha1"
)

// reconcileStorage will ensure that the storage options for the ArgoCDExport are present.
func (r *ReconcileArgoCDExport) reconcileStorage(cr *argoprojv1a1.ArgoCDExport) error {
	if cr.Spec.Storage == nil {
		cr.Spec.Storage = &argoprojv1a1.ArgoCDExportStorageSpec{
			Local: &argoprojv1a1.ArgoCDExportLocalStorageSpec{}, // Local is the default
		}
		return r.client.Update(context.TODO(), cr)
	}

	// Local storage
	if err := r.reconcileLocalStorage(cr); err != nil {
		return err
	}

	return nil
}