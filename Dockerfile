# Build the app with "npm run build" before running docker
FROM node
RUN npm i -g serve
ADD ./dist ./app
EXPOSE 5000
CMD serve ./app