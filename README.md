# Docs

## Layout example
https://github.com/snykk/go-rest-boilerplate

## Building the application
The application is split into two seperate parts, the migrator which should run before the application starts, and the API itself. To build the API image run the following command from the root directory
```
docker build --file ./deploy/dashboard.Dockerfile --tag covet_dashboard:latest .
```

## Running the application
The following commands should be run *in order* to ensure the proper startup of the application.

### Run the API
Run the following command to start the API
```
docker run --name covetdashboard -d -p 8000:1234 \
    -e environment=prod \
    -e db_username=db_usr \
    -e db_password=mysecretpassword \
    -e db_host=192.168.87.64 \
    -e db_port=5432 \
    -e db_name=app \
    -e external_address=192.168.87.64 \
    -e external_port=8000 \
    covet_dashboard:latest
```

## Project TODOs
- [x] Example