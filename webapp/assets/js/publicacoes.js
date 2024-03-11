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
