var loginApp = angular.module("loginApp", []);
loginApp.service("loginService", function(){
    this.login = function(user, password){
        if(user && password)
            $('#login-info').html("<p>Chúc mừng <b>" + user + "</b> đã đăng nhập thành công.</p>");
        else
            $('#login-info').html("<p>Không được để trống username và/hoặc password</p>");
    };
});
loginApp.controller("loginController", function($scope, loginService){
    $scope.login = function(){
        loginService.login($scope.user, $scope.password);
    };
});