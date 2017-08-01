FROM ubuntu
COPY scripts/ /scripts/
WORKDIR /scripts/

ENTRYPOINT ./bash-printfirstname.sh name.txt; ./golang-printfirstname name.txt;
