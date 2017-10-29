angular.module('app').controller('ConnectionsController', ['MainService', '$window', 'Notification', 'connections', function(MainService, $window, Notification, connections) {
    var _self = this
    
    _self.connections = connections.Data
    console.log(_self.connections)
}])
