package main

import (
	"helm.sh/helm/v3/pkg/kube"
	"os"

	"github.com/golang/glog"
	"helm.sh/helm/v3/pkg/action"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type KubeInformation struct {
	AimNamespace string
	AimContext   string
	AimConfig    string
	ApiServer    string
	BearerToken  string
}

func InitKubeInformation(namespace, context, config, server, token string) *KubeInformation {
	return &KubeInformation{
		AimNamespace: namespace,
		AimContext:   context,
		AimConfig:    config,
		ApiServer:    server,
		BearerToken:  token,
	}
}

func actionConfigInit(kubeInfo *KubeInformation) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	clientConfig := new(genericclioptions.ConfigFlags)
	var insecure = true
	clientConfig.Insecure = &insecure
	clientConfig.Namespace = &kubeInfo.AimNamespace
	if kubeInfo.ApiServer != "" && kubeInfo.BearerToken != "" {
		clientConfig.APIServer = &kubeInfo.ApiServer
		clientConfig.BearerToken = &kubeInfo.BearerToken
	} else {
		if kubeInfo.AimContext == "" {
		}
		if kubeInfo.AimConfig == "" {
			clientConfig = kube.GetConfig(settings.KubeConfig, kubeInfo.AimContext, kubeInfo.AimNamespace)
		} else {
			clientConfig = kube.GetConfig(kubeInfo.AimConfig, kubeInfo.AimContext, kubeInfo.AimNamespace)
		}
		if settings.KubeToken != "" {
			clientConfig.BearerToken = &settings.KubeToken
		}
		if settings.KubeAPIServer != "" {
			clientConfig.APIServer = &settings.KubeAPIServer
		}
	}
	err := actionConfig.Init(clientConfig, kubeInfo.AimNamespace, os.Getenv("HELM_DRIVER"), glog.Infof)
	if err != nil {
		glog.Errorf("%+v", err)
		return nil, err
	}

	return actionConfig, nil
}
