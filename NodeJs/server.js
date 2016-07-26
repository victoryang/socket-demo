const http = require("http");
var id = 0;
var servlist = null;

class server {
    constructor(){
        console.log("server created!");
        this._id = id++;
        this._sock = -1;
        this._nick = "nick";
        this._set_default();
    }

    connect(hostname, port, no_login){
        
    };
    disconnect(){};
    cleanup(){};
    flush_queue(){};
    auto_reconnect(){};
    finish(){console.log("finished!");};
    _set_default() {
        if(this.encoding) 
            _set_encoding();
    };
    _set_encoding() {
        this.encoding = "UTF-8";
    };
}

var serv = new server();
serv.finish();
console.log(serv._id);
serv._set_default();
module.export = server;
