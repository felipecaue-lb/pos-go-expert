# golang-migrate: Comando `migrate create`

Documentação detalhada sobre o comando utilizado para criar arquivos de migração de banco de dados.

## Comando

```bash
migrate create -ext=sql -dir=sql/migrations -seq init
```

## O que é o golang-migrate?

O [golang-migrate](https://github.com/golang-migrate/migrate) é uma ferramenta de migração de banco de dados escrita em Go. Ela permite versionar e gerenciar alterações no schema do banco de dados de forma incremental, segura e reproduzível.

A CLI (`migrate`) é o componente de linha de comando que permite executar migrações diretamente pelo terminal.

**Repositório da CLI:** https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

## Instalação

```bash
# Via Go
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Via Homebrew (macOS/Linux)
brew install golang-migrate

# Via Scoop (Windows)
scoop install migrate

# Via curl (Linux)
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$os-$arch.tar.gz | tar xvz
```

## Anatomia do Comando

```
migrate create -ext=sql -dir=sql/migrations -seq init
```

| Parte       | Descrição                                                        |
|-------------|------------------------------------------------------------------|
| `migrate`   | Binário da CLI do golang-migrate                                 |
| `create`    | Subcomando que cria um novo par de arquivos de migração          |
| `-ext=sql`  | Define a extensão dos arquivos gerados                           |
| `-dir=sql/migrations` | Define o diretório onde os arquivos serão criados      |
| `-seq`      | Usa numeração sequencial em vez de timestamp                     |
| `init`      | Nome/descrição da migração                                       |

## Detalhamento das Flags

### `-ext=sql`

Define a extensão dos arquivos de migração gerados.

- **Valor utilizado:** `sql`
- **Comportamento:** Gera arquivos com extensão `.sql`
- **Outros valores possíveis:** Qualquer extensão válida (ex: `.json`, `.cql` para Cassandra). O mais comum em projetos com bancos relacionais é `.sql`.

### `-dir=sql/migrations`

Especifica o diretório de destino onde os arquivos de migração serão criados.

- **Valor utilizado:** `sql/migrations`
- **Comportamento:** Cria os arquivos dentro do diretório `sql/migrations/` relativo ao diretório atual
- **Importante:** O diretório deve existir previamente. O comando **não** cria o diretório automaticamente.

### `-seq`

Controla o formato de numeração/versionamento dos arquivos.

- **Com `-seq`:** Usa numeração sequencial com zero-padding de 6 dígitos
  - Exemplo: `000001`, `000002`, `000003`
- **Sem `-seq`:** Usa timestamp Unix como versionamento
  - Exemplo: `1709913600`, `1709913700`

A numeração sequencial é mais legível e facilita a visualização da ordem das migrações.

### `init` (argumento posicional)

O nome descritivo da migração. Aparece no nome do arquivo após o número de versão.

- É o último argumento do comando (posicional, sem flag)
- Deve descrever brevemente o propósito da migração
- Exemplos comuns: `init`, `add_users_table`, `create_orders`, `add_index_to_email`

## Arquivos Gerados

O comando gera **dois arquivos** — um par de migração (up/down):

```
sql/migrations/
  000001_init.up.sql      # Migração "up" - aplica as alterações
  000001_init.down.sql    # Migração "down" - reverte as alterações
```

### Formato do nome

```
{versao}_{nome}.{direcao}.{extensao}
```

| Componente  | Exemplo    | Descrição                                              |
|-------------|------------|--------------------------------------------------------|
| `versao`    | `000001`   | Número sequencial (6 dígitos com zero-padding)         |
| `nome`      | `init`     | Nome descritivo passado como argumento                 |
| `direcao`   | `up`/`down`| Indica se aplica ou reverte a migração                 |
| `extensao`  | `sql`      | Extensão definida pela flag `-ext`                     |

### Arquivo `.up.sql`

Contém as instruções SQL para **aplicar** a migração (criar tabelas, adicionar colunas, etc).

```sql
-- Exemplo de conteúdo para 000001_init.up.sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Arquivo `.down.sql`

Contém as instruções SQL para **reverter** a migração (desfazer tudo que o `.up.sql` fez).

```sql
-- Exemplo de conteúdo para 000001_init.down.sql
DROP TABLE IF EXISTS users;
```

## Exemplo Prático Completo

```bash
# 1. Criar o diretório de migrações
mkdir -p sql/migrations

# 2. Criar a primeira migração
migrate create -ext=sql -dir=sql/migrations -seq init

# 3. Criar migrações subsequentes
migrate create -ext=sql -dir=sql/migrations -seq add_users_table
migrate create -ext=sql -dir=sql/migrations -seq add_orders_table
```

Resultado no diretório:

```
sql/migrations/
  000001_init.up.sql
  000001_init.down.sql
  000002_add_users_table.up.sql
  000002_add_users_table.down.sql
  000003_add_orders_table.up.sql
  000003_add_orders_table.down.sql
```

## Comandos Relacionados

Após criar e preencher os arquivos de migração, use os seguintes comandos para gerenciá-los:

```bash
# Aplicar todas as migrações pendentes
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

# Reverter a última migração
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down 1

# Reverter TODAS as migrações
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down

# Aplicar até uma versão específica
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" goto 2

# Ver a versão atual
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" version

# Forçar uma versão (útil para corrigir estado "dirty")
migrate -path=sql/migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" force 1
```

## Comparação: `-seq` vs Timestamp

| Aspecto             | `-seq` (sequencial)         | Sem `-seq` (timestamp)       |
|---------------------|-----------------------------|------------------------------|
| Formato             | `000001_init.up.sql`        | `1709913600_init.up.sql`     |
| Legibilidade        | Alta                        | Baixa                        |
| Conflito em equipe  | Maior risco de conflito     | Menor risco de conflito      |
| Ordenação           | Intuitiva                   | Cronológica por criação      |
| Uso recomendado     | Projetos solo ou equipes pequenas | Equipes grandes com branches paralelos |

---

# golang-migrate: Comando `migrate up`

Documentação detalhada sobre o comando utilizado para aplicar migrações de banco de dados.

## Comando

```bash
migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose up
```

## O que faz?

O subcomando `up` **aplica todas as migrações pendentes** em ordem sequencial (da menor versão para a maior). Cada arquivo `.up.sql` que ainda não foi executado será aplicado ao banco de dados.

## Anatomia do Comando

```
migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose up
```

| Parte                  | Descrição                                                      |
|------------------------|----------------------------------------------------------------|
| `migrate`              | Binário da CLI do golang-migrate                               |
| `-path=sql/migrations` | Caminho do diretório contendo os arquivos de migração          |
| `-database "..."`      | String de conexão com o banco de dados                         |
| `-verbose`             | Ativa logs detalhados de cada migração executada               |
| `up`                   | Subcomando que aplica as migrações pendentes                   |

## Detalhamento das Flags

### `-path=sql/migrations`

Especifica o diretório onde estão os arquivos de migração (`.up.sql` e `.down.sql`).

- **Valor utilizado:** `sql/migrations`
- **Comportamento:** O migrate lê todos os arquivos de migração deste diretório e determina quais ainda não foram aplicados.

### `-database`

String de conexão (DSN) que identifica o banco de dados de destino.

```
mysql://root:admin@tcp(mysql-db:3306)/goexpert
```

| Componente      | Valor         | Descrição                                              |
|-----------------|---------------|--------------------------------------------------------|
| `mysql://`      | —             | Scheme que identifica o driver MySQL                   |
| `root`          | usuário       | Usuário de conexão com o banco                         |
| `admin`         | senha         | Senha do usuário                                       |
| `tcp(mysql-db:3306)` | host:porta | Endereço do servidor MySQL via protocolo TCP      |
| `goexpert`      | database      | Nome do banco de dados                                 |

- **`mysql-db`** é o hostname do container MySQL (usado em ambientes Docker/Docker Compose).
- **`3306`** é a porta padrão do MySQL.

### `-verbose`

Ativa o modo verboso, exibindo logs detalhados durante a execução.

- Mostra qual migração está sendo aplicada no momento
- Exibe o tempo de execução de cada migração
- Útil para debugging e acompanhamento em ambientes de desenvolvimento

**Exemplo de saída:**
```
2024/03/08 10:00:00 Start buffering 1/u init
2024/03/08 10:00:00 Read and execute 1/u init
2024/03/08 10:00:00 Finished 1/u init (read 5.123ms, ran 12.456ms)
```

### `up` (subcomando)

Aplica migrações pendentes no banco de dados.

- **Sem argumento:** Aplica **todas** as migrações pendentes
- **Com argumento numérico (`up N`):** Aplica apenas as próximas **N** migrações pendentes

```bash
# Aplicar todas as migrações pendentes
migrate -path=sql/migrations -database "..." -verbose up

# Aplicar apenas as 2 próximas migrações
migrate -path=sql/migrations -database "..." -verbose up 2
```

## Como funciona internamente

1. O migrate lê os arquivos do diretório especificado em `-path`
2. Consulta a tabela `schema_migrations` no banco para verificar a versão atual
3. Identifica quais migrações ainda não foram aplicadas
4. Executa cada arquivo `.up.sql` pendente, em ordem crescente de versão
5. Atualiza a tabela `schema_migrations` com a nova versão após cada migração

### Tabela `schema_migrations`

O migrate cria automaticamente uma tabela de controle no banco:

| Coluna    | Tipo    | Descrição                                          |
|-----------|---------|----------------------------------------------------|
| `version` | bigint  | Número da versão da última migração aplicada       |
| `dirty`   | boolean | Indica se a última migração falhou no meio         |

## Possíveis Erros

| Erro                        | Causa                                              | Solução                                    |
|-----------------------------|----------------------------------------------------|--------------------------------------------|
| `no change`                 | Todas as migrações já foram aplicadas              | Nenhuma ação necessária                    |
| `dirty database version X`  | Uma migração anterior falhou parcialmente          | Corrija o SQL e use `migrate force X`      |
| `dial tcp: connection refused` | Banco de dados inacessível                      | Verifique se o MySQL está rodando          |
| `file does not exist`       | Diretório de migrações não encontrado              | Verifique o caminho em `-path`             |

---

# golang-migrate: Comando `migrate down`

Documentação detalhada sobre o comando utilizado para reverter migrações de banco de dados.

## Comando

```bash
migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose down
```

## O que faz?

O subcomando `down` **reverte migrações aplicadas** executando os arquivos `.down.sql` em ordem decrescente (da maior versão para a menor). Quando executado sem argumento numérico, **reverte TODAS as migrações**, retornando o banco ao estado inicial.

> **Atenção:** Executar `down` sem argumento remove **todas** as tabelas e dados criados pelas migrações. O CLI exibirá uma confirmação interativa antes de prosseguir.

## Anatomia do Comando

```
migrate -path=sql/migrations -database "mysql://root:admin@tcp(mysql-db:3306)/goexpert" -verbose down
```

| Parte                  | Descrição                                                      |
|------------------------|----------------------------------------------------------------|
| `migrate`              | Binário da CLI do golang-migrate                               |
| `-path=sql/migrations` | Caminho do diretório contendo os arquivos de migração          |
| `-database "..."`      | String de conexão com o banco de dados                         |
| `-verbose`             | Ativa logs detalhados de cada migração revertida               |
| `down`                 | Subcomando que reverte as migrações aplicadas                  |

## Detalhamento das Flags

As flags `-path`, `-database` e `-verbose` são idênticas às documentadas no comando `migrate up` acima.

### `down` (subcomando)

Reverte migrações já aplicadas no banco de dados.

- **Sem argumento:** Reverte **TODAS** as migrações (requer confirmação)
- **Com argumento numérico (`down N`):** Reverte apenas as últimas **N** migrações aplicadas

```bash
# Reverter TODAS as migrações (pede confirmação)
migrate -path=sql/migrations -database "..." -verbose down

# Reverter apenas a última migração
migrate -path=sql/migrations -database "..." -verbose down 1

# Reverter as últimas 3 migrações
migrate -path=sql/migrations -database "..." -verbose down 3
```

## Como funciona internamente

1. O migrate consulta a tabela `schema_migrations` para obter a versão atual
2. Identifica quais migrações precisam ser revertidas
3. Executa cada arquivo `.down.sql` em ordem **decrescente** de versão
4. Atualiza a tabela `schema_migrations` após cada reversão

**Exemplo de saída com `-verbose`:**
```
2024/03/08 10:05:00 Start buffering 1/d init
2024/03/08 10:05:00 Read and execute 1/d init
2024/03/08 10:05:00 Finished 1/d init (read 3.456ms, ran 8.789ms)
```

## Diferença entre `up` e `down`

| Aspecto           | `up`                                     | `down`                                    |
|-------------------|------------------------------------------|-------------------------------------------|
| Direção           | Aplica migrações (avança)                | Reverte migrações (retrocede)             |
| Arquivos usados   | `*.up.sql`                               | `*.down.sql`                              |
| Ordem de execução | Crescente (1 → 2 → 3)                   | Decrescente (3 → 2 → 1)                  |
| Sem argumento     | Aplica **todas** as pendentes            | Reverte **todas** as aplicadas            |
| Com argumento N   | Aplica as próximas N                     | Reverte as últimas N                      |
| Risco             | Baixo (adiciona estruturas)              | Alto (remove dados e estruturas)          |

## Possíveis Erros

| Erro                        | Causa                                              | Solução                                    |
|-----------------------------|----------------------------------------------------|--------------------------------------------|
| `no change`                 | Não há migrações para reverter                     | Nenhuma ação necessária                    |
| `dirty database version X`  | Uma migração anterior falhou parcialmente          | Corrija o SQL e use `migrate force X`      |
| `dial tcp: connection refused` | Banco de dados inacessível                      | Verifique se o MySQL está rodando          |

## Boas Práticas

- **Sempre escreva o `.down.sql`** correspondente ao criar uma migração. Sem ele, o `down` falhará.
- **Use `down 1`** em vez de `down` (sem argumento) para reverter apenas a última migração, evitando perda acidental de dados.
- **Teste o ciclo completo** (`up` → `down` → `up`) em ambiente de desenvolvimento antes de aplicar em produção.
- **Nunca execute `down` em produção** sem um plano claro de recuperação, pois dados podem ser perdidos permanentemente.

## Referência

- Repositório principal: https://github.com/golang-migrate/migrate
- CLI: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Bancos suportados: PostgreSQL, MySQL, SQLite, MongoDB, CockroachDB, e outros

---

# sqlc: Comando `sqlc generate`

Documentação detalhada sobre o comando utilizado para gerar código Go type-safe a partir de queries SQL.

## Comando

```bash
sqlc generate
```

## O que é o sqlc?

O [sqlc](https://sqlc.dev/) é uma ferramenta que gera código Go **type-safe** a partir de queries SQL puras. Em vez de escrever código de acesso ao banco manualmente (com `rows.Scan`, structs, etc.), você escreve apenas SQL e o sqlc gera todo o código Go correspondente automaticamente.

**Filosofia:** "Você escreve SQL. O sqlc gera código Go. Simples assim."

**Repositório:** https://github.com/sqlc-dev/sqlc

## Instalação

```bash
# Via Go
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Via Homebrew (macOS/Linux)
brew install sqlc

# Via Docker
docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

# Via snap (Linux)
sudo snap install sqlc
```

## O que faz?

O comando `sqlc generate` lê o arquivo de configuração (`sqlc.yaml`), analisa o schema do banco e as queries SQL definidas, e gera automaticamente três tipos de arquivo Go:

1. **`models.go`** — Structs Go que representam as tabelas do banco
2. **`db.go`** — Interface de conexão e struct `Queries` (o "repositório")
3. **`query.sql.go`** — Funções Go type-safe para cada query SQL definida

## Pré-requisitos

O comando exige dois elementos configurados antes de ser executado:

### 1. Arquivo de configuração `sqlc.yaml`

Deve existir na raiz do projeto (ou no diretório onde o comando é executado):

```yaml
version: "2"
sql:
- schema: "sql/migrations"
  queries: "sql/queries"
  engine: "mysql"
  gen:
    go:
      package: "db"
      out: "internal/db"
```

### 2. Arquivos SQL de queries com anotações

Arquivo(s) `.sql` no diretório de queries com comentários especiais do sqlc:

```sql
-- name: ListCategories :many
SELECT * FROM categories;
```

## Arquivo de Configuração (`sqlc.yaml`)

```yaml
version: "2"
sql:
- schema: "sql/migrations"
  queries: "sql/queries"
  engine: "mysql"
  gen:
    go:
      package: "db"
      out: "internal/db"
```

| Campo               | Valor               | Descrição                                                          |
|---------------------|---------------------|--------------------------------------------------------------------|
| `version`           | `"2"`               | Versão do formato de configuração do sqlc                          |
| `sql`               | lista               | Lista de configurações de geração (permite múltiplos bancos)       |
| `schema`            | `"sql/migrations"`  | Diretório contendo os arquivos de schema (DDL/migrações)           |
| `queries`           | `"sql/queries"`     | Diretório contendo os arquivos de queries SQL anotadas             |
| `engine`            | `"mysql"`           | Engine do banco de dados (`mysql`, `postgresql` ou `sqlite`)       |
| `gen.go.package`    | `"db"`              | Nome do package Go dos arquivos gerados                            |
| `gen.go.out`        | `"internal/db"`     | Diretório de destino dos arquivos Go gerados                       |

### Detalhamento dos campos

#### `schema`

Aponta para o diretório que contém os arquivos de definição do schema (DDL). O sqlc lê esses arquivos para entender a estrutura das tabelas, colunas e tipos. Neste projeto, reutiliza os mesmos arquivos de migração do golang-migrate.

#### `queries`

Aponta para o diretório com as queries SQL anotadas. Cada query deve ter um comentário especial que o sqlc usa para gerar a função Go correspondente.

#### `engine`

Define qual dialeto SQL o sqlc deve usar para interpretar o schema e as queries. Opções disponíveis:

| Engine       | Descrição                |
|--------------|--------------------------|
| `mysql`      | MySQL / MariaDB          |
| `postgresql` | PostgreSQL               |
| `sqlite`     | SQLite                   |

#### `gen.go.out`

Define o diretório de saída dos arquivos gerados. O diretório é criado automaticamente caso não exista.

## Anotações de Queries

O sqlc utiliza comentários especiais para entender como gerar o código Go para cada query.

### Formato

```sql
-- name: <NomeDaFuncao> :<tipo_de_retorno>
<QUERY SQL>;
```

### Tipos de retorno

| Anotação     | Retorno em Go                              | Uso                                          |
|--------------|--------------------------------------------|----------------------------------------------|
| `:one`       | `(Model, error)`                           | Queries que retornam exatamente uma linha     |
| `:many`      | `([]Model, error)`                         | Queries que retornam múltiplas linhas         |
| `:exec`      | `error`                                    | Comandos que não retornam dados (INSERT sem RETURNING, UPDATE, DELETE) |
| `:execresult`| `(sql.Result, error)`                      | Comandos que precisam do resultado (RowsAffected, LastInsertId)       |
| `:execrows`  | `(int64, error)`                           | Retorna apenas a quantidade de linhas afetadas |

### Exemplos de queries anotadas

```sql
-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description) VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ? WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;

-- name: CreateCourse :execresult
INSERT INTO courses (id, category_id, name, description, price) VALUES (?, ?, ?, ?, ?);
```

## Arquivos Gerados

Ao executar `sqlc generate`, três arquivos são criados no diretório `internal/db/`:

```
internal/db/
  db.go            # Interface DBTX, struct Queries e construtor New()
  models.go        # Structs Go correspondentes às tabelas do banco
  query.sql.go     # Funções Go geradas a partir das queries SQL
```

### `db.go` — Interface e struct de conexão

Gerado automaticamente. Contém a interface `DBTX` que abstrai a conexão com o banco, permitindo usar tanto `*sql.DB` quanto `*sql.Tx`.

```go
// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
```

| Elemento     | Descrição                                                                    |
|--------------|------------------------------------------------------------------------------|
| `DBTX`       | Interface que aceita tanto `*sql.DB` quanto `*sql.Tx`                        |
| `New(db)`    | Construtor que cria uma instância de `Queries`                               |
| `Queries`    | Struct que agrupa todas as funções de query geradas                          |
| `WithTx(tx)` | Cria uma nova instância de `Queries` associada a uma transação               |

### `models.go` — Structs das tabelas

Gerado a partir do schema SQL. Cada tabela vira uma struct Go com os tipos correspondentes.

```go
// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Category struct {
	ID          string
	Name        string
	Description sql.NullString
}

type Course struct {
	ID          string
	CategoryID  string
	Name        string
	Description sql.NullString
	Price       string
}
```

**Mapeamento de tipos SQL → Go:**

| Tipo SQL              | Tipo Go             | Observação                                      |
|-----------------------|---------------------|-------------------------------------------------|
| `VARCHAR(36) NOT NULL`| `string`            | Campos NOT NULL viram tipos primitivos          |
| `TEXT NOT NULL`       | `string`            | Idem                                             |
| `TEXT` (nullable)     | `sql.NullString`    | Campos nullable usam tipos `sql.Null*`          |
| `DECIMAL(10,2)`      | `string`            | MySQL DECIMAL é mapeado para string             |

### `query.sql.go` — Funções de query

Gerado a partir das queries anotadas. Cada query vira uma função Go type-safe na struct `Queries`.

Para a query:
```sql
-- name: ListCategories :many
SELECT * FROM categories;
```

O sqlc gera:
```go
// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package db

import (
	"context"
)

const listCategories = `-- name: ListCategories :many
SELECT id, name, description FROM categories
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
```

**Pontos importantes:**

- O `SELECT *` é expandido automaticamente para todas as colunas (`SELECT id, name, description`)
- O `rows.Scan` já mapeia cada coluna para o campo correto da struct
- O tratamento de erros é completo (erro na query, erro no scan, erro no close)
- A assinatura da função reflete a anotação `:many` → retorna `[]Category`

## Uso no código Go

```go
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"seu-modulo/internal/db"
)

func main() {
	conn, err := sql.Open("mysql", "root:admin@tcp(mysql-db:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	categories, err := queries.ListCategories(context.Background())
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Printf("ID: %s | Nome: %s\n", c.ID, c.Name)
	}
}
```

## Fluxo completo de trabalho

```
┌─────────────────────┐     ┌──────────────────────┐     ┌────────────────────┐
│  sql/migrations/    │     │  sql/queries/         │     │  sqlc.yaml         │
│  000001_init.up.sql │     │  query.sql            │     │  (configuração)    │
│  (schema DDL)       │     │  (queries anotadas)   │     │                    │
└────────┬────────────┘     └──────────┬───────────┘     └─────────┬──────────┘
         │                             │                           │
         └─────────────────┬───────────┘───────────────────────────┘
                           │
                    sqlc generate
                           │
              ┌────────────┼────────────┐
              ▼            ▼            ▼
      ┌──────────┐  ┌──────────┐  ┌──────────────┐
      │  db.go   │  │models.go │  │query.sql.go  │
      │(conexão) │  │(structs) │  │(funções)     │
      └──────────┘  └──────────┘  └──────────────┘
                    internal/db/
```

## Possíveis Erros

| Erro                                         | Causa                                              | Solução                                              |
|----------------------------------------------|----------------------------------------------------|----------------------------------------------------|
| `no queries found`                           | Nenhum arquivo `.sql` no diretório de queries      | Crie arquivos com queries anotadas em `sql/queries/` |
| `query references unknown table`             | Query usa tabela não definida no schema            | Verifique se o schema em `sql/migrations/` está correto |
| `invalid query annotation`                   | Comentário de anotação com formato incorreto       | Use o formato `-- name: NomeFuncao :tipo`            |
| `sqlc.yaml: file not found`                  | Arquivo de configuração ausente                    | Crie o `sqlc.yaml` na raiz do projeto                |
| `column "x" is not part of table "y"`        | Query referencia coluna inexistente                | Verifique o schema da tabela                         |

## sqlc vs código manual

| Aspecto                | sqlc                              | Código manual                      |
|------------------------|-----------------------------------|------------------------------------|
| Tipo de segurança      | Type-safe (verificado em build)   | Erros descobertos em runtime       |
| Manutenção             | Regenera automaticamente          | Atualização manual em cada mudança |
| Quantidade de código   | Gerado (zero boilerplate)         | Muito boilerplate repetitivo       |
| Performance            | Queries puras (sem ORM overhead)  | Queries puras                      |
| Curva de aprendizado   | Baixa (escreva SQL normal)        | Baixa                              |
| Flexibilidade          | Limitada às features do sqlc      | Total                              |

## Referência

- Site oficial: https://sqlc.dev/
- Repositório: https://github.com/sqlc-dev/sqlc
- Documentação: https://docs.sqlc.dev/
- Playground online: https://play.sqlc.dev/
