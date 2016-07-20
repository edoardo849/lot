ROOT=functions
5W=what who when where why
clean:
	@for folder in $(5W) ; do \
    rm -rf ./$(ROOT)/$$folder ; \
	done
duplicate:
	@for folder in $(5W) ; do \
		mkdir ./$(ROOT)/$$folder && \
		find ./src -name '*.go' -exec cp -prv '{}' "./$(ROOT)/$$folder/main.go" ';' && \
		cp ./data/$$folder.json ./$(ROOT)/$$folder/$$folder.json && \
		sed -i '' "s/masterfunc/$$folder/g" ./$(ROOT)/$$folder/main.go ; \
	done
deploy: clean duplicate
	apex deploy
invoke:
	@echo '{ "value": "What if" }' | apex invoke who | apex invoke what | apex invoke who | apex invoke when | apex invoke where | apex invoke why
