# EVAT - EMC

# Document

Database Schema* (https://dbdiagram.io)
* evat.dbml 

Example Flow (https://plantuml.com/)
* flow-payment.puml
* flow-register-team.puml

## RUN

<b>Prerequisites</b>
 * docker 
 * docker compose
 * goland 1.18

extract  evat-emc.gz
copy voloums to evat-emc-back-end
copy resources/key to evat-emc-back-end/resources

Environment Production

 * docker-compose-evat-emc-develop.yaml

```
make develop
```

Environment Develop

 * docker-compose-evat-emc-production.yaml

```
make production
```

