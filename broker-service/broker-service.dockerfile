# copy executable to this
from alpine:latest

run mkdir /app

copy brokerApp /app

cmd ["/app/brokerApp"]
