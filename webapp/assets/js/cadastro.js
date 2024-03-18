$(document).ready(function() {

    $('#cadastro').submit(function(event) {
        event.preventDefault()
        criarUsuario()
    })
})

function criarUsuario() {
    if ($('#senha').val() != $('#confirma-senha').val()) {
        Swal.fire({
            'title': 'Erro!',
            'text': 'As senhas não conferem!',
            'icon': 'error',
            'timer': 2000,
        })
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        Swal.fire({
            'title': 'Sucesso!',
            'text': 'Usuario criado com sucesso!',
            'icon': 'success',
            'timer': 2000,
        }).then((result) => {
            if (result.isConfirmed) {
                window.location = "/login";
            }
        })
    }).fail(function(erro) {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Erro ao criar o usuário!',
            'icon': 'error',
            'timer': 2000,
        })
    });
}
