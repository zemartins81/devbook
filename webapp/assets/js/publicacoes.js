$('#nova-publicacao').on('submit', criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

$('#atualizar-publicacao').on('submit', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(event) {
    event.preventDefault();

    const $titulo = $('#titulo');
    const $conteudo = $('#conteudo');

    if ($titulo.length === 0 || $conteudo.length === 0) {
        console.error('Elementos de formulário ausentes');
        return;
    }

    const titulo = $titulo.val();
    const conteudo = $conteudo.val();

    if (!titulo || !conteudo) {
        console.error('Campos de formulário vazios');
        return;
    }

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo,
            conteudo,
        }
    }).done(function() {
        Swal.fire({
            'title': 'Sucesso!',
            'text': 'Publicação criada com sucesso!',
            'icon': 'success',
            'timer': 2000,
        }).then((result) => {
            if (result.isConfirmed) {
                window.location = "/home";
            }
        })
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Erro ao criar publicação!',
            'icon': 'error',
            'timer': 2000,
        })
    });
}

function curtirPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Erro ao curtir publicação!',
            'icon': 'error',
            'timer': 2000,
        })
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function descurtirPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Erro ao descurtir publicação!',
            'icon': 'error',
            'timer': 2000,
        })
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function atualizarPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    elementoClicado.prop('disabled', true);

    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');
    const publicacaoTitulo = $('#titulo').val();
    const publicacaoConteudo = $('#conteudo').val();

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: publicacaoTitulo,
            conteudo: publicacaoConteudo
        }
    }).done(function() {
        Swal.fire({
            'title': 'Sucesso!',
            'text': 'Publicação atualizada com sucesso!',
            'icon': 'success',
            'confirmButtonText': 'Ok',
            'confirmButtonColor': '#0d6efd'
        }).then((result) => {
            window.location = "/home";
        })
    }).fail(function() {
        Swal.fire({
            'title': 'Erro!',
            'text': 'Erro ao atualizar publicação',
            'icon': 'error',
            'confirmButtonText': 'Ok',
            'confirmButtonColor': '#dc3545'
        })
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function deletarPublicacao(event) {
    event.preventDefault();

    Swal.fire({
        'title': 'Atenção!',
        'text': 'Tem certeza que deseja deletar esta publicação?',
        'icon': 'warning',
        'showCancelButton': true,
        'confirmButtonColor': '#0d6efd',
        'cancelButtonColor': '#dc3545',
        'confirmButtonText': 'Sim',
        'cancelButtonText': 'Cancelar'
    }).then((result) => {
        if (!result.isConfirmed) return;

        const elementoClicado = $(event.target);
        const publicacao = elementoClicado.closest('div');
        const publicacaoId = publicacao.data('publicacao-id');

        elementoClicado.prop('disabled', true);

        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function() {
            publicacao.fadeOut("slow", function() {
                $(this).remove();
            });
        }).fail(function() {
            Swal.fire({
                'title': 'Erro!',
                'text': 'Erro ao deletar publicação',
                'icon': 'error',
                'confirmButtonText': 'Ok',
                'confirmButtonColor': '#dc3545'
            })
        }).always(function() {
            elementoClicado.prop('disabled', false);
        });
    })
}
