angular.module('app').controller('MainController', ['$scope', '$transitions', '$state', 'AuthService', 'Notification', function($scope, $transitions, $state, AuthService, Notification) {
	var _self = this

    _self.showSignUp = false
    _self.showLogIn = true

    _self.toggleSignUp = function() {
        _self.showSignUp = true
        _self.showLogIn = false
    }

    _self.toggleLogIn = function() {
        _self.showLogIn = true
        _self.showSignUp = false    
    }

    _self.signUpFormData = {
        'Name': '',
        'Email': '',
        'Password': ''
    }

    _self.signInFormData = {
        'Email': '',
        'Password': ''
    }

    

    _self.signIn = function() {
        AuthService.signIn(_self.signInFormData).then(function(result) {
			if(result.Success) {
                window.location.reload()
            } else {
                Notification(result.Data)                
            }
		})
    }

    _self.signUp = function() {
        AuthService.signUp(_self.signUpFormData).then(function(result) {
			if(result.Success) {
				Notification(result.Data)
            }
		})
    }
}])
