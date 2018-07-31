#!/bin/bash

az login 
az extension add --name webapp
az group create -n MultiContainerExample -l centralus
az appservice plan create -g MultiContainerExample -n bjdlinuxasp --is-linux --sku S1 -l centralus
az webapp create -g MultiContainerExample -n bjdweb003 -p bjdlinuxasp --multicontainer-config-file ./deployment.yml --multicontainer-config-type KUBE

