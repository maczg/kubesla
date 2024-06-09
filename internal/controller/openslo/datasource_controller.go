/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package openslo

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	openslov1alpha1 "github.com/maczg/kubesla/api/openslo/v1alpha1"
)

// DatasourceReconciler reconciles a Datasource object
type DatasourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

const urlPattern = `^https?://.*`

//+kubebuilder:rbac:groups=openslo.kubesla.com,resources=datasources,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=openslo.kubesla.com,resources=datasources/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=openslo.kubesla.com,resources=datasources/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *DatasourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	dataSource := &openslov1alpha1.Datasource{}
	if err := r.Get(ctx, req.NamespacedName, dataSource); err != nil {
		if errors.IsNotFound(err) {
			// Resource not found
			return ctrl.Result{}, nil
		}
		// Error reading the object
		return ctrl.Result{}, err
	}

	logger.Info("Reconciling Datasource", "Name", dataSource.Name)

	err := r.handleDataSource(ctx, dataSource)
	if err != nil {
		logger.Error(err, "Failed to handle datasource")
		return ctrl.Result{RequeueAfter: 1 * time.Minute}, err
	}

	return ctrl.Result{RequeueAfter: 1 * time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatasourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&openslov1alpha1.Datasource{}).
		Complete(r)
}

func (r *DatasourceReconciler) handleDataSource(ctx context.Context, dataSource *openslov1alpha1.Datasource) error {
	switch dataSource.Spec.Type {
	case "Prometheus":
		url, urlOk := dataSource.Spec.ConnectionDetails["url"]
		if !urlOk {
			return fmt.Errorf("missing required connection detail: url")
		}

		isHealthy, statusMessage := probePrometheus(url)

		dataSource.Status.LastChecked = metav1.Now()
		dataSource.Status.IsHealthy = isHealthy
		dataSource.Status.StatusMessage = statusMessage

		if err := r.Status().Update(ctx, dataSource); err != nil {
			return err
		}
		return nil

	case "Datadog":
		return fmt.Errorf("unsupported data source type: %s", dataSource.Spec.Type)
	case "CloudWatch":
		return fmt.Errorf("unsupported data source type: %s", dataSource.Spec.Type)
	default:
		return fmt.Errorf("unsupported data source type: %s", dataSource.Spec.Type)
	}
}

func probePrometheus(url string) (bool, string) {
	resp, err := http.Get(url + "/-/ready")
	if err != nil {
		return false, fmt.Sprintf("Failed to reach Prometheus: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, "Prometheus is healthy"
	}

	return false, fmt.Sprintf("Prometheus is unhealthy: %s", resp.Status)
}
