ROOT=functions
5W=what who when where why
FILE=ascii
setup:
	@apex init
deps:
	@go get ./src/...
test: deps
	@go test ./src/...
clean:
	@for folder in $(5W) ; do \
    rm -rf ./$(ROOT)/$$folder ; \
	done
duplicate:
	@for folder in $(5W) ; do \
		mkdir ./$(ROOT)/$$folder && \
		find ./src -name '*.go' -exec cp -prv '{}' "./$(ROOT)/$$folder/" ';' && \
		cp ./data/$$folder.json ./$(ROOT)/$$folder/$$folder.json && \
		echo "*.go" > ./$(ROOT)/$$folder/.apexignore && \
		sed -i '' "s/masterfunc/$$folder/g" ./$(ROOT)/$$folder/main.go ; \
	done
deploy: deps clean duplicate
	@apex deploy
art:
	@cat $(FILE)
plot: art
	@echo '{ "value": "What if" }' | \
	apex invoke who | \
	apex invoke what | \
	apex invoke who | \
	apex invoke when | \
	apex invoke where | \
	apex invoke why | \
	json value && echo "\n"
