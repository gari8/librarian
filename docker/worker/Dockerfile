FROM public.ecr.aws/lambda/provided:al2 as build

RUN yum install -y golang && \
    go env -w GOPROXY=direct

ADD go.mod go.sum ./
RUN go mod download && \
    go mod tidy

ADD . .
RUN go build -o /main ./cmd/worker

FROM public.ecr.aws/lambda/provided:al2

COPY --from=build /main /main
ADD https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie /usr/bin/aws-lambda-rie
COPY entry.sh /

RUN chmod 755 /usr/bin/aws-lambda-rie /entry.sh
ENTRYPOINT ["/entry.sh"]
CMD ["/main"]


