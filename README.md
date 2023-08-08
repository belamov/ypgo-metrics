[![Coverage Status](https://coveralls.io/repos/github/belamov/ypgo-metrics/badge.svg?branch=main)](https://coveralls.io/github/belamov/ypgo-metrics?branch=main)

## Практические задания на курсе [продвинутый go-разработчик](https://practicum.yandex.ru/go-advanced/)

практика была оформлена в виде инкрементов, каждый инкремент открывался последовательно

## Список инкрементов

<details>
  <summary>Инкремент 1</summary>

Разработайте сервер для сбора рантайм-метрик, который будет собирать репорты от агентов по протоколу HTTP. 

Агент вам предстоит реализовать в следующем инкременте — в качестве источника метрик вы будете использовать пакет `runtime`.


Сервер должен быть доступен по адресу `http://localhost:8080`, а также:
* Принимать и хранить произвольные метрики двух типов:
* Тип `gauge`, `float64` — новое значение должно замещать предыдущее.
* Тип `counter`, `int64` — новое значение должно добавляться к предыдущему, если какое-то значение уже было известно серверу.
* Принимать метрики по протоколу HTTP методом `POST`.
* Принимать данные в формате `http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>, Content-Type: text/plain`.
* При успешном приёме возвращать `http.StatusOK`.
* При попытке передать запрос без имени метрики возвращать `http.StatusNotFound`.
* При попытке передать запрос с некорректным типом метрики или значением возвращать `http.StatusBadRequest`.

Редиректы не поддерживаются.

Для хранения метрик объявите тип `MemStorage`. 
Рекомендуем использовать тип `struct` с полем-коллекцией внутри (`slice` или `map`). 
В будущем это позволит добавлять к объекту хранилища новые поля, например логер или мьютекс, чтобы можно было использовать их в методах. 
Опишите интерфейс для взаимодействия с этим хранилищем.

</details>

<details>
  <summary>Инкремент 2</summary>

Разработайте агент (HTTP-клиент) для сбора рантайм-метрик и их последующей отправки на сервер по протоколу HTTP.

Агент должен собирать метрики двух типов:
- Тип `gauge`, `float64`.
- Тип `counter`, `int64`.

В качестве источника метрик используйте пакет `runtime`.

Нужно собирать следующие метрики типа `gauge`:
- `Alloc`
- `BuckHashSys`
- `Frees`
- `GCCPUFraction`
- `GCSys`
- `HeapAlloc`
- `HeapIdle`
- `HeapInuse`
- `HeapObjects`
- `HeapReleased`
- `HeapSys`
- `LastGC`
- `Lookups`
- `MCacheInuse`
- `MCacheSys`
- `MSpanInuse`
- `MSpanSys`
- `Mallocs`
- `NextGC`
- `NumForcedGC`
- `NumGC`
- `OtherSys`
- `PauseTotalNs`
- `StackInuse`
- `StackSys`
- `Sys`
- `TotalAlloc`

К метрикам пакета `runtime` добавьте ещё две:
- `PollCount` (тип `counter`) — счётчик, увеличивающийся на 1 при каждом 
    обновлении метрики из пакета `runtime` (на каждый `pollInterval` — см. ниже).
- `RandomValue` (тип `gauge`) — обновляемое произвольное значение.

По умолчанию приложение должно:

- Обновлять метрики из пакета `runtime` с заданной частотой: `pollInterval` — 2 секунды.
- Отправлять метрики на сервер с заданной частотой: `reportInterval` — 10 секунд.

Чтобы приостанавливать работу функции на заданное время, используйте вызов `time.Sleep(n * time.Second)`. 
Подробнее о пакете time и его возможностях вы узнаете в третьем спринте.

Метрики нужно отправлять по протоколу `HTTP` методом `POST`:
* Формат данных — `http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>`.
* Адрес сервера — *.
* Заголовок — `Content-Type: text/plain`.

Покройте код агента и сервера юнит-тестами.
</details>
<details>
  <summary>Инкремент 3</summary>

Вы написали приложение с помощью пакета стандартной библиотеки `net/http`. 
Используя любой внешний пакет (роутер или фреймворк), совместимый с `net/http`, перепишите ваш код.

Доработайте сервер так, чтобы в ответ на запрос `GET http://<АДРЕС_СЕРВЕРА>/value/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>` он 
возвращал текущее значение метрики в текстовом виде со статусом `http.StatusOK`.

При попытке запроса неизвестной метрики сервер должен возвращать `http.StatusNotFound`.

По запросу `GET http://<АДРЕС_СЕРВЕРА>/` сервер должен отдавать HTML-страницу со списком имён и значений всех 
известных ему на текущий момент метрик.

Хендлеры должны взаимодействовать с экземпляром `MemStorage` при помощи соответствующих интерфейсных методов.
</details>
<details>
  <summary>Инкремент 4</summary>
Доработайте код, чтобы он умел принимать аргументы с использованием флагов.

Аргументы сервера:
* Флаг `-a=<ЗНАЧЕНИЕ>` отвечает за адрес эндпоинта HTTP-сервера (по умолчанию `localhost:8080`).

Аргументы агента:
* Флаг `-a=<ЗНАЧЕНИЕ>` отвечает за адрес эндпоинта HTTP-сервера (по умолчанию `localhost:8080`).
* Флаг `-r=<ЗНАЧЕНИЕ>` позволяет переопределять `reportInterval` — частоту отправки метрик на сервер (по умолчанию 10 секунд).
* Флаг `-p=<ЗНАЧЕНИЕ>` позволяет переопределять `pollInterval` — частоту опроса метрик из пакета runtime (по умолчанию 2 секунды).

При попытке передать приложению незвестные флаги оно должно завершаться с сообщением о соответствующей ошибке.

Значения интервалов времени должны задаваться в секундах.

Во всех случаях должны присутствовать значения по умолчанию.
</details>