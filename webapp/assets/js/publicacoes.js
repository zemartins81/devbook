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
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar publicação!");
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
        alert("Erro ao curtir a publicação");
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
        alert("Erro ao descurtir a publicação");
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
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao atualizar publicação");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}

function deletarPublicacao(event) {
    event.preventDefault();

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
        alert("Erro ao deletar publicação");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });
}
