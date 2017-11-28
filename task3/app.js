function addTask(task) {
    $("#list").append('<li><span>' + task + '</span>' +
        '<button onclick="this.parentNode.remove()">Удалить</button></li>');
}


$(document).ready(function () {
    $("#root").append('<ul id="list"></ul><input id="add_task_input">' +
        '<button id="add_task">Добавить задание</button>');
    addTask('Сделать задание #3 по web-программированию');

    $("#add_task").click(function () {
        var input = $("#add_task_input");
        addTask(input.val());
        input.val('');
    });
});