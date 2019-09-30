westwallet-golang-api
=====================
.. image:: https://godoc.org/github.com/WestWallet/westwallet-golang-api?status.svg
    :alt: godoc
    :target: http://godoc.org/github.com/WestWallet/westwallet-golang-api

westwallet-golang-api is a `WestWallet Public API <https://westwallet.info/api_docs>`_ wrapper for Go programming language. Use it for building payment solutions.

Installing
----------

Install with go get:

.. code-block:: text

    go get github.com/westwallet/westwallet-golang-api


Create withdrawal example
-------------------------

.. code-block:: Go

    package main

    import (
        "fmt"
        westwallet "github.com/westwallet/westwallet-golang-api"
    )
    // Sending 0.1 ETH to 0x57689002367b407f031f1BB5Ef2923F103015A32
    client := westwallet.APIClient{
        Key:      "your_public_key",
        Secret:   "your_private_key",
    }
    transaction, err := client.CreateWithdrawal(
        "ETH", "0.1", "0x57689002367b407f031f1BB5Ef2923F103015A32", "", ""
    )
    fmt.Println(err)
    if err != nil {
        panic(err)
    }
    fmt.Println(transaction)


Generate address example
-------------------------

.. code-block:: Go

    package main

    import (
        "fmt"
        westwallet "github.com/westwallet/westwallet-golang-api"
    )
    client := westwallet.APIClient{
        Key:      "your_public_key",
        Secret:   "your_private_key",
    }
    address, err := client.GenerateAddress("BTC", "", "")
    if err != nil {
        panic(err)
    }
    fmt.Println(address.Address)

Documentation
-------------
* API: https://westwallet.info/api_docs

Other languages
---------------
* Python: https://github.com/WestWallet/westwallet-python-api
* JavaScript: https://github.com/WestWallet/westwallet-js-api
* PHP: https://github.com/WestWallet/westwallet-php-api
