Writing tests for infrastructure is very important, and to ensure I got the docker image build right, I wrote tests using Terratest. 

Using the test I wrote, I was able to ensure that the Docker image was built right. However, with all attempts to make this test pass, the port to access the webpage was denied.

## Test Process

Using go, I defined the test functions to build the image, and run a container that destroys itself when stopped. See the code [here](https://github.com/LuluNwenyi/SCA-Cloud-School-Application/tree/start/go_test).

To run the test locally,

- Clone this repository;
    
    ```bash
    git clone 
    ```
    
- Go to the “go_test” directory;
    
    ```bash
    cd go_test
    ```
    
- Initiate the go environment and requirements(sum) file;
    
    ```bash
    # to initiate the environment
    go mod init docker_test.go
    
    # to create sum file
    go mod tidy
    ```
    
- Run the test;
    
    ```bash
    go test
    ```
    

## Build Process

After running this test and I proceeded to build the first version of the image for this task. Using the command:

```bash
sudo docker build -t sca-task:v0 .
```

By doing this, an image called “sca-task” with the tag “v0” was created. To run a container with this image on port 5001, I used the command:

```bash
sudo docker run --name sca-task -p 5001:5001 sca-task:v0
```

After running this command, we now have a container named “sca-task” using the “sca-task” image on port 5001.

To update the app content as mentioned in the assessment, I also had to rebuild the image using a different tag:

```bash
sudo docker build -t sca-task:v1 .
```

This builds an image from the updated version of the app, that helps to track changes.

### Output

My flask app ran on the signified port — **5001**. 

***v0:***



***v1:***


## Dockerhub
[Link to image on Dockerhub](https://hub.docker.com/r/lulunwenyi/sca-task)
