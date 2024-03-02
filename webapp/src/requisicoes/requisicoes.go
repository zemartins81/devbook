package requisicoes

import (
	"io"
	"net/http"

	"github.com/zemartins81/devbookWebApp/src/cookies"
)

func RequisicoesComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(metodo, url, dados)
	if err != nil {
		return nil, err
	}

	cookie, err := cookies.Ler(r)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
