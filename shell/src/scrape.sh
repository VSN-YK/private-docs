#!/bin/bash
set -ex

usage_exit() {
    echo "Usage: $0 [-f] [-s]"
    echo "[-s]:single player scrape mode"
    echo "[-f]:multiple player scrape mode you can set opt_file"
    exit 1
}

# variable 
FILE_PATH=""
PLAYER_NAME=""
player_lists=()
p_id=""
kd=""

while getopts f:s:h OPT
do
    case $OPT in 
        f) FILE_PATH=$OPTARG
           ;;
        s) PLAYER_NAME=$OPTARG
           ;;
        h) usage_exit
           ;;
       \?) usage_exit
           ;;
    esac
done

shift $((OPTIND - 1))

which jq  1>/dev/null
if [[ ! $? = 0 ]]; then
    echo "you must install jq before execute this script"
    echo "like this 「brew install jq」"
    exit 1
fi

if [ ! "${FILE_PATH}" = '' ]; then
    if [ -f "${FILE_PATH}" ]; then

        IFS=$'\n'
        player_lists=(`cat "${FILE_PATH}"`)
        echo "Result of Player Kd"
	rm avg.txt
        for each_player_info in "${player_lists[@]}" ; do
            # get player id
            p_id=$(curl -s "https://r6tab.com/api/search.php?platform=psn&search=${each_player_info}" | jq '.results[].p_id')
            # scrape by player kd
            kd=$(curl -s "https://r6tab.com/api/search.php?platform=psn&search=${each_player_info}"  | jq '.results[].kd')
	   
	    echo "scale=2;var=${kd};var/=100;var" | bc >> avg.txt
            kd=$(echo "scale=2;var=${kd};var/=100;var" | bc)
	 
            echo "${each_player_info}":"${kd}"
        done

	rst=$(awk '{sum+=$1}END{print sum/NR}' avg.txt)
	echo "CLAN_AVG=${rst}"

        

    fi    
else
    p_id=$(curl -s "https://r6tab.com/api/search.php?platform=psn&search=${PLAYER_NAME}" | jq '.results[].p_id')
    kd=$(curl -s "https://r6tab.com/api/search.php?platform=psn&search=${PLAYER_NAME}" | jq '.results[].kd')
    #echo "${p_id}"
    kd=$(echo "scale=2;var=${kd};var/=100;var" | bc)
    echo "${PLAYER_NAME}":"${kd}"
fi
