# GoCloudCamp

## Описание задания

Требуется разработать модуль для обеспечения работы с плейлистом. Модуль должен обладать следующими возможностями:

Play - начинает воспроизведение
Pause - приостанавливает воспроизведение
AddSong - добавляет в конец плейлиста песню
Next воспроизвести след песню
Prev воспроизвести предыдущую песню

Реализовать сервис, который позволит управлять музыкальным плейлистом. Доступ к сервису должен осуществляться с помощью API.

## Реализация проекта

1. Реализовал сервис при помощи gRPC и gRPC-gateway
2. Для сборки использовал Docker
3. Тестовый пример использования: request.http

## Функционал

1. AddSong - Добавляет песню в плейлист
2. PlaySong - Начинает воспроизведение
3. NextSong - Перелистывает песню
4. PrevSong - Возвращается на предыдущую песню
5. GetSong - Возвращает значение текущей песни
6. UpdateSong - Меняет название песни
7. DeleteSong - Удаляет песню из плейлиста

## Запуск проекта

<pre><code>make run</code></pre>

## Установка пакетов

<pre><code>make all</code></pre>

## Генерация proto

<pre><code>make build</code></pre>
