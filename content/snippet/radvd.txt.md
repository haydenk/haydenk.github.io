+++
title = 'radvd.txt'
date = 2025-05-05T14:37:02Z
type = "snippet"
+++

This helps pihole to be the authority for ipv6 dhcpcd

```text
interface eth0 {
    AdvSendAdvert on;
    AdvLinkMTU 3600;
    AdvHomeAgentFlag off;
    AdvManagedFlag off;
    AdvOtherConfigFlag off;
    MinRtrAdvInterval 30;
    MaxRtrAdvInterval 100;
    prefix 2600:1700:7d48:b210::/64 {
        AdvOnLink on;
        AdvAutonomous on;
        AdvRouterAddr on;
    };

    RDNSS 2600:1700:7d48:b210:9858:17e0:6e9c:9e8 {
        AdvRDNSSLifetime 3600;
    };
};
```