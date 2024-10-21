# arquitetura-computadores-GA
Repo para o trabalho do grau A da cadeira de arquitetura de computadores da Unisinos

# Como rodar o projeto
Para rodar este projeto só necessário rodar o comando ```go run pkg/main.go```, 
isso irá executar o projeto sem a predição de branch, para usar a predição é necessário
habilitar setando uma váriavel de ambiente, e pode ser feito com o seguinte comando
```PREDICTION="" go run pkg/main.go```, com isso a predição de branch será habilitada,
para mudar o arquivo de teste é necessário alterar a string dentro do arquivo `pkg/main.go`
na chamada para a função `os.open()` com o nome do arquivo desejado.