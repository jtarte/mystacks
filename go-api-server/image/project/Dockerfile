# Dockerfile for building your application.
# Defines the final image that contains content from both the image and template.
# Step 1 build the application
FROM golang:1.13.7 as buildEnv
#get dep util tool
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# Get the source code of application
COPY . /project
# Set env variable fo GO environment
ENV GOPATH=/project/userapp
# ensure all the dependencies are present
WORKDIR /project/userapp/src/myapp
RUN dep ensure
# build the executable application
WORKDIR /project/userapp
RUN go build -o /userapp myapp
# ensure execution priviledges for application
RUN chmod +x /userapp

# Step 2 build the final image
# use 
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /project
# get the application from the build container (buildEnv)
COPY --from=buildEnv /userapp ./
# expose the port used
ENV PORT=8080
EXPOSE ${PORT}
# Pass control your application
CMD ["/project/userapp"]