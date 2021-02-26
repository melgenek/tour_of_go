docker build -t gold-rush:0.6 .
docker run --rm gold-rush:0.6


docker tag gold-rush:0.5 stor.highloadcup.ru/rally/hypnotic_coral
docker push stor.highloadcup.ru/rally/hypnotic_coral
