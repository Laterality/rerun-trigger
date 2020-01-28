# rerun-trigger

Simple web application which update running docker image.

## Usage

### Configuration

Add yaml configuration file into same directory.

Example config:

```yaml
server:
  port: 8000 # Application port
spec:
  # IMPORTANT: Note that both image name and path does not include image tag.
  # Tag will be specified in each update request.
  image: nginx # Image name
  imagePath: docker.io/library/nginx # Image path to pull
  containerName: webserver
  ports:
    - containerPort: 80
      hostPort: 8080
```

### Build and Run

```bash
go build
./rerun-trigger
```

### Pull and Restart container

`PUT /update`

Pull imgage specified in config file and restart container with given name, or start if container is not running.

> Note: This application does not remove the old images. It may needs to be removed.

#### query parameters

<dl>
    <dt>tag</dt>
    <dd><b>Required</b></dd>
    <dd>Specifies image's tag(e.g. <code>latest</code>, <code>1.17-alpine</code>).</dd>
</dl>

Example:

```bash
curl localhost:8080/update?tag=1.17 # Request run initially

curl localhost:8000 # Check if nginx is running
```
