# Python servers
for i in {1..5}; do python server.py server-$i 500$i &; done

# Go curls
for i in {1..20}; do curl localhost:8000; done
