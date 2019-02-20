dep-vendor:
	docker run -it --rm --name go-dep-vendor -e GO111MODULE=on -v `pwd`:/tmp/statuspage -w="/tmp/statuspage" golang go mod vendor

gofmt: 
	for i in `find ./* -not -path "./vendor*" -name '*.go'`; do gofmt -s $$i | sponge $$i ; done

test-all:
	docker-compose down
	docker-compose up -d mongo
	export MONGO_URI="localhost"
	for i in `find ./* -maxdepth 1 -type d -not -path "./vendor*" -not -path "./docs*"`; do go test $$i -coverprofile /dev/null ; done
	docker-compose down

build:
	docker-compose build statuspage

run: 
	docker-compose down
	docker-compose up -d 
	docker-compose logs -f
