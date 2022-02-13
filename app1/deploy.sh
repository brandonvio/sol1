npm run build
docker build -t brandonrvice/rythm-app1:latest .
docker push brandonrvice/rythm-app1:latest
kubectl rollout restart deployment rythm-app1
