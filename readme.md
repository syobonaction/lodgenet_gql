# Run the GQL server locally
1. Install the latest version of Docker desktop (https://www.docker.com/products/docker-desktop/)
2. From the root directory, run the following commands:
    - `go mod tidy`
    - `docker build -t docker-lodgenet-gql .`
        - This should publish an image of your container. Check the images section of your docker desktop app.
    - `docker run --publish 8080:8080 docker-lodgenet-gql`
        - This should run a container based off of the image you created
3. Check http://localhost:8080. You should see the igql app.
4. Run a query to test that data is retrieved.
5. You should now be able to get data in the FE app.