var apiUrl = {
    "createUserApi" : "/users/create",
    "getUserApi" : "/users/get",
};

var userManagementApp = angular.module("userManagementApp", []);
userManagementApp.service("userManagementService", function(){
    this.create = function(data){
        if(data.username && data.password) {
            $.ajax({
                type        : 'POST',
                url         : apiUrl.createUserApi,
                data        : data,
                dataType    : 'json',
                success     : function (response) {
                    console.log(response);
                },
                error       : function () {
                    console.log("error!")
                }
            });
            $('#create-info').html("<p>Chúc mừng <b>" + data.username + "</b> đã đăng ký thành công.</p>");
        }
        else
            $('#create-info').html("<p>Không được để trống username và/hoặc password</p>");
    };
});
userManagementApp.controller("createUserController", function($scope, userManagementService){
    $scope.create = function(){
        var data = {
            "username" : $scope.username,
            "password" : $scope.password
        };
        userManagementService.create(data);
    };
});