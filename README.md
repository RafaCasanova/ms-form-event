
# App Microsserviço de formulario 

Microsserviço de formulario para destribuir email dos formulario submetido corretamente, de acordo com a nessecidade e possivel escalar ou até implementar novos consumidores dos formularios submetidos.


## Arquitetura

Foi utilizado a arquitetura hexagonal, pelos seguintes motivos: facilidade de manuteção, flexibilidade de alteração, adptabilidade, facilidade na estruturação da regra de negocio sem precisar mexer em varios arquivos.

Com facilidade de um ponto, temos dificuldades em outro, a dificuldades desse projeto e estruturar os arquivos e camadas de acordo com sua responsabilidade, sem se preocupar em ficar repetititvo de mais, pois para cada um arquivo, ele tem que ter sua responsabilidade, com isso o projeto fica com varias pastas e arquivos, mas, cada um fazendo o que deve ser feito!

## Stack utilizada


**Back-end:** Go 

**Mensageria :** Kafka


## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

```env
PORT = <PORT>
HCAPTCHA_SITE_KEY = <SITE_CHAVE>
HCAPTCHA_SECRET_KEY = <CHAVE>
MAIL_HOST = <MAIL_HOST>
MAIL_PORT = <MAIL_PORT>
MAIL_AUTH_USER = <MAIL_AUTH_USER>
MAIL_AUTH_PASS = <MAIL_AUTH_PASS>
TEXT_MAIL_TITLE_COMPANY = <EMAIL_TITLE_TO_COMPANY>
TEXT_MAIL_TITLE_USER = <EMAIL_TITLE_TO_USER>
KAFKA_BROKER = <KAFKA_BROKER>
KAFKA_EMAIL_TOPIC = <TOPIC>
MONGODB_URL = <DB_URL>
MONGODB_PORT = <DB_PORT>
MONGODB_DATA_BASE = <DB_DATA_BASE>
MONGODB_USER_DB = <MONGODB_USER_DB>
MONGODB_PASS_DB = <MONGODB_PASS_DB>
```


## Uso

Para usar tem que esta rodando kafka e o banco de dados mongodb, com isso precisa esta configurado as variaveis de ambiente, para facilitar o uso do kafka e do mongodb já deixei o docker-compose.yml no projeto

    
## Submit

#### Erro de captcha:

 * request

```JSON
{
    "TypeErro": "my wrong captcha",
    "Title": "captcha is incorrect",
    "Detail": [
        {"Error: error class"}, 
        {"path: host and path"}
        ]
}
```

#### Erro de parametro:

 * request

```JSON
{
    "TypeErro": "something wrong",
    "Title": "what is wrong",
    "Detail": [
        {"Param: what is wrong"}, 
        {"OtherParam: what is wrong"}
        ]
}
```

#### Erro :

 * request

```JSON
{
    "TypeErro": "something wrong",
    "Title": "what is wrong",
    "Detail": [
        {"Error: method error"}, 
        {"path: host and path"}
        ]
}
```


