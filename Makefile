install-db:
	docker run \
		--name postgres-14 \
		--network micronet \
		--rm -e POSTGRES_USER=spadyprod \
		-e POSTGRES_PASSWORD=123456 \
		-p 5432:5432 -it \
		-d postgres:14

run:
	go run src/cmd/main.go -conf=./deploy/conf/develop.env -migration=true

prod:
	go run src/cmd/main.go -conf=./deploy/conf/production.env -migration=false

seed-dev:
	go run src/cmd/seed.go -conf=./deploy/conf/develop.env -acc=./deploy/conf/owner_account.json

seed-prod:
	go run src/cmd/seed.go -conf=./deploy/conf/production.env -acc=./deploy/conf/owner_account.json

gen-token:
	go run src/cmd/gen_token.go -conf=./deploy/conf/production.env -id=${id} -acc_type=${type}

doc:
	swag init -g src/cmd/main.go

container:
	docker build -t backend-spady-prod .
	docker run --network micronet \
            --log-driver=fluentd \
            --log-opt fluentd-address=192.168.228.2:24224 \
            --log-opt tag=backend.backend-spady-prod \
            --name backend-spady-prod -d -p 9991:9991 backend-spady-prod
prometheus:
	 docker run --rm --name prometheus --network micronet --publish 9090:9090 --volume $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml --detach prom/prometheus
test:
	go clean -testcache
	go test ./src/...