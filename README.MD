## Start Commands in bash

### Build the image and name it 'gatherchain-app-deploy'
docker build -t gatherchain-app-deploy -f Dockerfile .

### Runs the Docker image and connects the 8010 ports
docker run -it -p 8010:8010 gatherchain-app-deploy