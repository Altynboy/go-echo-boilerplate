# generate docs html from postman collection

docker build -t postmanerator .
docker run -v "$(pwd)/output:/app/postmanerator/output" postmanerator -output="/app/postmanerator/output/index.html" -collection="/app/postmanerator/mycollection.json" -environment="/app/to/environment.json"
