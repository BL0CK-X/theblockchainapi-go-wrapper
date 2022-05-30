generate:
	openapi-generator-cli generate -i api/openapi.yaml -g go -o . --git-repo-id theblockchainapi-go-wrapper --git-user-id BL0CK-X