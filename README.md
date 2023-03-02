# collector-data


Пример запроса сохранение данные об устройствах
```
curl --request POST \
--url http://localhost:8080/api/v1/ \
--header 'Content-Type: application/json' \
--data '{
"id": 1,
"nameOS": "macos",
"versionOS": "122",
"nameBrowser": "safari",
"versionBrowser": "10",
"IP": "122.312.21.31",
"brandPhone": "apple",
"modelPhone": "14 pro max",
"screenResolution": "1800 x 2001"
}'
```
Пример запроса предоставлять данные об устройствах по ID пользователя
```
curl --request GET \
  --url http://localhost:8080/api/v1/9
```
Пример запроса ТОП 100 наиболее популярных ОС
```
curl --request GET \
  --url http://localhost:8080/api/v1/top-name-os
```
Пример запроса ТОП 100 наиболее популярных браузеров
```
curl --request GET \
  --url http://localhost:8080/api/v1/top-name-browser
```
Пример запроса ТОП 100 наиболее популярных телефонов
```
curl --request GET \
  --url http://localhost:8080/api/v1/top-brand-phone
```
Пример запроса ТОП 100 наиболее разрешений экрана
```
curl --request GET \
  --url http://localhost:8080/api/v1/top-screen-resolution
```
Пример запроса ТОП 100 популярных версий ОС по названию ОС
```
curl --request GET \
  --url 'http://localhost:8080/api/v1/top-version-os?by=name_os&name=ubuntu'
```
Пример запроса ТОП 100 популярных версий браузера по названию браузера
```
curl --request GET \
  --url 'http://localhost:8080/api/v1/top-version-browser?by=name_browser&name=safari'
```
Пример запроса ТОП 100 популярных моделей телефона по бренду телефона
```
curl --request GET \
  --url 'http://localhost:8080/api/v1/top-model-phone?by=brand_phone&name=samsung'
```