&larr; [back to Commands](../README.md)

# `cf-mgmt update-org-users`

`update-org-users` command will:

- adds users from `ldap_groups` configured in orgConfig.yml assuming that ldap.yml is configured
- adds ldap users in `ldap_users` configured in orgConfig.yml assuming that ldap.yml is configured
- add internal `users` configured in orgConfig.yml (internal users must exist in uaa first)
- add `saml_users` configured in orgConfig.yml (internal users must exist in uaa first)
- will remove users from roles if `enable-remove-users` is set to `true` in orgConfig.yml

## Command Usage

```
Usage:
  cf-mgmt [OPTIONS] update-org-users [update-org-users-OPTIONS]

Help Options:
  -h, --help               Show this help message

[update-org-users command options]
  --config-dir=    Name of the config directory (default: config) [$CONFIG_DIR]
  --system-domain= system domain [$SYSTEM_DOMAIN]
  --user-id=       user id that has privileges to create/update/delete users, orgs and spaces [$USER_ID]
  --password=      password for user account [optional if client secret is provided] [$PASSWORD]
  --client-secret= secret for user account that has sufficient privileges to create/update/delete users,
                   orgs and spaces] [$CLIENT_SECRET]
  --ldap-password= LDAP password for binding [$LDAP_PASSWORD]
  --aad-tennantid= Azure AD Tennant id [$AAD_TENNANT_ID]
  --aad-clientid=  Azure AD Client Id [$AAD_CLIENT_ID]
  --aad-secret=    Azure AD Client secret [$AAD_SECRET]
  --aad-origin=    Azure AD Origin [$AAD_ORIGIN]
  --peek           Preview entities to change without modifying. [$PEEK]
```
