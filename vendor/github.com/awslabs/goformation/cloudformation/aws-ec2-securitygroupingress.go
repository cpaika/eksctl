package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2SecurityGroupIngress AWS CloudFormation Resource (AWS::EC2::SecurityGroupIngress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type AWSEC2SecurityGroupIngress struct {

	// CidrIp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidrip
	CidrIp *Value `json:"CidrIp,omitempty"`

	// CidrIpv6 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidripv6
	CidrIpv6 *Value `json:"CidrIpv6,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-description
	Description *Value `json:"Description,omitempty"`

	// FromPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-fromport
	FromPort *Value `json:"FromPort,omitempty"`

	// GroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupid
	GroupId *Value `json:"GroupId,omitempty"`

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupname
	GroupName *Value `json:"GroupName,omitempty"`

	// IpProtocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-ipprotocol
	IpProtocol *Value `json:"IpProtocol,omitempty"`

	// SourcePrefixListId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-securitygroupingress-sourceprefixlistid
	SourcePrefixListId *Value `json:"SourcePrefixListId,omitempty"`

	// SourceSecurityGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupid
	SourceSecurityGroupId *Value `json:"SourceSecurityGroupId,omitempty"`

	// SourceSecurityGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupname
	SourceSecurityGroupName *Value `json:"SourceSecurityGroupName,omitempty"`

	// SourceSecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupownerid
	SourceSecurityGroupOwnerId *Value `json:"SourceSecurityGroupOwnerId,omitempty"`

	// ToPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-toport
	ToPort *Value `json:"ToPort,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroupIngress) AWSCloudFormationType() string {
	return "AWS::EC2::SecurityGroupIngress"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2SecurityGroupIngress) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2SecurityGroupIngress
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2SecurityGroupIngress) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2SecurityGroupIngress
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEC2SecurityGroupIngress(*res.Properties)
	}

	return nil
}

// GetAllAWSEC2SecurityGroupIngressResources retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2SecurityGroupIngressResources() map[string]AWSEC2SecurityGroupIngress {
	results := map[string]AWSEC2SecurityGroupIngress{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2SecurityGroupIngress:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SecurityGroupIngress" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2SecurityGroupIngress{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2SecurityGroupIngressWithName retrieves all AWSEC2SecurityGroupIngress items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2SecurityGroupIngressWithName(name string) (AWSEC2SecurityGroupIngress, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2SecurityGroupIngress:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::SecurityGroupIngress" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2SecurityGroupIngress{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2SecurityGroupIngress{}, errors.New("resource not found")
}
