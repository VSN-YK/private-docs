#!/bin/bash

set -eu

# definate value 
select_target=""
validated_target=""
concourse_deploy_point=""
custom_trigger_job=""

# @Param: pipeline : exec pipeline whenever you want 
function trigger_job() {
  local pipeline=$1
  local job_name
  
  # getting first job name from pipeline.yml
  job_name=$(fly -t cc gp -p ${pipeline} | yaml2json | jq -r '.jobs[0].name')

  if [[ ! $(fly -t cc tj -j ${pipeline}/${job_name}) ]];then
    echo "Rejected! Becuase of ..."
    exit 1
  fi
}

function usage_exit() {
  echo "Usage: $0 [-c] [-t] [-j]"
  echo "[-c] : concourse deploy point"
  echo "[-t] : concourse target"
  echo "[-j] : custome execution trigger job"
  exit 1
}

while getopts t:d:j:h OPT
do
  case $OPT in
    t) select_target=$OPTARG
       ;;
    d) concourse_deploy_point=$OPTARG
       ;;
    j) custom_trigger_job=$OPTARG
       ;;
    h) usage_exit
       ;;
   \?) usage_exit
       ;;
  esac
done

shift $((OPTIND - 1))

if [[ ! $(which fly) ]]; then
  echo "fly binary isn\`t Installed"
  exit 1
fi

# list up concourse targets
targets=($(fly targets | awk '{print $1}'))

for tg in "${targets[@]}"
do
  if [[ ${tg} == ${select_target} ]]; then
    validated_target=${tg}
    break
  fi
done

# login check 
if [[ ! $(fly -t ${validated_target} status) ]]; then
  # getting login property from docker-compose.yml
  fly -t ${validated_target} login -c ${concourse_deploy_point} -u test -p test  
fi

# getting all pipelines
pipeline_list=($(fly -t cc pipelines --json | jq -r '.[].name'))  
for pipeline in "${pipeline_list[@]}"
do
  # call function of [trigger_job]
  trigger_job ${pipeline}
done


