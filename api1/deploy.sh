docker build -t brandonvio/rythm-api1:latest .
docker push brandonvio/rythm-api1:latest
kubectl rollout restart deployment rythm-api1
