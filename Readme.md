## Dependencies
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors   //we need to implement cors protection so that we can communicate from front end to backend

## Dependencies for Authentication service 
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors  
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/v4/stdlib


## Dependencies for logger service
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors 
mongo compass
mongodb://admin:password@localhost:27017/logs?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false

go get google.golang.org/grpc
go get google.golang.org/protobuf

## Dependencies for mailer service
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors 
go get github.com/vanng822/go-premailer/premailer
go get github.com/xhit/go-simple-mail/v2

## Dependencies for listner
go get github.com/rabbitmq/amqp091-go



go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto


## Docker swarm

Swarm initialized: current node (shdjygvknt4fdmiizg7875ziw) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-4x7zgquy14ftf3bn21ym2yove1bzg0yg5ifh5z7z9s41masrqa-82e9xcp9mu5m00ac6yoozw4em 192.168.65.3:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.

### To get back the token 

docker swarm join-token worker
docker swarm join-token manager 