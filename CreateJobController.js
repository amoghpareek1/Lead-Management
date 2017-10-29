angular.module('app').controller('CreateJobController', ['MainService', '$window', 'Notification', '$state', 'connections', function(MainService, $window, Notification, $state, connections) {
    var _self = this

    _self.connections = connections.Data

    _self.requestInProgress = false    

    _self.connectionNames = []
    for(var i = 0; i < _self.connections.length; i++) {
        if(_self.connections[i].SalesforceName !== "") {
            _self.connectionNames.push(_self.connections[i].SalesforceName)        
        }

        if(_self.connections[i].MySQLName !== "") {
            _self.connectionNames.push(_self.connections[i].MySQLName)        
        }
    }

    _self.selectedSourceConnection = ''
    _self.selectedTargetConnection = ''

    _self.jobFormData = {
        'Name': '',
        'SourceConnectionID': '',
        'SourceConnectionType': '',
        'TargetConnectionID': '',
        'TargetConnectionType': ''
    }

    _self.createJob = function() {
        for(var i = 0; i < _self.connections.length; i++) {
            
        }
        _self.requestInProgress = true
        MainService.postJob(_self.jobFormData).then(function(result) {
            _self.requestInProgress = false
            if(result.Success) {
                Notification(result.Data)
            } else {
                Notification(result.Data)
            }
        }) 
    }
}])
