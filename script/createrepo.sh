#criptomoeda/bin/bash
set -e
mkdir createrepo
cat > createrepo/Dockerfile << EOF
FROM fedora:32
RUN yum install -y createrepo_c
ENTRYPOINT "createrepo", "criptomoeda/packages"
EOF

docker build -t createrepo: createrepo/criptomoeda 
docker run --rm --volume "$PWD/dist":/packages createrepo
rm -rf createrepo
