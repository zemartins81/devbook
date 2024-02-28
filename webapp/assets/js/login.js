$(document).ready(function () {

    $('#login').submit(function (event) {
        event.preventDefault()
        fazerLogin()
    })
})

function fazerLogin() {

    $.ajax({
        url: "/login",
        method: "POST",
        data:  {
            email: $('email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        window.location.href = "/home"
    }).fail(function(erro){
        alert("Usuario ou senha invalidos!")
    })
}

