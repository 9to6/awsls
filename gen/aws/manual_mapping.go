// +build codegen

package aws

// AWS services that are excluded from generating code for.
var excludeServices = map[string]struct{}{
	"importexport": {},
}

// some resource types are excluded as they need be handled slightly differently
var ExcludedResourceTypes = map[string]bool{
	// not a resource
	"aws_api_gateway_integration": true,
	// is not a resource
	"aws_acm_certificate_validation": true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_alb_listener": true,
	//ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_alb_listener_rule": true,
	// A normal subnet with "DefaultForAz" attribute set to true
	"aws_default_subnet": true,
	// BadRequest: Must specify exactly one mutually exclusive parameter.
	"aws_acmpca_certificate_authority":  true,
	"aws_ami_from_instance":             true,
	"aws_appmesh_route":                 true,
	"aws_appmesh_virtual_node":          true,
	"aws_appmesh_virtual_router":        true,
	"aws_appmesh_virtual_service":       true,
	"aws_datapipeline_pipeline":         true,
	"aws_efs_mount_target":              true,
	"aws_emr_cluster":                   true,
	"aws_glue_classifier":               true,
	"aws_iam_openid_connect_provider":   true,
	"aws_iam_saml_provider":             true,
	"aws_inspector_assessment_target":   true,
	"aws_inspector_assessment_template": true,
	"aws_inspector_resource_group":      true,
	"aws_resourcegroups_group":          true,
	"aws_shield_protection":             true,
	"aws_vpn_connection":                true,
	"aws_xray_sampling_rule":            true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_lb_listener": true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_lb_listener_rule": true,
	// ValidationException: Please provide either one or more instance ID or one stack ID or one layer ID
	"aws_opsworks_instance":      true,
	"aws_default_route_table":    true,
	"aws_default_security_group": true,
	// A normal VPC with "IsDefault" attribute set to true
	"aws_default_vpc": true,
	// Error: failed to read current state of resource: rpc error: code = Canceled desc = context canceled
	"aws_kms_alias": true,
	// Not a resource. This are the instances registered with an ELB. Requires the ELB name.
	"aws_elb_attachment": true,
	// use aws_db_instance and filter by the cluster ID to get DB instances for a particular cluster
	"aws_rds_cluster_instance": true,
}

// manualMatchedListOps are list operations that could not be matched automatically
var ManualMatchedListOps = map[string]string{
	"aws_appautoscaling_target":           "DescribeScalableTargets",
	"aws_appautoscaling_policy":           "DescribeScalingPolicies",
	"aws_appautoscaling_scheduled_action": "DescribeScheduledActions",
	// The AWS API calls it images not AMI
	"aws_ami":                  "DescribeImages",
	"aws_ecs_service":          "DescribeServices",
	"aws_ecs_cluster":          "ListClusters",
	"aws_eip":                  "DescribeAddresses",
	"aws_elb":                  "DescribeLoadBalancers",
	"aws_route53_zone":         "ListHostedZones",
	"aws_cloudformation_stack": "DescribeStacks",
	// GetParameters has required input fields, so DescribeParameters it is
	"aws_ssm_parameter":               "DescribeParameters",
	"aws_ssm_resource_data_sync":      "ListResourceDataSync",
	"aws_lb":                          "DescribeLoadBalancers",
	"aws_cloudtrail":                  "DescribeTrails",
	"aws_workspaces_directory":        "DescribeWorkspaceDirectories",
	"aws_rds_cluster":                 "DescribeDBClusters",
	"aws_rds_cluster_endpoint":        "DescribeDBClusterEndpoints",
	"aws_rds_cluster_parameter_group": "DescribeDBClusterParameterGroups",
	"aws_redshift_parameter_group":    "DescribeClusterParameterGroups",
	"aws_redshift_security_group":     "DescribeClusterSecurityGroups",
	"aws_redshift_subnet_group":       "DescribeClusterSubnetGroups",
}

var ManualMatchedOutputFields = map[string]string{
	"aws_ecs_service": "ServicePkgNames",
}

var ManualMatchedResourceID = map[string]string{
	"aws_ami":                                       "ImageId",
	"aws_autoscaling_group":                         "AutoScalingGroupName",
	"aws_launch_configuration":                      "LaunchConfigurationName",
	"aws_s3_bucket":                                 "Name",
	"aws_elb":                                       "LoadBalancerName",
	"aws_db_instance":                               "DBInstanceIdentifier",
	"aws_route53_zone":                              "Id",
	"aws_ses_email_identity":                        "EmailAddress",
	"aws_ses_domain_identity":                       "Domain",
	"aws_simpledb_domain":                           "DomainName",
	"aws_workspaces_workspace":                      "WorkspaceId",
	"aws_subnet":                                    "SubnetId",
	"aws_imagebuilder_component":                    "Arn",
	"aws_imagebuilder_distribution_configuration":   "Arn",
	"aws_imagebuilder_infrastructure_configuration": "Arn",
	"aws_imagebuilder_image":                        "Arn",
	"aws_imagebuilder_image_pipeline":               "Arn",
	"aws_imagebuilder_image_recipe":                 "Arn",
	"aws_workspaces_directory":                      "DirectoryId",
	"aws_rds_cluster":                               "DBClusterIdentifier",
	"aws_rds_cluster_endpoint":                      "DBClusterEndpointIdentifier",
}

var Inputs = map[string]string{
	"aws_iam_policy": `
Scope: "Local",
`,
	"aws_ebs_snapshot": `
OwnerIds: []string{"self"},
`,
	"aws_ami": `
Owners: []string{"self"},
`,
}
