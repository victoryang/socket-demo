'use strict'
const net = require('net');

var id = 0;

class client {
    constructor() {
        console.log("client created!");
        this._id = id++;
        this._sock = net.Socket();
        this._sock.setEncoding('utf-8');
    }

    connect() {
        this._sock.connect({
            port:30000
        });
    }
}

var c = new client();
console.log(c.id);
