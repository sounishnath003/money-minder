DockerImageName=asia-south1-docker.pkg.dev/sounish-cloud-workstation/sounish-cloud-workstation/moneyminder

.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify
	cd web && npm install

.PHONY: build-backend
build-backend:
	gofmt -w cmd/
	gofmt -w internal/
	GOARCH=arm64 go build -o ./tmp/moneyminder cmd/*.go

.PHONY: run-backend
run-backend: build-backend
	export SITE_SECRET_PUBLIC_KEY=$(SITE_SECRET_PUBLIC_KEY)
	export SITE_SECRET_PRIVATE_KEY=$(SITE_SECRET_PRIVATE_KEY)
	./tmp/moneyminder

.PHONY: frontend-run-dev
frontend-run-dev:
	cd web && npm run dev

.PHONY: build-frontend
build-frontend:
	cd web && npm run build

.PHONY: build-all
build-all: build-backend build-frontend
	docker images
	echo "Removing old images"
	echo "Building image"
	docker build --build-arg VITE_COMMIT_HASH=$$(git rev-parse --short HEAD) -t $(DockerImageName):$$(git rev-parse --short HEAD) -f Dockerfile .
	echo "Pushing to GCP"
	gcloud auth configure-docker asia-south1-docker.pkg.dev
	docker push $(DockerImageName):$$(git rev-parse --short HEAD)


.PHONY: deploy-all
deploy-all: build-all
	gcloud run deploy moneyminder \
	--image=$(DockerImageName):$$(git rev-parse --short HEAD) \
	--allow-unauthenticated \
	--port=3000 \
	--service-account=797087556919-compute@developer.gserviceaccount.com \
	--max-instances=5 \
	--region=asia-south1 \
	--set-env-vars=JWT_SECRET=$$(JWT_SECRET),SITE_SECRET_PUBLIC_KEY=$$(SITE_SECRET_PUBLIC_KEY),SITE_SECRET_PRIVATE_KEY=$$(SITE_SECRET_PRIVATE_KEY)