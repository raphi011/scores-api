### DEPLOY ###
deploy: deploy-frontend deploy-backend

deploy-frontend:

deploy-backend:


### RUN ### 
dependencies:
	dep ensure
	cd frontend; npm install

run-frontend: 
	cd frontend; npm run dev

run-backend: install-backend
	${GOPATH}/bin/web -goauth client_secret.json

install-backend:
	go install ./cmd/web
