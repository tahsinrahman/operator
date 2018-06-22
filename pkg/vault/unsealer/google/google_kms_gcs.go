package google

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	api "github.com/kubevault/operator/apis/core/v1alpha1"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ModeGoogleCloudKmsGCS  = "google-cloud-kms-gcs"
	GoogleCredentialFile   = "/etc/vault/unsealer/google/creds/credential.json"
	GoogleCredentialEnv    = "GOOGLE_APPLICATION_CREDENTIALS"
	GoogleCredentialSecret = "vault-unsealer-google-credential"
	GoogleCredentialVolume = "vault-unsealer-google-credential"
)

type Options struct {
	api.GoogleKmsGcsSpec
}

func NewOptions(s api.GoogleKmsGcsSpec) (*Options, error) {
	return &Options{
		s,
	}, nil
}

func (o *Options) Apply(pt *corev1.PodTemplateSpec, cont *corev1.Container) error {
	var args []string

	args = append(args, fmt.Sprintf("--mode=%s", ModeGoogleCloudKmsGCS))

	if o.Bucket != "" {
		args = append(args, fmt.Sprintf("--google.storage-bucket=%s", o.Bucket))
	}
	if o.KmsProject != "" {
		args = append(args, fmt.Sprintf("--google.kms-project=%s", o.KmsProject))
	}
	if o.KmsLocation != "" {
		args = append(args, fmt.Sprintf("--google.kms-location=%s", o.KmsLocation))
	}
	if o.KmsKeyRing != "" {
		args = append(args, fmt.Sprintf("--google.kms-key-ring=%s", o.KmsKeyRing))
	}
	if o.KmsCryptoKey != "" {
		args = append(args, fmt.Sprintf("--google.kms-crypto-key=%s", o.KmsCryptoKey))
	}

	cont.Args = append(cont.Args, args...)

	if o.CredentialPath != "" {
		pt.Spec.Volumes = append(pt.Spec.Volumes, corev1.Volume{
			Name: GoogleCredentialVolume,
			VolumeSource: corev1.VolumeSource{
				Secret: &corev1.SecretVolumeSource{
					SecretName: GoogleCredentialSecret,
				},
			},
		})

		cont.VolumeMounts = append(cont.VolumeMounts, corev1.VolumeMount{
			Name:      GoogleCredentialVolume,
			MountPath: filepath.Dir(GoogleCredentialFile),
		})

		cont.Env = append(cont.Env, corev1.EnvVar{
			Name:  GoogleCredentialEnv,
			Value: GoogleCredentialFile,
		})
	}

	return nil
}

// GetRBAC returns required rbac roles
func (o *Options) GetRBAC(namespace string) []rbac.Role {
	return nil
}

func (o *Options) GetSecrets(namespace string) ([]corev1.Secret, error) {
	var secrets []corev1.Secret
	if o.CredentialPath != "" {
		data, err := ioutil.ReadFile(o.CredentialPath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read google credential file(%s)", o.CredentialPath)
		}

		secrets = append(secrets, corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name: GoogleCredentialSecret,
			},
			Data: map[string][]byte{
				filepath.Base(GoogleCredentialFile): data,
			},
		})
	}
	return secrets, nil
}
