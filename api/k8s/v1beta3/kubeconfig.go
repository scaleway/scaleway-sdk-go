package k8s

import (
	"io/ioutil"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"gopkg.in/yaml.v2"
)

// Kubeconfig represents a kubernetes kubeconfig file
type Kubeconfig struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Clusters   []struct {
		Name    string `yaml:"name"`
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
	} `yaml:"clusters"`
	Contexts []struct {
		Name    string `yaml:"name"`
		Context struct {
			Cluster   string `yaml:"cluster"`
			Namespace string `yaml:"namespace,omitempty"`
			User      string `yaml:"user"`
		} `yaml:"context"`
	} `yaml:"contexts"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			Token string `yaml:"token"`
		} `yaml:"user"`
	} `yaml:"users"`
}

// GetRaw returns the raw bytes of the kubeconfig
func (k *Kubeconfig) GetRaw() ([]byte, error) {
	return yaml.Marshal(k)
}

// GetServer returns the server URL of the cluster in the kubeconfig
func (k *Kubeconfig) GetServer() (string, error) {
	if len(k.Clusters) != 1 {
		return "", errors.New("wrong number of clusters in kubeconfig")
	}

	return k.Clusters[0].Cluster.Server, nil
}

// GetCertificateAuthorityData returns the server certificate authority data of the cluster in the kubeconfig
func (k *Kubeconfig) GetCertificateAuthorityData() (string, error) {
	if len(k.Clusters) != 1 {
		return "", errors.New("wrong number of clusters in kubeconfig")
	}

	return k.Clusters[0].Cluster.CertificateAuthorityData, nil
}

// GetToken returns the token for the cluster in the kubeconfig
func (k *Kubeconfig) GetToken() (string, error) {
	if len(k.Users) != 1 {
		return "", errors.New("wrong number of users in kubeconfig")
	}

	return k.Users[0].User.Token, nil
}

// GetClusterKubeConfigRequest is the requst for GetClusterKubeConfig
type GetClusterKubeConfigRequest struct {
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// GetClusterKubeConfig downloads the kubeconfig for the given cluster
func (s *API) GetClusterKubeConfig(req *GetClusterKubeConfigRequest, opts ...scw.RequestOption) (*Kubeconfig, error) {
	kubeconfigFile, err := s.getClusterKubeConfig(&getClusterKubeConfigRequest{
		Region:    req.Region,
		ClusterID: req.ClusterID,
	})
	if err != nil {
		return nil, err
	}

	kubeconfigContent, err := ioutil.ReadAll(kubeconfigFile.Content)
	if err != nil {
		return nil, err
	}

	var kubeconfig Kubeconfig
	err = yaml.Unmarshal(kubeconfigContent, &kubeconfig)
	if err != nil {
		return nil, err
	}

	return &kubeconfig, nil
}
