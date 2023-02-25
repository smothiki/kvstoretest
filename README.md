### instructions to execute the code.

#### Prereqs
docker required.
using docker you can run the below command once you clone the code to a folder. 
```
docker run  -p 8000:8000 -w /go/src/github.com/smothiki/testcode --cpus=1 --memory=1g -v $PWD:/go/src/github.com/smothiki/testcode golang go run *.go
```
This should start the server on port 8000.
you can request a key using the following curl syntax 
`curl localhost:8000/key/<uuid>`.

Every response is a 200 but if the key is not found its a not found resposne in the body.

The program takes almost 2 minutes to startup as it has to download go libraries and there is a setup time for DB using the example data.
