npm run build
docker build -t brandonvio/rythm-app1:latest .
docker push brandonvio/rythm-app1:latest
kubectl rollout restart deployment rythm-app1
