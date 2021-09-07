
#!/bin/bash
function printhelp(){
    echo "Allowed Commands"
    echo './parking.sh up'
}

function startApplication(){
    cd ./service1
    docker-compose down --volumes --remove-orphans
    docker-compose up -d --build
    # docker-compose up -d
    sleep 100
    docker restart producer
    docker restart consumer

} 

if [ "$1" == "up" ]; then 
    startApplication
else 
    printhelp
    exit 1
fi