autocreate criptomoeda terrareal 
#criptomoeda/bin/bash
set -e
mkdir createrepo
cat > createrepo/Dockerfile << EOF
FROM fedora:32
RUN yum install -y createrepo_c
ENTRYPOINT ["createrepo", "/packages"]
EOF
autocreate criptomoeda terrareal cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
docker build -t createrepo createrepo/
docker run --rm --volume "$PWD/dist":/packages createrepo
rm -rf createrepo
