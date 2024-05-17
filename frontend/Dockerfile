FROM node:16.13.2

WORKDIR /usr/src/app

ENV PATH /usr/src/app/node_modules/.bin:$PATH

RUN npm config set unsafe-perm true
RUN npm install -g serve@10.1.1


COPY package.json ./
RUN npm install --legacy-peer-deps

COPY . .

# RUN chown -Rh $user:$user /usr/src/app
USER $user
# USER node


RUN npm run build

EXPOSE 3000

CMD ["serve", "-s", "-l", "3000", "build"]