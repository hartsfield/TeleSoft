#!/bin/bash
git add .
git commit -m "hotfix" 
git push -u origin $1
gcloud compute ssh \
        --zone "us-central1-a" "instance-1" \
        --project "mysterygift" \
        --command "cd $2 && git pull && 
        sudo netstat -ltnp 
        | grep -w 'telesoft' 
        | awk '{ print $7 }' 
        | sed 's/\/.*//'" \
 
