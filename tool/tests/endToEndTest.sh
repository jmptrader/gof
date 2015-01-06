#!/bin/bash

fail(){
	echo "Failed... deleting bootstrap"
	deleteFiles
	exit 1
}

deleteFiles(){
	rm -f tests_suite_test.go

	for i in $( ls *.gof); do
		rm `basename $i .gof`.go
	done
}

# Create ginkgo bootstrap
ginkgo bootstrap
if [[ $? -ne 0 ]]; then 
	echo "Failed to ginkgo bootstrap" 
	fail
fi

# Generate Go from GoF
go run $GOPATH/src/github.com/apoydence/gof/tool/gof.go generate 
if [[ $? -ne 0 ]]; then
	echo "Generating go from gof failed" 
	fail
fi

# Run Tests
ginkgo -r
if [[ $? -ne 0 ]]; then
	echo "Tests failed" 
	fail
fi

deleteFiles
