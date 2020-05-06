var apiUrl = {
    "createUser" : "/users/create",
    "getUsers" : "/users/get",
    "getUser" : "/users/get/",
    "getUsersPage" : "/users"
};

var userManagementApp = angular.module("userManagementApp", []);
userManagementApp.controller("createUserController", function($scope, userManagementService){
    $scope.createUser = function(){
        var data = {
            "username" : $scope.username,
            "password" : $scope.password
        };
        userManagementService.createUser(data);
    };
    $scope.Users = userManagementService.getUsers();
    $scope.showCreateUser = function () {
        userManagementService.showCreateUser();
    }
});

userManagementApp.directive("ngCreateUserForm", function () {
    return {
        templateUrl : '/template/create.html'
    }
});

userManagementApp.directive("ngListUser", function () {
    return {
        templateUrl : '/template/list.html'
    };
});

userManagementApp.service("userManagementService", function(){
    this.createUser = function(data){
        if(data.username && data.password) {
            $.ajax({
                type        : 'POST',
                url         : apiUrl.createUser,
                data        : data,
                dataType    : 'json',
                success     : function (response) {
                    console.log(response);
                },
                error       : function () {
                    console.log("error when createUser!")
                }
            });
            $('#create-info').html("<p>Chúc mừng <b>" + data.username + "</b> đã đăng ký thành công.</p>");
        }
        else
            $('#create-info').html("<p>Không được để trống username và/hoặc password</p>");
    };
    this.getUsers = function () {
        $.ajax({
            type            : 'GET',
            url             : apiUrl.getUsers,
            success         : function (data) {
                console.log(data)
            },
            error           : function () {
                console.log("error when getUsers!")
            }
        });
    };
    this.showCreateUser = function () {
        $.ajax({
            type            : 'GET',
            url             : apiUrl.getUsersPage,
            success         : function (data) {
                console.log(data)
            },
            error           : function () {
                console.log("error when getUsers!")
            }
        });
    };
});
