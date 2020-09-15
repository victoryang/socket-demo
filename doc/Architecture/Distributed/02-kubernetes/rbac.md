# RBAC

[role & rolebinding & clusterrole & clusterrolebinding](https://blog.csdn.net/hxpjava1/article/details/103779148)

[k8s.io doc rbac](https://kubernetes.io/zh/docs/reference/access-authn-authz/rbac/#%E5%88%9D%E5%A7%8B%E5%8C%96%E4%B8%8E%E9%A2%84%E9%98%B2%E6%9D%83%E9%99%90%E5%8D%87%E7%BA%A7)

[k8s rbac](https://www.cnblogs.com/weiyiming007/p/10484763.html)

## Access Control

- ABAC
- RBAC
- Webhook
- Node
- AlwaysDeny
- AlwaysAllow

### RBAC

权限与角色关联，用户通过称为角色的成员来得到角色的权限。K8S 的RBAC使用 rbac.authorization.k8s.io/v1 API 组驱动认证策略，准许管理员通过API动态配置策略。为了启用 RBAC，需要在 apiserver 启动参数中添加 --authorization-mode=RBAC

