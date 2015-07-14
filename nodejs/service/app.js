(function(){
    'use strict'

    var hprose = require("hprose");
    var Person = require("./person");

    var service = hprose.Server.create('http://127.0.0.1:7080');
    service.addInstanceMethods(Person);
    service.start();
})();
