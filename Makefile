helm_upgrade:
	helm upgrade zip-fetcher -n zip-fetcher helm/go_zipper/


helm_uninstall:
	helm uninstall zip-fetcher -n zip-fetcher helm/go_zipper/

helm_install:
	helm install zip-fetcher -n zip-fetcher helm/go_zipper/