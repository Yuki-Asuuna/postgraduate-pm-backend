go build -o ./output/
PID=$(ps -ef | grep postgraduate-pm-backend | grep -v grep| awk '{print $2}')
kill -9 $PID
./output/postgraduate-pm-backend