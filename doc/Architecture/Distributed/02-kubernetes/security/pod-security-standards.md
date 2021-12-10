# Pod Security Standards

The Pod Security Standards define three different *policies* to broadly cover the security spectrum. These policies are cumulative and range from highly-permissive to highly-restrictive. 

|Profile|Description|
|-|-|
|Privileged|Unrestricted policy, providing the widest possible level of permissions. This policy allows for known privilege escalations|
|Baseline|Minimally restrictive policy which prevents known privilege escalations.Allow the dfault(minimally specified) Pod configuration|
|Restricted|Heavily restricted policy, following current Pod hardening best practises|