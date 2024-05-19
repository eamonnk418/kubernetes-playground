
kubernetes-deploy:
	kubectl apply -f kubernetes/manifests/namespace.yml
	kubectl apply -f kubernetes/manifests/deployment.yml

kubernetes-down:
	kubectl delete -f kubernetes/manifests/deployment.yml
	kubectl delete -f kubernetes/manifests/namespace.yml

healthcheck: clean
	gcc -std=c18 -Wall -Werror -Wpedantic -o healthcheck scripts/clang/healthcheck.c

.PHONY: docker-image
docker-image:
	docker build -t kubernetes-playground:1.0.0 .
	docker tag kubernetes-playground:1.0.0 localhost:5001/kubernetes-playground:1.0.0

.PHONY: docker-run
docker-run:
	docker run --rm -dit --name kubernetes-playground kubernetes-playground:1.0.0

docker-push:
	docker push localhost:5001/kubernetes-playground:1.0.0

local-dev-cluster:
	sh local-dev-cluster.sh

clean:
	rm -f healthcheck