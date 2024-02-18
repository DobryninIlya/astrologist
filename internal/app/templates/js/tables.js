document.querySelectorAll('font[color="#0000ff"]').forEach(function(font) {
    font.setAttribute('color', '');
    font.style = "font-family: 'StyreneAWeb-Regular', serif; font-size: 16px; ";
});
function showHiddenRows() {
    var rows = document.querySelectorAll('tr[style="display: none;"]');
    rows.forEach(function(row) {
        // row.style.display = 'table-row';
        row.style.display = '';
    });
}
function removeLastCell() {
    var table = document.getElementById('planetTable');
    var rows = table.getElementsByTagName('tr');

    for (var i = 0; i < rows.length; i++) {
        var cells = rows[i].getElementsByTagName('td');
        var lastCellIndex = cells.length - 1;
        if (lastCellIndex >= 0) {
            rows[i].deleteCell(lastCellIndex);
        }
    }
}


showHiddenRows();
removeLastCell();