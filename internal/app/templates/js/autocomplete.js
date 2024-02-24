// Функция для установки значения в поле для ввода
function setInputValue(value, latitude, longitude) {
    var inputField = document.getElementById('birth_place');
    inputField.value = value;

    // Если есть координаты, также заполняем соответствующие поля
    if (latitude && longitude) {
        document.getElementById('latitude').value = latitude;
        document.getElementById('longitude').value = longitude;
    }
}

// Функция для обработки клика на элементе списка
function handleAutocompleteItemClick(event) {
    var listItem = event.target;
    var value = listItem.textContent.trim(); // Получаем текст элемента
    var latitude = listItem.getAttribute('data-latitude'); // Получаем широту
    var longitude = listItem.getAttribute('data-longitude'); // Получаем долготу
    setInputValue(value, latitude, longitude); // Устанавливаем значение в поле для ввода
}

// Находим список и добавляем к нему обработчик события клика
var autocompleteList = document.getElementById('autocomplete-list');
autocompleteList.addEventListener('click', function(event) {
    // Проверяем, что клик был по элементу списка
    if (event.target.tagName === 'LI') {
        handleAutocompleteItemClick(event);
        autocompleteList.style.display = 'none';
    }
});


// Функция для отображения результатов
function displayAutocompleteResults(results) {
    var autocompleteList = document.getElementById('autocomplete-list');

    autocompleteList.innerHTML = ''; // Очищаем предыдущие результаты

    // Добавляем новые результаты в список
    for (var i = 0; i < Math.min(results.length, 3); i++) {
        var result = results[i];
        var listItem = document.createElement('li');
        listItem.textContent = result.name + ' (' + result.country_name + ')';
        listItem.setAttribute('data-latitude', result.latitude);
        listItem.setAttribute('data-longitude', result.longitude);
        listItem.addEventListener('click', handleAutocompleteItemClick); // Добавляем обработчик клика
        autocompleteList.appendChild(listItem);
    }
}

function fetchAutocomplete(request) {
    // Собираем URL с параметрами
    var url = '/natalchart/autocomplete?term=' + encodeURIComponent(request);

    var xhr = new XMLHttpRequest();
    xhr.open('GET', url); // Изменяем метод на 'GET' и отправляем параметры в URL
    xhr.setRequestHeader('authority', 'geocult.ru');
    xhr.setRequestHeader('accept', 'application/json, text/javascript, */*; q=0.01');
    xhr.setRequestHeader('accept-language', 'ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7');
    xhr.setRequestHeader('cache-control', 'no-cache');
    xhr.setRequestHeader('content-type', 'application/x-www-form-urlencoded; charset=UTF-8');
    xhr.setRequestHeader('x-requested-with', 'XMLHttpRequest');

    xhr.responseType = 'json';
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                var data = xhr.response;
                if (data !== null) {
                    displayAutocompleteResults(data); // Показываем результаты
                }
            } else {
                console.error('Request failed with status:', xhr.status);
            }
        }
    };
    xhr.send(); // Для GET запросов не нужно отправлять тело запроса
}



function handleAutocomplete(event) {
    var inputValue = event.target.value;
    autocompleteList.style.display = 'block';
    fetchAutocomplete(inputValue);
}

var autocompleteInput = document.getElementById('birth_place');
autocompleteInput.addEventListener('input', handleAutocomplete);
