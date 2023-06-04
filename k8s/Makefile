.PHONY: istio
istio:
	istioctl install -y

.PHONY: setup
setup: enable-inject
	kubectl apply -f setup/default-peer-authentication.yaml
	kubectl apply -f setup/allow-nothing.yaml

.PHONY: enable-inject
enable-inject:
	kubectl label namespace default istio-injection=enabled --overwrite
	kubectl delete pods --all --namespace default

.PHONY: strimzi
strimzi:
	helm install strimzi-operator strimzi/strimzi-kafka-operator

.PHONY: backend
backend:
	helm upgrade --install backend backend

.PHONY: collector
collector:
	helm upgrade --install collector collector -f collector/prod_values.yaml

.PHONY: kafka
kafka:
	helm upgrade --install kafka kafka
