ifdef release
docker_build_and_push:
	echo $(release)
	docker build \
		-t registry.baxterhome.net/cbaxter1988/zip-fetcher:$(release) .
	docker push registry.baxterhome.net/cbaxter1988/zip-fetcher:$(release)
endif

helm_upgrade:
	helm upgrade zip-fetcher -n zip-fetcher helm/go_zipper/


helm_uninstall:
	helm uninstall zip-fetcher -n zip-fetcher helm/go_zipper/

helm_install:
	helm install zip-fetcher -n zip-fetcher helm/go_zipper/