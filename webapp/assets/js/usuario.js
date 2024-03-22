$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)

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
