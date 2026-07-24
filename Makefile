VERSION ?= 0.1.0
PROVIDER := pulumi-resource-systeam
TFGEN := pulumi-tfgen-systeam
VERSION_PKG := github.com/systeampl/pulumi-systeam/provider/pkg/version
LDFLAGS := -s -w -X $(VERSION_PKG).Version=$(VERSION)

.PHONY: build generate_schema generate_sdk install clean

build: generate_sdk
	go build -ldflags "$(LDFLAGS)" -o bin/$(PROVIDER) ./provider/cmd/pulumi-resource-systeam/

generate_schema:
	go build -o bin/$(TFGEN) ./provider/cmd/pulumi-tfgen-systeam/
	bin/$(TFGEN) schema --out provider/cmd/pulumi-resource-systeam/

generate_sdk: generate_schema
	rm -rf sdk/python
	bin/$(TFGEN) python --out sdk/python/

install: build
	mkdir -p $(HOME)/.pulumi/plugins/resource-systeam-v$(VERSION)
	cp bin/$(PROVIDER) $(HOME)/.pulumi/plugins/resource-systeam-v$(VERSION)/

clean:
	rm -rf bin/ sdk/python/
