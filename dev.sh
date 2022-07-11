cd src
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -o ../.build/app . # CGO_ENABLED=0 - building app statically
cd ..

sudo docker build -f Dockerfile-dev -t blog-dev .
sudo docker-compose -f docker-compose-dev.yml up
