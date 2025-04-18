package infra

import (
	"github.com/pulumi/pulumi-awsx/sdk/v2/go/awsx/ec2"
	"github.com/pulumi/pulumi-eks/sdk/v3/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type EKSCluster struct {
}

func NewEKSCluster() *EKSCluster {
	return &EKSCluster{}
}

func (h *EKSCluster) Run(ctx *pulumi.Context) (*eks.Cluster, error) {
	cfg := config.New(ctx, "")
	minClusterSize, err := cfg.TryInt("minClusterSize")
	if err != nil {
		minClusterSize = 3
	}
	maxClusterSize, err := cfg.TryInt("maxClusterSize")
	if err != nil {
		maxClusterSize = 6
	}
	desiredClusterSize, err := cfg.TryInt("desiredClusterSize")
	if err != nil {
		desiredClusterSize = 3
	}
	eksNodeInstanceType, err := cfg.Try("eksNodeInstanceType")
	if err != nil {
		eksNodeInstanceType = "t3.medium"
	}
	vpcNetworkCidr, err := cfg.Try("vpcNetworkCidr")
	if err != nil {
		vpcNetworkCidr = "10.0.0.0/16"
	}

	// Create a new VPC, subnets, and associated infrastructure
	eksVpc, err := ec2.NewVpc(ctx, "eks-vpc", &ec2.VpcArgs{
		EnableDnsHostnames: pulumi.Bool(true),
		CidrBlock:          &vpcNetworkCidr,
	})
	if err != nil {
		return nil, err
	}

	apiAuthMode := eks.AuthenticationModeApi
	// Create a new EKS cluster
	eksCluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
		// Put the cluster in the new VPC created earlier
		VpcId: eksVpc.VpcId,
		// Use the "API" authentication mode to support access entries
		AuthenticationMode: &apiAuthMode,
		// Public subnets will be used for load balancers
		PublicSubnetIds: eksVpc.PublicSubnetIds,
		// Private subnets will be used for cluster nodes
		PrivateSubnetIds:   eksVpc.PrivateSubnetIds,
		CreateOidcProvider: pulumi.Bool(true),
		// Change configuration values above to change any of the following settings
		InstanceType:    pulumi.String(eksNodeInstanceType),
		DesiredCapacity: pulumi.Int(desiredClusterSize),
		MinSize:         pulumi.Int(minClusterSize),
		MaxSize:         pulumi.Int(maxClusterSize),
		// Do not give the worker nodes a public IP address
		NodeAssociatePublicIpAddress: pulumi.BoolRef(false),
		// Change these values for a private cluster (VPN access required)
		EndpointPrivateAccess: pulumi.Bool(false),
		EndpointPublicAccess:  pulumi.Bool(true),
	})
	if err != nil {
		return nil, err
	}

	// Export some values in case they are needed elsewhere
	ctx.Export("kubeconfig", eksCluster.Kubeconfig)
	ctx.Export("vpcId", eksVpc.VpcId)
	return eksCluster, nil

}
