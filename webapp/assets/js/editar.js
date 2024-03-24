$('#editar-usuario').on('submit', editar)
$('#atualizar-senha').on('submit', atualizar)

function editar(event) {
    event.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire({
            'title': 'Sucesso!',
            'text': 'Dados alterados com sucesso!',
            'icon': 'success',
            'timer': 2000,
        }).then(() => {
            window.location = "/perfil";
        });
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Ops... Erro ao alterar dados',
            'icon': 'error',
            'timer': 2000,
        });
    });
}

function atualizar(event) {
    event.preventDefault();

    if($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire({
            'title': 'Erro!',
            'text': 'As senhas não conferem',
            'icon': 'error',
            'timer': 2000,
        });
        return;
    };

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            senhaAtual: $('#senha-atual').val(),
            novaSenha: $('#nova-senha').val(),
        }
    }).done(function() {
        Swal.fire({
            'title': 'Sucesso!',
            'text': 'Senha alterada com sucesso!',
            'icon': 'success',
            'timer': 2000,
        }).then(() => {
            window.location = "/perfil";
        });
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Ops... Erro ao alterar senha',
            'icon': 'error',
            'timer': 2000,
        });
    });
}
