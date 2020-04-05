FROM node:lts-alpine


FROM node:lts-alpine


# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH
COPY /web/ .
# install and cache app dependencies
COPY web/package.json /app/package.json
RUN npm install
RUN npm install @vue/cli@3.7.0 -g
RUN npm rebuild node-sass
# start app
CMD ["npm", "run", "serve"]