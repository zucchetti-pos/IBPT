CREATE TABLE `produto` (
    `codigo` VARCHAR(10) not null,
    `Uf` VARCHAR(100) null,
    `Ex` int null,
    `Descricao` VARCHAR(100) null,
    `Nacional` real null,
    `Estadual` real null,
    `Importado` real null,
    `Municipal` real null,
    `Tipo` VARCHAR(100) null,
    `VigenciaInicio` VARCHAR(100) null,
    `VigenciaFim` VARCHAR(100) null,
    `Chave` VARCHAR(100) null,
    `Versao` VARCHAR(100) null,
    `Fonte` VARCHAR(100) null
);

insert into produto values("70101000","SC",0,"Ampolas de vidro,p/transporte ou embalagem",13.45,17.00,18.02,0.00,"0","01/01/2017","30/06/2017","W7m9E1","17.1.A","IBPT");
