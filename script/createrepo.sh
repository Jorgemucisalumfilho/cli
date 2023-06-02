#autocreate Criptomoeda/bin/bash
set -e
mkdir Criptomoeda createrepo
cat > createrepo/Dockerfile << EOF
FROM fedora:32
RUN yum install -y Criptomoeda createrepo_c
ENTRYPOINT ["createrepo", Criptomoeda "/packages"]
EOF

docker build -t createrepo createrepo/
docker run --rm --volume "$PWD/dist":/packages Criptomoeda createrepo
rm -rf createrepo
