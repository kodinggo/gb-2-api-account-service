FROM golang:latest

WORKDIR /app

# RUN git config --global url."https://${GITLAB_TOKEN}:@gitlab.com/".insteadOf "https://gitlab.com/"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# It's required to run server on Heroku
ENV PORT=1325

EXPOSE 1325

CMD [ "./main", "serve" ]