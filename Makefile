default: container
	echo "Container built"
static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blocktor .

clean:
	rm blocktor

test:
	go run ./blocktor.go --dryrun

container: clean static
	docker build  -t joemcmahon/blocktor .
