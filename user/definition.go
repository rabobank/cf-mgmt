package user

//go:generate counterfeiter -o fakes/fake_cf_client.go types.go CFClient
//go:generate counterfeiter -o fakes/fake_mgr.go types.go Manager
//go:generate counterfeiter -o fakes/fake_ldap_mgr.go types.go LdapManager
//go:generate counterfeiter -o fakes/fake_aad_mgr.go types.go AzureADManager
