$('#nova-publicacao').submit(function (event) {
    event.preventDefault()
    criarPublicacao()
})

function criarPublicacao() {

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function () {
        window.location = "/home"
    }).fail(function (erro) {
        console.log(erro)
        alert("Erro ao criar a publicação!")
    });
}

$('.curtir-publicacao').click(function (event) {
    event.preventDefault()
    var publicacaoId = $(this).parent().data('publicacao-id')
    curtirPublicacao(publicacaoId)
})

function curtirPublicacao(publicacaoId) {
    console.log(publicacaoId)
    $.ajax({
        url: "/publicacoes/" + publicacaoId + "/curtir",
        method: "POST"
    }).done(function () {
        window.location = "/home"
    }).fail(function (erro) {
        console.log(erro)
        alert("Erro ao curtir a publicação!")
    });
}

