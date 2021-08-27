# Metaheuristica

## Instalação
1. Baixar o arquivo no [site do golang](https://golang.org/doc/install#download)
2. Extrair o conteúdo
3. Utilizar os seguintes comandos

```
sudo cp -R .../go/ /usr/local/ 
```
sendo `.../go/` o caminho extraído

4. Criar links simbólicos

```
ln -s /usr/local/go/bin/go /usr/local/bin/go
ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
``` 

## Execução do projeto

Para executar o código usar no terminal

```
$ make all
```

### Gerar um executável

```
$ make build
```

Executar o programa:

```
$ ./Metaheuristica
```
