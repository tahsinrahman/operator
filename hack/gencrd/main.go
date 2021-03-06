package main

import (
	"io/ioutil"
	"os"

	"github.com/appscode/go/log"
	gort "github.com/appscode/go/runtime"
	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	"github.com/appscode/kutil/openapi"
	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	vaultinstall "github.com/kubevault/operator/apis/core/install"
	corev1alpha1 "github.com/kubevault/operator/apis/core/v1alpha1"
	extinstall "github.com/kubevault/operator/apis/extensions/install"
	extv1alpha1 "github.com/kubevault/operator/apis/extensions/v1alpha1"
	crd_api "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kube-openapi/pkg/common"
	"path/filepath"
)

func generateCRDDefinitions() {
	corev1alpha1.EnableStatusSubresource = true

	filename := gort.GOPath() + "/src/github.com/kubevault/operator/apis/core/v1alpha1/crds.yaml"
	os.Remove(filename)

	err := os.MkdirAll(filepath.Join(gort.GOPath(), "/src/github.com/kubevault/operator/api/crds"), 0755)
	if err != nil {
		log.Fatal(err)
	}

	crds := []*crd_api.CustomResourceDefinition{
		corev1alpha1.VaultServer{}.CustomResourceDefinition(),
	}
	for _, crd := range crds {
		filename := filepath.Join(gort.GOPath(), "/src/github.com/kubevault/operator/api/crds", crd.Spec.Names.Singular+".yaml")
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		crdutils.MarshallCrd(f, crd, "yaml")
		f.Close()
	}
}
func generateSwaggerJson() {
	var (
		Scheme = runtime.NewScheme()
		Codecs = serializer.NewCodecFactory(Scheme)
	)

	vaultinstall.Install(Scheme)
	extinstall.Install(Scheme)

	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
		Scheme: Scheme,
		Codecs: Codecs,
		Info: spec.InfoProps{
			Title:   "KubeVault",
			Version: "v0.1.0",
			Contact: &spec.ContactInfo{
				Name:  "AppsCode Inc.",
				URL:   "https://appscode.com",
				Email: "kubevault@appscode.com",
			},
			License: &spec.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			corev1alpha1.GetOpenAPIDefinitions,
			extv1alpha1.GetOpenAPIDefinitions,
		},
		Resources: []openapi.TypeInfo{
			{corev1alpha1.SchemeGroupVersion, corev1alpha1.ResourceVaultServers, corev1alpha1.ResourceKindVaultServer, true},
		},
		RDResources: []openapi.TypeInfo{
			{extv1alpha1.SchemeGroupVersion, extv1alpha1.ResourceVaultSecrets, extv1alpha1.ResourceKindVaultSecret, true},
		},
	})
	if err != nil {
		glog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/github.com/kubevault/operator/api/openapi-spec/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		glog.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(apispec), 0644)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	generateCRDDefinitions()
	generateSwaggerJson()
}
