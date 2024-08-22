ARG WORKINGPATH=/app
ARG ENTRYPATH=/app
ARG PORT=80
ARG CONFIG=development
ARG BASEHREF=/chat/
ARG DEPLOYMENT_IMAGE=nginx:1.25.4-alpine 

# Stage: BUILD
# Install dependencies first to maximize Docker layer caching.
FROM node:22 AS build
ARG WORKINGPATH
ARG CONFIG
ARG BASEHREF
WORKDIR ${WORKINGPATH}

# Install packages.
COPY package.json package-lock.json ./
RUN npm ci && npm cache clean --force

# Install Chrome for running UI tests.
# RUN apt-get install -y gnupg wget curl unzip --no-install-recommends; \
#     wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | \
#     gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/google.gpg --import; \
#     chmod 644 /etc/apt/trusted.gpg.d/google.gpg; \
#     echo "deb https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list; \
#     apt-get update -y; \
#     apt-get install -y google-chrome-stable;

# Build from source code.
COPY . .
RUN npm run build

# Stage: DEPLOY
FROM $DEPLOYMENT_IMAGE
ARG ENTRYPATH
ARG PORT
WORKDIR ${ENTRYPATH}

COPY nginx/default.conf /etc/nginx/conf.d/
COPY --from=build ${ENTRYPATH}/build /usr/share/nginx/html/chatapp
EXPOSE ${PORT}
CMD ["nginx", "-g", "daemon off;"]