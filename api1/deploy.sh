docker build -t brandonrvice/rythm-api1:latest .
docker push brandonrvice/rythm-api1:latest
kubectl rollout restart deployment rythm-api1
