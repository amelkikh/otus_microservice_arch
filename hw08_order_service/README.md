## Домашнее задание №8 «Идемпотетность и коммутативность API в HTTP и очередях»

### Цель: Сделать [сервис "Заказ"](./hw06_stream_processing) из прошлого ДЗ идемпотетным

---

На выходе должно быть:

1) описание того, какой паттерн для реализации идемпотентности использовался
2) команда установки приложения (из helm-а или из манифестов). Обязательно указать в каком namespace нужно устанавливать
   и команду создания namespace, если это важно для сервиса.
3) тесты в postman

---

В тестах обязательно:

- наличие `{{baseUrl}}` для урла
- использование домена `arch.homework` в качестве initial значения `{{baseUrl}}`
- отображение данных запроса и данных ответа при запуске из командной строки с помощью newman
