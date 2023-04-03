.PHONY: knative
knative:
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.9.2/serving-crds.yaml
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.9.2/serving-core.yaml

.PHONY: istio
istio:
	istioctl install -y
	kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.9.2/net-istio.yaml
	kubectl label namespace knative-serving istio-injection=enabled
	kubectl apply -f setup/peer-authentication.yaml

.PHONY: setup
setup: enable-inject
	kubectl apply -f setup/default-peer-authentication.yaml
	kubectl apply -f setup/knative-peer-authentication.yaml
	kubectl apply -f setup/scale-to-zero.yaml
	kubectl apply -f setup/allow-nothing.yaml

.PHONY: enable-inject
enable-inject:
	kubectl label namespace default istio-injection=enabled --overwrite

.PHONY: disable-inject
disable-inject:
	kubectl label namespace default istio-injection=disabled --overwrite

.PHONY: local-registry
local-registry:
	docker run -d -p 5005:5005 --restart=always --name registry registry:2
