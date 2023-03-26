.PHONY: knative
knative:
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.9.2/serving-crds.yaml
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.9.2/serving-core.yaml

.PHONY: istio
istio:
	kubectl apply -l knative.dev/crd-install=true -f https://github.com/knative/net-istio/releases/download/knative-v1.9.2/istio.yaml
	kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.9.2/istio.yaml
	kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.9.2/net-istio.yaml

.PHONY: istio2
istio2:
	istioctl install -y
	kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v1.9.2/net-istio.yaml
	kubectl label namespace knative-serving istio-injection=enabled
	kubectl apply -f setup/peer-authentication.yaml

.PHONY: local-registry
local-registry:
	docker run -d -p 5005:5005 --restart=always --name registry registry:2
