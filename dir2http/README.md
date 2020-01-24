# dir2http: Directory to a HTTP Server

Makes a local directory accessible via HTTP.

## Usage
```
dir2http 1234 ./test

curl localhost:1234 
curl localhost:1234/page/
curl localhost:1234/page/next.html
curl localhost:1234/image.jpg --output image.jpg
```