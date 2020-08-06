main: main.go 
		go build


clean:
		go clean
		
		
		
before:
	echo "make before"
	
build:
	echo "make build"
	
test:
	echo "make test"
	
deploy:
	echo "make deploy"
	
after:
	echo "make after"
		

