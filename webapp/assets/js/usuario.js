$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#deletar-usuario').on('click', deletarUsuario)

function pararDeSeguir(event) {
    event.preventDefault();
    const usuarioId = $(this).data('usuario-id')

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST",
    }).done(function() {
        $(this).fadeOut("slow", function() {
            window.location = `/usuarios/${usuarioId}`;
        });
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Ops..., Erro ao parar de seguir o usuário',
            'icon': 'error',
            'timer': 2000,
        });
    });
}

function seguir(event) {
    event.preventDefault();
    const usuarioId = $(this).data('usuario-id')

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST",
    }).done(function() {
        $(this).fadeOut("slow", function() {
            window.location = `/usuarios/${usuarioId}`;
        });
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Ops..., Erro ao seguir o usuário',
            'icon': 'error',
            'timer': 2000,
        });
    });

}

function deletarUsuario() {
    Swal.fire({
        'title': 'Atenção!',
        'text': 'Tem certeza que deseja deletar esse usuário?',
        'icon': 'warning',
        'showCancelButton': true,
        'cancelButtonText': 'Cancelar',
    }).then(function(confirmado) {
        if (!confirmado.isConfirmed) return;
        $.ajax({
            url: `/deletar-usuario`,
            method: "DELETE"
        }).done(function() {
            Swal.fire({
                'title': 'Sucesso!',
                'text': 'Usuario deletado com sucesso!',
                'icon': 'success',
                'timer': 2000,
            }).then(() => {
                window.location = "/logout";
            })

        }).fail(function() {
            Swal.fire({
                'title': 'Erro!',
                'text': 'Ops... Erro ao deletar o usuário',
                'icon': 'error',
                'timer': 2000,
            });
        });
    })
}


