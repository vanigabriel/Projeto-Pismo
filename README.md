# Projeto Piso
    API para implementação do desafio proposto pela Pismo.

    Cada cliente possue uma conta e cada transação realizada está vinculada à uma conta. Cada transação possui um tipo, um valor, um saldo e uma data de lançamento.

    O sistema deve ser capaz de criar, atualizar e listar contas. Também deve ser capaz de criar uma transação, trabalhando com o conceito do saldo em aberto das mesmas, sendo capaz de recalculado o saldo em aberto após um pagamento.
    Implementado em Golang e utilizando o framework Gin.

# Documentação API
https://documenter.getpostman.com/view/7995657/SVtbPkFJ

# Características
- Docker;
- Princípios de Clean Architecture;
- Govalidator para conferir campos obrigatórios;
- Log enviado para o Timber.io;
- TDD;
