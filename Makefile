ROOT=functions
5W=what who when where why
clean:
	for folder in $(5W) ; do \
    rm -rf ./$(ROOT)/$$folder ; \
	done
duplicate:
	for folder in $(5W) ; do \
    cp -r ./src ./$(ROOT)/$$folder && \
		cp ./data/$$folder.json ./$(ROOT)/$$folder/$$folder.json && \
		sed -i '' "s/masterfunc/$$folder/g" ./$(ROOT)/$$folder/main.go ; \
	done
deploy: clean duplicate
	apex deploy