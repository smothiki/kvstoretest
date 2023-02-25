### Instructions to execute the code.

#### Prereqs
docker required

#### Running the program.
using docker you can run the below command once you clone the code to a folder. 
```
docker run  -p 8000:8000 -w /go/src/github.com/smothiki/testcode --cpus=1 --memory=1g -v $PWD:/go/src/github.com/smothiki/testcode golang go run *.go
```
This should start the server on port 8000.
you can request a key using the following curl syntax 
`curl localhost:8000/key/<uuid>`.

Every response is a 200 but if the key is not found its a not found resposne in the body.

The program takes almost 2 minutes to startup as it has to download go libraries and there is a setup time for DB using the example data.

#### Observations on a 1cpu and 1GB memory node.
* Using plain hashmaps was able to serve the 300RPS with less than 1milli second latency for a million records and each record not exceeding 90bytes.
* Using pogreb DB with memory usage is less than what native hashmap used but for the same 300RPS pbserved average response times of 6 milli seconds.
* Latency definetely improved by adding LRU cache for podgeb DB usecase in the second iteration the latency is down to 3-4 milli seconds.

#### Failure patterns
* This is a single server and prone to unprecedented crashes and renders the app un usable.
