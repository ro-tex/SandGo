
## Build a static binary
`CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -o deployable .`    
Binary size: 5.6MB  

Build flags:  
`-a` Force rebuild  
`-o` Custom output name  

Linking flags:  
`-w` Omit the DWARF symbol table.  
`-s` Omit the symbol table and debug information.    

## Build the container
`docker build -t inovakov/scratchy .`  
Container image size: 5.8MB

## Run the container (expose at port 80)
`docker run -p 80:3000 --name scratchy inovakov/scratchy`

## Sources
### Articles
[Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07) - quite minimal, covers SSL as well  
[Create the smallest and secured golang docker image based on scratch](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324) - covers multi-stage docker builds

[Container Example](https://gist.github.com/PurpleBooth/ec81bad0a7b56ac767e0da09840f835a) - A Dockerfile that implements a two-stage build

### Docker docs
[FROM scratch](https://hub.docker.com/_/scratch/) - the image we're using + a super simple example  
[Use multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/) - what *are* those multi-stage builds, anyway?
