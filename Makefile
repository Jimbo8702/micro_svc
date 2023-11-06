user: 
	@go build -o bin/user_svc ./user_svc
	@./bin/user_svc

accounts: 
	@go build -o bin/account_svc ./account_svc
	@./bin/user_svc

.PHONY: user