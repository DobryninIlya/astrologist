// Функция для отображения модального окна с сообщением об ошибке
function showModal(message) {
    var modal = document.getElementById('modal');
    var modalMessage = document.getElementById('modal-message');
    modalMessage.textContent = message;
    modal.style.display = 'block';
}

// Функция для скрытия модального окна
function closeModal() {
    var modal = document.getElementById('modal');
    modal.style.display = 'none';
}

// Функция для валидации данных
function validateData() {
    var name = document.getElementById('name').value.trim();
    var day = document.getElementById('fd').value;
    var month = document.getElementById('month').value;
    var year = document.getElementById('year').value;
    var hour = document.getElementById('hour').value;
    var minute = document.getElementById('minute').value;

    if (name === '' || day === '' || month === '' || year === '' || hour === '' || minute === '' ) {
        showModal('Пожалуйста, заполните все поля формы.');
        return false;
    }
    return true;
}

var closeButton = document.querySelector('.close');

// Добавляем обработчик события для закрытия модального окна при клике на крестик
closeButton.addEventListener('click', function() {
    closeModal();
});


document.getElementById('natal_card_button').addEventListener('click', function(event) {
    if (!validateData()) {
        event.preventDefault();
        return
    }
    // Собираем данные из полей формы
    var name = encodeURIComponent(document.getElementById('name').value);
    var fd = encodeURIComponent(document.getElementById('fd').value);
    var fm = encodeURIComponent(document.getElementById('month').value);
    var fy = encodeURIComponent(document.getElementById('year').value);
    var c1 = encodeURIComponent(document.getElementById('birth_place').value);
    var timezone = encodeURIComponent(document.getElementById('timezone').value);
    var hour = encodeURIComponent(document.getElementById('hour').value);
    var minute = encodeURIComponent(document.getElementById('minute').value);
    var latitude = encodeURIComponent(document.getElementById('latitude').value);
    var longitude = encodeURIComponent(document.getElementById('longitude').value);
    var hs = encodeURIComponent(document.getElementById('hs').value);
    var city = encodeURIComponent(document.getElementById('birth_place').value);

    // Формируем строку URL-параметров
    var queryParams = `?fn=${name}&fd=${fd}&fm=${fm}&fy=${fy}&c1=${c1}&ttz=${timezone}&fh=${hour}&fmn=${minute}&lt=${latitude}&ln=${longitude}&hs=${hs}&as=1&sb=1&с1=${city}`;

    // Формируем полный URL для перехода
    var url = '/natalchart/' + queryParams;

    // Выполняем переход на новую страницу
    window.location.href = url;
});