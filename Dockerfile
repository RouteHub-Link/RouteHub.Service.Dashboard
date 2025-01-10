# Build stage
FROM node:16 AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

# RUN npx tailwindcss -i web/public/css/input.css -o web/public/css/output.css

FROM golang:alpine AS go-build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# RUN go install github.com/a-h/templ/cmd/templ@latest
# RUN templ generate

RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

# Create an empty config.json file {}
RUN echo "{}" > /config.json

# Release stage
FROM gcr.io/distroless/static-debian12

WORKDIR /publish

COPY --chown=nonroot:nonroot --from=go-build /entrypoint /publish/entrypoint
COPY --chown=nonroot:nonroot --from=build /app/web/public /publish/public
COPY --chown=nonroot:nonroot --from=build /app/node_modules/preline/dist /publish/node_modules/preline/dist
COPY --chown=nonroot:nonroot --from=go-build /config.json /publish/config.json

USER nonroot:nonroot

ENTRYPOINT ["/publish/entrypoint"]
