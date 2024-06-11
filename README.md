# tgbot_buy
Telegram-бот для покупок привилегий.

# Проект находится на стадии разработки.

Этот Telegram бот позволяет пользователям покупать привилегии на игровых серверах Counter-Strike 1.6.

Бот работает с базой данных MySQL, для работы бота необходим CS-Bans (https://dev-cs.ru/resources/156/) с помошью таблиц из CS-Bans бот будет добавлять админов на ваши сервера. Также, бот в базе данных сохраняет контекст пользователей, которые взаимодействовали с ботом.

Бот умеет оправлять RCON комманды на сервер, данные он берет с таблиц CS-Bans.

# Требования.
- ЮKassa (https://yookassa.ru/)
- CS-Bans (https://dev-cs.ru/resources/156/)
- Любой adminloader, который понимает команду amx_reloadadmins.

# Установка и настройка.

## lib

- go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
- go get -u github.com/go-sql-driver/mysql
- go get -u github.com/sirupsen/logrus
- go get github.com/joho/godotenv