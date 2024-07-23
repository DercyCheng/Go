.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook .
	@docker rmi -f flycash/webook:v0.0.1
	@docker build -t flycash/webook:v0.0.1 .

.PHONY: k8s
k8s:
	@kubectl delete deployment webook-record-mysql || true
	@kubectl delete deployment webook-record-redis || true
	@kubectl delete deployment webook-record-service || true
	@kubectl apply -f webook-deployment.yaml
	@kubectl apply -f webook-record-ingress.yaml
	@kubectl apply -f webook-record-mysql-deployment.yaml
	@kubectl apply -f webook-record-mysql-pv.yaml
	@kubectl apply -f webook-record-mysql-pvc.yaml
	@kubectl apply -f webook-record-mysql-service.yaml
	@kubectl apply -f webook-record-redis-deployment.yaml
	@kubectl apply -f webook-record-redis-service.yaml
	@kubectl apply -f webook-service.yaml

#	@kubectl get deployment