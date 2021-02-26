docker build -t gold-rush:0.7 .
docker run --rm gold-rush:0.7


docker tag gold-rush:0.7 stor.highloadcup.ru/rally/hypnotic_coral
docker push stor.highloadcup.ru/rally/hypnotic_coral
