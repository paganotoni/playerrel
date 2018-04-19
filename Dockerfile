# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:development as builder

RUN mkdir -p $GOPATH/src/github.com/paganotoni/playerrel
WORKDIR $GOPATH/src/github.com/paganotoni/playerrel

# this will cache the npm install step, unless package.json changes
ADD package.json .
ADD yarn.lock .
RUN yarn install --no-progress
ADD . .
RUN go get $(go list ./... | grep -v /vendor/)
RUN buffalo build --static -o /bin/playerrel

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/playerrel .

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the playerrel to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

# Uncomment to run the migrations before running the binary:
# CMD /bin/playerrel migrate; /bin/playerrel
CMD exec /bin/playerrel
