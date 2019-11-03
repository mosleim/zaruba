build:
	go build
	./build.sh
test:
	# tear down
	rm -Rf playground/projects/test-watch
	rm -Rf playground/projects/test-create
	# prepare
	cp -r playground/projects/example playground/projects/test-watch
	# test
	go test -race ./... -v -coverprofile=profile.out -count=1 -covermode=atomic
	go tool cover -html=profile.out