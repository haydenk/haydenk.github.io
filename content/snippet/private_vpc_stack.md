+++
title = 'private_vpc_stack.yml'
date = 2025-04-27T13:45:28Z
type = "snippet"
tags = ['aws', 'cloudformation']
+++

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: 'VPC Creation'
Resources:
  PrivateVpc:
    Type: 'AWS::EC2::VPC'
    Properties:
      CidrBlock: '172.16.0.0/16'
      EnableDnsSupport: true
      EnableDnsHostnames: false
      Tags:
        - Key: Name
          Value: !Sub 'Private-VPC'
  PrivateVpcMainRouteTable:
    Type: 'AWS::EC2::RouteTable'
    Properties:
      VpcId: !Ref PrivateVpc
      Tags:
        - Key: Name
          Value: !Sub 'Private-Main-RouteTable'

  PrivateVpcIPv6CidrBlock:
     Type: 'AWS::EC2::VPCCidrBlock'
     Properties:
         AmazonProvidedIpv6CidrBlock: true
         VpcId: !Ref PrivateVpc
  PrivateSubnetA:
    Type: 'AWS::EC2::Subnet'
    DependsOn: PrivateVpcIPv6CidrBlock
    Properties:
      AssignIpv6AddressOnCreation: true
      AvailabilityZone: !Select [0, Fn::GetAZs: !Ref 'AWS::Region']
      CidrBlock: !Select [0, !Cidr [!GetAtt PrivateVpc.CidrBlock, 4, 8]]
      Ipv6CidrBlock: !Select [0, !Cidr [ !Select [0, !GetAtt PrivateVpc.Ipv6CidrBlocks], 4, 64]]
      VpcId: !Ref PrivateVpc
      Tags:
        - Key: Name
          Value: !Sub 'Private-Subnet-A'
  PrivateSubnetB:
    Type: 'AWS::EC2::Subnet'
    DependsOn: PrivateVpcIPv6CidrBlock
    Properties:
      AssignIpv6AddressOnCreation: true
      AvailabilityZone: !Select [1, Fn::GetAZs: !Ref 'AWS::Region']
      CidrBlock: !Select [1, !Cidr [!GetAtt PrivateVpc.CidrBlock, 4, 8]]
      Ipv6CidrBlock: !Select [1, !Cidr [ !Select [0, !GetAtt PrivateVpc.Ipv6CidrBlocks], 4, 64]]
      VpcId: !Ref PrivateVpc
      Tags:
        - Key: Name
          Value: !Sub 'Private-Subnet-B'
  PrivateSubnetC:
    Type: 'AWS::EC2::Subnet'
    DependsOn: PrivateVpcIPv6CidrBlock
    Properties:
      AssignIpv6AddressOnCreation: true
      AvailabilityZone: !Select [2, Fn::GetAZs: !Ref 'AWS::Region']
      CidrBlock: !Select [2, !Cidr [!GetAtt PrivateVpc.CidrBlock, 4, 8]]
      Ipv6CidrBlock: !Select [2, !Cidr [ !Select [0, !GetAtt PrivateVpc.Ipv6CidrBlocks], 4, 64]]
      VpcId: !Ref PrivateVpc
      Tags:
        - Key: Name
          Value: !Sub 'Private-Subnet-C'
  PrivateSubnetD:
    Type: 'AWS::EC2::Subnet'
    DependsOn: PrivateVpcIPv6CidrBlock
    Properties:
      AssignIpv6AddressOnCreation: true
      AvailabilityZone: !Select [3, Fn::GetAZs: !Ref 'AWS::Region']
      CidrBlock: !Select [3, !Cidr [!GetAtt PrivateVpc.CidrBlock, 4, 8]]
      Ipv6CidrBlock: !Select [3, !Cidr [ !Select [0, !GetAtt PrivateVpc.Ipv6CidrBlocks], 4, 64]]
      VpcId: !Ref PrivateVpc
      Tags:
        - Key: Name
          Value: !Sub 'Private-Subnet-D'
Outputs:
  PrivateVpc:
    Value: !Ref PrivateVpc
  PrivateMainRouteTable:
    Value: !Ref PrivateVpcMainRouteTable
```
