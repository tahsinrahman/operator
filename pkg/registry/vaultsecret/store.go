package vaultsecret

import (
	"context"
	"fmt"

	api "github.com/kubevault/operator/apis/extensions/v1alpha1"
	"github.com/kubevault/operator/client/clientset/versioned"
	"github.com/pkg/errors"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/kubernetes"
	restconfig "k8s.io/client-go/rest"
)

type REST struct {
	stashClient versioned.Interface
	kubeClient  kubernetes.Interface
	config      *restconfig.Config
}

var _ rest.Getter = &REST{}
var _ rest.Lister = &REST{}
var _ rest.GracefulDeleter = &REST{}
var _ rest.Scoper = &REST{}
var _ rest.GroupVersionKindProvider = &REST{}

func NewREST(config *restconfig.Config) *REST {
	return &REST{
		stashClient: versioned.NewForConfigOrDie(config),
		kubeClient:  kubernetes.NewForConfigOrDie(config),
		config:      config,
	}
}

func (r *REST) New() runtime.Object {
	return &api.VaultSecret{}
}

func (r *REST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
	return api.SchemeGroupVersion.WithKind(api.ResourceKindVaultSecret)
}

func (r *REST) NamespaceScoped() bool {
	return true
}

func (r *REST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, errors.New("missing namespace")
	}
	if len(name) < 9 {
		return nil, errors.New("invalid secret name")
	}
	fmt.Println(ns)

	snapshot := &api.VaultSecret{}
	return snapshot, nil
}

func (r *REST) NewList() runtime.Object {
	return &api.VaultSecretList{}
}

func (r *REST) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	ns, ok := apirequest.NamespaceFrom(ctx)
	if !ok {
		return nil, errors.New("missing namespace")
	}
	fmt.Println(ns)

	objects := &api.VaultSecretList{}
	return objects, nil
}

func (r *REST) Delete(ctx context.Context, name string, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	return nil, false, nil
}
