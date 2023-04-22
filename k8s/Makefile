.PHONY: istio
istio:
	istioctl install

.PHONY: setup
setup: enable-inject
	kubectl apply -f setup/default-peer-authentication.yaml
	kubectl apply -f setup/allow-nothing.yaml

.PHONY: enable-inject
enable-inject:
	kubectl label namespace default istio-injection=enabled --overwrite
	kubectl delete pods --all --namespace default

