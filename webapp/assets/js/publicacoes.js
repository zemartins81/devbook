$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao)

function criarPublicacao(event) {
    event.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {
        alert("Erro ao criar publicação", "error")
    })
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
