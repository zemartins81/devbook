$(document).ready(function () {

    $('#cadastro').submit(function (event) {
        event.preventDefault()
        criarUsuario()
    })
})

function criarUsuario() {
    if ($('#senha').val() != $('#confirma-senha').val()) {
        alert('As senhas não conferem')
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            senha: $('#senha').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){
        alert("Usuario cadastrado com sucesso!");
    }).fail(function(erro){
        console.log(erro)
        alert("Erro ao cadastrar o usuário!")
    });
}
