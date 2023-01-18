In short, this API is a simple REST API that provides functionality to encrypt and decrypt strings, and saves the requests to a PostgreSQL database. It uses the "github.com/lib/pq" package to connect to the database, the "RestAPI/cryptLogic" package for encryption and decryption, and the "github.com/joho/godotenv" and "github.com/rs/zerolog/log" packages for loading environment variables and logging respectively.



>To build the Docker image posgreSQL, navigate to the directory where the Dockerfile is located, and run the following command:
>-         docker build -t my-postgres .
> To run the container with DB using the following command:
>-         docker run --name my-postgres -p 5432:5432 -d my-postgres
