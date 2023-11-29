build:
	GOOS=linux GOARCH=amd64 go build -o src/ src/cmd/api/main.go
clean:
	rm src/main
init:
	terraform -chdir=deployments init
apply:
	make build && terraform -chdir=deployments apply -var-file="variables.tfvars"
destroy:
	terraform -chdir=deployments destroy -var-file="variables.tfvars"
