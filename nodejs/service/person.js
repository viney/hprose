(function(){
    'use strict'

    var Person = {
        content: ''
    };

    Person.Say = function(content){
        this.content = content;
    }

    Person.Hello = function(){
        return 'Hello ' + this.content + '!';
    }

    module.exports = Person;
})();
