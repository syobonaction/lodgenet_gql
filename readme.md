# Run the GQL server locally
1. Install the latest version of Docker desktop (https://www.docker.com/products/docker-desktop/)
2. Install the latest version of Go (https://go.dev/)
3. From the root directory, run the following commands:
    - `go mod tidy`
    - (If you'd like to skip the next steps on running a docker container, you can just run `go run server.go`)
    - `docker build -t docker-lodgenet-gql .`
        - This should publish an image of your container. Check the images section of your docker desktop app.
    - `docker run --publish 8080:8080 docker-lodgenet-gql`
        - This should run a container based off of the image you created
4. Check http://localhost:8080. You should see the igql app.
5. Run a query to test that data is retrieved.
6. You should now be able to get data in the FE app.
