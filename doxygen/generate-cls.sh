#!/bin/bash

set -e

PROJECT_DIR_PATH=""

function usage_exit() {
    echo "Usage: $0 [-d]"
    echo "[-d]:set your root project dir"
    exit 1
}

while getopts d:h OPT
do
    case $OPT in
        d) PROJECT_DIR_PATH=$OPTARG
           ;;
        h) usage_exit
           ;;
       \?) usage_exit
           ;;
    esac
done

shift $((OPTIND - 1))

if [[ ! $( which dot && which doxygen ) ]]; then
        echo "Invalid"
        return 1
fi

if [[ -z ${PROJECT_DIR_PATH} ]]; then
    echo [ERROR] "You Must Set Your Project Dir"
    exit 1
fi


cd ${PROJECT_DIR_PATH}

#generate Doxyfile
if [[ -f ./DoxyFile ]]; then
    rm Doxyfile
    #create_custom_doxy_template
fi

doxygen -g
doxygen && echo "" &&  echo "if you are gen cls please flow at this tempalte of your DoxyFile"
           echo "https://www.chihayafuru.jp/tech/index.php/archives/1974"
