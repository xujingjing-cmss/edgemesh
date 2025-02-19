package config

import (
	"k8s.io/klog/v2"

	"github.com/kubeedge/edgemesh/common/acl"
	meshConstants "github.com/kubeedge/edgemesh/common/constants"
	"github.com/kubeedge/edgemesh/common/util"
)

const (
	defaultListenPort = 20004
)

// TunnelServerConfig indicates networking module config
type TunnelServerConfig struct {
	// Enable indicates whether Tunnel is enabled,
	// if set to false (for debugging etc.), skip checking other Networking configs.
	// default true
	Enable bool `json:"enable,omitempty"`
	// TunnelACLConfig indicates the set of tunnel server config about acl
	TunnelACLConfig acl.TunnelACLConfig `json:"ACL,omitempty"`
	// NodeName indicates the node name of tunnel server
	NodeName string `json:"nodeName,omitempty"`
	// ListenPort indicates the listen port of tunnel server
	// default 20004
	ListenPort int `json:"listenPort,omitempty"`
	// EnableSecurity indicates whether to use the ca acl and security transport
	// default false
	EnableSecurity bool `json:"enableSecurity"`
	// AdvertiseAddress sets the IP address for the edgemesh-server to advertise
	AdvertiseAddress []string `json:"advertiseAddress,omitempty"`
}

func NewTunnelServerConfig() *TunnelServerConfig {
	// fetch the public IP auto and append it to the advertiseAddress
	publicIP := util.FetchPublicIP()
	advertiseAddress := make([]string, 0)
	if publicIP != "" {
		klog.Infof("Fetch public IP: %s", publicIP)
		advertiseAddress = append(advertiseAddress, publicIP)
	} else {
		klog.Infof("Unable to fetch public IP")
	}

	return &TunnelServerConfig{
		Enable: true,
		TunnelACLConfig: acl.TunnelACLConfig{
			TLSPrivateKeyFile: meshConstants.ServerDefaultKeyFile,
			TLSCAFile:         meshConstants.ServerDefaultCAFile,
			TLSCertFile:       meshConstants.ServerDefaultCertFile,
		},
		ListenPort:       defaultListenPort,
		EnableSecurity:   false,
		AdvertiseAddress: advertiseAddress,
	}
}
