#!/bin/bash


template="template-service"


# Проверка аргументов
if [ "$#" -ne 1 ]; then
    echo "Использование: $0 <new_service>"
    exit 1
fi

SERVICE_NAME=$1
SERVICE_PATH="services/${SERVICE_NAME}"

OLD_MODULE=$(cat ./${template}/go.mod | head -1 |  awk '{print $2}')
NEW_MODULE="${OLD_MODULE}/${SERVICE_NAME}"

echo "old module: ${OLD_MODULE}"
echo "new module: ${NEW_MODULE}"

rm -rf ${SERVICE_PATH}
cp -R ${template} ${SERVICE_PATH}

rm -rf "${SERVICE_PATH}/bin"
rm -rf "${SERVICE_PATH}/dist"

echo "${SERVICE_PATH}/bin/" >> .gitignore
echo "${SERVICE_PATH}/dist/" >> .gitignore

find ${SERVICE_PATH} -type f \
  ! -path "./.git/*" \
  ! -path "./vendor/*" \
  \( \
    -name "*.go" \
    -o -name "go.mod" \
    -o -name "*.yml" \
    -o -name "*.yaml" \
  \) \
  -print0 | while IFS= read -r -d '' file; do
    echo $file
    sed -i '' "s|$OLD_MODULE|$NEW_MODULE|g" "$file"
  done

sed -i '' "s/project_name:.*/project_name: ${SERVICE_NAME}/g" "${SERVICE_PATH}/.goreleaser.yaml"
sed -i '' "s/appName.*/appName = \"${SERVICE_NAME}\"/g" "${SERVICE_PATH}/internal/config/build/vars.go"


echo "Замена завершена."
