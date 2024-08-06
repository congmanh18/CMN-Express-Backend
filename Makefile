pro:
	docker rmi -f cmn-express-be:1.0
	docker-compose up
dev:
	cd cmd/dev; go run main.go