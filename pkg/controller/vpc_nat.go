package controller

import (
	"fmt"

	"k8s.io/klog/v2"

	"github.com/kubeovn/kube-ovn/pkg/util"
)

var (
	vpcNatImage             = ""
	vpcNatGwBgpSpeakerImage = ""
	vpcNatAPINadProvider    = ""
)

func (c *Controller) resyncVpcNatConfig() {
	cm, err := c.configMapsLister.ConfigMaps(c.config.PodNamespace).Get(util.VpcNatConfig)
	if err != nil {
		err = fmt.Errorf("failed to get ovn-vpc-nat-config, %w", err)
		klog.Error(err)
		return
	}

	// Image we're using to provision the NAT gateways
	image, exist := cm.Data["image"]
	if !exist {
		err = fmt.Errorf("%s should have image field", util.VpcNatConfig)
		klog.Error(err)
		return
	}
	vpcNatImage = image

	// Image for the BGP sidecar of the gateway (optional)
	vpcNatGwBgpSpeakerImage = cm.Data["bgpSpeakerImage"]

	// NetworkAttachmentDefinition provider for the BGP speaker to call the API server
	vpcNatAPINadProvider = cm.Data["apiNadProvider"]
}
