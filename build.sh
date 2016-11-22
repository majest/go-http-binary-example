img=$(basename `pwd`)
docker run --rm  -v "$GOPATH":/go/ -w /go/src/github.com/majest/$img golang:alpine go build -v
docker build -t majest/$img:latest .
docker login
docker push majest/$img:latest 
