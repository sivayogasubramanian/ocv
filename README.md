# OCV Assignment

Name: Ruppa Nagarajan Sivayoga Subramanian

### Getting Started Locally

1. Install go 1.18
1. Install PostgreSQL
1. Create the .env file in the src directory of the project. You can refer to the `./src/.env.example` for the required environment variables.
1. Install dependencies using `go mod tidy` in the src directory.
1. Run `go run main.go` to start the server in the src directory.

The server will be running on `localhost:8080`.

### Production

The application is deployed on a digital ocean droplet. NGINX API gateway is used to manage the docker deployment setup.

Please make requests to "https://ocv.sivarn.com"

Example: "https://ocv.sivarn.com/api/register"

Note: HTTPS is strictly enforced. Please make sure to use https in the request.
