angular.module('app').controller('CreateConnectionController', ['MainService', '$window', 'Notification', '$state', function(MainService, $window, Notification, $state) {
    var _self = this

    _self.connectionType = ['MySQL', 'Salesforce']
    _self.selectedConnection = ''
    
    _self.salesforceFormData = {
        'SalesforceName': '',
        'SalesforceUsername': '',
        'SalesforcePassword': '',
        'SalesforceSecurityToken': '',
        'SalesforceAPIVersion': '',
        'Type': 'Salesforce'
    }

    _self.requestInProgress = false

    _self.createSalesforceConnection = function() {
        _self.requestInProgress = true

        MainService.postSalesforceConnection(_self.salesforceFormData).then(function(result) {
            _self.requestInProgress = false
			if(result.Success) {
                Notification(result.Data)
                $state.go('connections')
            } else {
                Notification(result.Data)                
            }
		})
    }

    _self.mySQLFormData = {
        'MySQLName': '',
        'MySQLServer': '',
        'MySQLPort': '',
        'MySQLUserID': '',
        'MySQLPassword': '',
        'MySQLDatabase': '',
        'Type': 'MySQL'
    }

    _self.createMySQLConnection = function() {
        _self.requestInProgress = true

        MainService.postMySQLConnection(_self.mySQLFormData).then(function(result) {
            _self.requestInProgress = false
			if(result.Success) {
                Notification(result.Data)
                $state.go('connections')
            } else {
                Notification(result.Data)                
            }
		})
    }
}])
