+++
title = 'dns_host_management_stack.yml'
date = 2025-04-27T13:47:28Z
type = "snippet"
tags = ['aws', 'cloudformation']
+++

<!--more-->
```yaml
AWSTemplateFormatVersion: '2010-09-09'
Description: DNS Host Management
Parameters:
  PrimaryDomainName:
    Type: String
Resources:
  PrimaryDomain:
    Type: 'AWS::Route53::HostedZone'
    Properties:
      Name: !Ref PrimaryDomainName
  PrimaryDomainMx:
    Type: 'AWS::Route53::RecordSetGroup'
    Properties:
      HostedZoneId: !GetAtt PrimaryDomain.Id
      Comment: 'Creating records for mail server'
      RecordSets:
      - Name: !Ref PrimaryDomainName
        Type: MX
        TTL: 300
        ResourceRecords:
        - '10 in1-smtp.messagingengine.com'
        - '20 in2-smtp.messagingengine.com'
      - Name: !Sub "fm1._domainkey.${PrimaryDomainName}"
        Type: CNAME
        TTL: 300
        ResourceRecords:
        - 'fm1.domain.com.dkim.fmhosted.com'
      - Name: !Sub "fm2._domainkey.${PrimaryDomainName}"
        Type: CNAME
        TTL: 300
        ResourceRecords:
        - 'fm2.domain.com.dkim.fmhosted.com'
      - Name: !Sub "fm3._domainkey.${PrimaryDomainName}"
        Type: CNAME
        TTL: 300
        ResourceRecords:
        - 'fm3.domain.com.dkim.fmhosted.com'
      - Name: !Ref PrimaryDomainName
        Type: TXT
        TTL: 300
        ResourceRecords:
        - '"v=spf1 include:spf.messagingengine.com ~all"'
Outputs:
  PrimaryDomainNameservers:
    Value: !GetAtt PrimaryDomain.NameServers
```
