package google

import (
	"testing"

	api "github.com/kubevault/operator/apis/core/v1alpha1"
	"github.com/kubevault/operator/pkg/vault/util"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestOptions_Apply(t *testing.T) {
	expected := []string{
		"--mode=google-cloud-kms-gcs",
		"--google.storage-bucket=test",
		"--google.kms-project=pro",
		"--google.kms-location=loc",
		"--google.kms-key-ring=r-key",
		"--google.kms-crypto-key=c-key",
	}
	cont := corev1.Container{
		Name: util.VaultUnsealerImageName(),
	}
	pt := &corev1.PodTemplateSpec{
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{cont},
		},
	}

	opts, err := NewOptions(api.GoogleKmsGcsSpec{
		KmsCryptoKey: "c-key",
		KmsKeyRing:   "r-key",
		KmsLocation:  "loc",
		KmsProject:   "pro",
		Bucket:       "test",
	})
	assert.Nil(t, err)

	err = opts.Apply(pt)
	assert.Nil(t, err)

	assert.Equal(t, expected, pt.Spec.Containers[0].Args)
}
