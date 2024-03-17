$('#nova-publicacao').onSubmit(criarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);

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
        window.location = "/home";
    }).fail(function(jqXHR, textStatus, errorThrown) {
        console.error('Erro ao criar publicação:', textStatus, errorThrown);
        alert("Erro ao criar publicação", "error");
    });
}

function curtirPublicacao(event) {
    event.preventDefault();

    const elementoClicado = $(event.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next('span');
            const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

            contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
    }).fail(function() {
        alert("Erro ao curtir a publicação");
    })
}
