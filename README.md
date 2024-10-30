# Gerador de Perfis de Usuário

Bem-vindo ao projeto **Gerador de Perfis de Usuário**! Este é um aplicativo web desenvolvido em Go (Golang) para consolidar o conhecimento sobre templates HTML. O projeto permite a criação e listagem de perfis de usuário em uma interface simples.

## 🚀 Tecnologias

- **Go (Golang)**: Linguagem de programação usada para o back-end.
- **HTML e Templates Go**: Utilizados para renderizar as páginas de forma dinâmica.
- **FuncMap**: Função customizada `ToUppercase` para transformar textos no template em maiúsculas.

## 📋 Funcionalidades

- **Listagem de Perfis**: Exibe todos os perfis de usuário já cadastrados.
- **Cadastro de Novo Perfil**: Permite ao usuário inserir um novo perfil com nome, idade e biografia.
- **Conversão de Texto para Maiúsculas**: Exibe a bio do usuário em maiúsculas na listagem.

## 🧩 Estrutura do Projeto

- **main.go**: Arquivo principal contendo a lógica da aplicação.
- **main_test.go**: Arquivo contendo os testes da aplicação.
- **templates/**: Pasta que contém os arquivos HTML para as páginas de listagem (`index.html`) e formulário (`form.html`).

## 📝 Pré-requisitos

Certifique-se de ter o Go instalado. Para instalar o Go, acesse: [Golang Downloads](https://golang.org/dl/).

## 📂 Estrutura de Arquivos

```plaintext
├── templates
│   ├── index.html
│   └── form.html
├── .dockerignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── main_test.go
├── main.go
└── README.md
