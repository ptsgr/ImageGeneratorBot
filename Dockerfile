FROM alpine AS image_generator
WORKDIR /src/
ADD ./assets/* /src/assets/
COPY ./build/imageGenerator /src/imageGenerator
ENTRYPOINT ["/src/imageGenerator"]
