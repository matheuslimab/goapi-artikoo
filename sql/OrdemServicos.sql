CREATE TABLE OrdemServicos (
    id_ordem_servico INT AUTO_INCREMENT PRIMARY KEY,
    nome_servico VARCHAR(255),
    descricao_servico TEXT,
    id_cliente INT,
    id_fornecedor INT,
    data_inicio DATE,
    data_termino DATE,
    preco_servico DECIMAL(10, 2),
    condicoes_pagamento VARCHAR(255),
    status_servico VARCHAR(50),
    responsavel VARCHAR(255),
    observacoes TEXT,
    FOREIGN KEY (id_cliente) REFERENCES Clients(id_client)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;