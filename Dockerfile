# Build stage
FROM golang:1.19.5-alpine3.17 AS builder
ARG BITBUCKET_USER
ARG BITBUCKET_PASSWORD

# git is required to fetch go dependencies
RUN apk add --no-cache ca-certificates git

# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Create a netrc file using the credentials specified using --build-arg
RUN printf "machine bitbucket.org\n\
    login ${BITBUCKET_USER}\n\
    password ${BITBUCKET_PASSWORD}\n\
    \n\
    machine api.bitbucket.org\n\
    login ${BITBUCKET_USER}\n\
    password ${BITBUCKET_PASSWORD}\n"\
    >> /root/.netrc

RUN chmod 600 /root/.netrc

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /src

# Import the code from the context.
COPY . .

# Download dependency
RUN go mod download

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

# Run stage
FROM alpine:3.17 AS run

# Assing env value
ARG ENV
ENV ENV=$ENV

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Import localize files
COPY --from=builder /src/locales /locales

# Perform any further action as an unprivileged user.
USER nobody:nobody

EXPOSE 3000

# Run the compiled binary.
ENTRYPOINT ["/app"]
