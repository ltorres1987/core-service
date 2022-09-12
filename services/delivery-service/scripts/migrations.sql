DROP DATABASE IF EXISTS deliverydb;
CREATE DATABASE IF NOT EXISTS deliverydb;
USE deliverydb;
CREATE TABLE users
(
    id            INT NOT NULL AUTO_INCREMENT,
    name          VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    application   VARCHAR(100) NOT NULL,
    created_user  VARCHAR(100) NOT NULL,
    created_at    DATETIME    NOT NULL,
    updated_user  VARCHAR(100) NOT NULL,
    updated_at    DATETIME    NOT NULL,
    status   VARCHAR(1)   NOT NULL,
    PRIMARY KEY (id)
)ENGINE=InnoDB CHARACTER SET utf8;

CREATE TABLE product
(
    id            INT NOT NULL AUTO_INCREMENT,
    name         VARCHAR(100) NOT NULL ,
    short_description VARCHAR(100) NOT NULL,
    created_user  VARCHAR(100) NOT NULL,
    created_at    DATETIME    NOT NULL,
    updated_user  VARCHAR(100) NOT NULL,
    updated_at    DATETIME    NOT NULL,
    status   VARCHAR(1)   NOT NULL,
    PRIMARY KEY (id)
)ENGINE=InnoDB CHARACTER SET utf8;

CREATE TABLE shipping_order
(
    id            INT NOT NULL AUTO_INCREMENT,
    idSender        VARCHAR(200) NOT NULL ,
    fullNameSender  VARCHAR(200) NOT NULL ,
    phoneSender     VARCHAR(200) NOT NULL,
    emailSender     VARCHAR(200) NOT NULL,
    idRecipient         VARCHAR(200) NOT NULL ,
    fullNameRecipient   VARCHAR(200) NOT NULL ,
    phoneRecipient      VARCHAR(200) NOT NULL ,
    emailRecipient      VARCHAR(200) NOT NULL ,
    latOrigin       VARCHAR(50) NOT NULL ,
    lngOrigin       VARCHAR(50) NOT NULL ,
    addressOrigin   VARCHAR(200) NOT NULL ,
    countryOrigin   VARCHAR(200) NOT NULL ,
    zipcodeOrigin   VARCHAR(200) NOT NULL ,
    referenceOrigin VARCHAR(200) NOT NULL ,
    latDestination       VARCHAR(50) NOT NULL ,
    lngDestination       VARCHAR(50) NOT NULL ,
    addressDestination   VARCHAR(200) NOT NULL ,
    countryDestination   VARCHAR(200) NOT NULL ,
    zipcodeDestination   VARCHAR(200) NOT NULL ,
    referenceDestination VARCHAR(200) NOT NULL ,
    packageSize     VARCHAR(1) NOT NULL ,
    quantityProduct INT NOT NULL,
    weightProduct   INT NOT NULL,
    orderStatus  VARCHAR(200) NOT NULL,
    created_user  VARCHAR(200) NOT NULL,
    created_at    DATETIME    NOT NULL,
    updated_user  VARCHAR(200) NOT NULL,
    updated_at    DATETIME    NOT NULL,
    status   VARCHAR(1)   NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET utf8;

CREATE TABLE package_size
(
    id              INT NOT NULL AUTO_INCREMENT,
    name            VARCHAR(100) NOT NULL ,
    nemo            VARCHAR(1) NOT NULL,
    limitvalue      INT NOT NULL,
    created_user    VARCHAR(100) NULL,
    created_at      DATETIME    NULL,
    updated_user    VARCHAR(100) NULL,
    updated_at      DATETIME    NULL,
    status          VARCHAR(1)   NULL,
    PRIMARY KEY (id)
)ENGINE=InnoDB CHARACTER SET utf8;

-- Create DML.
INSERT INTO package_size(id, name, nemo, limitvalue, created_user, created_at, updated_user,updated_at, status) VALUES
(1, '0 hasta 5kg', 'S', 5, "luis.torres", UTC_TIMESTAMP(), "luis.torres", UTC_TIMESTAMP(), "A"),
(2, 'hasta 15kg', 'M', 15, "luis.torres", UTC_TIMESTAMP(), "luis.torres", UTC_TIMESTAMP(), "A"),
(3, 'hasta 25kg', 'L', 25, "luis.torres", UTC_TIMESTAMP(), "luis.torres", UTC_TIMESTAMP(), "A");