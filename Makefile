setup:
	npm install json-to-env -g

run:
	json-to-env ./env.json ./env.sh
	source ./env.sh && rm ./env.sh && go run main.go