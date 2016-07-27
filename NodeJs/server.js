'use strict'
const net = require('net');
const fs = require('fs');

var id = 0;
var servlist = null;

class server {
    constructor(){
        console.log("server created!");
        this._id = id++;
        this._nick = "nick";
        this._use_ssl = false;
        this._encoding = 'utf-8';
        this._server = new net.createServer((c) => {
            console.log("server created!");
            c.on('end', () => {
                console.log ("server ends!");
            })
        });
    }

    listen (port,hostname,no_login) {
        console.log ("listen!");
        this._port = port;
        this._host = hostname;
        this._server.listen(port,hostname,() => {
            console.log(this._server.address());
        });
    }

    close(){
        this._server.close();
    }
}

var serv = new server();
console.log(serv._id);
serv.listen(30000);
setTimeout(() => {serv.close()},3000);

