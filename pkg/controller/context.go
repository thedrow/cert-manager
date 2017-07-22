package controller

import (
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"github.com/jetstack/cert-manager/pkg/client"
	"github.com/jetstack/cert-manager/pkg/informers/externalversions"
	"github.com/jetstack/cert-manager/pkg/log"
)

type Context struct {
	Client            *kubernetes.Clientset
	CertManagerClient *client.Clientset

	InformerFactory            informers.SharedInformerFactory
	CertManagerInformerFactory externalversions.SharedInformerFactory

	Namespace string
	Logger    log.Logger
}